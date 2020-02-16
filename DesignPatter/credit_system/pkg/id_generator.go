package pkg

import (
	"bytes"
	"strconv"
)

//   生成channel_id
// 2. channel_id: 积分赚取或消费的途径（1位） + id
//3. event_id: 事件类型（两位）+ id
func GenerateChannleId(id, way, ctype int)  string {
	var channelId bytes.Buffer
	channelId.WriteString(strconv.Itoa(ctype))
	channelId.WriteString(strconv.Itoa(way))
	channelId.WriteString(strconv.Itoa(id))
	return channelId.String()
}

// 生成eventid
func GenerateEventId(id, ctype int) string {
	var eventId bytes.Buffer
	eventId.WriteString(strconv.Itoa(ctype))
	eventId.WriteString(strconv.Itoa(id))
	return eventId.String()
}

// 通过eventID获取消费事件类型
func GetType(id string) (ctype int) {

	ids := []rune(id)
	ctype, _ = strconv.Atoi(string(ids[0]))

	return ctype
}

