package core

func New() *PVF {
	a := new(PVF)

	return a
}

type PVF struct {
	SupportingTracks [2]struct { // SupportingTracks 代表伴奏音轨，通常包含一些辅助演奏的乐器。
		Notes []struct { // Notes 包含了该音轨上的所有音符信息。
			Midi     int     `json:"midi"`     // Midi 音符的 MIDI 值 (0-127)。
			Time     float64 `json:"time"`     // Time 音符开始的时间 (通常是相对于歌曲开始的秒数)。
			Velocity float64 `json:"velocity"` // Velocity 音符的力度 (0-1)。
			Duration float64 `json:"duration"` // Duration 音符的持续时间 (通常是秒数)。
		} `json:"notes"`
		MyInstrument    int `json:"myInstrument"`    // MyInstrument  当前用户所演奏的乐器编号
		TheirInstrument int `json:"theirInstrument"` // TheirInstrument  伴奏音轨中乐器编号
	} `json:"supportingTracks"`
	StartTime  int        `json:"start_time"`  // StartTime 歌曲开始的时间 (通常是 0)。
	SongLength int        `json:"song_length"` // SongLength 歌曲的总长度 (通常是毫秒或秒)。
	Resolution int        `json:"resolution"`  // Resolution 每四分音符的 ticks 数，用于将时间转换为音符长度。
	Tempos     []struct { // Tempos 歌曲中的速度变化。
		Bpm   int `json:"bpm"`   // Bpm 每分钟的拍数。
		Ticks int `json:"ticks"` // Ticks 此速度开始的 tick 位置。
		Time  int `json:"time"`  // Time 此速度开始的时间 (通常是秒数)。
	} `json:"tempos"`
	KeySignatures []struct { // KeySignatures 歌曲中的调号变化。
		Key   string `json:"key"`   // Key 调号 (例如 "C", "G", "Am")。
		Scale string `json:"scale"` // Scale 调式 (例如 "major", "minor")。
		Ticks int    `json:"ticks"` // Ticks 此调号开始的 tick 位置。
	} `json:"keySignatures"`
	TimeSignatures []struct { // TimeSignatures 歌曲中的拍号变化。
		Ticks         int   `json:"ticks"`         // Ticks 此拍号开始的 tick 位置。
		TimeSignature []int `json:"timeSignature"` // TimeSignature 拍号 (例如 [4, 4] 表示 4/4 拍)。
		Measures      int   `json:"measures"`      // Measures  在对应节拍下的总小节数
	} `json:"timeSignatures"`
	Measures []struct { // Measures 歌曲中各个小节的信息。
		Time            float64 `json:"time"`            // Time 小节开始的时间 (通常是秒数)。
		TimeSignature   []int   `json:"timeSignature"`   // TimeSignature 小节的拍号。
		TicksPerMeasure int     `json:"ticksPerMeasure"` // TicksPerMeasure 每个小节的 ticks 数。
		TicksStart      float64 `json:"ticksStart"`      // TicksStart 小节开始的 tick 位置。
		TotalTicks      float64 `json:"totalTicks"`      // TotalTicks 小节结束的 tick 位置。
		Type            int     `json:"type"`            // Type 小节的类型 (可能代表不同的标记或演奏指示)。
	} `json:"measures"`
	TracksV2 struct { // TracksV2 更详细的左右手音轨信息。
		Right []struct { // Right 右手音轨信息。
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
		} `json:"right"`
		Left []struct { // Left 左手音轨信息 (结构与 Right 相同)。
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
				Start          int     `json:"start"`          // Start 组开始的时间 (通常是秒数)。
				End            float64 `json:"end"`            // End 组结束的时间 (通常是秒数)。
				NoteLengthType string  `json:"noteLengthType"` // NoteLengthType 组的长度类型 (例如 "quarter", "eighth")。
				NoteInds       []int   `json:"noteInds"`       // NoteInds 组中的音符索引。
				GroupId        int     `json:"groupId"`        // GroupId 组的唯一标识符。
				BarY           int     `json:"barY"`           // BarY  在五线谱上的位置
			} `json:"groups"`
		} `json:"left"`
	} `json:"tracksV2"`
	Original struct { // Original 原始 MIDI 文件的数据。
		Header struct { // Header MIDI 文件头信息。
			KeySignatures []struct { // KeySignatures 歌曲中的调号变化 (原始 MIDI 文件中的数据)。
				Key   string `json:"key"`   // Key 调号。
				Scale string `json:"scale"` // Scale 调式。
				Ticks int    `json:"ticks"` // Ticks 此调号开始的 tick 位置。
			} `json:"keySignatures"`
			Meta []struct { // Meta  元数据，例如歌名，作者
				Text  string `json:"text"`  // Text 元数据文本。
				Ticks int    `json:"ticks"` // Ticks 元数据所在的 tick 位置。
				Type  string `json:"type"`  // Type 元数据的类型。
			} `json:"meta"`
			Name   string     `json:"name"` // Name MIDI 文件的名称。
			Ppq    int        `json:"ppq"`  // Ppq 每四分音符的 ticks 数。
			Tempos []struct { // Tempos 歌曲中的速度变化 (原始 MIDI 文件中的数据)。
				Bpm   int `json:"bpm"`   // Bpm 每分钟的拍数。
				Ticks int `json:"ticks"` // Ticks 此速度开始的 tick 位置。
			} `json:"tempos"`
			TimeSignatures []struct { // TimeSignatures 歌曲中的拍号变化 (原始 MIDI 文件中的数据)。
				Ticks         int   `json:"ticks"`         // Ticks 此拍号开始的 tick 位置。
				TimeSignature []int `json:"timeSignature"` // TimeSignature 拍号。
				Measures      int   `json:"measures"`      // Measures  在对应节拍下的总小节数
			} `json:"timeSignatures"`
		} `json:"header"`
		Tracks []struct { // Tracks MIDI 文件中的所有音轨。
			Channel        int      `json:"channel"` // Channel 音轨的 MIDI 通道。
			ControlChanges struct { // ControlChanges 音轨上的控制变化信息 (例如音量、混响)。
				Field1 []struct { // Field1 控制变化 #0 (Bank Select)。
					Number int `json:"number"` // Number 控制变化编号。
					Ticks  int `json:"ticks"`  // Ticks 控制变化所在的 tick 位置。
					Time   int `json:"time"`   // Time 控制变化发生的时间。
					Value  int `json:"value"`  // Value 控制变化的值。
				} `json:"0"`
				Field2 []struct { // Field2 控制变化 #10 (Pan)。
					Number int     `json:"number"` // Number 控制变化编号。
					Ticks  int     `json:"ticks"`  // Ticks 控制变化所在的 tick 位置。
					Time   int     `json:"time"`   // Time 控制变化发生的时间。
					Value  float64 `json:"value"`  // Value 控制变化的值。
				} `json:"10"`
				Field3 []struct { // Field3 控制变化 #32 (Bank Select LSB)。
					Number int `json:"number"` // Number 控制变化编号。
					Ticks  int `json:"ticks"`  // Ticks 控制变化所在的 tick 位置。
					Time   int `json:"time"`   // Time 控制变化发生的时间。
					Value  int `json:"value"`  // Value 控制变化的值。
				} `json:"32"`
				Field4 []struct { // Field4 控制变化 #91 (Reverb Send Level)。
					Number int     `json:"number"` // Number 控制变化编号。
					Ticks  int     `json:"ticks"`  // Ticks 控制变化所在的 tick 位置。
					Time   int     `json:"time"`   // Time 控制变化发生的时间。
					Value  float64 `json:"value"`  // Value 控制变化的值。
				} `json:"91"`
			} `json:"controlChanges"`
			PitchBends []interface{} `json:"pitchBends"` // PitchBends 音轨上的弯音信息。
			Instrument struct {      // Instrument 音轨使用的乐器信息。
				Family string `json:"family"` // Family 乐器家族 (例如 "Piano", "Guitar")。
				Number int    `json:"number"` // Number 乐器编号 (MIDI Program Change)。
				Name   string `json:"name"`   // Name 乐器名称。
			} `json:"instrument"`
			Name  string     `json:"name"` // Name 音轨的名称。
			Notes []struct { // Notes 音轨上的所有音符信息。
				Duration      float64 `json:"duration"`      // Duration 音符的持续时间 (通常是秒数)。
				DurationTicks int     `json:"durationTicks"` // DurationTicks 音符的持续时间 (以 ticks 为单位)。
				Midi          int     `json:"midi"`          // Midi 音符的 MIDI 值。
				Name          string  `json:"name"`          // Name 音符的名称 (例如 "C4", "G#5")。
				Ticks         int     `json:"ticks"`         // Ticks 音符开始的 tick 位置。
				Time          float64 `json:"time"`          // Time 音符开始的时间 (通常是秒数)。
				Velocity      float64 `json:"velocity"`      // Velocity 音符的力度。
			} `json:"notes"`
			EndOfTrackTicks int `json:"endOfTrackTicks"` // EndOfTrackTicks 音轨结束的 tick 位置。
		} `json:"tracks"`
	} `json:"original"`
	AccompanyingInstruments []int         `json:"accompanyingInstruments"` // AccompanyingInstruments 伴奏中使用的乐器编号列表。
	AccompanyingChannels    []int         `json:"accompanyingChannels"`    // AccompanyingChannels 伴奏中使用的 MIDI 通道列表。
	AccompanyingTracks      []interface{} `json:"accompanyingTracks"`      // AccompanyingTracks 伴奏音轨 (可能包含一些处理后的数据)。
	Sections                []interface{} `json:"sections"`                // Sections 歌曲的章节信息 (例如 "Verse", "Chorus")。
	PositionGroups          []interface{} `json:"positionGroups"`          // PositionGroups  位置分组，可能和显示相关
	TechnicalGroups         []interface{} `json:"technicalGroups"`         // TechnicalGroups  技术分组，可能和谱面难度相关
	MaxSimplification       int           `json:"maxSimplification"`       // MaxSimplification 最大简化程度 (可能用于简化谱面)。
	Name                    string        `json:"name"`                    // Name 歌曲的名称。
	Artist                  string        `json:"artist"`                  // Artist 歌曲的艺术家。
}
