package core

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
		println(k, v)
	}
}
