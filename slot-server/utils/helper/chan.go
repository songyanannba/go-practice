package helper

import (
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

// WaitTimeout 等待超时
func WaitTimeout(c <-chan struct{}, timeout int) bool {
	select {
	case <-c:
		return false
	case <-time.After(time.Duration(timeout) * time.Second):
		return true
	}
}

// Parallel 并行执行
func Parallel[T any](num, size int, fn func() (T, error)) (chan T, chan error, *ants.Pool) {
	if size <= 0 {
		size = 1
	}
	ch := make(chan T, size)
	errCh := make(chan error)
	p, _ := ants.NewPool(size)
	wg := sync.WaitGroup{}
	go func(p *ants.Pool, ch chan T, errCh chan error) {
		defer func() {
			close(ch)
			p.Release()
		}()

		for i := 0; i < num; i++ {
			wg.Add(1)
			err := p.Submit(func() {
				defer wg.Done()
				v, err := fn()
				if err != nil {
					errCh <- err
					return
				}
				ch <- v
			})
			if err != nil {
				errCh <- err
			}
		}
		wg.Wait()
	}(p, ch, errCh)
	return ch, errCh, p
}
