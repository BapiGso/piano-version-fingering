package core

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

// 因为要把midi解析完才能把小节信息补充完整，所以最后一部分放置在了parseTrackV2中完成了
func (p *PVF) parseMetaMeasures(ev smf.Event, absTicks uint64, msg smf.Message, index int) {
	var numerator, denominator, clocksPerClick, demiSemiQuaverPerQuarter uint8
	msg.GetMetaTimeSig(&numerator, &denominator, &clocksPerClick, &demiSemiQuaverPerQuarter)
	//todo 解析p.Measures字段，但是小节是怎么分的没有搞懂
	_, lastTimeSignature := float64(0), int(0)
	//for 循环思路，当前的delta/小节拍分子xMetricTicks比如1536/256*3=2 循环2次，最少循环一次
	// 思路改一下，loop>1多次循环重复的最后在正常aapend一次变化的
	ticksPerMeasure := int(numerator) * p.Resolution
	if len(p.Measures) > 0 {
		_ = p.Measures[len(p.Measures)-1].TicksStart
		lastTimeSignature = p.Measures[len(p.Measures)-1].TimeSignature[0]
		if loop := int(ev.Delta) / lastTimeSignature / p.Resolution; loop > 1 {
			for k := 1; k < loop; k++ {
				tmp := p.Measures[len(p.Measures)-1]
				tmp.TicksStart += tmp.TotalTicks
				tmp.Time += float64(tmp.TicksPerMeasure / p.Resolution)
				tmp.Type = 2
				p.Measures = append(p.Measures, tmp)
			}
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
		Time:            float64(int(absTicks) / p.Resolution),   //maybe有问题
		TimeSignature:   []int{int(numerator), int(denominator)}, //这个没问题不要动
		TicksPerMeasure: ticksPerMeasure,                         //这个没问题不要动
		TicksStart:      float64(absTicks),                       //这一次的开始等于上一次的开始+上一次小节xMetricTicks
		TotalTicks:      float64(ticksPerMeasure),                //这个没问题不要动
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
