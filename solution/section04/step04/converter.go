package main

type converter struct {
	name     string
	calc     func(v float64) float64
	fromUnit string
	toUnit   string
}

func (c *converter) convert(v float64) float64 {
	return c.calc(v)
}

func celsiusToFahrenheit() *converter {
	return &converter{
		name: "摂氏[°C] -> 華氏[°F]",
		calc: func(v float64) float64 {
			return v*1.8 + 32
		},
		fromUnit: "°C",
		toUnit:   "°F",
	}
}

func fahrenheitToCelsius() *converter {
	return &converter{
		name: "華氏[°F] -> 摂氏[°C]",
		calc: func(v float64) float64 {
			return (v - 32) / 1.8
		},
		fromUnit: "°F",
		toUnit:   "°C",
	}
}

func calToJoule() *converter {
	return &converter{
		name: "カロリー[cal] -> ジュール[J]",
		calc: func(v float64) float64 {
			return v * 4.18
		},
		fromUnit: "cal",
		toUnit:   "J",
	}
}

func jouleToCal() *converter {
	return &converter{
		name: "ジュール[J] -> カロリー[cal]",
		calc: func(v float64) float64 {
			return v * 0.239
		},
		fromUnit: "J",
		toUnit:   "cal",
	}
}
