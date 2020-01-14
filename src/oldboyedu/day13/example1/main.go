package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpResponse struct {
	r   *http.Response
	err error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{
		Transport: tr,
	}
	//req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("new request failed, err :", err)
		return
	}
	var httpResponseChan = make(chan *HttpResponse)

	go func() {
		response, err := client.Do(req)
		httpResponseChan <- &HttpResponse{
			r:   response,
			err: err,
		}
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-httpResponseChan
		fmt.Println("time out !!!")
	case res := <-httpResponseChan:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("%s", out)
	}

}
