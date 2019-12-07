package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	process int
}

func (wavPlayer *WAVPlayer) Play(source string) {
	wavPlayer.process = 0
	fmt.Println("playing wav music :", source)
	for wavPlayer.process < 100 {
		time.Sleep(time.Millisecond * 100)
		fmt.Print("-")
		wavPlayer.process += 10
	}
	fmt.Println("\nfinished play :", source)
}
