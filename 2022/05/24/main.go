package main

import (
	"context"
	"fmt"
	"time"
)


func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}



const shortDuration  = 4 * time.Second

func main11() {

	d := time.Now().Add(shortDuration)

	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()

	for  {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("overslept")
		case <-time.After(3 * time.Second):
			fmt.Println("233")
		case <-ctx.Done():
			fmt.Println("1234")
			fmt.Println(ctx.Err())
		}
		fmt.Println("hahah")
		time.Sleep(1 * time.Second)
	}


}


func main1 () {

	gen := func (ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func () {
			for  {
				fmt.Println("for...")
				select {
				case <-ctx.Done():
					fmt.Println("done...")
					return
				case dst <- n:
					fmt.Println("n++...")
					n++
				}
			}
		} ()
		return dst
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	for nn :=range gen(ctx) {
		fmt.Println("nnnn ...")
		fmt.Println(nn)
		if nn == 10 {
			break
		}
	}

}
