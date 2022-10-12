package pgxpool

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
)

func TestPgx(t *testing.T) {
	n := 5
	pool := NewPgxPool("postgres://user:pw@ip:5432/db", n)
	for i := 0; i < n; i++ {
		conn, err := pool.GetConn()
		fmt.Println(conn, err)
		if err != nil {
			fmt.Println("error")
		}
		go func(c *pgx.Conn) {
			time.Sleep(time.Second * 5)
			pool.Release(c)
		}(conn)
	}
	fmt.Println(pool.GetConn())

}

func TestWalk(t *testing.T) {
	n := 5
	pool := NewPgxPool("postgres://user:pw@ip:5432/db", n)
	foo := []WalkFunc{}

	for i := 0; i < 20; i++ {
		foo = append(foo, func(ctx context.Context, conn *pgx.Conn) context.Context {
			num := ctx.Value("num").(int)
			fmt.Println(num)
			if num == 0 {
				ctx, cancel := context.WithCancel(ctx)
				cancel()
				return ctx

			}
			ctx = context.WithValue(ctx, "num", num-1)
			return ctx
		})
	}

	fmt.Println(pool.Walk(context.WithValue(context.TODO(), "num", 10), foo))

}
