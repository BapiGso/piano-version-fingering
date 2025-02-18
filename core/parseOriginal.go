package core

// parseOriginal 做一些收尾的工作，填充original字段
func (p *PVF) parseOriginal() {
	//todo
	p.Original.Header.Name = ""
	p.Original.Header.Ppq = p.Resolution
}
