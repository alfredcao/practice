package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source string, musicType string) {
	var player Player
	switch musicType {
	case "mp3":
		player = &MP3Player{}
	case "wav":
		player = &WAVPlayer{}
	default:
		fmt.Println("不支持的音乐格式！")
		return
	}

	player.Play(source)
}
