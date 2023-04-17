package ticker

func Ticker() {
	// 每天提醒
	go MasterTicker()

}
