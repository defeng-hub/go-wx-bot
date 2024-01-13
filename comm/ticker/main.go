package ticker

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-wxbot/openwechat/comm/global"
	"go-wxbot/openwechat/comm/tools"
)

// MasterTicker 每天提醒自己一些事
func MasterTicker() {
	for {
		select {
		case t := <-time.After(1 * time.Minute):
			nowTime := t.Format("15:04")
			//fmt.Printf(nowTime, global.Conf.Keys.MasterAccount)
			if nowTime == "10:00" {
				message := fmt.Sprintf("盛年不重来，一日难再晨。及时当勉励，岁月不待人。\n今年还剩 %d 天。", tools.RemainingDays())
				err := global.WxFriends.
					SearchByRemarkName(1, global.Conf.Keys.MasterAccount).
					SendText(message)

				err = global.WxGroups.
					SearchByNickName(1, global.Conf.Keys.MasterGroup).
					SendText(message)

				if err != nil {
					err = errors.Wrapf(err, "SendMessageToMasterAccout err")
					logrus.Error(err.Error())
				}
			}
		}
	}
}
