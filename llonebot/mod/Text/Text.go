package Text

import (
	"encoding/base64"
	"fmt"
)

// Cq 表情
/*
data 表情id
*/
func Cq(data string) string {
	r := fmt.Sprintf("[CQ:face,id=%s]" + data)
	return r
}

// Img 图片 data图片文件
func Img(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	r := fmt.Sprintf("[CQ:image,file=base64://%s]", encoded)
	return r
}

// Record 语音 data 语音文件
func Record(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	r := fmt.Sprintf("[CQ:record,file=base64://%s]", encoded)
	return r
}

// Video 短视频 data视频文件
func Video(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	r := fmt.Sprintf("[CQ:video,file=base64://%s]", encoded)
	return r
}

// At @at data QQ 艾特某人
func At(data int64) string {
	r := fmt.Sprintf("[CQ:at,qq=%d]", data)
	return r
}

// AtAll @all 艾特全体
func AtAll() string {
	r := "[CQ:at,qq=all]"
	return r
}

// Rps 猜拳魔法表情
func Rps() string {
	r := "[CQ:rps]"
	return r
}

// Dice 掷骰子魔法表情
func Dice() string {
	r := "[CQ:dice]"
	return r
}

// Shake 窗口抖动（戳一戳）
func Shake() string {
	r := "[CQ:shake]"
	return r
}

// Poke 戳一戳
func Poke() string {
	r := "[CQ:poke,type=126,id=2003]"
	return r
}

// Anonymous 匿名发消息
func Anonymous() string {
	r := "[CQ:anonymous]"
	return r
}

// Share 链接分享
func Share(url string) string {
	r := fmt.Sprintf("[CQ:share,url=%s,title=百度]", url)
	return r
}

// Contact 推荐好友 QQ 推荐好友QQ
func ContactQq(QQ int64) string {
	r := fmt.Sprintf("[CQ:contact,type=qq,id=%d]", QQ)
	return r
}

// ContactGroup 推荐群
func ContactGroup(group int) string {
	r := fmt.Sprintf("[CQ:contact,type=group,id=%d]", group)
	return r
}

// Location 位置 lat 纬度 long 经度
func Location(lat string, long string) string {
	r := fmt.Sprintf("[CQ:location,lat=%s,lon=%s]", lat, long)
	return r
}

//Music 音乐分享
/*
data 类型 qq 163 xm
id 歌曲id
*/
func Music(data string, id string) string {
	r := fmt.Sprintf("[CQ:music,type=%s,id=%s]", data, id)
	return r
}

//MusicCustom 音乐自定义分享
/*
url 点击后跳转
audio 音乐url
title 标题
*/
func MusicCustom(url string, audio string, title string) string {
	r := fmt.Sprintf("[CQ:music,type=custom,url=%s,audio=%s,title=%s]", url, audio, title)
	return r
}

//Reply 回复
/*
id 回复时引用消息id
*/
func Reply(id string) string {
	r := fmt.Sprintf("[CQ:reply,id=%s]", id)
	return r
}

//Forward 合并转发
/*
id 需通过 get_forward_msg API 获取具体内容
*/
func Forward(id string) string {
	r := fmt.Sprintf("[CQ:forward,id=%s]", id)
	return r
}

//Node 合并转发节点
/*
id 转发的消息 ID
*/
func Node(id string) string {
	r := fmt.Sprintf("[CQ:node,id=%s]", id)
	return r
}

//ForwardNode 合并转发自定义节点
/*
user QQ
nickname 发送者名称
content 消息内容，支持发送消息时的 message 数据类型，见 API 的参数
*/
func ForwardNode(user int64, nickname string, content string) string {
	r := fmt.Sprintf("[CQ:node,user_id=%d,nickname=%s,content=%s]", user, nickname, content)
	return r
}

//DataXml 消息
/*
data xml数据
*/
func DataXml(data string) string {
	r := fmt.Sprintf("[CQ:xml,data=<?xml %s]", data)
	return r
}

//DataJson 消息
/*
data json数据
*/
func DataJson(data string) string {
	r := fmt.Sprintf("[CQ:json,data={\"app\":%s]", data)
	return r
}
