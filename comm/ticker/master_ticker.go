package ticker

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-wxbot/openwechat/comm/funcs"
	"go-wxbot/openwechat/comm/global"
)

// MasterTicker 每天提醒自己一些事
func MasterTicker() {
	for {
		select {
		case t := <-time.After(1 * time.Minute):
			nowTime := t.Format("15:04")
			if nowTime == "10:00" {
				message := fmt.Sprintf("盛年不重来，一日难再晨。及时当勉励，岁月不待人。\n今年还剩 %d 天。", funcs.RemainingDays())
				err := global.WxFriends.
					SearchByRemarkName(1, global.Conf.Keys.MasterAccount).
					SendText(message)
				if err != nil {
					err = errors.Wrapf(err, "SendMessageToMasterAccout err")
					logrus.Error(err.Error())
				}
			}
		}
	}
}
