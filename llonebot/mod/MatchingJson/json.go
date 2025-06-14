package MatchingJson

type Marcgung struct {
	PostType string `json:"post_type"`
}
type Message struct {
	MessageType string `json:"message_type"`
}

// 群消息json
type Group struct {
	SelfId      int    `json:"self_id"`
	UserId      int64  `json:"user_id"`
	Time        int    `json:"time"`
	MessageId   int    `json:"message_id"`
	MessageSeq  int    `json:"message_seq"`
	MessageType string `json:"message_type"`
	Sender      struct {
		UserId   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Role     string `json:"role"`
		Title    string `json:"title"`
	} `json:"sender"`
	RawMessage    string `json:"raw_message"`
	Font          int    `json:"font"`
	SubType       string `json:"sub_type"`
	Message       string `json:"message"`
	MessageFormat string `json:"message_format"`
	PostType      string `json:"post_type"`
	GroupId       int    `json:"group_id"`
}

// 私聊消息json
type Private struct {
	SelfId      int    `json:"self_id"`
	UserId      int64  `json:"user_id"`
	Time        int    `json:"time"`
	MessageId   int    `json:"message_id"`
	MessageSeq  int    `json:"message_seq"`
	MessageType string `json:"message_type"`
	Sender      struct {
		UserId   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
	} `json:"sender"`
	RawMessage    string `json:"raw_message"`
	Font          int    `json:"font"`
	SubType       string `json:"sub_type"`
	Message       string `json:"message"`
	MessageFormat string `json:"message_format"`
	PostType      string `json:"post_type"`
}
