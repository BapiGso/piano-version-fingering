package core

import (
	"fmt"
	"math"
	"sort"
)

func (p *PVF) parseTrackV2() {
	//先把Measures补全
	loop := (p.SongLength - int(p.Measures[len(p.Measures)-1].Time)) / (p.Measures[len(p.Measures)-1].TimeSignature[0])
	for i := 0; i < loop; i++ {
		tmp := p.Measures[len(p.Measures)-1]
		tmp.TicksStart += tmp.TotalTicks
		tmp.Time += float64(tmp.TicksPerMeasure / p.Resolution)
		tmp.Type = 2
		p.Measures = append(p.Measures, tmp)
	}

	//正式开始，用两次for循环，外面的for循环Measures
	//里面的for循环Channel
	for k, v := range p.Measures {
		//先创建一个tmp
		tmp := MusicSection{}
		tmp.Time = v.Time
		tmp.TimeEnd = v.Time + float64(v.TimeSignature[0])
		tmp.TimeSignature = v.TimeSignature
		tmp.MeasureTicksStart = v.TicksStart
		tmp.MeasureTicksEnd = v.TicksStart + float64(v.TimeSignature[0]*p.Resolution)
		tmp.Max = 200 //todo
		tmp.Min = 0

		for k2, note := range p.SupportingTracks[0].Notes { //判断什么音符在这个小节里
			if tmp.Time <= note.Time && note.Time < tmp.TimeEnd {
				noteName, octave, notePitch := midiToNoteName(note.Midi)
				tmp.Notes = append(tmp.Notes, struct {
					Note            int         `json:"note"`
					DurationTicks   int         `json:"durationTicks"`
					NoteOffVelocity int         `json:"noteOffVelocity"`
					TicksStart      int         `json:"ticksStart"`
					Velocity        float64     `json:"velocity"`
					MeasureBars     float64     `json:"measureBars"`
					Duration        float64     `json:"duration"`
					NoteName        string      `json:"noteName"`
					Octave          int         `json:"octave"`
					NotePitch       string      `json:"notePitch"`
					Start           float64     `json:"start"`
					End             float64     `json:"end"`
					NoteLengthType  string      `json:"noteLengthType"`
					Group           int         `json:"group"`
					MeasureInd      int         `json:"measureInd"`
					NoteMeasureInd  int         `json:"noteMeasureInd"`
					Id              string      `json:"id"`
					Finger          int         `json:"finger"`
					Smp             interface{} `json:"smp"`
				}{
					Note:            note.Midi,
					DurationTicks:   int(note.Duration * float64(p.Resolution)),
					NoteOffVelocity: 0, //todo unknown,看别的也是0
					TicksStart:      int(note.Time * float64(p.Resolution)),
					Velocity:        note.Velocity,
					MeasureBars:     note.Time / float64(v.TimeSignature[0]),
					Duration:        note.Duration,
					NoteName:        noteName,
					Octave:          octave,
					NotePitch:       notePitch,
					Start:           note.Time,
					End:             note.Time + note.Duration,
					NoteLengthType:  "", //todo原json解析为dottedthirtysecond
					Group:           0,  //todo不知道什么意思有0，-1
					MeasureInd:      k,
					NoteMeasureInd:  len(tmp.Notes),
					Id:              fmt.Sprintf("r%v", k2),
					Finger:          0,
					Smp:             nil,
				})
			}
		}
		tmp.addRestsAndGroups(p.Resolution)
		tmp.guessStemDirection()
		p.TracksV2.Right = append(p.TracksV2.Right, tmp)
		tmp.clean()
		for k2, note := range p.SupportingTracks[1].Notes { //判断什么音符在这个小节里
			if tmp.Time <= note.Time && note.Time < tmp.TimeEnd {
				noteName, octave, notePitch := midiToNoteName(note.Midi)
				tmp.Notes = append(tmp.Notes, struct {
					Note            int         `json:"note"`
					DurationTicks   int         `json:"durationTicks"`
					NoteOffVelocity int         `json:"noteOffVelocity"`
					TicksStart      int         `json:"ticksStart"`
					Velocity        float64     `json:"velocity"`
					MeasureBars     float64     `json:"measureBars"`
					Duration        float64     `json:"duration"`
					NoteName        string      `json:"noteName"`
					Octave          int         `json:"octave"`
					NotePitch       string      `json:"notePitch"`
					Start           float64     `json:"start"`
					End             float64     `json:"end"`
					NoteLengthType  string      `json:"noteLengthType"`
					Group           int         `json:"group"`
					MeasureInd      int         `json:"measureInd"`
					NoteMeasureInd  int         `json:"noteMeasureInd"`
					Id              string      `json:"id"`
					Finger          int         `json:"finger"`
					Smp             interface{} `json:"smp"`
				}{
					Note:            note.Midi,
					DurationTicks:   int(note.Duration * float64(p.Resolution)),
					NoteOffVelocity: 0, //todo unknown,看别的也是0
					TicksStart:      int(note.Time * float64(p.Resolution)),
					Velocity:        note.Velocity,
					MeasureBars:     note.Time / float64(v.TimeSignature[0]),
					Duration:        note.Duration,
					NoteName:        noteName,
					Octave:          octave,
					NotePitch:       notePitch,
					Start:           note.Time,
					End:             note.Time + note.Duration,
					NoteLengthType:  "", //todo原json解析为dottedthirtysecond
					Group:           0,  //todo不知道什么意思有0，-1
					MeasureInd:      k,
					NoteMeasureInd:  len(tmp.Notes),
					Id:              fmt.Sprintf("r%v", k2),
					Finger:          0,
					Smp:             nil,
				})
			}
		}
		tmp.addRestsAndGroups(p.Resolution)
		tmp.guessStemDirection()
		p.TracksV2.Left = append(p.TracksV2.Left, tmp)
	}
}

