package core

import (
	"fmt"
	"gitlab.com/gomidi/midi/v2/smf"
	"math"
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
	//todo 解析p.Measures字段，但是小节是怎么分的没有搞懂
	// 比如4*256<19456
	lastTicksStart, lastTimeSignature := float64(0), int(0)

	//for循环思路，当前的delta/小节拍分子xMetricTicks比如1536/256*3=2 循环2次，最少循环一次
	ticksPerMeasure := int(numerator) * p.Resolution
	for k := range int(math.Max(1, float64(int(ev.Delta)/ticksPerMeasure-1))) {
		if len(p.Measures) > 0 {
			lastTicksStart = p.Measures[len(p.Measures)-1].TicksStart
			lastTimeSignature = p.Measures[len(p.Measures)-1].TimeSignature[0]
		}
		if lastTicksStart != 0 && ev.Delta == 0 { // 这个事件很特殊，节拍没变但就是解析到了一次，所以跳过
			continue
		}
		if int(ev.Delta)/ticksPerMeasure > 1 { //这一部分写的脑溢血
			fmt.Println(k, int(ev.Delta)/ticksPerMeasure, ev.Message)
			ticksPerMeasure = p.Measures[len(p.Measures)-1].TicksPerMeasure
			numerator = uint8(lastTimeSignature)
		}
		p.Measures = append(p.Measures, struct {
			Time            float64 `json:"time"`
			TimeSignature   []int   `json:"timeSignature"`
			TicksPerMeasure int     `json:"ticksPerMeasure"`
			TicksStart      float64 `json:"ticksStart"`
			TotalTicks      float64 `json:"totalTicks"`
			Type            int     `json:"type"`
		}{
			Time:            float64(int(lastTicksStart+float64(ev.Delta)) / p.Resolution),
			TimeSignature:   []int{int(numerator), int(denominator)},
			TicksPerMeasure: ticksPerMeasure,
			TicksStart:      lastTicksStart + float64(lastTimeSignature*p.Resolution),
			TotalTicks:      float64(ticksPerMeasure),
			Type:            calcMeasuresType(numerator, lastTimeSignature),
		})
	}

}

func calcMeasuresType(numerator uint8, lastNumerator int) int {
	if lastNumerator == 0 {
		return 0
	}
	if numerator == uint8(lastNumerator) {
		return 2
	}
	return 1
}
