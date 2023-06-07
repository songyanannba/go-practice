package main

import (
	"06/dbtest"
	"fmt"
	"time"
)

func init() {
	//mysql
	//dbtest.InitDB()
	//cassandra
	//dbtest.InitCassandra()
}

func Close() {
	dbtest.Session.Close()
	fmt.Println("关闭资源")
}

// 读30000 条数据 mysql  比cassandra 几乎快一倍
func readBiJiao() {
	mysqlUnixStart := time.Now().Unix()
	for a := 0; a < 3; a++ {
		for i := 40005; i < 70004; i++ {
			dbtest.Read(i)
		}
		fmt.Println("mysql 第几次 =  ", a)
	}
	mysqlUnixEnd := time.Now().Unix()

	fmt.Println("mysqlUnixStart == ", mysqlUnixStart)
	fmt.Println("mysqlUnixEnd == ", mysqlUnixEnd)
	fmt.Println("mysqlUnixStart-mysqlUnixEnd == ", mysqlUnixEnd-mysqlUnixStart)

	CassUnixStart := time.Now().Unix()
	for b := 0; b < 3; b++ {
		for j := 0; j < 30000; j++ {
			dbtest.ReadCassandra(j)
		}
		fmt.Println("Cass 第几次 =  ", b)
	}

	//time.Sleep(time.Second * 1)
	cassUnixEnd := time.Now().Unix()
	fmt.Println("CassUnixStart == ", CassUnixStart)
	fmt.Println("cassUnixEnd == ", cassUnixEnd)
	fmt.Println("cassUnixEnd -CassUnixStart == ", cassUnixEnd-CassUnixStart)
}

// 写30000 条数据 cassandra 比mysql 几乎快一倍
func wBiJiao() {
	mysqlUnixStart := time.Now().Unix()
	for i := 0; i < 30000; i++ {
		dbtest.WriteMsqUser()
		if i == 10000 {
			fmt.Println("WriteMsqUser 一半")
		}
	}
	mysqlUnixEnd := time.Now().Unix()
	fmt.Println("mysqlUnixStart == ", mysqlUnixStart)
	fmt.Println("mysqlUnixEnd == ", mysqlUnixEnd)
	fmt.Println("mysqlUnixStart-mysqlUnixEnd == ", mysqlUnixEnd-mysqlUnixStart)

	CassUnixStart := time.Now().Unix()
	for j := 0; j < 30000; j++ {
		dbtest.WCassandra(j)
		if j == 10000 {
			fmt.Println("WriteMsqUser 一半")
		}
	}
	cassUnixEnd := time.Now().Unix()
	fmt.Println("cassUnixEnd -CassUnixStart == ", cassUnixEnd-CassUnixStart)
}

func main() {
	//mysql 和 cassandra 读写比较
	//defer Close()
	/*dbtest.Read(40005)
	dbtest.ReadCassandra(1)*/

	//readBiJiao()
	//wBiJiao()
	//
	//fmt.Println("main")


	//================================================================
	//2 范型 实验
	/*tag := TestTag()
	fmt.Println(tag)*/

	var m1 map[int]string
	var m2 map[string]int

	m1 = make(map[int]string)
	m2 = make(map[string]int)

	m1[1] = "a"
	m1[2] = "b"
	m1[3] = "c"
	m2["s"] = 1
	m2["sq"] = 2
	m2["sc"] = 3

	a := test1[int, string](m1)
	b := test1[string, int](m2)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}
