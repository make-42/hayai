package waves

import (
	"hayai/utils"
	"image/color"
	"sort"

	"gonum.org/v1/gonum/interp"
)

type TravelTime struct {
	Travel []float64 // deg
	Time   []float64 // minutes
}

type ShadowZone struct {
	Start float64 // deg
	End   float64 // deg
}

type Wave struct {
	ShadowZones          []ShadowZone
	TravelTime           TravelTime // TODO: calculate min delay to not calculate when impossible
	CanBeDouble          bool
	DoubleTravelIndex    int
	DoubleStart          bool
	AkimaSpline          interp.AkimaSpline
	InvertedAkimaSpline  interp.AkimaSpline
	AkimaSplineA         interp.AkimaSpline
	InvertedAkimaSplineA interp.AkimaSpline
	AkimaSplineB         interp.AkimaSpline
	InvertedAkimaSplineB interp.AkimaSpline
	Color                color.NRGBA
	Label                string
}

type Waves struct {
	Rayleigh Wave
	Love     Wave
	SSS      Wave
	SS       Wave
	S        Wave
	ScS      Wave
	SKS      Wave
	PcS      Wave
	PKP      Wave
	PcP      Wave
	PPP      Wave
	PP       Wave
	P        Wave
}

// Using this https://www.geo.arizona.edu/saso/Earthquakes/Recent/ttc.html
var WaveTable = Waves{
	Rayleigh: Wave{
		ShadowZones: []ShadowZone{{0, 7.}},
		TravelTime: TravelTime{
			Travel: []float64{7, 10, 20, 40, 60, 80, 100, 180},
			Time:   []float64{3.5, 5, 10, 19.8, 27, 38, 47.5, 85.5},
		},
		Color: color.NRGBA{0xFF, 0xAF, 0x87, 192},
		Label: "Rayleigh",
	},
	Love: Wave{
		ShadowZones: []ShadowZone{{0, 7.}},
		TravelTime: TravelTime{
			Travel: []float64{7, 10, 20, 40, 60, 80, 100, 120, 180},
			Time:   []float64{3.5, 5, 9.8, 17.8, 26, 34, 42.5, 51, 76.5},
		},
		Color: color.NRGBA{0xFF, 0x9F, 0x7D, 192},
		Label: "Love",
	},
	SSS: Wave{
		ShadowZones: []ShadowZone{{0, 48.5}},
		TravelTime: TravelTime{
			Travel: []float64{48.5, 60, 80, 100, 120, 140, 160, 180},
			Time:   []float64{20.2, 24.7, 31, 36.4, 41.2, 46, 50, 53.7},
		},
		Color: color.NRGBA{0xFF, 0x8E, 0x72, 192},
		Label: "SSS",
	},
	SS: Wave{
		ShadowZones: []ShadowZone{{0, 15.}},
		TravelTime: TravelTime{
			Travel: []float64{15, 20, 40, 60, 80, 100, 120, 140, 160, 180},
			Time:   []float64{6.1, 9, 16, 22.4, 27.5, 32.5, 37.2, 41.2, 45, 48.2},
		},
		Color: color.NRGBA{0xF6, 0x7C, 0x68, 192},
		Label: "SS",
	},
	S: Wave{
		ShadowZones: []ShadowZone{{105., 180.}},
		TravelTime: TravelTime{
			Travel: []float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 105},
			Time:   []float64{0, 4.9, 8, 11.5, 14, 16.5, 18.2, 21, 22.3, 24, 25.8},
		},
		Color: color.NRGBA{0xED, 0x6A, 0x5E, 192},
		Label: "S",
	},
	ScS: Wave{
		ShadowZones: []ShadowZone{{78.5, 180.}},
		TravelTime: TravelTime{
			Travel: []float64{0, 5, 10, 20, 30, 40, 50, 60, 78.5},
			Time:   []float64{15.8, 15.7, 15.8, 16.2, 16.8, 17.8, 18, 20, 22.4},
		},
		CanBeDouble:       true,
		DoubleTravelIndex: 1,
		DoubleStart:       true,
		Color:             color.NRGBA{0xD9, 0x79, 0x69, 192},
		Label:             "ScS",
	},
	SKS: Wave{
		ShadowZones: []ShadowZone{{0., 78.5}},
		TravelTime: TravelTime{
			Travel: []float64{78.5, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180},
			Time:   []float64{22.4, 23.8, 24.8, 25.2, 26, 26.5, 27, 27.3, 27.4, 27.5, 27.4},
		},
		CanBeDouble:       true,
		DoubleTravelIndex: 9,
		DoubleStart:       false,
		Color:             color.NRGBA{0xC5, 0x88, 0x74, 192},
		Label:             "SKS",
	},
	PcS: Wave{
		ShadowZones: []ShadowZone{{0, 1}, {69., 180.}},
		TravelTime: TravelTime{
			Travel: []float64{1, 5, 10, 20, 30, 40, 50, 60, 69},
			Time:   []float64{12.3, 12.2, 12.3, 12.5, 13, 13.5, 13.7, 15, 15.8},
		},
		CanBeDouble:       true,
		DoubleTravelIndex: 1,
		DoubleStart:       true,
		Color:             color.NRGBA{0x9D, 0xA5, 0x89, 192},
		Label:             "PcS",
	},
	PKP: Wave{
		ShadowZones: []ShadowZone{{0., 109.}},
		TravelTime: TravelTime{
			Travel: []float64{109., 120, 130, 140, 150, 160, 170, 180},
			Time:   []float64{18.5, 19, 19.5, 19.8, 19.9, 20, 20.1, 20.2},
		},
		Color: color.NRGBA{0x75, 0xC3, 0x9E, 192},
		Label: "PKP",
	}, PcP: Wave{
		ShadowZones: []ShadowZone{{0, 1}, {151., 180.}},
		TravelTime: TravelTime{
			Travel: []float64{1, 10, 20, 30, 40, 50, 60, 70, 80, 151},
			Time:   []float64{8.5, 8.9, 9, 9.4, 9.9, 10.3, 11, 11.8, 12.4, 17.6},
		},
		Color: color.NRGBA{0x4C, 0xE0, 0xB3, 192},
		Label: "PcP",
	}, PPP: Wave{
		ShadowZones: []ShadowZone{},
		TravelTime: TravelTime{
			Travel: []float64{0, 20, 40, 60, 80, 100, 120, 140, 160, 180},
			Time:   []float64{0, 5, 9.9, 14, 17.4, 20.2, 22.8, 25.8, 28.7, 30.5},
		},
		Color: color.NRGBA{0x47, 0xC6, 0xA3, 192},
		Label: "PPP",
	}, PP: Wave{
		ShadowZones: []ShadowZone{},
		TravelTime: TravelTime{
			Travel: []float64{0, 20, 40, 60, 80, 100, 120, 140, 160, 180},
			Time:   []float64{0, 5, 9.2, 13, 15.7, 18.5, 20.5, 23, 24.8, 26},
		},
		Color: color.NRGBA{0x42, 0xAC, 0x92, 192},
		Label: "PP",
	}, P: Wave{
		ShadowZones: []ShadowZone{{151., 180.}},
		TravelTime: TravelTime{
			Travel: []float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 151},
			Time:   []float64{0, 2.4, 4.2, 6.2, 7.5, 9, 10.1, 11.5, 12.35, 17.6},
		},
		Color: color.NRGBA{0x37, 0x77, 0x71, 192},
		Label: "P",
	},
}

