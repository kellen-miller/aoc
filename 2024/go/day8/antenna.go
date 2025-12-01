package day8

import "github.com/kellen-miller/aoc/languages/go/pkg/grid"

type Antenna struct {
	Position grid.Point
	Value    rune
}

func (a *Antenna) Equals(other *Antenna) bool {
	return a.Position == other.Position && a.Value == other.Value
}

type UniqueAntennaPair struct {
	Antenna1 Antenna
	Antenna2 Antenna
}

func (ap *UniqueAntennaPair) Equals(other *UniqueAntennaPair) bool {
	return (ap.Antenna1.Position == other.Antenna1.Position && ap.Antenna2.Position == other.Antenna2.Position) ||
		(ap.Antenna1.Position == other.Antenna2.Position && ap.Antenna2.Position == other.Antenna1.Position)
}

func (ap *UniqueAntennaPair) XYDistance() (int, int) {
	return ap.Antenna2.Position.XYDistance(ap.Antenna1.Position)
}
