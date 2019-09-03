package pubsub

import (
    "sync"
    "time"
)

type (
    subscriber chan interface{}        // 订阅者为一个管道
    topicFunc func(v interface{}) bool // 主题过滤器
)
type Publisher struct {
    mutex sync.RWMutex  // 互斥量
    buffer int     // 缓冲区大小
    timeout time.Duration  // 超时时间
    subscribers map[subscriber]topicFunc // 订阅者信息
}

func NewPublisher(timeout time.Duration, buffer int) *Publisher {
    return &Publisher{
        buffer:      buffer,
        timeout:     timeout,
        subscribers: make(map[subscriber]topicFunc),
    }
}

func (p *Publisher) subscribe() subscriber {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    ch := make(chan interface{}, p.buffer)
    p.subscribers[ch] = nil
    return ch
}

func (p *Publisher) subscribeTopic(topicFunc topicFunc) subscriber {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    ch := make(chan interface{}, p.buffer)
    p.subscribers[ch] = topicFunc
    return ch
}

func (p *Publisher) evict(sub subscriber) {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    delete(p.subscribers, sub)
    close(sub)
}

func (p *Publisher) publish(v interface{}) {
    p.mutex.RLock()
    defer p.mutex.RUnlock()

    wg := sync.WaitGroup{}
    for sub, topic := range p.subscribers {
        wg.Add(1)
        go p.sendTopic(sub, topic, v, &wg)
    }

    wg.Wait()
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
    defer wg.Done()

    if topic != nil && !topic(v) {
        return
    }

    select {
    case sub <- v:
    case <-time.After(p.timeout):
    }
}

func (p *Publisher) Close() {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    for sub := range p.subscribers {
        delete(p.subscribers, sub)
        close(sub)
    }
}