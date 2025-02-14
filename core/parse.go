package core

import (
	"bytes"
	"fmt"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
	"log"
)

func (p *PVF) Parse(file []byte) {

	// read the bytes
	s, err := smf.ReadFrom(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Sscanf(s.TimeFormat.String(), "%d MetricTicks", &p.Resolution)
	fmt.Printf("got %v tracks\n", s.NumTracks())

	for no, track := range s.Tracks {

		// it might be a good idea to go from delta ticks to absolute ticks.
		var absTicks uint64

		var trackname string
		var channel, program uint8
		var gm_name string

		for _, ev := range track {
			absTicks += uint64(ev.Delta)
			msg := ev.Message

			if msg.Type() == smf.MetaEndOfTrackMsg {
				// ignore
				continue
			}

			switch {
			case msg.GetMetaTempo(nil): //获取BPM
				p.Tempos = []struct { // Tempos 歌曲中的速度变化。
					Bpm   int `json:"bpm"`   // Bpm 每分钟的拍数。
					Ticks int `json:"ticks"` // Ticks 此速度开始的 tick 位置。
					Time  int `json:"time"`  // Time 此速度开始的时间 (通常是秒数)。
				}{{Bpm: 0, Ticks: 0, Time: 0}}
				_, err = fmt.Sscanf(ev.Message.String(), "MetaTempo bpm: %d.00", &p.Tempos[0].Bpm)
				//fmt.Println(p.Tempos)
			case msg.GetChannel(nil): //获取钢琴音轨
				//fmt.Println(len(p.SupportingTracks))
				var Midi int
				var Velocity float64
				_, err = fmt.Sscanf(ev.Message.String(), "NoteOn channel: 0 key: %d velocity: %d", &Midi, &Velocity)
				p.SupportingTracks[no-1].MyInstrument = -5
				p.SupportingTracks[no-1].TheirInstrument = 0
				p.SupportingTracks[no-1].Notes = append(p.SupportingTracks[no-1].Notes, struct {
					Midi     int     `json:"midi"`
					Time     float64 `json:"time"`
					Velocity float64 `json:"velocity"`
					Duration float64 `json:"duration"`
				}{Midi: Midi, Time: 0, Velocity: Velocity / 255, Duration: float64(ev.Delta / uint32(p.Resolution))})
				fmt.Println(ev.Message)
				//fmt.Printf("track %v %s %s @%v %s\n", no, trackname, gm_name, absTicks, ev.Message)
			case msg.IsMeta():
				fmt.Println(123, no, ev.Message)
			case msg.GetMetaTrackName(&trackname):
				trackname = decodeTrackName(trackname) // 自动检测并解码 解码音轨名称 (假设为 Shift-JIS)
			case msg.GetMetaInstrument(&trackname):
				trackname = decodeTrackName(trackname) // 自动检测并解码 同样解码乐器名称
			case msg.GetProgramChange(&channel, &program):
				gm_name = fmt.Sprintf("(%v)", gm.Instr(program).String())
			default:
				//fmt.Println(ev.Message)
				fmt.Printf("track %v %s %s @%v %s\n", no, trackname, gm_name, absTicks, ev.Message)
			}
		}
	}
}
