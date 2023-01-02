package main

import (
	"flag"
	"fmt"
)

func main() {

	txt := flag.String("s", "", "md5 txt")

	flag.Usage = func() {
		fmt.Println("usage : [-s abc]")
		flag.PrintDefaults()
	}

	if *txt == "" {
		flag.Usage()
	}

}