var WaveList = []*Wave{&WaveTable.Rayleigh, &WaveTable.Love, &WaveTable.SSS, &WaveTable.SS, &WaveTable.S, &WaveTable.ScS, &WaveTable.SKS, &WaveTable.PcS, &WaveTable.PKP, &WaveTable.PcP, &WaveTable.PPP, &WaveTable.PP, &WaveTable.P}

func Init() {
	for _, wave := range WaveList {
		err := wave.AkimaSpline.Fit(wave.TravelTime.Travel, wave.TravelTime.Time)
		utils.CheckError(err)
		if !wave.CanBeDouble {
			err = wave.InvertedAkimaSpline.Fit(wave.TravelTime.Time, wave.TravelTime.Travel)
			utils.CheckError(err)
		} else {
			travelA := make([]float64, wave.DoubleTravelIndex+1)
			for i := range travelA {
				travelA[i] = wave.TravelTime.Travel[i]
			}

			timeA := make([]float64, wave.DoubleTravelIndex+1)
			for i := range timeA {
				timeA[i] = wave.TravelTime.Time[i]
			}

			travelB := make([]float64, len(wave.TravelTime.Travel)-wave.DoubleTravelIndex)
			for i := range travelB {
				travelB[i] = wave.TravelTime.Travel[i+wave.DoubleTravelIndex]
			}

			timeB := make([]float64, len(wave.TravelTime.Time)-wave.DoubleTravelIndex)
			for i := range timeB {
				timeB[i] = wave.TravelTime.Time[i+wave.DoubleTravelIndex]
			}

			sort.Float64s(travelA)
			sort.Float64s(timeA)
			sort.Float64s(travelB)
			sort.Float64s(timeB)

			err := wave.AkimaSplineA.Fit(travelA, timeA)
			utils.CheckError(err)
			err = wave.InvertedAkimaSplineA.Fit(timeA, travelA)
			utils.CheckError(err)

			err = wave.AkimaSplineB.Fit(travelB, timeB)
			utils.CheckError(err)
			err = wave.InvertedAkimaSplineB.Fit(timeB, travelB)
			utils.CheckError(err)
		}
	}
}

func TimeToTravel(wave Wave, time float64) (float64, float64, float64, bool, bool) {
	if time < 0 {
		return 0, 0, 0, false, false
	}
	travel := 0.
	travelA := 0.
	travelB := 0.
	exists := true
	double := false

	for _, shadowZone := range wave.ShadowZones {
		shadowZoneTimeStart := wave.AkimaSpline.Predict(shadowZone.Start)
		shadowZoneTimeEnd := wave.AkimaSpline.Predict(shadowZone.End)
		if shadowZone.Start == 0 {
			shadowZoneTimeStart = 0
		}
		if shadowZone.End == 180 {
			shadowZoneTimeEnd = 10000000000
		}
		if time >= shadowZoneTimeStart && time <= shadowZoneTimeEnd {
			exists = false
		}
	}

	if !wave.CanBeDouble {
		travel = wave.InvertedAkimaSpline.Predict(time)
	} else {
		if wave.DoubleStart {
			maxDoubleTime := wave.TravelTime.Time[0]
			if time < maxDoubleTime {
				travelA = wave.InvertedAkimaSplineA.Predict(time)
				travelB = wave.InvertedAkimaSplineB.Predict(time)
				double = true
			} else {
				travel = wave.InvertedAkimaSplineB.Predict(time)
			}
			if time < wave.TravelTime.Time[wave.DoubleTravelIndex] {
				exists = false
			}
		} else {
			minDoubleTime := wave.TravelTime.Time[len(wave.TravelTime.Time)-1]
			if time > minDoubleTime {
				travelA = wave.InvertedAkimaSplineA.Predict(time)
				travelB = wave.InvertedAkimaSplineB.Predict(time)
				double = true
			} else {
				travel = wave.InvertedAkimaSplineA.Predict(time)
			}
		}
	}
	if travel >= 180 {
		exists = false
	}
	return travel, travelA, travelB, double, exists
}
