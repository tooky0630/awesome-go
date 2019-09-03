package limit

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

func TestLimiter(t *testing.T) {
    limiter := NewLimiter(3, time.Millisecond)
    defer limiter.Close()

    work := func() {
        time.Sleep(time.Second)
    }
    wg := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            fmt.Println("done...", i, limiter.do(work))
            wg.Done()
        }(i)
    }
    wg.Wait()
}
