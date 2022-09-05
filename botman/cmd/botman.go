package main

import (
	"log"
	"os"

	"gitlab.ozon.dev/DrompiX/homework-2/botman/config"
	"gitlab.ozon.dev/DrompiX/homework-2/botman/internal/telegram"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
	"gitlab.ozon.dev/DrompiX/homework-2/botman/pb"
	"google.golang.org/grpc"
)

func main() {
	confPath := os.Getenv("CONFIG_PATH")
	if confPath == "" {
		confPath = "./config/config.yml"
	}

	f := config.Read(confPath)
	conf := config.Parse(f)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(conf.ServerConfig.URL, opts...)
	if err != nil {
		log.Fatalf("could not establish connection: %s", err)
	}
	defer conn.Close()

	termitClient := &tm.GrpcTermitClient{
		TermServ: pb.NewTermServiceClient(conn),
		TaskServ: pb.NewTaskServiceClient(conn),
	}

	log.Printf("Starting telegram bot polling...")
	bc := telegram.NewTGBotClient(conf.BotConfig.Token, termitClient, conf.BotConfig.Debug)
	bc.PollForUpdates()
}
