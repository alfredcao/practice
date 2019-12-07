package library

import "testing"

func TestOps(t *testing.T) {
	musicManager := NewMusicManager()
	if musicManager == nil {
		t.Error("新建音乐管理器失败！")
		return
	}

	if musicManager.Len() != 0 {
		t.Error("新建的音乐管理器不为空！")
		return
	}

	music1 := &Music{"1", "My Heart Will Go On", "Celion Dion",
		"http://qbox.me/24501234", "MP4"}
	musicManager.Add(music1)

	if musicManager.Len() != 1 {
		t.Error("添加音乐失败！")
		return
	}

	music2 := musicManager.Find("My Heart Will Go On")
	if music2 == nil || music1.Id != music2.Id {
		t.Error("查找音乐[My Heart Will Go On]失败！")
	}

	music3, err := musicManager.Get(0)
	if err != nil || music3 == nil || music1.Id != music3.Id {
		t.Error("查找音乐[0]失败！")
	}
	musicManager.Add(&Music{"2", "晴天", "Jay Chou",
		"http://qbox.me/24501235", "MP3"})
	if musicManager.Len() != 2 {
		t.Error("添加音乐失败！")
		return
	}
	musicManager.Add(&Music{"3", "雨一直下", "张宇",
		"http://qbox.me/24501236", "MP3"})

	if musicManager.Len() != 3 {
		t.Error("添加音乐失败！")
		return
	}

	music4, err := musicManager.Remove(0)
	if err != nil {
		t.Error("删除音乐失败！")
	}
	if music4.Id != "1" {
		t.Error("删除音乐错误！")
	}
}
