package deliver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostBack() error {
	data := RequestData{
		Name: "John",
		Age:  30,
	}

	//json编码
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("JSON 编码错误：%v", err)
	}

	//发送注册的数据
	response, err := http.POST("http://39.103.157.137:8090/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求错误：%v", err)
	}

	//获得响应
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("读取响应错误：%v", err)
	}

	fmt.Println("响应内容：", string(body))

	return nil
}
