package library

import "errors"

// 音乐
type Music struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

// 音乐管理器
type MusicManager struct {
	musics []Music
}

// 新建一个音乐管理器
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

// 通过索引获取音乐
func (musicManager *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 || index > (musicManager.Len()-1) {
		err = errors.New("索引超出范围！")
		return
	}

	music = &musicManager.musics[index]
	return
}

// 往音乐管理器添加音乐
func (musicManager *MusicManager) Add(music *Music) {
	musicManager.musics = append(musicManager.musics, *music)
}

// 通过音乐名查找音乐
func (musicManager *MusicManager) Find(name string) (music *Music) {
	length := musicManager.Len()
	if length == 0 {
		return
	}

	for i := 0; i < length; i++ {
		v := &musicManager.musics[i]
		if v.Name == name {
			music = v
			return
		}
	}
	return
}

// 音乐管理器里的音乐数量
func (musicManager *MusicManager) Len() int {
	return len(musicManager.musics)
}

func (musicManager *MusicManager) Remove(index int) (music *Music, err error) {
	length := musicManager.Len()
	if index < 0 || index > (length-1) {
		err = errors.New("索引超出范围！")
		return
	}

	removedMusic := musicManager.musics[index]

	if index == 0 {
		// 第一个
		musicManager.musics = musicManager.musics[1:]
	} else if index == (length - 1) {
		// 最后一个
		musicManager.musics = musicManager.musics[:index]
	} else {
		// 中间
		musicManager.musics = append(musicManager.musics[:index], musicManager.musics[index+1:]...)
	}

	music = &removedMusic
	return
}

func (musicManager *MusicManager) RemoveByName(name string) (music *Music, err error) {
	index := -1
	for i, v := range musicManager.musics {
		if v.Name == name {
			index = i
			break
		}
	}
	if index != -1 {
		return musicManager.Remove(index)
	}
	return
}
