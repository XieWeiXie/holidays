package holidays

import (
	"errors"
	"time"
)

func getDate(value string) (time.Time, error) {
	valueTime, err := time.Parse("2006/01/02", value)
	if err != nil {
		return time.Time{}, err
	}
	return valueTime, nil
}

func countOneHoliday(value OneCollection) int {
	start, _ := time.Parse("2006/01/02", value.Start)
	end, _ := time.Parse("2006/01/02", value.End)
	return int(end.Sub(start).Hours()/24) + 1
}

type OneCollection struct {
	Start  string `json:"start"`
	End    string `json:"end"`
	ChName string `json:"ch_name"`
	EnName string `json:"en_name"`
}

// CollectionOption 函数选项模式
type CollectionOption func(c *OneCollection)

func NewOneCollection(chName string, enName string, start string, end string) *OneCollection {
	return &OneCollection{
		Start:  start,
		End:    end,
		ChName: chName,
		EnName: enName,
	}
}

func WithStart(start string) CollectionOption {
	return func(c *OneCollection) {
		c.Start = start
	}
}

func WithEnd(end string) CollectionOption {
	return func(c *OneCollection) {
		c.End = end
	}
}

func WithChName(chName string) CollectionOption {
	return func(c *OneCollection) {
		c.ChName = chName
	}
}

func WithEnName(enName string) CollectionOption {
	return func(c *OneCollection) {
		c.EnName = enName
	}
}

type IHoliday interface {
	Set(data []OneCollection) error
	Get() []OneCollection
}

type holiday struct {
	data []OneCollection
}

func newHoliday(data []OneCollection) *holiday {
	return &holiday{
		data: data,
	}
}

func (h *holiday) Set(data []OneCollection) (err error) {
	if len(data) == 0 {
		return errors.New("data is empty")
	}
	h.data = data
	return
}
func (h *holiday) Get() []OneCollection {
	return h.data
}

type Compose struct {
	Holidays []IHoliday
}

func (c Compose) Add(iHoliday IHoliday) {
	if len(c.Holidays) == 0 {
		c.Holidays = make([]IHoliday, 0)
		c.Holidays = append(c.Holidays, iHoliday)
	}
	c.Holidays = append(c.Holidays, iHoliday)
}

func NewCompose() *Compose {
	return &Compose{
		Holidays: []IHoliday{
			newHoliday(holiday2022),
			newHoliday(holiday2021),
			newHoliday(holiday2020),
			newHoliday(holiday2019),
			newHoliday(holiday2018),
			newHoliday(holiday2017),
			newHoliday(holiday2016),
			newHoliday(holiday2015),
			newHoliday(holiday2014),
			newHoliday(holiday2013),
			newHoliday(holiday2012),
			newHoliday(holiday2011),
			newHoliday(holiday2020),
		},
	}
}
