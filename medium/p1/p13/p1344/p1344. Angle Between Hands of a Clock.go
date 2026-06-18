package p1344

import "math"

const (
	HourHandDegreesPerMinute   = 0.5  // Скорость часовой стрелки (360° / 720 мин)
	MinuteHandDegreesPerMinute = 6.0  // Скорость минутной стрелки (360° / 60 мин)
	HourHandDegreesPerHour     = 30.0 // Скорость часовой стрелки без учета минут (360° / 12 часов)
)

func angleClock(hour int, minutes int) float64 {
	hourDegrees := HourHandDegreesPerHour*float64(hour) + HourHandDegreesPerMinute*float64(minutes)
	minuteDegrees := MinuteHandDegreesPerMinute * float64(minutes)
	res := math.Abs(hourDegrees - minuteDegrees) // угол между стрелками часов
	res = min(res, 360.0-res)                    // угол между стрелками часов в привычном виде
	return res
}
