package core

import (
	"bytes"
	"fmt"
	_ "gitlab.com/gomidi/midi/v2/drivers/midicat" // autoregisters driver
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
	"log"
	"math"
)

func (p *PVF) Parse(file []byte) {

	// read the bytes
	s, err := smf.ReadFrom(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Sscanf(s.TimeFormat.String(), "%d MetricTicks", &p.Resolution)
	fmt.Printf("got %v tracks\n", s.NumTracks())
	println(len(s.Tracks[1]))
	for no, track := range s.Tracks {

		// it might be a good idea to go from delta ticks to absolute ticks.
		var absTicks uint64

		var trackname string
		var channel, program uint8
		var gm_name string

		for key, ev := range track {
			absTicks += uint64(ev.Delta)
			msg := ev.Message

			if msg.Type() == smf.MetaEndOfTrackMsg {
				// ignore
				continue
			}

			//这个库的写法有点怪，是传参数指针赋值，所以判断用空值调用一次，子函数里在调用一次来获取值
			switch {
			case msg.GetMetaTempo(nil): //获取BPM
				p.parseMetaTempo(ev, msg)
			case msg.GetChannel(nil): //获取钢琴音轨
				p.parseChannel(ev, no, absTicks, trackname, gm_name)
			case msg.GetMetaKeySig(new(uint8), new(uint8), new(bool), new(bool)): //todo只判断了C大调的写法
				p.parseMetaKeySig(ev, msg)
			case msg.GetMetaTimeSig(new(uint8), new(uint8), new(uint8), new(uint8)):
				p.parseMetaTimeSig(ev, absTicks, msg, key)
				if ev.Delta != 0 || absTicks == 0 {
					p.parseMetaMeasures(ev, absTicks, msg, key)
				}
			case msg.GetMetaTrackName(nil):
				msg.GetMetaTrackName(&trackname)
				fmt.Printf("钢琴TrackName: %v\n", decodeTrackName(trackname)) // 自动检测并解码 解码音轨名称 (假设为 Shift-JIS)
			case msg.GetMetaInstrument(nil):
				fmt.Printf("钢琴Instrument: %v\n", trackname) // 自动检测并解码 同样解码乐器名称
			case msg.GetMetaCopyright(nil):
				fmt.Printf("钢琴Copyright: %v\n", ev.Message)
			case msg.GetSysEx(new([]byte)):
				fmt.Printf("也许无关紧要的系统信息 %v\n", ev.Message)
			case msg.GetProgramChange(&channel, &program):
				gm_name = gm.Instr(program).String()
			case msg.IsMeta():
				fmt.Printf("元信息track %v %s %s @%v %s\n", no, trackname, gm_name, absTicks, ev.Message)
			default:
				fmt.Printf("其他信息track %v %s %s @%v %s\n", no, trackname, gm_name, absTicks, ev.Message)
			}
		}
		p.SongLength = int(math.Max(float64(p.SongLength), float64(int(absTicks)/p.Resolution)))
		//fmt.Scanln()
	}
	p.parseTrackV2()
	p.ParseEnd()
}
