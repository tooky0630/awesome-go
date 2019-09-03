package pubsub

import (
    "fmt"
    "strings"
    "testing"
    "time"
)

func TestPubSub(t *testing.T) {
    p := NewPublisher(time.Millisecond * 100, 10)
    defer p.Close()

    all := p.subscribe()
    golang := func(v interface{}) bool {
        if str, ok := v.(string); ok && strings.Contains(str, "golang") {
            return true
        }
        return false
    }

    goSub := p.subscribeTopic(golang)

    p.publish("hello, world")
    p.publish("hello, golang")

    go func(sub subscriber) {
        for msg := range sub {
            fmt.Println("all:", msg)
        }
    }(all)

    go func(sub subscriber) {
        for msg := range sub {
            fmt.Println("golang:", msg)
        }
    }(goSub)

    time.Sleep(time.Second)
}