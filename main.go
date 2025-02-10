package piano_version_fingering

type T struct {
	SupportingTracks []struct {
		Notes []struct {
			Midi     int     `json:"midi"`
			Time     float64 `json:"time"`
			Velocity float64 `json:"velocity"`
			Duration float64 `json:"duration"`
		} `json:"notes"`
		MyInstrument    int `json:"myInstrument"`
		TheirInstrument int `json:"theirInstrument"`
	} `json:"supportingTracks"`
	StartTime  int `json:"start_time"`
	SongLength int `json:"song_length"`
	Resolution int `json:"resolution"`
	Tempos     []struct {
		Bpm   int `json:"bpm"`
		Ticks int `json:"ticks"`
		Time  int `json:"time"`
	} `json:"tempos"`
	KeySignatures []struct {
		Key   string `json:"key"`
		Scale string `json:"scale"`
		Ticks int    `json:"ticks"`
	} `json:"keySignatures"`
	TimeSignatures []struct {
		Ticks         int   `json:"ticks"`
		TimeSignature []int `json:"timeSignature"`
		Measures      int   `json:"measures"`
	} `json:"timeSignatures"`
	Measures []struct {
		Time            float64 `json:"time"`
		TimeSignature   []int   `json:"timeSignature"`
		TicksPerMeasure int     `json:"ticksPerMeasure"`
		TicksStart      float64 `json:"ticksStart"`
		TotalTicks      float64 `json:"totalTicks"`
		Type            int     `json:"type"`
	} `json:"measures"`
	TracksV2 struct {
		Right []struct {
			Direction     string  `json:"direction"`
			Time          float64 `json:"time"`
			TimeEnd       float64 `json:"timeEnd"`
			TimeSignature []int   `json:"timeSignature"`
			Notes         []struct {
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
			} `json:"notes"`
			Max               int     `json:"max"`
			Min               int     `json:"min"`
			MeasureTicksStart float64 `json:"measureTicksStart"`
			MeasureTicksEnd   float64 `json:"measureTicksEnd"`
			Rests             []struct {
				Time           float64 `json:"time"`
				NoteLengthType string  `json:"noteLengthType"`
			} `json:"rests"`
			Groups []struct {
				Start          float64 `json:"start"`
				End            float64 `json:"end"`
				NoteLengthType string  `json:"noteLengthType"`
				NoteInds       []int   `json:"noteInds"`
				GroupId        int     `json:"groupId"`
				BarY           int     `json:"barY"`
			} `json:"groups"`
		} `json:"right"`
		Left []struct {
			Direction     string  `json:"direction"`
			Time          float64 `json:"time"`
			TimeEnd       float64 `json:"timeEnd"`
			TimeSignature []int   `json:"timeSignature"`
			Notes         []struct {
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
			} `json:"notes"`
			Max               int     `json:"max"`
			Min               int     `json:"min"`
			MeasureTicksStart float64 `json:"measureTicksStart"`
			MeasureTicksEnd   float64 `json:"measureTicksEnd"`
			Rests             []struct {
				Time           float64 `json:"time"`
				NoteLengthType string  `json:"noteLengthType"`
			} `json:"rests"`
			Groups []struct {
				Start          int     `json:"start"`
				End            float64 `json:"end"`
				NoteLengthType string  `json:"noteLengthType"`
				NoteInds       []int   `json:"noteInds"`
				GroupId        int     `json:"groupId"`
				BarY           int     `json:"barY"`
			} `json:"groups"`
		} `json:"left"`
	} `json:"tracksV2"`
	Original struct {
		Header struct {
			KeySignatures []struct {
				Key   string `json:"key"`
				Scale string `json:"scale"`
				Ticks int    `json:"ticks"`
			} `json:"keySignatures"`
			Meta []struct {
				Text  string `json:"text"`
				Ticks int    `json:"ticks"`
				Type  string `json:"type"`
			} `json:"meta"`
			Name   string `json:"name"`
			Ppq    int    `json:"ppq"`
			Tempos []struct {
				Bpm   int `json:"bpm"`
				Ticks int `json:"ticks"`
			} `json:"tempos"`
			TimeSignatures []struct {
				Ticks         int   `json:"ticks"`
				TimeSignature []int `json:"timeSignature"`
				Measures      int   `json:"measures"`
			} `json:"timeSignatures"`
		} `json:"header"`
		Tracks []struct {
			Channel        int `json:"channel"`
			ControlChanges struct {
				Field1 []struct {
					Number int `json:"number"`
					Ticks  int `json:"ticks"`
					Time   int `json:"time"`
					Value  int `json:"value"`
				} `json:"0"`
				Field2 []struct {
					Number int     `json:"number"`
					Ticks  int     `json:"ticks"`
					Time   int     `json:"time"`
					Value  float64 `json:"value"`
				} `json:"10"`
				Field3 []struct {
					Number int `json:"number"`
					Ticks  int `json:"ticks"`
					Time   int `json:"time"`
					Value  int `json:"value"`
				} `json:"32"`
				Field4 []struct {
					Number int     `json:"number"`
					Ticks  int     `json:"ticks"`
					Time   int     `json:"time"`
					Value  float64 `json:"value"`
				} `json:"91"`
			} `json:"controlChanges"`
			PitchBends []interface{} `json:"pitchBends"`
			Instrument struct {
				Family string `json:"family"`
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"instrument"`
			Name  string `json:"name"`
			Notes []struct {
				Duration      float64 `json:"duration"`
				DurationTicks int     `json:"durationTicks"`
				Midi          int     `json:"midi"`
				Name          string  `json:"name"`
				Ticks         int     `json:"ticks"`
				Time          float64 `json:"time"`
				Velocity      float64 `json:"velocity"`
			} `json:"notes"`
			EndOfTrackTicks int `json:"endOfTrackTicks"`
		} `json:"tracks"`
	} `json:"original"`
	AccompanyingInstruments []int         `json:"accompanyingInstruments"`
	AccompanyingChannels    []int         `json:"accompanyingChannels"`
	AccompanyingTracks      []interface{} `json:"accompanyingTracks"`
	Sections                []interface{} `json:"sections"`
	PositionGroups          []interface{} `json:"positionGroups"`
	TechnicalGroups         []interface{} `json:"technicalGroups"`
	MaxSimplification       int           `json:"maxSimplification"`
	Name                    string        `json:"name"`
	Artist                  string        `json:"artist"`
}
