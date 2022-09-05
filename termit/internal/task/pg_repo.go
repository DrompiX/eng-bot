package task

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrCollectionIsEmpty = errors.New("term collection is empty")
	ErrTaskAlreadyExist  = errors.New("task with same id already exists")
	ErrNotFound          = errors.New("not found")
)

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, t *task) error {
	const query = "INSERT INTO tasks (id, uid, term, expected, success) VALUES ($1, $2, $3, $4, $5);"
	cmd, err := r.pool.Exec(ctx, query, t.ID, t.Uid, t.Term, t.Expected, t.Success)
	if err == nil && cmd.RowsAffected() == 0 {
		err = ErrTaskAlreadyExist
	}
	return err
}

func (r *postgresRepository) Update(ctx context.Context, t *task) error {
	const query = "UPDATE tasks SET success = $2 WHERE id = $1"
	cmd, err := r.pool.Exec(ctx, query, t.ID, t.Success)
	if err == nil && cmd.RowsAffected() == 0 {
		err = ErrNotFound
	}
	return err
}

func (r *postgresRepository) GetById(ctx context.Context, tid TaskID) (t task, err error) {
	const query = "SELECT id, uid, term, expected, success FROM tasks WHERE id = $1;"
	err = r.pool.QueryRow(ctx, query, tid).Scan(&t.ID, &t.Uid, &t.Term, &t.Expected, &t.Success)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}
	return
}
