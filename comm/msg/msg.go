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
	//if msg.IsSendBySelf() { // 自己的消息不处理
	//	return
	//}
	//if msg.IsSendByFriend() { // 好友的消息不处理
	//	return
	//}
	if !msg.IsSendByGroup() {
		// 自己的消息不处理
		// 好友的消息不处理
		// 处理群消息
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
		//msg.ReplyText(err.Error())
		return
	}
	fmt.Printf("\n\n\n%#v", sender)
	fmt.Printf("\n\n\n%#v", sender2)

	if msg.IsText() { // 处理文本消息
		// 去除空格
		contentText = trimMsgContent(msg.Content, " ")

		//获取返回的消息
		reply := contextTextBypass(contentText, sender.ID())
		//		reply = fmt.Sprintf(`收到信息
		//%v:%v
		//收到你的信息了~`, sender2.NickName, contentText)

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
	if txt == "打赏" {
		img, err := os.Open("reword.png")
		defer img.Close()
		if err != nil {
			err = errors.Wrapf(err, "reword open file err")
			logrus.Error(err.Error())
			_, err = msg.ReplyText("学雷锋，视钱财如粪土，不用打赏。")
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
	if txt == "123455" {
		return `123`
	}

	// todo 其他的一些
	return ""
}
