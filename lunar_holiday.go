package holidays

import "time"

type Traditional struct {
	Name   string
	Detail string
}

func (l Traditional) Desc() string {
	return l.Name
}

func (l Traditional) When() string {
	return l.Detail
}

type TraditionalDay interface {
	When() string
	Desc() string
}

var (
	chineseNewYear Traditional = Traditional{
		Name:   "Chinese New Year(新年)",
		Detail: "1st - 15th of the first lunar month",
	}
	lanternFestival Traditional = Traditional{
		Name:   "LanternFestival(元宵节)",
		Detail: "15th day of the first lunar month",
	}
	qingMingFestival Traditional = Traditional{
		Name:   "Qingming Festival(清明节)",
		Detail: "April 4th or 5th of the solar calendar",
	}
	dragonBoatFestival Traditional = Traditional{
		Name:   "Dragon Boat Festival(端午节)",
		Detail: "5th day of the 5th lunar month",
	}
	doubleSeventhDay Traditional = Traditional{
		Name:   "Double Seventh Festival(七夕节)",
		Detail: "7th day of seventh lunar month",
	}
	midAutumnDay Traditional = Traditional{
		Name:   "Mid-Autumn Festival(中秋节)",
		Detail: "15th of the 8th lunar month",
	}
	doubleNinthDay Traditional = Traditional{
		Name:   "Double Ninth Festival(重阳节)",
		Detail: "9th day of the 9th lunar month",
	}
	winterSolstice Traditional = Traditional{
		Name:   "Winter Solstice(冬至)",
		Detail: "Dec.21st,22nd or 23rd in solar month",
	}
	labaFestival Traditional = Traditional{
		Name:   "Laba Festival(腊八节)",
		Detail: "8th day of the 12th lunar month",
	}
)

type lunar struct {
	Days []TraditionalDay
}

func (l lunar) TraditionDays() []TraditionalDay {
	return l.Days
}

func (l lunar) GuessWithMonth(month int) []TraditionalDay {
	switch time.Month(month) {
	case time.January:
		return []TraditionalDay{chineseNewYear, lanternFestival}
	case time.April:
		return []TraditionalDay{qingMingFestival}
	case time.May:
		return []TraditionalDay{dragonBoatFestival}
	case time.June:
		return []TraditionalDay{doubleSeventhDay}
	case time.August:
		return []TraditionalDay{midAutumnDay}
	case time.September:
		return []TraditionalDay{doubleNinthDay}
	case time.December:
		return []TraditionalDay{winterSolstice, labaFestival}
	default:
		return nil
	}
}

func (l lunar) CurrentMonth() []TraditionalDay {
	month := time.Now().Month()
	return l.GuessWithMonth(int(month))
}

func LunarDay() lunar {
	return lunar{
		Days: []TraditionalDay{
			chineseNewYear,
			lanternFestival,
			qingMingFestival,
			dragonBoatFestival,
			doubleSeventhDay,
			midAutumnDay,
			doubleNinthDay,
			winterSolstice,
			labaFestival,
		},
	}
}
