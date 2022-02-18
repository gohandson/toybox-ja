package main

func main() {
	converters := []converter{
		celsiusToFahrenheit(), // 摂氏[°C] -> 華氏[°F]
		fahrenheitToCelsius(), // 華氏[°F] -> 摂氏[°C]
		calToJoule(),          // カロリー[cal] -> ジュール[J]
		jouleToCal(),          // ジュール[J] -> カロリー[cal]
	}
	start(converters)
}
