package tailf

import (
	"cargo/log"
	"fmt"
	"github.com/astaxie/beego/logs"
	tail2 "github.com/hpcloud/tail"
	"oldboyedu/day12/logagent/conf"
	"time"
)

type TailMsg struct {
	Text  string
	Topic string
}

type TailObj struct {
	tail *tail2.Tail
}

type TailObjMgr struct {
	tailObjs []TailObj
	MsgChan  chan *TailMsg
}

var (
	TailObjManager *TailObjMgr
)

func InitTailf() (err error) {
	if len(conf.AppConfig.CollectConfs) == 0 {
		err = fmt.Errorf("no collect conf")
		return
	}

	TailObjManager = &TailObjMgr{
		MsgChan: make(chan *TailMsg, 100),
	}

	for _, collectConf := range conf.AppConfig.CollectConfs {
		filePath := collectConf.LogPath
		tail, tailErr := tail2.TailFile(filePath, tail2.Config{
			ReOpen:    true,
			MustExist: false,
			Follow:    true,
			Poll:      true,
			//Location: &tail2.SeekInfo{
			//	Offset: 0,
			//	Whence: 2,
			//},
		})

		if tailErr != nil {
			log.Errorf("init log file[%s] tail failed, err : %v", filePath, err)
			continue
		}

		tailObj := TailObj{
			tail: tail,
		}
		TailObjManager.tailObjs = append(TailObjManager.tailObjs, tailObj)

		go func() {
			for {
				msg, ok := <-tail.Lines
				if !ok {
					log.Warnf("tail log file[%s] close reopen", filePath)
					time.Sleep(time.Second)
					continue
				}
				TailObjManager.MsgChan <- &TailMsg{
					Text:  msg.Text,
					Topic: collectConf.LogTopic,
				}
			}
		}()

		logs.Info("init log file[%s] tail success", filePath)
	}

	return
}
