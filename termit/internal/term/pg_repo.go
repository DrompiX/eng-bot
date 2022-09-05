package term

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
)

var (
	TermAlreadyExists = errors.New("specified term already exists")
)

// connTx is a minimal shared interface for pgx.Tx, pgx.Conn and pgxpool.Pool
// can be used when you need to dynamically store these in the same variable
type connTx interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type postgresRepository struct {
	pool *pgxpool.Pool
	conn connTx
}

func NewPostgresRepository(pool *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{pool: pool, conn: pool}
}

func (r *postgresRepository) Atomic(ctx context.Context, fn func(r Repository) error) (err error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("transaction startup failed: %s", err.Error())
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit(ctx)
		}
	}()

	injectedRepo := &postgresRepository{
		pool: r.pool,
		conn: tx,
	}
	err = fn(injectedRepo)
	return
}

func (r *postgresRepository) AddTerm(ctx context.Context, t *Term) (err error) {
	const query = "INSERT INTO terms (uid, term, translation) VALUES ($1, $2, $3);"
	cmd, err := r.conn.Exec(ctx, query, t.Uid, t.Data, t.Translation)
	if err == nil && cmd.RowsAffected() == 0 {
		err = TermAlreadyExists
	}
	return
}

func (r *postgresRepository) GetAllTerms(ctx context.Context, uid user.UserID) ([]*Term, error) {
	const query = "SELECT * FROM terms WHERE uid = $1;"
	rows, err := r.conn.Query(ctx, query, uid)
	if err != nil {
		return nil, err
	}

	terms := make([]*Term, 0)
	for rows.Next() {
		var t Term
		err := rows.Scan(&t.Uid, &t.Data, &t.Translation)
		if err != nil {
			return nil, err
		}
		terms = append(terms, &t)
	}

	return terms, nil
}

func (r *postgresRepository) GetTermCount(ctx context.Context, uid user.UserID) (cnt int, err error) {
	const query = "SELECT COUNT(*) FROM terms WHERE uid = $1;"
	err = r.conn.QueryRow(ctx, query, uid).Scan(&cnt)
	return
}

func (r *postgresRepository) GetTermByName(ctx context.Context, uid user.UserID, termName string) (*Term, error) {
	const query = "SELECT * FROM terms WHERE uid = $1 AND term = $2;"
	var t Term
	err := r.conn.QueryRow(ctx, query, uid, termName).Scan(&t.Uid, &t.Data, &t.Translation)
	if errors.Is(err, pgx.ErrNoRows) {
		err = errors.New("not found")
		return nil, err
	}
	return &t, err
}

func (r *postgresRepository) UpdateTerm(ctx context.Context, t *Term) error {
	panic("out of MVP scope")
}

func (r *postgresRepository) DeleteTerm(ctx context.Context, t *Term) error {
	panic("out of MVP scope")
}
