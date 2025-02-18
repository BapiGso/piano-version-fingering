package core

import (
	"encoding/json"
	"fmt"
)

type output struct {
	Right []MusicSectionLite
	Left  []MusicSectionLite
}

func (p *PVF) Export4LLM() (string, error) { // 返回 JSON 字符串和 error
	o := new(output)
	//o.BPM = p.Tempos[0].Bpm

	// 初始化 slice，避免 panic
	o.Right = make([]MusicSectionLite, len(p.TracksV2.Right))
	o.Left = make([]MusicSectionLite, len(p.TracksV2.Left))

	for i, section := range p.TracksV2.Right {
		o.Right[i] = MusicSectionLite(section) // 假设 MusicSectionLite 可以直接从 MusicSection 转换
	}
	for i, section := range p.TracksV2.Left {
		o.Left[i] = MusicSectionLite(section) // 假设 MusicSectionLite 可以直接从 MusicSection 转换
	}

	// 将结构体编码为 JSON 字符串
	jsonData, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return "", err // 返回空字符串和 error
	}

	return string(jsonData), nil // 返回 JSON 字符串和 nil error
}

type MusicSectionLite struct { // Right 右手音轨信息。
	Direction     string     `json:"-"`  // Direction  方向，一般为 right 或者 left
	Time          float64    `json:"-"`  // Time 音轨开始的时间 (通常是秒数)。
	TimeEnd       float64    `json:"-"`  // TimeEnd 音轨结束的时间 (通常是秒数)。
	TimeSignature []int      `json:"ts"` // TimeSignature 音轨的拍号。
	Notes         []struct { // Notes 音轨上的所有音符信息。
		Note            int         `json:"-"` // Note 音符的 MIDI 值。
		DurationTicks   int         `json:"-"` // DurationTicks 音符的持续时间 (以 ticks 为单位)。
		NoteOffVelocity int         `json:"-"` // NoteOffVelocity 音符结束时的力度。
		TicksStart      int         `json:"-"` // TicksStart 音符开始的 tick 位置。
		Velocity        float64     `json:"-"` // Velocity 音符的力度。
		MeasureBars     float64     `json:"-"` // MeasureBars 音符在小节中的位置 (以小节为单位)。
		Duration        float64     `json:"-"` // Duration 音符的持续时间 (通常是秒数)。
		NoteName        string      `json:"n"` // NoteName 音符的名称 (例如 "C", "D#", "Bb")。
		Octave          int         `json:"-"` // Octave 音符的八度音阶。
		NotePitch       string      `json:"-"` // NotePitch 音符的音高 (例如 "C4", "G#5")。
		Start           float64     `json:"-"` // Start 音符开始的时间 (通常是秒数)。
		End             float64     `json:"-"` // End 音符结束的时间 (通常是秒数)。
		NoteLengthType  string      `json:"-"` // NoteLengthType 音符的长度类型 (例如 "quarter", "eighth")。
		Group           int         `json:"-"` // Group 音符所属的组 (用于连音等)。
		MeasureInd      int         `json:"-"` // MeasureInd 音符所在的小节的索引。
		NoteMeasureInd  int         `json:"-"` // NoteMeasureInd 音符在小节中的索引。
		Id              string      `json:"-"` // Id 音符的唯一标识符。
		Finger          int         `json:"f"` // Finger  指法建议
		Smp             interface{} `json:"-"` // Smp  未知，可能和音色采样相关
	} `json:"notes"`
	Max               int        `json:"-"` // Max 音轨中最高音符的 MIDI 值。
	Min               int        `json:"-"` // Min 音轨中最低音符的 MIDI 值。
	MeasureTicksStart float64    `json:"-"` // MeasureTicksStart 音轨开始的小节的 ticks 位置。
	MeasureTicksEnd   float64    `json:"-"` // MeasureTicksEnd 音轨结束的小节的 ticks 位置。
	Rests             []struct { // Rests 音轨上的休止符信息。
		Time           float64 `json:"-"` // Time 休止符开始的时间 (通常是秒数)。
		NoteLengthType string  `json:"-"` // NoteLengthType 休止符的长度类型 (例如 "quarter", "eighth")。
	} `json:"-"`
	Groups []struct { // Groups 音轨上的音符分组信息。
		Start          float64 `json:"-"` // Start 组开始的时间 (通常是秒数)。
		End            float64 `json:"-"` // End 组结束的时间 (通常是秒数)。
		NoteLengthType string  `json:"-"` // NoteLengthType 组的长度类型 (例如 "quarter", "eighth")。
		NoteInds       []int   `json:"-"` // NoteInds 组中的音符索引。
		GroupId        int     `json:"-"` // GroupId 组的唯一标识符。
		BarY           int     `json:"-"` // BarY  在五线谱上的位置
	} `json:"-"`
}
