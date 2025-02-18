package core

type MusicSection struct { // Right 右手音轨信息。
	Direction     string     `json:"direction"`     // Direction  方向，一般为 right 或者 left
	Time          float64    `json:"time"`          // Time 音轨开始的时间 (通常是秒数)。
	TimeEnd       float64    `json:"timeEnd"`       // TimeEnd 音轨结束的时间 (通常是秒数)。
	TimeSignature []int      `json:"timeSignature"` // TimeSignature 音轨的拍号。
	Notes         []struct { // Notes 音轨上的所有音符信息。
		Note            int         `json:"note"`            // Note 音符的 MIDI 值。
		DurationTicks   int         `json:"durationTicks"`   // DurationTicks 音符的持续时间 (以 ticks 为单位)。
		NoteOffVelocity int         `json:"noteOffVelocity"` // NoteOffVelocity 音符结束时的力度。
		TicksStart      int         `json:"ticksStart"`      // TicksStart 音符开始的 tick 位置。
		Velocity        float64     `json:"velocity"`        // Velocity 音符的力度。
		MeasureBars     float64     `json:"measureBars"`     // MeasureBars 音符在小节中的位置 (以小节为单位)。
		Duration        float64     `json:"duration"`        // Duration 音符的持续时间 (通常是秒数)。
		NoteName        string      `json:"noteName"`        // NoteName 音符的名称 (例如 "C", "D#", "Bb")。
		Octave          int         `json:"octave"`          // Octave 音符的八度音阶。
		NotePitch       string      `json:"notePitch"`       // NotePitch 音符的音高 (例如 "C4", "G#5")。
		Start           float64     `json:"start"`           // Start 音符开始的时间 (通常是秒数)。
		End             float64     `json:"end"`             // End 音符结束的时间 (通常是秒数)。
		NoteLengthType  string      `json:"noteLengthType"`  // NoteLengthType 音符的长度类型 (例如 "quarter", "eighth")。
		Group           int         `json:"group"`           // Group 音符所属的组 (用于连音等)。
		MeasureInd      int         `json:"measureInd"`      // MeasureInd 音符所在的小节的索引。
		NoteMeasureInd  int         `json:"noteMeasureInd"`  // NoteMeasureInd 音符在小节中的索引。
		Id              string      `json:"id"`              // Id 音符的唯一标识符。
		Finger          int         `json:"finger"`          // Finger  指法建议
		Smp             interface{} `json:"smp"`             // Smp  未知，可能和音色采样相关
	} `json:"notes"`
	Max               int        `json:"max"`               // Max 音轨中最高音符的 MIDI 值。
	Min               int        `json:"min"`               // Min 音轨中最低音符的 MIDI 值。
	MeasureTicksStart float64    `json:"measureTicksStart"` // MeasureTicksStart 音轨开始的小节的 ticks 位置。
	MeasureTicksEnd   float64    `json:"measureTicksEnd"`   // MeasureTicksEnd 音轨结束的小节的 ticks 位置。
	Rests             []struct { // Rests 音轨上的休止符信息。
		Time           float64 `json:"time"`           // Time 休止符开始的时间 (通常是秒数)。
		NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 休止符的长度类型 (例如 "quarter", "eighth")。
	} `json:"rests"`
	Groups []struct { // Groups 音轨上的音符分组信息。
		Start          float64 `json:"start"`          // Start 组开始的时间 (通常是秒数)。
		End            float64 `json:"end"`            // End 组结束的时间 (通常是秒数)。
		NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 组的长度类型 (例如 "quarter", "eighth")。
		NoteInds       []int   `json:"noteInds"`       // NoteInds 组中的音符索引。
		GroupId        int     `json:"groupId"`        // GroupId 组的唯一标识符。
		BarY           int     `json:"barY"`           // BarY  在五线谱上的位置
	} `json:"groups"`
}

// clean 	clear(m.Notes) clear(m.Rests) clear(m.Groups)
func (m *MusicSection) clean() {
	m.Notes = m.Notes[:0]
	m.Rests = m.Rests[:0]
	m.Groups = m.Groups[:0]
}

func (m *MusicSection) guessStemDirection() {
	// 如果已经提供了方向，直接返回
	if m.Direction != "" && m.Direction != "auto" {
		return
	}
	//首先，按组处理
	if len(m.Groups) > 0 {
		groupDirections := make(map[int]string) //groupId -> direction
		for _, group := range m.Groups {
			groupNotePitches := []int{}
			for _, noteIndex := range group.NoteInds {
				if noteIndex < len(m.Notes) {
					groupNotePitches = append(groupNotePitches, m.Notes[noteIndex].Note)
				}
			}
			groupDirections[group.GroupId] = guessDirectionByPitches(groupNotePitches)

		}
		//统计group的方向
		directionCounts := make(map[string]int)
		for _, dir := range groupDirections {
			directionCounts[dir]++
		}

		//找到最常见的group方向
		maxCount := 0
		dominantDirection := ""
		for dir, count := range directionCounts {
			if count > maxCount {
				maxCount = count
				dominantDirection = dir
			}
		}
		m.Direction = dominantDirection

	}

	// 如果没有分组信息, 根据音高来
	notePitches := []int{}
	for _, note := range m.Notes {
		notePitches = append(notePitches, note.Note)
	}
	m.Direction = guessDirectionByPitches(notePitches)

}
func guessDirectionByPitches(pitches []int) string {
	if len(pitches) == 0 {
		return "unknown"
	}

	// 计算平均音高
	sum := 0
	for _, pitch := range pitches {
		sum += pitch
	}
	avgPitch := float64(sum) / float64(len(pitches))

	// B4 (MIDI note number 71) 作为中间线
	if avgPitch > 71 {
		return "down" // 符干朝下
	} else {
		return "up" // 符干朝上
	}

}
