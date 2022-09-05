package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Save(ctx context.Context, u *User) error {
	const query = "INSERT INTO users (id, username, password) VALUES ($1, $2, $3);"
	cmd, err := r.pool.Exec(ctx, query, u.ID, u.Username, u.Password)
	if err == nil && cmd.RowsAffected() == 0 {
		err = ErrAlreadyExists
	}
	return err
}

func (r *postgresRepository) Find(ctx context.Context, username string) (*User, error) {
	const query = "SELECT * FROM users WHERE username = $1;"
	u := User{}
	err := r.pool.QueryRow(ctx, query, username).Scan(&u.ID, &u.Username, &u.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return nil, err
	}
	return &u, err
}
