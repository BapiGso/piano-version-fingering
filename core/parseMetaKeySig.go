package core

import (
	"fmt"
	"gitlab.com/gomidi/midi/v2/smf"
)

func (p *PVF) parseMetaKeySig(ev smf.Event, msg smf.Message) {
	var key, num uint8
	var isMajor, isFlat bool
	msg.GetMetaKeySig(&key, &num, &isMajor, &isFlat)
	//fmt.Println(key, num, isMajor, isFlat)
	//_, err := fmt.Sscanf(ev.Message.String(), "MetaKeySig key: %v", &key)
	//if err != nil {
	//	fmt.Println(err)
	//}
	p.KeySignatures = append(p.KeySignatures, struct {
		Key   string `json:"key"`
		Scale string `json:"scale"`
		Ticks int    `json:"ticks"`
	}{}) //{Key: key[:1], Scale: key[1:], Ticks: 0}) //todo 未完成并且解析的不规范
	fmt.Printf("钢琴调MetaKeySig: %v\n", ev.Message)
}