func midiToNoteName(midi int) (noteName string, octave int, notePitch string) {
	// MIDI 范围检查（0-127）
	if midi < 0 || midi > 127 {
		return "", 0, ""
	}

	// 音高名称数组（C, C#, D, D#, E, F, F#, G, G#, A, A#, B）
	noteNames := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	// 计算音高和八度（MIDI 21 对应 A0，中央 C 是 MIDI 60 → C4）
	noteIndex := (midi - 12) % 12 // 调整基准为 C（例如 72-12=60 → C4）
	octave = (midi - 12) / 12     // 计算八度（60 → 4）
	notePitch = noteNames[noteIndex]
	noteName = fmt.Sprintf("%s%d", notePitch, octave)
	return
}

func (tmp *MusicSection) addRestsAndGroups(resolution int) {
	// 处理 rests
	sort.Slice(tmp.Notes, func(i, j int) bool {
		return tmp.Notes[i].Start < tmp.Notes[j].Start
	})

	var rests []struct {
		Time           float64 `json:"time"`           // Time 休止符开始的时间 (通常是秒数)。
		NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 休止符的长度类型 (例如 "quarter", "eighth")。
	}
	prevEnd := tmp.Time
	for _, note := range tmp.Notes {
		if note.Start > prevEnd {
			gap := note.Start - prevEnd
			restType := getNoteLengthType(gap, resolution)
			rests = append(rests, struct {
				Time           float64 `json:"time"`           // Time 休止符开始的时间 (通常是秒数)。
				NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 休止符的长度类型 (例如 "quarter", "eighth")。
			}{
				Time:           prevEnd,
				NoteLengthType: restType,
			})
		}
		if note.End > prevEnd {
			prevEnd = note.End
		}
	}
	if prevEnd < tmp.TimeEnd {
		gap := tmp.TimeEnd - prevEnd
		restType := getNoteLengthType(gap, resolution)
		rests = append(rests, struct {
			Time           float64 `json:"time"`           // Time 休止符开始的时间 (通常是秒数)。
			NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 休止符的长度类型 (例如 "quarter", "eighth")。
		}{
			Time:           prevEnd,
			NoteLengthType: restType,
		})
	}
	tmp.Rests = rests

	// 处理 groups
	if len(tmp.Notes) == 0 {
		tmp.Groups = nil
		return
	}
	var groups []struct { // Groups 音轨上的音符分组信息。
		Start          float64 `json:"start"`          // Start 组开始的时间 (通常是秒数)。
		End            float64 `json:"end"`            // End 组结束的时间 (通常是秒数)。
		NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 组的长度类型 (例如 "quarter", "eighth")。
		NoteInds       []int   `json:"noteInds"`       // NoteInds 组中的音符索引。
		GroupId        int     `json:"groupId"`        // GroupId 组的唯一标识符。
		BarY           int     `json:"barY"`           // BarY  在五线谱上的位置
	}
	currentGroup := &struct { // Groups 音轨上的音符分组信息。
		Start          float64 `json:"start"`          // Start 组开始的时间 (通常是秒数)。
		End            float64 `json:"end"`            // End 组结束的时间 (通常是秒数)。
		NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 组的长度类型 (例如 "quarter", "eighth")。
		NoteInds       []int   `json:"noteInds"`       // NoteInds 组中的音符索引。
		GroupId        int     `json:"groupId"`        // GroupId 组的唯一标识符。
		BarY           int     `json:"barY"`           // BarY  在五线谱上的位置
	}{
		Start:          tmp.Notes[0].Start,
		End:            tmp.Notes[0].End,
		NoteLengthType: tmp.Notes[0].NoteLengthType,
		NoteInds:       []int{0},
		GroupId:        0,
	}
	for i := 1; i < len(tmp.Notes); i++ {
		note := tmp.Notes[i]
		if note.NoteLengthType == currentGroup.NoteLengthType && note.Start <= currentGroup.End+0.001 {
			currentGroup.End = math.Max(currentGroup.End, note.End)
			currentGroup.NoteInds = append(currentGroup.NoteInds, i)
		} else {
			groups = append(groups, *currentGroup)
			currentGroup = &struct { // Groups 音轨上的音符分组信息。
				Start          float64 `json:"start"`          // Start 组开始的时间 (通常是秒数)。
				End            float64 `json:"end"`            // End 组结束的时间 (通常是秒数)。
				NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 组的长度类型 (例如 "quarter", "eighth")。
				NoteInds       []int   `json:"noteInds"`       // NoteInds 组中的音符索引。
				GroupId        int     `json:"groupId"`        // GroupId 组的唯一标识符。
				BarY           int     `json:"barY"`           // BarY  在五线谱上的位置
			}{
				Start:          note.Start,
				End:            note.End,
				NoteLengthType: note.NoteLengthType,
				NoteInds:       []int{i},
				GroupId:        len(groups),
			}
		}
	}
	groups = append(groups, *currentGroup)

	// 设置 barY 为组内第一个音符的音高
	for i := range groups {
		if len(groups[i].NoteInds) > 0 {
			groups[i].BarY = tmp.Notes[groups[i].NoteInds[0]].Note
		}
	}
	tmp.Groups = groups
}

func getNoteLengthType(gapInBeats float64, resolution int) string {
	gapInTicks := gapInBeats * float64(resolution)
	typeThresholds := []struct {
		name  string
		ticks float64
	}{
		{"dottedhalf", 3.0 * float64(resolution)},
		{"dottedquarter", 1.5 * float64(resolution)},
		{"dottedeighth", 0.75 * float64(resolution)},
		{"dottedsixteenth", 0.375 * float64(resolution)},
		{"dottedthirtysecond", 0.1875 * float64(resolution)},
	}

	closestType := ""
	minDiff := math.MaxFloat64
	for _, t := range typeThresholds {
		diff := math.Abs(gapInTicks - t.ticks)
		if diff < minDiff {
			minDiff = diff
			closestType = t.name
		}
	}
	return closestType
}
