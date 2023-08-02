package tests

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"testing"
	"time"
)

type B struct {
	C string `json:"c"`
	D int    `json:"d"`
	E []int  `json:"e"`
}

type A struct {
	A string `json:"a"`
	B B
	F map[string]any `json:"f"`
}

// go test -v -run JsonTemp json_test.go
func JsonTemp(t *testing.T) {
	data := A{
		A: "abbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbdd",
		B: B{
			C: "abbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbddabbdd",
			D: 1312311,
			E: []int{1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				1, 2, 3, 4, 5, 6, 7, 8, 685, 85, 858, 5, 57, 46, 63, 535, 35, 3, 52,
				131, 4141, 41412, 3123131, 4141, 4343,
			},
		},
		F: map[string]any{
			"1adadada": 1231231,
			"2adadada": 123123113131,
			"1a4adada": "1231231",
			"dadad":    []any{1, "dad", "dada", "fafa", 4141, []int{1, 2, 3, 4, 5, 6}},
		},
	}
	i := testTime(func() error {
		_, err := sonic.Marshal(data)
		return err
	})
	t.Log("sonic ", i)
	i = testTime(func() error {
		_, err := jsoniter.ConfigFastest.Marshal(data)
		return err
	})
	t.Log("jsoniter ", i)
	i = testTime(func() error {
		_, err := gojson.Marshal(data)
		return err
	})
	t.Log("gojson ", i)
	i = testTime(func() error {
		_, err := json.Marshal(data)
		return err
	})
	t.Log("json ", i)
	v, _ := json.Marshal(data)

	var a = &A{}

	t.Log("---------------")
	i = testTime(func() error {
		return sonic.Unmarshal(v, a)
	})
	t.Log("sonic ", i)
	i = testTime(func() error {
		return jsoniter.ConfigFastest.Unmarshal(v, a)
	})
	t.Log("jsoniter ", i)
	i = testTime(func() error {
		return gojson.Unmarshal(v, a)
	})
	t.Log("gojson ", i)
	i = testTime(func() error {
		return json.Unmarshal(v, a)
	})
	t.Log("json ", i)
}

func testTime(fn func() error) int64 {
	s := time.Now()
	for i := 0; i < 10000; i++ {
		err := fn()
		if err != nil {
			panic(err)
		}
	}
	e := time.Now()
	i := e.Sub(s).Microseconds()
	return i
}
