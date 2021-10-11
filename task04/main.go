package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgconn/stmtcache"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	PG_HOST     = "127.0.0.1"
	PG_PORT     = "5432"
	PG_USER     = "sanya"
	PG_PASSWORD = "passwd"
	PG_DB_NAME  = "sauna"
)

type Config struct {
	Mode            int
	MaxConns        int32
	MinConns        int32
	AttackMS        time.Duration
	GoroutinesCount int
}

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config, err := getConfig()
	if err != nil {
		return err
	}
	switch config.Mode {
	// case 1:
	// 	return runSQLInterfaceExample()
	// case 2:
	// 	return runPGXExample()
	// case 3:
	// 	return runPoolExample()
	// case 4:
	// 	return generateDBLoad(config)
	// case 5:
	// 	return runUpdateWithTX()
	case 6:
		return generateDBLoadWithPgBouncer(config)
	default:
		return fmt.Errorf("got an unexpected running mode %d", config.Mode)
	}
}

func getConfig() (*Config, error) {
	var mode, maxConns, minConns, goroutinesCount int
	var attackMS int64
	flag.IntVar(&mode, "mode", -1,
		`running mode of the application:
			- 1: with database/sql
			- 2: with jackc/pgx
			- 3: with pgxpool
			- 4: generate DB load
			- 5: run update with TX
			- 6: generate DB load with PgBouncer`)
	flag.IntVar(&maxConns, "maxConns", 1, "max pool connections")
	flag.IntVar(&minConns, "minConns", 1, "min pool connections")
	flag.Int64Var(&attackMS, "attackMS", 100, "attack duration in MS")
	flag.IntVar(&goroutinesCount, "goroutines", 10, "goroutines count")
	flag.Parse()
	if mode == -1 {
		return nil, fmt.Errorf("running mode must be defined")
	}
	switch mode {
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	default:
		return nil, fmt.Errorf("unknown mode %d", mode)
	}
	return &Config{
		Mode:            mode,
		MaxConns:        int32(maxConns),
		MinConns:        int32(minConns),
		AttackMS:        time.Duration(attackMS),
		GoroutinesCount: goroutinesCount,
	}, nil
}

func generateDBLoadWithPgBouncer(cfg *Config) error {
	pool, err := createPGXPoolModeDescribe(cfg.MaxConns, cfg.MinConns)
	if err != nil {
		return fmt.Errorf("failed to create a PGX pool: %w", err)
	}
	res, err := Attack(
		context.Background(),
		time.Millisecond*cfg.AttackMS,
		cfg.GoroutinesCount,
		pool,
	)
	if err != nil {
		return fmt.Errorf("attack failed: %w", err)
	}
	log.Println("Attack result: ", res)
	return nil
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

type PersonalService struct {
	serviceID *int64
	service   *string
	price     *float64
}

func SearchPersonalServiceByName(
	ctx context.Context,
	dbpool *pgxpool.Pool,
	fName string,
	lName string,
) ([]PersonalService, error) {
	const sql = `
	-- информация о навыках конкретного сотрудника
	select
		service_id
		, service
		, price
	from
		personal_info
	where
		fname = $1
		and lname = $2
	order by
		service;`
	rows, err := dbpool.Query(ctx, sql, fName, lName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the personal service: %w", err)
	}
	defer rows.Close()
	hints := make([]PersonalService, 0)
	for rows.Next() {
		h := PersonalService{}
		if err := rows.Scan(&h.serviceID, &h.service, &h.price); err != nil {
			return nil, fmt.Errorf("failed to scan the received data row: %w", err)
		}
		hints = append(hints, h)
	}
	return hints, nil
}

func getNextName(ctx context.Context, dbpool *pgxpool.Pool) (string, string, error) {
	const sql = `
	-- выбор случайного сотрудника
	select fname, lname from personal
	order by random()
	limit 1;`

	var fname, lName string
	err := dbpool.QueryRow(ctx, sql).Scan(&fname, &lName)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch row: %w", err)
	}

	return fname, lName, nil
}

func Attack(
	ctx context.Context,
	duration time.Duration,
	threads int,
	dbpool *pgxpool.Pool,
) (*AttackResults, error) {

	var queries uint64
	attacker := func(stopAt time.Time) {
		for {
			fName, lName, err := getNextName(ctx, dbpool)
			if err != nil {
				log.Printf("an error occurred getting next Persona: %v", err)
				continue
			}
			_, err = SearchPersonalServiceByName(ctx, dbpool, fName, lName)
			if err != nil {
				log.Printf("an error occurred while searching Name: %v", err)
				continue
			}
			atomic.AddUint64(&queries, 1)
			if time.Now().After(stopAt) {
				return
			}
		}
	}

	startAt := time.Now()
	stopAt := startAt.Add(duration)

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}

	wg.Wait()

	return &AttackResults{
		Duration:         time.Now().Sub(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}, nil
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
