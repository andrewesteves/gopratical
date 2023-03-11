package heartbeats

import (
	"context"
	"fmt"
	"time"
)

func ProcessLongTask(
	ctx context.Context,
	letters chan rune,
	interval time.Duration,
) (<-chan struct{}, <-chan string) {
	heartbeat := make(chan struct{}, 1)
	words := make(chan string)

	go func() {
		defer close(heartbeat)
		defer close(words)

		beat := time.NewTicker(interval)
		defer beat.Stop()

		for letter := range letters {
			select {
			case <-ctx.Done():
				return
			case <-beat.C:
				select {
				case heartbeat <- struct{}{}:
				default:
				}
			case words <- alphabet[letter]:
				word := alphabet[letter]
				fmt.Println(word)
				// if word == "oceano" {
				// 	words = nil
				// }
			}
		}
	}()

	return heartbeat, words
}
