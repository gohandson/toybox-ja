package main

type converter struct {
	name     string
	calc     func(v float64) float64
	fromUnit string
	toUnit   string
}

// 摂氏[°C] -> 華氏[°F]を行う変換器を作成
func celsiusToFahrenheit() converter {
	return converter{
		name: "摂氏[°C] -> 華氏[°F]",
		calc: func(v float64) float64 {
			return v*1.8 + 32
		},
		fromUnit: "°C",
		toUnit:   "°F",
	}
}

// TODO: 華氏[°F] -> 摂氏[°C]を行う変換器を作成する関数を定義
// 関数名：fahrenheitToCelsius

// カロリー[cal] -> ジュール[J]を行う変換器を作成
func calToJoule() converter {
	return converter{
		name: "カロリー[cal] -> ジュール[J]",
		calc: func(v float64) float64 {
			return v * 4.18
		},
		fromUnit: "cal",
		toUnit:   "J",
	}
}

// ジュール[J] -> カロリー[cal]を行う変換器を作成
func jouleToCal() converter {
	return converter{
		name: "ジュール[J] -> カロリー[cal]",
		calc: func(v float64) float64 {
			return v * 0.239
		},
		fromUnit: "J",
		toUnit:   "cal",
	}
}
