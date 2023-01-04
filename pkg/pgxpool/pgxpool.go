package pgxpool

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/juju/errors"
)

type PgxPool struct {
	connChan chan *pgx.Conn
	connStr  string
	len      int
	cap      int
	timeout  time.Duration
}

type WalkFunc func(ctx context.Context, conn *pgx.Conn) context.Context

func NewPgxPool(connStr string, cap int, timeout time.Duration) *PgxPool {
	if timeout == 0 {
		timeout = 5 * time.Second
	}
	p := &PgxPool{
		connChan: make(chan *pgx.Conn, cap),
		connStr:  connStr,
		len:      0,
		cap:      cap,
		timeout:  timeout,
	}
	return p
}

func (p *PgxPool) GetConn() (*pgx.Conn, error) {
	select {
	case c := <-p.connChan:
		return c, nil
	default:
		if p.len < p.cap {
			conn, err := pgx.Connect(context.Background(), p.connStr)
			if err != nil {
				return nil, err
			}
			p.len++
			return conn, nil
		}

		select {
		case c := <-p.connChan:
			return c, nil
		case <-time.After(p.timeout):
			return nil, errors.Timeout
		}

	}
}

func (p *PgxPool) Release(conn *pgx.Conn) {
	if conn != nil {
		p.connChan <- conn
	}
}

func (p *PgxPool) Walk(ctx context.Context, foo []WalkFunc) error {
	conn, err := p.GetConn()
	defer p.Release(conn)
	if err != nil {
		return err
	}
	// context.
	for _, f := range foo {
		ctx = f(ctx, conn)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// continue to do
		}
	}
	return nil
}
