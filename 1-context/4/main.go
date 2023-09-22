package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//ctx1 := context.WithValue(context.Background(), "key_1", "value_1")
	//ctx2 := context.WithValue(ctx1, "key_2", "value_2")
	//fmt.Println(ctx2.Value("key_1"))
	//
	context := context.WithValue(context.Background(), "TraceID", "6a1c8fc6-2d12-43bf-9f83-7f23d0f23295")
	//client := http.Client{}

	go func() {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			log.Println("Server: trace id is", request.Header.Get("x-trace-id"))
		})
		http.ListenAndServe(":8081", nil)
	}()
	time.Sleep(time.Second)
	doRequest(context)

}

func doRequest(ctx context.Context) {
	var client httpClient
	client = TracebleCient{&http.Client{}} //&http.Client{}

	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8081", nil)

	fmt.Println(client.Do(request))
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type TracebleCient struct {
	client *http.Client
}

func (c TracebleCient) Do(req *http.Request) (*http.Response, error) {
	traceID := req.Context().Value("TraceID").(string)
	log.Println("Client: TraceID is", traceID)

	req.Header.Add("x-trace-id", traceID)
	return c.client.Do(req)
}
