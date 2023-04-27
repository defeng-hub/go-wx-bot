package global

import (
	"github.com/eatmoreapple/openwechat"
)

var (
	Conf      *Config
	WxSelf    *openwechat.Self
	WxFriends openwechat.Friends // 可能有缓存
	WxGroups  openwechat.Groups
)
