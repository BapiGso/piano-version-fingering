package main

import (
	"bytes"
	"fmt"
	"github.com/saintfish/chardet"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese" // 例如，对于 Shift-JIS
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"log"
	"os"
)

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

func main() { // we write a SMF file into a buffer and read it back

	//var (
	//	bf                   bytes.Buffer
	//	clock                = smf.MetricTicks(96) // resolution: 96 ticks per quarternote 960 is also a common choice
	//	general, piano, bass smf.Track             // our tracks
	//)
	//
	//// first track must have tempo and meter informations
	//general.Add(0, smf.MetaTrackSequenceName("general"))
	//general.Add(0, smf.MetaMeter(3, 4))
	//general.Add(0, smf.MetaTempo(140))
	//general.Add(clock.Ticks4th()*6, smf.MetaTempo(130))
	//general.Add(clock.Ticks4th(), smf.MetaTempo(135))
	//general.Close(0) // don't forget to close a track
	//
	//piano.Add(0, smf.MetaInstrument("Piano"))
	//piano.Add(0, midi.ProgramChange(0, gm.Instr_HonkytonkPiano.Value()))
	//piano.Add(0, midi.NoteOn(0, 76, 120))
	//// duration: a quarter note (96 ticks in our case)
	//piano.Add(clock.Ticks4th(), midi.NoteOff(0, 76))
	//piano.Close(0)
	//
	//bass.Add(0, smf.MetaInstrument("Bass"))
	//bass.Add(0, midi.ProgramChange(1, gm.Instr_AcousticBass.Value()))
	//bass.Add(clock.Ticks4th(), midi.NoteOn(1, 47, 64))
	//bass.Add(clock.Ticks4th()*3, midi.NoteOff(1, 47))
	//bass.Close(0)
	//
	//// create the SMF and add the tracks
	s := smf.New()
	//s.TimeFormat = clock
	//s.Add(general)
	//s.Add(piano)
	//s.Add(bass)
	//
	//// write the bytes to the buffer
	//_, err := s.WriteTo(&bf)
	file, err := os.ReadFile("Mitsuha's Theme.mid")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		return
	}

	// read the bytes
	s, err = smf.ReadFrom(bytes.NewReader(file))

	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		return
	}

	fmt.Printf("got %v tracks\n", len(s.Tracks))

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

			case msg.GetMetaTrackName(&trackname):
				trackname = decodeTrackName(trackname) // 自动检测并解码 解码音轨名称 (假设为 Shift-JIS)
			case msg.GetMetaInstrument(&trackname):
				trackname = decodeTrackName(trackname) // 自动检测并解码 同样解码乐器名称
			case msg.GetProgramChange(&channel, &program):
				gm_name = "(" + gm.Instr(program).String() + ")"
			default:
				fmt.Printf("track %v %s %s @%v %s\n", no, trackname, gm_name, absTicks, ev.Message)
			}
		}
	}
}

// decodeTrackName 尝试自动检测并解码音轨名称
func decodeTrackName(trackname string) string {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest([]byte(trackname))
	if err != nil {
		// 检测失败，返回原始字符串
		return trackname
	}

	fmt.Printf("Detected encoding: %s (Confidence: %f)\n", result.Charset, result.Confidence)

	var enc encoding.Encoding
	switch result.Charset {
	case "UTF-8":
		// 已经是 UTF-8，无需转换
		return trackname
	case "Shift_JIS":
		enc = japanese.ShiftJIS
	case "GB-18030": //更广范围的gbk
		enc = simplifiedchinese.GB18030
	case "GBK": //或者 GBK
		enc = simplifiedchinese.GBK
	// 添加其他可能需要处理的编码...
	default:
		// 未知或不支持的编码，返回原始字符串
		return trackname
	}

	decoder := enc.NewDecoder()
	decoded, _, err := transform.String(decoder, trackname)
	if err != nil {
		return trackname // 解码失败
	}
	return decoded
}
