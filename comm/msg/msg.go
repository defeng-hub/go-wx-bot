package msg

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-wxbot/openwechat/comm/global"
	"os"
	"strings"
)

func HandleMsg(msg *openwechat.Message) {
	if msg.IsSendBySelf() { // 自己的消息不处理
		//global.WxGroups.
		//	SearchByNickName(1, global.Conf.Keys.MasterGroup).
		//	SendText("message - 测试")
		return
	}

	if msg.IsSendByFriend() { // 好友的消息不处理
		return
	}
	if msg.IsSendByGroup() { // 群消息不处理
		return
	}
	var (
		contentText = ""
		err         error
		sender      *openwechat.User
	)
	//获取发送者(群或者人)
	sender, err = msg.Sender()
	sender2, _ := msg.SenderInGroup()
	if err != nil {
		err = errors.Wrapf(err, "%s获取发送人信息失败", global.Conf.Keys.BotName)
		return
	}
	//fmt.Printf("\n\n\n%#v", sender)
	//fmt.Printf("\n\n\n%#v", sender2)
	//sender 是群的基础信息

	// sender2是发送人的信息
	//sender2.NickName 是群内发信息人的名称
	//contentText 是发送的消息

	if msg.IsText() { // 处理文本消息
		var reply string
		// 去除空格
		contentText = trimMsgContent(msg.Content, " ")

		//获取返回的消息
		reply = contextTextBypass(contentText, sender.ID())
		reply = messageTest(contentText, sender2.NickName)

		reply = trimMsgContent(reply, "\n")
		_, err = msg.ReplyText(reply)
		if err != nil {
			err = errors.Wrapf(err, "reply text msg err,contentText: %s", contentText)
			logrus.Error(err.Error())
		}
		return

	}
}

// 回复图片信息
func handleTextReplyBypass(msg *openwechat.Message, txt string) {
	if txt == "图标" {
		img, err := os.Open("./reword.png")
		defer img.Close()
		if err != nil {
			err = errors.Wrapf(err, "reword open file err")
			logrus.Error(err.Error())
			_, err = msg.ReplyText("https://github.com/defeng-hub")
			return
		}
		_, err = msg.ReplyImage(img)
		return
	}
}

func trimMsgContent(content string, cutset string) string {
	content = strings.TrimLeft(content, cutset)
	content = strings.TrimRight(content, cutset)
	return content
}

func contextTextBypass(txt, userID string) (retMsg string) {
	if txt == "ping" {
		return `pong`
	}

	return ""
}

func messageTest(txt, name string) (reply string) {
	// name 为群组人员名
	if txt == "测试" {
		return fmt.Sprintf(`收到信息
%v:%v
收到你的信息了~`, name, txt)
	}
	return ""
}
