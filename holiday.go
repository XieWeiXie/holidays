package holidays

import (
	"strings"
	"time"
)

// FetchAll get all holidays in China
func FetchAll() []IHoliday {
	return NewCompose().Holidays
}

// FetchByYear get holidays by year in China
func FetchByYear(year int) IHoliday {
	var index int
	nowYear, _, _ := time.Now().Date()
	if year > nowYear+1 {
		return nil
	}
	index = nowYear - year
	return NewCompose().Holidays[index]

}

// FetchByMonth get holidays by year and month in China
func FetchByMonth(year int, month int) []OneCollection {
	if month < 1 || month > 12 {
		return nil

	}
	collections := FetchByYear(year)
	var data []OneCollection
	for _, collection := range collections.Get() {
		collectionTime, _ := time.Parse("2006/01/02", collection.End)
		if int(collectionTime.Month()) == month {
			data = append(data, collection)
		}
	}
	return data
}

// FetchByChName get holidays by year and chinese name in China
func FetchByChName(year int, name string) []OneCollection {
	collections := FetchByYear(year)
	var data []OneCollection
	for _, collection := range collections.Get() {
		if strings.Contains(collection.ChName, name) {
			data = append(data, collection)
		}
	}
	return data

}

// FetchByEnName get holidays by year and english name in China
func FetchByEnName(year int, name string) []OneCollection {
	collections := FetchByYear(year)
	var data []OneCollection
	for _, collection := range collections.Get() {
		if strings.Contains(collection.EnName, name) {
			data = append(data, collection)
		}
	}
	return data
}

// IsHoliday judge date is holiday or not
func IsHoliday(value string) bool {
	collectionTime, err := time.Parse("2006/01/02", value)
	if err != nil {
		return false
	}
	nowYear, _, _ := time.Now().Date()
	if collectionTime.Year() > nowYear+1 {
		return false
	}
	collections := FetchByYear(collectionTime.Year())
	for _, collection := range collections.Get() {
		startDate, _ := getDate(collection.Start)
		endDate, _ := getDate(collection.End)
		if collectionTime.Unix() >= startDate.Unix() && collectionTime.Unix() <= endDate.Unix() {
			return true
		}
	}
	return false

}

// IsWorkDay judge date is work day or not
func IsWorkDay(value string) bool {
	if IsHoliday(value) {
		return false
	}
	collectionTime, err := time.Parse("2006/01/02", value)
	if err != nil {
		return false
	}
	isWorkDay := int(collectionTime.Weekday())
	if isWorkDay == 0 || isWorkDay == 6 {
		return false
	}
	return true
}

// IsWeekDay judge date is week day or not
func IsWeekDay(value string) bool {
	return !IsWorkDay(value) && !IsHoliday(value)
}

// FetchYearHolidayCount count day of one year holidays
func FetchYearHolidayCount(year int) int {
	collections := FetchByYear(year)
	var count int
	for _, collection := range collections.Get() {
		count += countOneHoliday(collection)
	}
	return count
}

// FetchMonthHolidayCount count day in one month of one year holidays
func FetchMonthHolidayCount(year int, month int) int {
	collections := FetchByMonth(year, month)
	var count int
	for _, collection := range collections {
		count += countOneHoliday(collection)
	}
	return count
}

// FetchLunarDay fetch all lunar day in China
func FetchLunarDay() []TraditionalDay {
	return LunarDay().Days
}

// GuessLunarByMonth guess lunar day by month in China
func GuessLunarByMonth(month int) []TraditionalDay {
	if month <= 0 || month > 12 {
		return LunarDay().CurrentMonth()
	}
	return LunarDay().GuessWithMonth(month)
}
