package helper

import (
	"fmt"
	"testing"
)

func TestAES_Decrypt(t *testing.T) {
	a := GenAesSecret()
	fmt.Println(a)
	return
	key := []byte("nXI8pG9asyIK1WhA")
	s := []byte("kqsRmEJ3u5mK1mfN")
	data, err := AesEncode([]byte("123dadada"), key, s)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(data))
	res, err := AesDecode(data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
