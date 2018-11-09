package util

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

// EncodeJSON json序列化(禁止 html 符号转义)
func EncodeJSON(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//StringToInt string 类型转 int
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("agent 类型转换失败, 请检查配置文件中 agentid 配置是否为纯数字(%v)", err)
		return 0
	}
	return n
}

// HandleContent [P2][PROBLEM][10-13-33-153][][测试 all(#1) net.port.listen port=2 0==0][O3 2017-06-06 16:46:00]
func HandleContent(content string) string {
	//content = strings.Replace(content, "][", "\n", -1)
	if content[0] == '[' {
		content = content[1:]
	}

	if content[len(content)-1] == ']' {
		content = content[:len(content)-1]
	}

	content = strings.Replace(content, "][", " ", -1)
        tmpcontent := strings.Split(content," ")
        fmt.Println(tmpcontent)
        wantcontent :=  "告警级别 : " + tmpcontent[0] + "\n告警状态 : " +  tmpcontent[1] + "\n主机名   : " + tmpcontent[2] + "\n告警内容 : " +   tmpcontent[4] + "\n监控项   : "  + tmpcontent[6] + "\n监控tag  : " + tmpcontent[7] + "\n监控值   : " + tmpcontent[8] + "\n告警次数 : " + tmpcontent[9] + "\n告警时间 : " + tmpcontent[10] +" "+ tmpcontent[11]
        fmt.Println(wantcontent)

	return wantcontent

}
