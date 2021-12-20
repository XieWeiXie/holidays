package holidays

import (
	"fmt"
	"testing"
)

func TestFetchAll(t *testing.T) {
	FetchAll()
}

func TestFetchByYear(t *testing.T) {
	tests := []struct {
		year int
	}{
		{
			year: 2021,
		}, {
			year: 2019,
		},
	}
	for _, test := range tests {
		fmt.Println(FetchByYear(test.year))
	}
}

func TestFetchByMonth(t *testing.T) {
	tests := []struct {
		year  int
		month int
	}{
		{
			year:  2019,
			month: 12,
		}, {
			year:  2018,
			month: 1,
		}, {
			year:  2019,
			month: 10,
		}, {
			year:  2018,
			month: 6,
		},
	}
	for _, test := range tests {
		fmt.Println(FetchByMonth(test.year, test.month))
	}
}

func TestFetchByChName(t *testing.T) {
	tests := []struct {
		year   int
		chName string
	}{
		{
			year:   2018,
			chName: "国",
		}, {
			year:   2017,
			chName: "中秋",
		},
	}
	for _, test := range tests {
		fmt.Println(FetchByChName(test.year, test.chName))
	}
}

func TestFetchByEnName(t *testing.T) {
	tests := []struct {
		year   int
		enName string
	}{
		{
			year:   2018,
			enName: "Mid",
		}, {
			year:   2017,
			enName: "Spring",
		},
	}
	for _, test := range tests {
		fmt.Println(FetchByEnName(test.year, test.enName))
	}
}

func TestFetchYearHolidayCount(t *testing.T) {
	tests := []struct {
		year int
	}{
		{
			year: 2019,
		},
		{
			year: 2018,
		},
		{
			year: 2017,
		}, {
			year: 2016,
		},
	}
	for _, test := range tests {
		fmt.Println(FetchYearHolidayCount(test.year))
	}
}

func TestFetchMonthHolidayCount(t *testing.T) {
	tests := []struct {
		year  int
		month int
	}{
		{
			year:  2018,
			month: 10,
		},
		{
			year:  2017,
			month: 5,
		},
	}
	for _, test := range tests {
		fmt.Println(FetchMonthHolidayCount(test.year, test.month))
	}

}

func TestIsHoliday(t *testing.T) {
	tests := []struct {
		value string
	}{
		{
			value: "2019/10/01",
		},
		{
			value: "2018/10/01",
		},
		{
			value: "2018/01/01",
		}, {
			value: "2018/11/28",
		},
		{
			value: " 2018/11/25",
		},
	}
	for _, test := range tests {
		fmt.Println(IsHoliday(test.value))
		fmt.Println(IsWorkDay(test.value))
		fmt.Println(IsWeekDay(test.value))
	}
}

func TestFetchLunarDay(t *testing.T) {
	for _, i := range FetchLunarDay() {
		fmt.Println(i.When(), " || ", i.Desc())
	}
}

func TestGuessLunarByMonth(t *testing.T) {
	tests := []struct {
		month int
	}{
		{
			month: 1,
		},
		{
			month: 12,
		},
	}
	for _, i := range tests {
		res := GuessLunarByMonth(i.month)
		for _, j := range res {
			fmt.Println(j.When(), " || ", j.Desc())
		}
	}
}
