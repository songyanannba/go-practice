package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)




func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err:= scanner.Err() ; err != nil {
		fmt.Fprintln(os.Stderr , "readig standard input:" , err)
	}
}

func main1() {
	const input = "1,2,3,4,"

	scanner := bufio.NewScanner(strings.NewReader(input))


	onComma := func(data []byte , atEOF bool) (advance int ,token []byte , err error) {

		fmt.Println(data)
		fmt.Println(string(data))

		for i:= 0 ; i < len(data) ; i++ {
			if data[i] == ',' {
				fmt.Println(i ,data[:i] )
				return i +1 ,data[:i] , nil
			}
		}

		if !atEOF {
			return 0 , nil ,nil
		}
		return 0 ,data , bufio.ErrFinalToken
	}

	scanner.Split(onComma)

	for scanner.Scan() {
		fmt.Printf("%q " , scanner.Text())
	}
	if err := scanner.Err ; err != nil {
		fmt.Fprintln(os.Stderr , "reading input:" ,err)
	}


}
