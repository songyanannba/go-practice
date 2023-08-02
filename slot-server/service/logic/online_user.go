package logic

import (
	"github.com/lonng/nano/session"
	"slot-server/enum"
	"sync"
	"time"
)

// UserSession userId -> session
type UserSession struct {
	sesMap map[uint]*session.Session
	sync.RWMutex
}

var registry *UserSession

func init() {
	registry = NewUserSession()
}

func NewUserSession() *UserSession {
	object := &UserSession{
		sesMap: make(map[uint]*session.Session),
	}
	return object
}

func SetUserSession(userId uint, ses *session.Session) {
	registry.Lock()
	registry.sesMap[userId] = ses
	registry.Unlock()
	ses.Set(enum.SessionDataUserId, userId)
}

func GetUserSession(userId uint) *session.Session {
	registry.RLock()
	defer registry.RUnlock()
	return registry.sesMap[userId]
}

func GetAllSession() map[uint]*session.Session {
	registry.RLock()
	defer registry.RUnlock()

	mp := make(map[uint]*session.Session)
	for uid, ses := range registry.sesMap {
		mp[uid] = ses
	}

	return mp
}

func DeleteUserSession(ses *session.Session, userId uint) {
	registry.Lock()
	defer registry.Unlock()

	if userId != 0 {
		delete(registry.sesMap, userId)
	}
	for uid, item := range registry.sesMap {
		if item == ses {
			delete(registry.sesMap, uid)
			break
		}
	}
}

// SyncOnlineUser 同步在线用户到DB
func SyncOnlineUser() {
	for {
		select {
		case <-time.After(10 * time.Second):

		}
	}
}
