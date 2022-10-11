// 比赛模块
package game

import (
	"fmt"
	"sync"
	"time"
)

const (
	MATCH_STAGE_WAIT   = 0
	MATCH_STAGE_SIGNUP = 1
	MATCH_STAGE_32     = 2
	MATCH_STAGE_16     = 3
	MATCH_STAGE_8      = 4
	MATCH_STAGE_4      = 5
	MATCH_STAGE_2      = 6
	MATCH_STAGE_1      = 7
	MATCH_STAGE_END    = 8
)


const (
	TEST_TIME = time.Duration(10) * time.Second  // 1 天 1 场，比赛
)


var manageMatch *ManageMatch


type MatchCounter struct {
	EndTime int64
}

type ManageMatch struct {
	lock               *sync.RWMutex
	Stage              int            // 比赛的阶段
	MatchCounterWait   *MatchCounter  // 等待
	MatchCounterSignUp *MatchCounter  // 报名
	MatchCounter32     *MatchCounter  // 32 进 16
	MatchCounter16     *MatchCounter
	MatchCounter8      *MatchCounter
	MatchCounter4      *MatchCounter
	MatchCounter2      *MatchCounter
	MatchCounter1      *MatchCounter
	MatchCounterEnd    *MatchCounter  // 结束
}

func GetManageMatch() *ManageMatch {
	if manageMatch == nil {
		manageMatch = new(ManageMatch)
		manageMatch.lock = new(sync.RWMutex)
	}

	return manageMatch
}


func (mm *ManageMatch) Run() {
	mm.ChangeStage(MATCH_STAGE_WAIT)
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	
	for {
		select {
		case <- ticker.C:  // 定时器每秒检查状态机
			mm.OnTimer()
		}
	}

	ticker.Stop()
}


func (mm *ManageMatch) OnTimer() {
	nowTime := time.Now().Unix()

	switch mm.Stage {
	case MATCH_STAGE_WAIT:
		// 当前时间 > 等待阶段的结束时间
		if nowTime > mm.MatchCounterWait.EndTime {
			mm.ChangeStage(MATCH_STAGE_SIGNUP)
		}

	case MATCH_STAGE_SIGNUP:
		if nowTime > mm.MatchCounterSignUp.EndTime {
			mm.ChangeStage(MATCH_STAGE_WAIT)
		}
		
	case MATCH_STAGE_32:
	case MATCH_STAGE_16:
	case MATCH_STAGE_8:
	case MATCH_STAGE_4:
	case MATCH_STAGE_2:
	case MATCH_STAGE_1:
	case MATCH_STAGE_END:
	}
}


func (mm *ManageMatch) ChangeStage(stage int) {
	mm.Stage = stage

	switch mm.Stage {
	case MATCH_STAGE_WAIT:
		mm.InitWait()

	case MATCH_STAGE_SIGNUP:
		mm.InitSignUp()

	case MATCH_STAGE_32:
	case MATCH_STAGE_16:
	case MATCH_STAGE_8:
	case MATCH_STAGE_4:
	case MATCH_STAGE_2:
	case MATCH_STAGE_1:
	case MATCH_STAGE_END:
	}

}

func (mm *ManageMatch) InitWait() {
	mm.lock.RLock()
	defer mm.lock.RUnlock()

	mm.MatchCounterWait = new(MatchCounter)
	mm.MatchCounterWait.EndTime = time.Now().Add(TEST_TIME).Unix()  // 等待 10 S

	timerTxt := fmt.Sprintf("%d 秒后，开始报名。\r\n", mm.MatchCounterWait.EndTime - time.Now().Unix())
	msg := []byte(timerTxt)
	GetManagePlayer().BoardCast(msg)
}


func (mm *ManageMatch) InitSignUp() {
	mm.lock.RLock()
	defer mm.lock.RUnlock()

	mm.MatchCounterSignUp = new(MatchCounter)
	mm.MatchCounterSignUp.EndTime = time.Now().Add(TEST_TIME).Unix()  // 等待 10 S

	timerTxt := fmt.Sprintf("%d 秒后，开始 32 进 16。\r\n", mm.MatchCounterSignUp.EndTime - time.Now().Unix())
	msg := []byte(timerTxt)
	GetManagePlayer().BoardCast(msg)
}
/*

	启动 10 s 内，玩家上线
	广播准备打比赛





*/