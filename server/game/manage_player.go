package game

import (
	"sync"

	"golang.org/x/net/websocket"
)


var managePlayer *ManagePlayer


// 玩家管理模块
type ManagePlayer struct {
	Players map[int64]*Player  // 所有玩家
	Id     int64
	lock   *sync.RWMutex
}


func GetManagePlayer() *ManagePlayer {
	if managePlayer == nil {
		managePlayer = new(ManagePlayer)
		managePlayer.Players = make(map[int64]*Player)
		managePlayer.lock = new(sync.RWMutex)
	}

	return managePlayer
}


func (mp *ManagePlayer) PlayerLogin(ws *websocket.Conn) *Player {
	mp.lock.Lock()
	defer mp.lock.Unlock()

	mp.Id++
	playerInfo := NewTestPlayer(ws, mp.Id)

	// 加入管理器
	mp.Players[playerInfo.UserId] = playerInfo

	return playerInfo
}


func (mp *ManagePlayer) BoardCast(msg []byte) {
	mp.lock.RLock()
	defer mp.lock.RUnlock()
	
	for _, p := range mp.Players {
		p.SendNotice(msg)
	}
}