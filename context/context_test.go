package context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}

func TestTimeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	wg.Wait()
}

func TestDeadline(t *testing.T) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, i int) error {
	defer wg.Done()
	start := time.Now().Unix()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("quit %d, ts:  %d \n", i, time.Now().Unix()-start)
			return ctx.Err()
		default:
			fmt.Println("hello ", i)
		}
	}
}
