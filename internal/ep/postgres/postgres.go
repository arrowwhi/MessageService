package postgres

import (
	"MessageService/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"go.uber.org/zap"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultMaxPoolSize = 1

type PgxPool interface {
	Close()
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Ping(ctx context.Context) error
}

type Postgres struct {
	maxPoolSize int
	Builder     squirrel.StatementBuilderType
	Pool        PgxPool
}

func New(ctx context.Context, cfg *config.Postgres, zapLogger *zap.Logger, opts ...Option) (*Postgres, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Database)
	pg := &Postgres{
		maxPoolSize: defaultMaxPoolSize,
	}

	for _, opt := range opts {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("pgdb - New - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		zapLogger.Error("Failed to connect to Postgres", zap.Error(err))
		return nil, fmt.Errorf("pgdb - New - pgxpool.NewWithConfig: %w", err)
	}

	if err = pg.Pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgdb - Pool.Ping: %w", err)
	}
	zapLogger.Info("Successfully connected to Postgres")

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
