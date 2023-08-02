package tests

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"slot-server/utils/helper"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	p, _ := ants.NewPool(100)
	ch := make(chan int, 100)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()

	group := sync.WaitGroup{}
	go func(pool *ants.Pool, ch chan int) {
		defer func() {
			if _, beforeClosed := <-ch; beforeClosed {
				close(ch)
			}
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			p.Release()
		}()
		for i := 0; i < 100; i++ {
			group.Add(1)
			err := pool.Submit(func() {
				defer group.Done()
				ch <- GetInt(i)
			})
			if err != nil {
				return
			}
		}
		group.Wait()
	}(p, ch)
	count := 0
	for {
		select {
		case v, beforeClosed := <-ch:
			count++
			if count >= 10 {
				close(ch)
				p.Release()
				continue
			}
			fmt.Printf(" %d %d %v ", count, v, beforeClosed)
			if !beforeClosed {
				return
			}
		}
	}
}

func GetInt(i int) int {
	return i
}

func TestPool2(t *testing.T) {
	ch, cherr, p := helper.Parallel(100, 10, func() (int, error) {
		return 5, nil
	})

	count := 0
	for {
		select {
		case v, beforeClosed := <-cherr:
			fmt.Printf("err %d %d %v ", count, v, beforeClosed)
		case v, beforeClosed := <-ch:
			fmt.Printf("ch %d %d %v ", count, v, beforeClosed)
			if count >= 10 {
				//close(ch)
				p.Release()
				return
			}
			count++

			if !beforeClosed {
				return
			}
		}
	}
}

func Parallel[T any](num, size int, fn func() (T, error)) (chan T, chan error, *ants.Pool) {
	p, _ := ants.NewPool(size)
	tChan := make(chan T, 100)
	errorCh := make(chan error, 100)

	defer func() {
		if err := recover(); err != nil {
			errorCh <- fmt.Errorf("%v", err)
		}
	}()

	group := sync.WaitGroup{}
	go func(pool *ants.Pool, ch chan T, errorCh chan error) {
		defer func() {
			if _, beforeClosed := <-ch; beforeClosed {
				close(ch)
			}
			if err := recover(); err != nil {
				errorCh <- fmt.Errorf("%v", err)
			}
			p.Release()
		}()
		for i := 0; i < num; i++ {
			group.Add(1)
			err := pool.Submit(func() {
				defer group.Done()
				t, err := fn()
				if err != nil {
					errorCh <- err
				}
				ch <- t
			})
			if err != nil {
				return
			}
		}
		group.Wait()
	}(p, tChan, errorCh)
	return tChan, errorCh, p
}
