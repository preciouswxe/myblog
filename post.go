package post

import (
	"encoding/json"
	"fmt"
)

type PostData struct {
	Type    int `json:"type"`
	Content struct {
		DevEui    string `json:"devEui"`
		Data      string `json:"data"`
		FPort     int    `json:"fPort"`
		Timestamp int64  `json:"timestamp"`
	} `json:"content"`
	ProjectId  string `json:"projectId"`
	ServerTime int64  `json:"serverTime"`
}

func PostFunction(data []byte) {

	//接受data注册信息后进行json解析
	var pData PostData
	err := json.Unmarshal(data, &pData)
	if err != nil {
		fmt.Println("JSON解析失败:", err)
		return
	}

	//连接数据库
	initDB()

}
