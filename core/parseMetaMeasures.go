package core

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

func (p *PVF) parseMetaMeasures(ev smf.Event, absTicks uint64, msg smf.Message, index int) {
	var numerator, denominator, clocksPerClick, demiSemiQuaverPerQuarter uint8
	msg.GetMetaTimeSig(&numerator, &denominator, &clocksPerClick, &demiSemiQuaverPerQuarter)
	//todo 解析p.Measures字段，但是小节是怎么分的没有搞懂
	// 比如4*256<19456
	lastTicksStart, lastTimeSignature := float64(0), int(0)
	//for 循环思路，当前的delta/小节拍分子xMetricTicks比如1536/256*3=2 循环2次，最少循环一次
	// 思路改一下，loop>1多次循环重复的最后在正常aapend一次变化的
	ticksPerMeasure := int(numerator) * p.Resolution
	if loop := int(ev.Delta) / ticksPerMeasure; loop > 1 {
		for k := 0; k < loop-1; k++ {
			tmp := p.Measures[len(p.Measures)-2]
			tmp.Time += float64(tmp.TicksPerMeasure / p.Resolution)
			tmp.TicksStart += tmp.TotalTicks
			tmp.Type = 2
			p.Measures = append(p.Measures, tmp)
		}
	}
	p.Measures = append(p.Measures, struct {
		Time            float64 `json:"time"`
		TimeSignature   []int   `json:"timeSignature"`
		TicksPerMeasure int     `json:"ticksPerMeasure"`
		TicksStart      float64 `json:"ticksStart"`
		TotalTicks      float64 `json:"totalTicks"`
		Type            int     `json:"type"`
	}{
		Time:            float64(int(lastTicksStart+float64(ticksPerMeasure)) / p.Resolution),
		TimeSignature:   []int{int(numerator), int(denominator)},                  //这个没问题不要动
		TicksPerMeasure: ticksPerMeasure,                                          //这个没问题不要动
		TicksStart:      lastTicksStart + float64(lastTimeSignature*p.Resolution), //这一次的开始等于上一次的开始+上一次小节xMetricTicks
		TotalTicks:      float64(ticksPerMeasure),                                 //这个没问题不要动
		Type:            calcMeasuresType(numerator, lastTimeSignature),
	})

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
