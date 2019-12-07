package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	process int
}

func (mp3Player *MP3Player) Play(source string) {
	mp3Player.process = 0
	fmt.Println("playing mp3 music :", source)
	for mp3Player.process < 100 {
		time.Sleep(time.Millisecond * 100)
		fmt.Print(".")
		mp3Player.process += 10
	}
	fmt.Println("\nfinished play :", source)
}
