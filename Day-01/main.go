package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Processing complited.")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)

	defer cancel()

	if err := process(ctx); err != nil {
		fmt.Println("error:", err)
	}
}
