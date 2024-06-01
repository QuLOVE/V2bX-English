package core

import (
	"github.com/QuLOVE/V2bX-English/api/panel"
	"github.com/QuLOVE/V2bX-English/conf"
)

type AddUsersParams struct {
	Tag   string
	Users []panel.UserInfo
	*panel.NodeInfo
}

type Core interface {
	Start() error
	Close() error
	AddNode(tag string, info *panel.NodeInfo, config *conf.Options) error
	DelNode(tag string) error
	AddUsers(p *AddUsersParams) (added int, err error)
	GetUserTraffic(tag, uuid string, reset bool) (up int64, down int64)
	DelUsers(users []panel.UserInfo, tag string) error
	Protocols() []string
	Type() string
}
