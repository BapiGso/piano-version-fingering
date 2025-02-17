package core

import (
	"fmt"
	"gitlab.com/gomidi/midi/v2/smf"
)

func (p *PVF) parseMetaTimeSig(ev smf.Event, absTicks uint64, msg smf.Message, index int) {
	var numerator, denominator, clocksPerClick, demiSemiQuaverPerQuarter uint8
	msg.GetMetaTimeSig(&numerator, &denominator, &clocksPerClick, &demiSemiQuaverPerQuarter)
	//fmt.Println(numerator, denominator, clocksPerClick, demiSemiQuaverPerQuarter)
	var measures, lastMeasures, lastTicks int
	if len(p.TimeSignatures) > 0 { //跳过第一次解析
		lastMeasures = p.TimeSignatures[len(p.TimeSignatures)-1].Measures
		lastTicks = p.TimeSignatures[len(p.TimeSignatures)-1].Ticks
		lastNumerator := p.TimeSignatures[len(p.TimeSignatures)-1].TimeSignature[0]
		measuresDelta := (int(absTicks) - lastTicks) / (lastNumerator * p.Resolution) //计算有几个小节 =tickDelta/节拍分子xMetricTicks
		measures = measuresDelta + lastMeasures                                       //总的小节=上一个小节+本次小节
	}

	p.TimeSignatures = append(p.TimeSignatures, struct {
		Ticks         int   `json:"ticks"`
		TimeSignature []int `json:"timeSignature"`
		Measures      int   `json:"measures"`
	}{
		Ticks:         int(absTicks),
		TimeSignature: []int{int(numerator), int(denominator)},
		Measures:      measures,
	})
	fmt.Printf("钢琴拍号变化MetaTimeSig: %v ,delta:%v,ticktotal:%v\n", ev.Message.String(), ev.Delta, absTicks)

}
