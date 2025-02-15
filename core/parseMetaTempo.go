package core

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

func (p *PVF) parseMetaTempo(ev smf.Event, msg smf.Message) {
	var bpm float64
	msg.GetMetaTempo(&bpm)
	p.Tempos = []struct { // Tempos 歌曲中的速度变化。
		Bpm   int `json:"bpm"`   // Bpm 每分钟的拍数。
		Ticks int `json:"ticks"` // Ticks 此速度开始的 tick 位置。
		Time  int `json:"time"`  // Time 此速度开始的时间 (通常是秒数)。
	}{{Bpm: int(bpm), Ticks: 0, Time: 0}}
}
