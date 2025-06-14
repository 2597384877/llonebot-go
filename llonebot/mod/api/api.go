package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var Host string

// SendGroupMsg 发送群消息
/*group 群号，message 内容*/
func SendGroupMsg(group int, message string) string {
	u := struct {
		Group      int    `json:"group_id"`
		Message    string `json:"message"`
		AutoEscape bool   `json:"auto_escape"`
	}{
		Group:      group,
		Message:    message,
		AutoEscape: false,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/send_group_msg",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SendPrivateMsg 发送私聊消息
/*user QQ,message 内容*/
func SendPrivateMsg(user int, message string) string {

	u := struct {
		User       int    `json:"user_id"`
		Message    string `json:"message"`
		AutoEscape bool   `json:"auto_escape"`
	}{
		User:       user,
		Message:    message,
		AutoEscape: false,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/send_group_msg",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// DeleteMsg 撤回
/*messageId 消息ID*/
func DeleteMsg(messageId int) string {

	u := struct {
		MessageId int `json:"message_id"`
	}{
		MessageId: messageId,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/delete_msg",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// GetMsg 获取消息 messageId 消息ID
/*
响应数据
字段名	数据类型	说明
time	number (int32)	发送时间
message_type	string	消息类型，同 消息事件
message_id	number (int32)	消息 ID
real_id	number (int32)	消息真实 ID
sender	object	发送人信息，同 消息事件
message	message	消息内容
*/
func GetMsg(messageId int) string {

	u := struct {
		MessageId int `json:"message_id"`
	}{
		MessageId: messageId,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_msg",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// GetForwardMsg 获取合并转发消息 id 合并转发 ID
/*
字段名	类型	说明
message	message	消息内容，使用 消息的数组格式 表示，数组中的消息段全部为 node 消息段
*/
func GetForwardMsg(id int) string {

	u := struct {
		MessageId int `json:"message_id"`
	}{
		MessageId: id,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_forward_msg",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SendLike 发送好友赞
/*
userId QQ,times 赞数量
*/
func SendLike(userId int, times int) string {

	u := struct {
		UserId int `json:"user_id"`
		Times  int `json:"times"`
	}{
		UserId: userId,
		Times:  times,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/send_like",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupKick 群组踢人
/*
group 群号，user QQ，rejectAddRequest 是否踢出bool
*/
func SetGroupKick(group int, user int, rejectAddRequest bool) string {

	u := struct {
		Group            int  `json:"group_id"`
		User             int  `json:"user_id"`
		RejectAddRequest bool `json:"reject_add_request"`
	}{
		Group:            group,
		User:             user,
		RejectAddRequest: rejectAddRequest,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_kick",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupBan 群组单人禁言
/*
group 群号，user QQ，duration 禁言时长
*/
func SetGroupBan(group int, user int, duration int) string {

	u := struct {
		Group    int `json:"group_id"`
		User     int `json:"user_id"`
		Duration int `json:"duration"`
	}{
		Group:    group,
		User:     user,
		Duration: duration,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_ban",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupWholeBan 群组全员禁言
/*
group 群号，enable 是否禁言 true false
*/
func SetGroupWholeBan(group int, enable bool) string {

	u := struct {
		Group  int  `json:"group_id"`
		Enable bool `json:"enable"`
	}{
		Group:  group,
		Enable: enable,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_whole_ban",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupAdmin 群组设置管理员
/*
group 群号，user QQ，enable 管理员true false
*/
func SetGroupAdmin(group int, user int, enable bool) string {

	u := struct {
		Group  int  `json:"group_id"`
		User   int  `json:"user_id"`
		Enable bool `json:"enable"`
	}{
		Group:  group,
		User:   user,
		Enable: enable,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_admin",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupAnonymous 群组匿名
/*
group 群号，enable 匿名是否 true false
*/
func SetGroupAnonymous(group int, enable bool) string {

	u := struct {
		Group  int  `json:"group_id"`
		Enable bool `json:"enable"`
	}{
		Group:  group,
		Enable: enable,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_anonymous",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupCard 设置群名片（群备注）
/*
group 群号，user QQ，card 设置群名片
*/
func SetGroupCard(group int, user int, card string) string {

	u := struct {
		Group int    `json:"group_id"`
		User  int    `json:"user_id"`
		Card  string `json:"enable"`
	}{
		Group: group,
		User:  user,
		Card:  card,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_card",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupName 设置群名
/*
group 群号，groupName 设置群名
*/
func SetGroupName(group int, groupName string) string {

	u := struct {
		Group     int    `json:"group_id"`
		GroupName string `json:"group_name"`
	}{
		Group:     group,
		GroupName: groupName,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_name",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupLeave 退出群组
/*
group 群号，isDismiss 是否解散，如果登录号是群主，则仅在此项为 true 时能够解散
*/
func SetGroupLeave(group int, isDismiss bool) string {

	u := struct {
		Group     int  `json:"group_id"`
		IsDismiss bool `json:"is_dismiss"`
	}{
		Group:     group,
		IsDismiss: isDismiss,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_leave",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

// SetGroupSpecialTitle 设置群组专属头衔
/*
group 群号，user QQ,specialTitle 头衔名称
*/
func SetGroupSpecialTitle(group int, user int, specialTitle string) string {

	u := struct {
		Group        int    `json:"group_id"`
		User         int    `json:"user_id"`
		SpecialTitle string `json:"special_title"`
		Duration     int    `json:"duration"`
	}{
		Group:        group,
		User:         user,
		SpecialTitle: specialTitle,
		Duration:     -1,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_special_title",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//SetFriendAddRequest 处理加好友请求
/*
flag 加好友请求的 flag（需从上报的数据中获得）
approve 是否同意请求
remark 添加后的好友备注（仅在同意时有效） 可忽略
*/
func SetFriendAddRequest(flag string, approve bool, remark string) string {

	u := struct {
		Flag    string `json:"flag"`
		Approve bool   `json:"approve"`
		Remark  string `json:"remark"`
	}{
		Flag:    flag,
		Approve: approve,
		Remark:  remark,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_friend_add_request",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//SetGroupAddRequest 处理加群请求／邀请
/*
flag 加群请求的 flag（需从上报的数据中获得）
sub_type 或 type	string	-	add 或 invite，请求类型（需要和上报消息中的 sub_type 字段相符）
approve	boolean	true	是否同意请求／邀请
reason	string	空	拒绝理由（仅在拒绝时有效）*/
func SetGroupAddRequest(flag string, subType string, approve bool, reason string) string {

	u := struct {
		Flag    string `json:"flag"`
		SubType string `json:"sub_type"`
		Approve bool   `json:"approve"`
		Reason  string `json:"reason"`
	}{
		Flag:    flag,
		SubType: subType,
		Approve: approve,
		Reason:  reason,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/set_group_add_request",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetLoginInfo 获取登录号信息
/*响应数据
字段名	数据类型	说明
user_id	number (int64)	QQ 号
nickname	string	QQ 昵称*/
func GetLoginInfo() string {

	u := struct {
	}{}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_login_info",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetStrangerInfo 获取陌生人信息
/*

user QQ

响应数据
字段名	数据类型	说明
user_id	number (int64)	QQ 号
nickname	string	昵称
sex	string	性别，male 或 female 或 unknown
age	number (int32)	年龄*/
func GetStrangerInfo(user int) string {

	u := struct {
		UserId  int  `json:"user_id"`
		NoCache bool `json:"no_cache"`
	}{
		UserId:  user,
		NoCache: false,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_stranger_info",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetFriendList 获取好友列表
/*
响应数据
响应内容为 JSON 数组，每个元素如下：

字段名	数据类型	说明
user_id	number (int64)	QQ 号
nickname	string	昵称
remark	string	备注名*/
func GetFriendList() string {

	u := struct {
	}{}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_friend_list",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetGroupInfo 获取群信息
/*
group_id 群号

响应数据
字段名	数据类型	说明
group_id	number (int64)	群号
group_name	string	群名称
member_count	number (int32)	成员数
max_member_count	number (int32)	最大成员数（群容量）
*/
func GetGroupInfo(group int) string {

	u := struct {
		Group   int  `json:"group_id"`
		NoCache bool `json:"no_cache"`
	}{
		Group:   group,
		NoCache: false,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_group_info",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetGroupList 获取群列表
/*
响应数据
响应内容为 JSON 数组，每个元素和上面的 get_group_info 接口相同。*/
func GetGroupList() string {

	u := struct {
	}{}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_group_list",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetGroupMemberInfo 获取群成员信息
/*
group 群号
user QQ

响应数据
字段名	数据类型	说明
group_id	number (int64)	群号
user_id	number (int64)	QQ 号
nickname	string	昵称
card	string	群名片／备注
sex	string	性别，male 或 female 或 unknown
age	number (int32)	年龄
area	string	地区
join_time	number (int32)	加群时间戳
last_sent_time	number (int32)	最后发言时间戳
level	string	成员等级
role	string	角色，owner 或 admin 或 member
unfriendly	boolean	是否不良记录成员
title	string	专属头衔
title_expire_time	number (int32)	专属头衔过期时间戳
card_changeable	boolean	是否允许修改群名片*/
func GetGroupMemberInfo(group int, user int) string {

	u := struct {
		Group   int  `json:"group_id"`
		User    int  `json:"user_id"`
		NoCache bool `json:"no_cache"`
	}{
		Group:   group,
		User:    user,
		NoCache: false,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_group_member_info",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetGroupMemberList 获取群成员列表
/*
group 群号

响应数据
响应内容为 JSON 数组，每个元素的内容和上面的 get_group_member_info 接口相同，但对于同一个群组的同一个成员，获取列表时和获取单独的成员信息时，某些字段可能有所不同，例如 area、title 等字段在获取列表时无法获得，具体应以单独的成员信息为准。
*/
func GetGroupMemberList(group int, user int) string {

	u := struct {
		Group int `json:"group_id"`
	}{
		Group: group,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_group_member_list",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}

//GetGroupHonorInfo 获取群荣誉信息
/*
group 群号
type 要获取的群荣誉类型，可传入 talkative performer legend strong_newbie emotion 以分别获取单个类型的群荣誉数据，或传入 all 获取所有数据

响应数据
字段名	数据类型	说明
group_id	number (int64)	群号
current_talkative	object	当前龙王，仅 type 为 talkative 或 all 时有数据
talkative_list	array	历史龙王，仅 type 为 talkative 或 all 时有数据
performer_list	array	群聊之火，仅 type 为 performer 或 all 时有数据
legend_list	array	群聊炽焰，仅 type 为 legend 或 all 时有数据
strong_newbie_list	array	冒尖小春笋，仅 type 为 strong_newbie 或 all 时有数据
emotion_list	array	快乐之源，仅 type 为 emotion 或 all 时有数据
其中 current_talkative 字段的内容如下：

字段名	数据类型	说明
user_id	number (int64)	QQ 号
nickname	string	昵称
avatar	string	头像 URL
day_count	number (int32)	持续天数
其它各 *_list 的每个元素是一个 JSON 对象，内容如下：

字段名	数据类型	说明
user_id	number (int64)	QQ 号
nickname	string	昵称
avatar	string	头像 URL
description	string	荣誉描述*/
func GetGroupHonorInfo(group int, typeID string) string {

	u := struct {
		Group  int    `json:"group_id"`
		TypeID string `json:"honor_info"`
	}{
		Group:  group,
		TypeID: typeID,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://"+Host+"/get_group_honor_info",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() { _ = r.Body.Close() }()
	text, _ := io.ReadAll(r.Body)
	return string(text)
}
