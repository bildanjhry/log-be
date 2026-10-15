package repository

import (
	"context"

	"github.com/bildanjhry/log-be/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindAll(ctx context.Context) ([]model.UserRes, error) {
	rows, err := r.db.Query(ctx,
		`SELECT "users"."id", "users"."email", "user_profile"."name" FROM "users"
		LEFT JOIN "user_profile" ON "user_profile"."id_user" = "users"."id" `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]model.UserRes, 0)

	for rows.Next() {
		var user model.UserRes

		if err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Name,
			&user.Phone,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
