package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "traceId", 123456)
	ctx = context.WithValue(ctx, "sessionId", 987654)

	process(ctx)
}

func process(ctx context.Context) {
	fmt.Println("traceId :", ctx.Value("traceId"))
	fmt.Println("sessionId :", ctx.Value("sessionId"))
	fmt.Println("transId :", ctx.Value("transId"))
}
