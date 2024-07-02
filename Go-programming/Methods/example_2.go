package Methods

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
