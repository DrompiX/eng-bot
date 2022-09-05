package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/config"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/middleware"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/task"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/term"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TermItServer struct {
	OrderedInterceptors []grpc.UnaryServerInterceptor
	TermServer          pb.TermServiceServer
	TaskServer          pb.TaskServiceServer
}

func (s *TermItServer) run(listener net.Listener) error {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(s.OrderedInterceptors...),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTermServiceServer(grpcServer, s.TermServer)
	pb.RegisterTaskServiceServer(grpcServer, s.TaskServer)

	log.Printf("Starting GRPC TermItServer at %s", listener.Addr())
	return grpcServer.Serve(listener)
}

func dbConnPool(ctx context.Context, conf config.TermitConfig) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Db.Host, conf.Db.Port, conf.Db.User, conf.Db.Password, conf.Db.Dbname,
	)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}

func startHttpProxy(conf config.TermitConfig) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcAddr := fmt.Sprintf("localhost:%d", conf.Server.GrpcPort)
	err := pb.RegisterTermServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		return err
	}

	log.Printf("Starting HTTP TermItServer at localhost:%d", conf.Server.HTTPPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", conf.Server.HTTPPort), mux)
}

func main() {
	// Read configurations
	confPath := os.Getenv("CONFIG_PATH")
	if confPath == "" {
		confPath = "./config/config.yml"
	}
	f := config.Read(confPath)
	conf := config.Parse(f)

	// Prepare database connection pool
	pool := dbConnPool(context.Background(), conf)

	// Term functionality-related inits
	termRepo := term.NewPostgresRepository(pool)
	termServer := term.NewServer(term.NewService(termRepo, 20))

	// User functionality-related inits
	userRepo := user.NewPostgresRepository(pool)
	userService := user.NewService(userRepo)

	// Task functionality-related inits
	taskRepo := task.NewPostgresRepository(pool)
	taskServer := task.NewServer(task.NewService(taskRepo, termRepo))

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", conf.Server.GrpcPort))
	if err != nil {
		log.Fatalf("could not start server: %s", err)
	}

	server := TermItServer{
		OrderedInterceptors: []grpc.UnaryServerInterceptor{
			middleware.NewLogInterceptor().Unary(),
			middleware.NewAuthInterceptor(userService).Unary(),
		},
		TermServer: termServer,
		TaskServer: taskServer,
	}

	// Run GRPC server
	go func() {
		if err := server.run(listener); err != nil {
			log.Fatalf("could not start grpc server: %s", err)
		}
	}()

	// Wait a bit for grpc server launch
	time.Sleep(2 * time.Second)

	// Run HTTP server as a proxy to GRPC
	if err := startHttpProxy(conf); err != nil {
		log.Fatalf("could not start http server: %s", err)
	}

}
