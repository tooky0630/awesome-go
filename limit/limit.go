package limit

import (
    "errors"
    "time"
)

type Limiter struct {
    limit int
    timeout time.Duration
    ch chan interface{}
}

func NewLimiter(limit int, timeout time.Duration) *Limiter {
    return &Limiter{
        limit:     limit,
        timeout:   timeout,
        ch:        make(chan interface{}, limit),
    }
}

func (p *Limiter) do(work func()) error {
    select {
        case p.ch <- 1:
        case <-time.After(p.timeout):
            return errors.New("wait timeout...")
    }
    work()
    <- p.ch
    return nil
}

func (p *Limiter) Close() {
    close(p.ch)
}