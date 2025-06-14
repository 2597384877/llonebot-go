package main

import (
	"encoding/json"
	"fmt"
	"llonebot/mod/MatchingJson"
	"llonebot/mod/api"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 打印连接信息
	fmt.Printf("接收到来自 %s 的连接\n", conn.RemoteAddr().String())

	// 可以在这里添加处理连接的逻辑
	// 例如，读取客户端发送的数据
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("读取数据时出错:", err)
		return
	}

	message := string(buffer[:n])
	result := textGetRight(message, "\n")
	//fmt.Printf("接收到消息: %s\n", result)
	event(result)

}
func event(text string) {
	var payload MatchingJson.Marcgung
	err := json.Unmarshal([]byte(text), &payload)
	if err != nil {
		fmt.Println("event解析JSON到map时出错:", err)
		return
	}

	if payload.PostType == "message" { //群 私 消息

		var payload MatchingJson.Message
		err := json.Unmarshal([]byte(text), &payload)
		if err != nil {
			fmt.Println("群分类解析JSON到map时出错:", err)
			return
		}

		if payload.MessageType == "group" {
			fmt.Printf("群消息: %s\n", text)

			var data MatchingJson.Group
			err := json.Unmarshal([]byte(text), &data)
			if err != nil {
				fmt.Println("group解析JSON到map时出错:", err)
				return
			}
			Group(data)

		} else if payload.MessageType == "private" {
			fmt.Printf("私消息: %s\n", text)
			var data MatchingJson.Private
			err := json.Unmarshal([]byte(text), &data)
			if err != nil {
				fmt.Println("private解析JSON到map时出错:", err)
				return
			}
			Private(data)
		}

	} else if payload.PostType == "request" { //请求, 例如, 好友申请

		fmt.Printf("事件消息: %s\n", text)

	} else if payload.PostType == "notice" { //通知, 例如, 群成员增加,撤回

		fmt.Printf("通知消息: %s\n", text)

	}

}
func textGetRight(str, separator string) string {
	lastIndex := strings.LastIndex(str, separator)
	if lastIndex == -1 {
		// 没找到分隔符，返回原字符串
		return str
	}
	return str[lastIndex+len(separator):]
}

func main() {
	//发送地址+端口
	api.Host = "192.168.10.10:3000"
	// 要监听的地址和端口
	address := ":5000"

	// 监听指定端口
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("监听端口时出错:", err)
	}
	defer listener.Close()

	fmt.Printf("正在监听端口 %s...\n", address)

	// 循环接受连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("接受连接时出错:", err)
			continue
		}

		// 处理连接
		go handleConnection(conn)
	}
}
