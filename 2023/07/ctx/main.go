package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//var (
	//	minMul  int
	//	maxMul  int
	//	initNum int
	//)
	//minMul = 0
	//fmt.Println(minMul, maxMul, initNum)

	//context.Background()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		fmt.Println("go ")
		time.Sleep(time.Second * 3)
		fmt.Println("go end")
		cancelFunc()
	}()

	<-ctx.Done()
	fmt.Println("end ... ")

}
