package basetest

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "https://www.baidu.com"
	successStatusCode := 200
	t.Log("测试访问百度首页")
	t.Logf("测试目标：访问%s响应码%d\n", url, successStatusCode)

	resp, err := http.Get(url)
	if err != nil {
		t.Fatal("访问百度首页成功", ballotX, err)
	}
	t.Log("访问百度首页成功", checkMark)

	if resp.StatusCode == successStatusCode {
		t.Log("响应码正确", checkMark)
	} else {
		t.Error("响应码正确", ballotX, resp.StatusCode)

	}
}
