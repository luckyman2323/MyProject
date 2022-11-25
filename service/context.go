package service

import (
	"context"
	"fmt"
	"time"
)

func ContextTest() {
	mainContext := context.Background()
	ctx, cancelFunc := context.WithCancel(mainContext)

	fmt.Println("====", ctx.Value("test"))
	ctx1 :=  setValue(ctx)

	time.Sleep(1*time.Second)

	fmt.Println("====", ctx1.Value("test"))
	defer cancelFunc()
}

func setValue(ctx context.Context)context.Context{
	return context.WithValue(ctx, "test", "ttttt")
}