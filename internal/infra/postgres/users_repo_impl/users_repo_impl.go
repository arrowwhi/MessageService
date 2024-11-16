package users_repo_impl

import (
	postgreslib "MessageService/internal/ep/postgres"
	"MessageService/internal/interfaces/infra/postgres"
	"MessageService/internal/interfaces/service"
	"context"
	"fmt"
	"go.uber.org/zap"
)

var _ postgres.Users = (*impl)(nil)

type impl struct {
	log *zap.Logger
	pg  *postgreslib.Postgres
}

func New(log *zap.Logger, pg *postgreslib.Postgres) postgres.Users {
	return &impl{log: log, pg: pg}
}

func (i impl) GetStatusInfo(ctx context.Context, ids []int64) ([]*service.User, error) {
	query := `select id, username, status, EXTRACT(EPOCH FROM last_active_at)::BIGINT
	from users where id in ($1)`
	args := []interface{}{ids}
	rows, err := i.pg.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	users := make([]*service.User, 0)

	for rows.Next() {
		var elem service.User
		if err = rows.Scan(
			&elem.UserId, &elem.Username,
			&elem.Status,
			&elem.LastActive,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		users = append(users, &elem)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}
	return users, nil

}

func (i impl) UpdateStatus(ctx context.Context, userId int64, online bool) error {
	query := `
		UPDATE users 
		SET status = $1, last_active_at = NOW()
		WHERE id = $2;
	`
	args := []interface{}{online, userId}

	_, err := i.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		// todo error no data
		return err
	}
	return nil
}
