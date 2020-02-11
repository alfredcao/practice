package mocktest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"
const mockResponseContent = `
<html>
	<head>
		<title>百度</title>
	</head>
	<body>
		百度一下，你就知道
	</body>
</html>
`

func TestMockService(t *testing.T) {
	successStatusCode := http.StatusOK
	server := server()
	t.Log("测试访问百度首页")
	t.Logf("测试目标：访问%s响应码%d\n", server.URL, successStatusCode)

	resp, err := http.Get(server.URL)
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

func server() *httptest.Server {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprint(w, mockResponseContent)
	}
	return httptest.NewServer(http.HandlerFunc(handleFunc))
}
