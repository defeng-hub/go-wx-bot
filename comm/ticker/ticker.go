package ticker

import "go-wxbot/openwechat/comm/apps/weibo"

func Ticker() {
	// 每天提醒
	go MasterTicker()

	go weibo.WeiboRun()
}
