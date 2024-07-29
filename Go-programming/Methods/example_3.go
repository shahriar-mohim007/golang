package Methods

type Duration int64

func (d Duration) Hours() float64 {
	return float64(d) / float64(3600)
}
