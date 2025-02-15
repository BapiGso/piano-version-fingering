package core

// ParseEnd 做一些收尾的工作，填充original字段
func (p *PVF) ParseEnd() {
	p.Original = struct {
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
	}{}
}
