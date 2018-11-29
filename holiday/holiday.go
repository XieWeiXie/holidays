package holidays

import (
	"holidays/history"
	"strings"
	"time"
)

func FetchAll() history.CollectionYearHistory {
	return history.FetchCollectionYearHistory()

}
func FetchByYear(year int) []history.OneCollection {
	var index int
	nowYear, _, _ := time.Now().Date()
	if year > nowYear+1 {
		return nil
	}
	index = nowYear + 1 - year
	return history.FetchCollectionYearHistory().Data[index]

}
func FetchByMonth(year int, month int) []history.OneCollection {
	if month < 1 || month > 12 {
		return nil

	}
	collections := FetchByYear(year)
	var data []history.OneCollection
	for _, collection := range collections {
		collectionTime, _ := time.Parse("2006/01/02", collection.End)
		if int(collectionTime.Month()) == month {
			data = append(data, collection)
		}
	}
	return data
}
func FetchByChName(year int, name string) []history.OneCollection {
	collections := FetchByYear(year)
	var data []history.OneCollection
	for _, collection := range collections {
		if strings.Contains(collection.ChName, name) {
			data = append(data, collection)
		}
	}
	return data

}
func FetchByEnName(year int, name string) []history.OneCollection {
	collections := FetchByYear(year)
	var data []history.OneCollection
	for _, collection := range collections {
		if strings.Contains(collection.EnName, name) {
			data = append(data, collection)
		}
	}
	return data
}
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
	for _, collection := range collections {
		startDate, _ := GetDate(collection.Start)
		endDate, _ := GetDate(collection.End)
		if collectionTime.Unix() >= startDate.Unix() && collectionTime.Unix() <= endDate.Unix() {
			return true
		}
	}
	return false

}
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

func IsWeekDay(value string) bool {
	return !IsWorkDay(value) && !IsHoliday(value)
}
func FetchYearHolidayCount(year int) int {
	collections := FetchByYear(year)
	var count int
	for _, collection := range collections {
		count += CountOneHoliday(collection)
	}
	return count
}

func FetchMonthHolidayCount(year int, month int) int {
	collections := FetchByMonth(year, month)
	var count int
	for _, collection := range collections {
		count += CountOneHoliday(collection)
	}
	return count
}
