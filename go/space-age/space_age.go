package space

type Planet string

const (
	earthYearInSec = 31557600
)

var planetsYearSec = map[Planet]float64{
	"Earth":   earthYearInSec,
	"Mercury": earthYearInSec * 0.2408467,
	"Venus":   earthYearInSec * 0.61519726,
	"Mars":    earthYearInSec * 1.8808158,
	"Jupiter": earthYearInSec * 11.862615,
	"Saturn":  earthYearInSec * 29.447498,
	"Uranus":  earthYearInSec * 84.016846,
	"Neptune": earthYearInSec * 164.79132,
}

func Age(sec float64, planet Planet) float64 {
	return sec / planetsYearSec[planet]
}
