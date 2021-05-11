package httpRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)


type ReportResponse struct {
	ErrCode int `json:"errCode"`
	Message string `json:"message"`
	RequestId string `json:"requestId"`
}

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "request failed. err: " + err.Error()
	}
	
	respBody := resp.Body
	statusCode := resp.StatusCode
	defer resp.Body.Close()
	
	respBodyByte, err := ioutil.ReadAll(respBody)
	//fmt.Println("statusCode: ", statusCode)
	//fmt.Println("respBody: ", string(respBodyByte))
	
	return "request successful. statusCode: " + strconv.Itoa(statusCode) +". respBody: " + string(respBodyByte)
}

// 上报检测内容
func ReportTrace(reportUrl string, reportData []byte) {
	var reportResponse ReportResponse
	resp, err := http.Post(reportUrl, "application/json", bytes.NewBuffer(reportData))
	
	if err != nil {
		fmt.Println("上报失败, err: ", err.Error())
		return
	}
	
	respBody := resp.Body
	
	defer resp.Body.Close()
	
	respBodyByte, err := ioutil.ReadAll(respBody)
	if err != nil {
		fmt.Println("读取上报结果失败, err: ", err.Error())
		return
	}
	
	err = json.Unmarshal(respBodyByte, &reportResponse)
	if err != nil {
		fmt.Println("解析上报信息Response失败, err: ", err.Error())
		return
	}
	
	fmt.Println("上报信息结果: ", string(respBodyByte))
	
}