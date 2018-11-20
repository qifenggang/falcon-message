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

	//content = strings.Replace(content, "][", " ", -1)
    tmpcontent := strings.Split(content,"][")
    fmt.Println(tmpcontent)
    alertcontent := strings.Split(tmpcontent[4]," ")
    fmt.Println(alertcontent)
    alertstring := strings.Join(alertcontent[0:len(alertcontent)-4], " ")
    timecontent := strings.Split(tmpcontent[5]," ")
    wantcontent :=  "============falcon============" + "\n告警级别 : " + tmpcontent[0] + "\n告警状态 : " +  tmpcontent[1] + "\n主机名   : " + tmpcontent[2] + "\n告警内容 : " +  alertstring + "\n监控项   : "  + alertcontent[len(alertcontent)-3] + "\n监控tag  : " + alertcontent[len(alertcontent)-2] + "\n监控值   : " + alertcontent[len(alertcontent)-1] + "\n告警次数 : " + timecontent[0] + "\n告警时间 : " + timecontent[1] +" "+ timecontent[2]
    fmt.Println(wantcontent)


	return wantcontent

}
