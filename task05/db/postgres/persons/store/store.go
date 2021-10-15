package store

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgconn/stmtcache"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sanya-spb/Go-Postgres/app/repos/persons"
)

const (
	PG_HOST     = "127.0.0.1"
	PG_PORT     = "5432"
	PG_USER     = "sanya"
	PG_PASSWORD = "passwd"
	PG_DB_NAME  = "sauna"

	MaxConns = 10
	MinConns = 2
)

var _ persons.PersonsStore = &Persons{}

type Persons struct {
	Pool *pgxpool.Pool
}

func NewPersons() *Persons {
	pgpool, err := createPGXPoolModeDescribe(MaxConns, MinConns)
	if err != nil {
		log.Fatalf("failed to create a PGX pool: %s", err.Error())
	}
	return &Persons{
		Pool: pgpool,
	}
}

func (p *Persons) GetPerson(ctx context.Context, fName string, lName string) (*persons.TPerson, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	const sql = `
	-- информация о навыках конкретного сотрудника
	select
		id,
		fname,
		lname,
		phone,
		email
	from
		personal
	where
		fname = $1
		and lname = $2;`
	result := persons.TPerson{}
	err := p.Pool.QueryRow(ctx, sql, fName, lName).Scan(
		&result.ID,
		&result.FNname,
		&result.LName,
		&result.Phone,
		&result.Email,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the personal service: %w", err)
	}

	return &result, nil
}

func createPGXPoolModeDescribe(maxConns int32, minConns int32) (*pgxpool.Pool, error) {
	cfg, err := getPoolConfig(maxConns, minConns)
	if err != nil {
		return nil, fmt.Errorf("failed to get the pool config: %w", err)
	}
	cfg.ConnConfig.BuildStatementCache = func(conn *pgconn.PgConn) stmtcache.Cache {
		mode := stmtcache.ModeDescribe
		capacity := 512
		return stmtcache.New(conn, mode, capacity)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize a connection pool: %w", err)
	}
	return pool, nil
}

func getPoolConfig(maxConns int32, minConns int32) (*pgxpool.Config, error) {
	connStr := ComposeConnectionString()
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create a pool config from connection string %s: %w", connStr, err,
		)
	}
	cfg.MaxConns = maxConns
	cfg.MinConns = minConns

	// HealthCheckPeriod - частота проверки работоспособности
	// соединения с Postgres
	cfg.HealthCheckPeriod = 1 * time.Minute

	// MaxConnLifetime - сколько времени будет жить соединение.
	// Так как большого смысла удалять живые соединения нет,
	// можно устанавливать большие значения
	cfg.MaxConnLifetime = 24 * time.Hour

	// MaxConnIdleTime - время жизни неиспользуемого соединения,
	// если запросов не поступало, то соединение закроется.
	cfg.MaxConnIdleTime = 30 * time.Minute

	// ConnectTimeout устанавливает ограничение по времени
	// на весь процесс установки соединения и аутентификации.
	cfg.ConnConfig.ConnectTimeout = 1 * time.Second

	// Лимиты в net.Dialer позволяют достичь предсказуемого
	// поведения в случае обрыва сети.
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: cfg.HealthCheckPeriod,
		// Timeout на установку соединения гарантирует,
		// что не будет зависаний при попытке установить соединение.
		Timeout: cfg.ConnConfig.ConnectTimeout,
	}).DialContext
	return cfg, nil
}

func ComposeConnectionString() string {
	userspec := fmt.Sprintf("%s:%s", PG_USER, PG_PASSWORD)
	hostspec := fmt.Sprintf("%s:%s", PG_HOST, PG_PORT)
	return fmt.Sprintf("postgresql://%s@%s/%s", userspec, hostspec, PG_DB_NAME)
}
