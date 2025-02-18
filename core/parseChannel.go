package core

import (
	"fmt"
	"gitlab.com/gomidi/midi/v2/smf"
)

func (p *PVF) parseChannel(ev smf.Event, no int, absTicks uint64, trackname string, gm_name string) {
	var Channel, Midi int
	var Velocity float64
	//fmt.Println(ev.Message.String())
	_, err := fmt.Sscanf(ev.Message.String(), "NoteOn channel: %d key: %d velocity: %f", &Channel, &Midi, &Velocity)
	if err == nil && Velocity != 0 {
		p.SupportingTracks[no].MyInstrument = -5
		p.SupportingTracks[no].TheirInstrument = 0
		p.SupportingTracks[no].Notes = append(p.SupportingTracks[no].Notes, struct {
			Midi     int     `json:"midi"`
			Time     float64 `json:"time"`
			Velocity float64 `json:"velocity"`
			Duration float64 `json:"duration"`
		}{Midi: Midi, Time: float64(absTicks) / float64(p.Resolution), Velocity: Velocity / 127, Duration: 0})
	} else if err == nil && Velocity == 0 { //用于计算按键持续时间
		p.SupportingTracks[no].Notes[len(p.SupportingTracks[no].Notes)-1].Duration = float64(ev.Delta) / float64(p.Resolution)
	}
	fmt.Printf("钢琴音轨track %v %s %s @%v %s\n", no, trackname, gm_name, absTicks, ev.Message)
}
