package service

import (
	"log"
	"recService/data"
)

type Calc struct {

}

//定义add 方法
func (c *Calc) Add(request *data.CalcRequest  ,response *data.CalcResponse) error {
	log.Printf("[+] call add\n")
	response.Result = request.Left + request.Right
	return nil
}
