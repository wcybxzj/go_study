package lib

import "go_study/13.oldboy/day9/chat/chat_server/model"

var (
	mgr *model.UserMgr
)

func InitUserMgr() {
	mgr = model.NewUserMgr(pool)
}
