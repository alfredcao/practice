package test

import (
	"gowebprograming/chitchat/dao"
	_ "gowebprograming/chitchat/dao"
	"testing"
)

func TestGetThreadList(t *testing.T) {
	threads, err := dao.GetThreadList()
	if err != nil {
		t.Errorf("获取帖子信息失败！错误信息：%v", err)
		return
	}

	t.Logf("获取帖子信息成功，共 %d 个帖子", len(threads))
	for index, thread := range threads {
		t.Logf("第 %d 个帖子内容 -> id : %d, uuid : %s, 标题 : %s, 内容 : %s, 用户id : %d, 创建时间 : %v, 更新时间 : %v",
			index, thread.Id, thread.Uuid, thread.Title, thread.Content, thread.UserId, thread.CreatedAt, thread.UpdatedAt)
	}
}
