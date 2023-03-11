package heartbeats

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestProcessLongTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	letters := make(chan rune)
	go func() {
		defer close(letters)
		for i := 'a'; i <= 'z'; i++ {
			letters <- i
		}
	}()

	heartbeat, words := ProcessLongTask(ctx, letters, time.Second)
	for {
		select {
		case <-ctx.Done():
			t.FailNow()
		case <-heartbeat:
			fmt.Println("beating...")
		case word, ok := <-words:
			if !ok {
				return
			}
			if _, has := alphabet[rune(word[0])]; !has {
				t.Errorf("unknown word: %s", word)
			}
		}
	}
}
