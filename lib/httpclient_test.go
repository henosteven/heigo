package lib

import (
	"testing"
	"context"
	"fmt"
)

func TestHttpGet(t *testing.T) {
	ctx := context.Background()
	url := "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9391898625033851859%22%7D&n_type=0&p_from=1"
	params := map[string]interface{} {
	}

	headers := map[string]interface{} {
	}

	data, err := HttpGet(ctx, url, params, headers)
	if err != nil {
		t.Errorf("error:%s", err)
	}

	fmt.Print(data)
}