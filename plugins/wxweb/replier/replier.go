/*
Copyright 2017 wechat-go Authors. All Rights Reserved.
MIT License

Copyright (c) 2017

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package replier

import (
	"math/rand"

	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
)

// register plugin
func Register(session *wxweb.Session) {
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(autoReply), "text-replier")
	if err := session.HandlerRegister.Add(wxweb.MSG_IMG, wxweb.Handler(autoReply), "img-replier"); err != nil {
		logs.Error(err)
	}

	if err := session.HandlerRegister.EnableByName("text-replier"); err != nil {
		logs.Error(err)
	}

	if err := session.HandlerRegister.EnableByName("img-replier"); err != nil {
		logs.Error(err)
	}

}
func autoReply(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	st := map[int]string{
		1: "小小虫子，max爱你哟",
		2: "很爱很爱你，所以愿意为你",
		3: "笔芯♥️",
		4: "故事很长,我长话短说,我喜欢你,很久了",
		5: "说片面是熬夜,说实在是失眠,说实话是想你,晚安",
		6: "见过这么多猪, 还是你最可爱",
		7: "我很想你很想你, 是想穿过屏幕抱抱你的那种 ",
		8: "想抱抱你, 就是那种两只手搂住你的腰, 把头埋进你胸口的那种抱抱",
		9: "我想和你一起走很远的路,趁四处没人时偷亲你一下,等天色渐暗 再亲久一点",
	}
	if !msg.IsGroup {
		logs.Error("frome:%v, to:%v, who:%v, content:%v\n", msg.FromUserName, msg.ToUserName, msg.Who, msg.Content)
		session.SendText(st[rand.Intn(10)], session.Bot.UserName, msg.FromUserName)
	}
}
