package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/url"
	"slot-server/utils/helper"
	"testing"
	"time"
)

// go test -v -run TestParallel util_test.go
func TestParallel(t *testing.T) {
	var (
		arr []int
		fn  = func() (int, error) {
			time.Sleep(1 * time.Second)
			if 1 == helper.RandInt(10) {
				return 0, errors.New("error")
			}
			return 1, nil
		}
		errNum int
	)
	ch, errCh, _ := helper.Parallel(100, 100, fn)

	for {
		select {
		case <-errCh:
			errNum++
		case v, beforeClosed := <-ch:
			if !beforeClosed {
				goto end
			}
			arr = append(arr, v)
		}
	}
end:
	t.Log("error num ", errNum)
	t.Log("success num ", len(arr))
}

func TestUrl(t *testing.T) {
	_, err := url.ParseRequestURI("123.123.123.123")
	assert.Equal(t, nil, err)
	_, err = url.ParseRequestURI("https://xaxa-da")
	assert.Equal(t, nil, err)
	_, err = url.ParseRequestURI("http://xxx")
	assert.Equal(t, nil, err)
	_, err = url.ParseRequestURI("https://localhost")
	assert.Equal(t, nil, err)
	_, err = url.ParseRequestURI("http://123.123.123.123")
	assert.Equal(t, nil, err)
	_, err = url.ParseRequestURI("https://123.123.123.123")
	assert.Equal(t, nil, err)
	_, err = url.ParseRequestURI("daxa-vasva")
	assert.Equal(t, nil, err)
	_, err = url.ParseRequestURI("")
	assert.Equal(t, nil, err)
}

func generateArray(n int, num int) []int {

	arr := make([]int, n)
	// 首先为数组的每个元素分配一个值1
	for i := 0; i < n; i++ {
		arr[i] = 1
		num--
	}

	// 随机选择一个元素并增加它的值，直到我们达到所需的总和
	for num > 0 {
		i := helper.RandInt(n) // 随机选择一个元素的索引
		if arr[i] < 8 {        // 确保元素的值不大于8
			arr[i]++
			num--
		}
	}
	return arr
}
