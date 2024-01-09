package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"MetroLab/internal/entity"
)

const getUserQuery = `SELECT id, name FROM users.user WHERE tg_id = $1;`

func (s *PGStorage) GetUser(ctx context.Context, userID int64) (*entity.User, error) {
	rows, err := s.db.Query(ctx,
		getUserQuery,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readUser(rows)
}

func readUser(rows pgx.Rows) (*entity.User, error) {
	user := &entity.User{}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Name,
		); err != nil {
			return nil, errors.Wrap(err, "scan user from row")
		}
	}

	return user, nil
}
