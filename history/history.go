package history

const (
	NewYearDay = iota
	SpringFestivalDay
	TombSweepingDay
	LaborDay
	DragonBoatFestivalDay
	NationalDay
	MidAutumnFestivalDay
)

var ChHolidays = [...]string{
	"元旦",
	"春节",
	"清明节",
	"劳动节",
	"端午节",
	"中秋节",
	"国庆节",
}
var EnHolidays = [...]string{
	"New Year\\'s Day",
	"Spring Festival",
	"Tomb-sweeping Day",
	"Labour Day",
	"Dragon Boat Festival",
	"Mid-autumn Festival",
	"National Day",
}

type CollectionYearHistory struct {
	Data [][]OneCollection `json:"data"`
}

func (h *CollectionYearHistory) Add(one []OneCollection) {
	h.Data = append(h.Data, one)
}

type OneCollection struct {
	Start  string `json:"start"`
	End    string `json:"end"`
	ChName string `json:"ch_name"`
	EnName string `json:"en_name"`
}

func NewOneCollection(chName string, enName string, start string, end string) *OneCollection {
	return &OneCollection{
		Start:  start,
		End:    end,
		ChName: chName,
		EnName: enName,
	}
}

type YearCollection struct {
	Data []OneCollection `json:"data"`
}

func (y *YearCollection) Add(one OneCollection) {
	y.Data = append(y.Data, one)
}

func FetchCollectionYearHistory() CollectionYearHistory {
	return CollectionYearHistory{
		Data: [][]OneCollection{
			holiday2021,
			holiday2020,
			holiday2019,
			holiday2018,
			holiday2017,
			holiday2016,
			holiday2015,
			holiday2014,
			holiday2013,
			holiday2012,
			holiday2011,
			holiday2010,
		},
	}
}

var holiday2021 = []OneCollection{
	{
		Start:  "2021/01/01",
		End:    "2021/01/03",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2021/02/11",
		End:    "2021/02/17",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2021/04/03",
		End:    "2021/04/05",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2021/05/01",
		End:    "2021/05/05",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2021/06/12",
		End:    "2021/06/14",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2021/09/19",
		End:    "2021/09/21",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2021/10/01",
		End:    "2021/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}


var holiday2020 = []OneCollection{
	{
		Start:  "2020/01/01",
		End:    "2020/01/01",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2020/01/24",
		End:    "2020/01/30",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2020/04/04",
		End:    "2020/04/06",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2020/05/01",
		End:    "2020/05/05",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2020/06/25",
		End:    "2020/06/27",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2020/10/01",
		End:    "2020/10/01",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2020/10/01",
		End:    "2020/10/08",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2019 = []OneCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2018 = []OneCollection{
	{
		Start:  "2017/12/30",
		End:    "2018/01/01",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2018/02/15",
		End:    "2018/02/21",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2018/04/05",
		End:    "2018/04/07",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2018/04/29",
		End:    "2018/05/01",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2018/06/16",
		End:    "2018/06/18",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2018/09/22",
		End:    "2018/09/24",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2018/10/01",
		End:    "2018/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2017 = []OneCollection{
	{
		Start:  "2016/12/31",
		End:    "2017/01/02",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2017/01/27",
		End:    "2017/02/02",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2017/04/02",
		End:    "2017/04/04",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2017/04/29",
		End:    "2017/05/01",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2017/05/28",
		End:    "2017/05/30",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2017/10/01",
		End:    "2017/10/08",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2016 = []OneCollection{
	{
		Start:  "2016/01/01",
		End:    "2016/01/03",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2016/02/07",
		End:    "2016/02/13",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2016/04/02",
		End:    "2016/04/04",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2016/04/30",
		End:    "2016/05/02",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2016/06/09",
		End:    "2016/06/11",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2016/09/15",
		End:    "2016/09/17",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2016/10/01",
		End:    "2016/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2015 = []OneCollection{
	{
		Start:  "2015/01/01",
		End:    "2015/01/03",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2015/02/18",
		End:    "2015/02/24",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2015/04/04",
		End:    "2015/04/06",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2015/05/01",
		End:    "2015/05/03",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2015/06/20",
		End:    "2015/06/22",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2015/09/26",
		End:    "2015/09/27",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2015/10/01",
		End:    "2015/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2014 = []OneCollection{
	{
		Start:  "2014/01/01",
		End:    "2014/01/01",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2014/01/31",
		End:    "2014/02/06",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2014/04/05",
		End:    "2014/04/07",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2014/05/01",
		End:    "2014/05/03",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2014/05/31",
		End:    "2014/06/02",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2014/09/06",
		End:    "2014/09/08",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2014/10/01",
		End:    "2014/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2013 = []OneCollection{
	{
		Start:  "2013/01/01",
		End:    "2013/01/03",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2013/02/09",
		End:    "2013/02/15",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2013/04/04",
		End:    "2013/04/06",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2013/04/29",
		End:    "2013/05/01",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2013/06/10",
		End:    "2013/06/12",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2013/09/19",
		End:    "2013/09/21",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2013/10/01",
		End:    "2013/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2012 = []OneCollection{
	{
		Start:  "2012/01/01",
		End:    "2012/01/03",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2012/01/22",
		End:    "2012/01/28",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2012/04/02",
		End:    "2012/04/04",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2012/04/29",
		End:    "2012/05/01",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2012/06/22",
		End:    "2012/06/24",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2012/09/30",
		End:    "2012/09/30",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2012/10/01",
		End:    "2012/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2011 = []OneCollection{
	{
		Start:  "2011/01/01",
		End:    "2011/01/03",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2011/02/02",
		End:    "2011/02/08",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2011/04/03",
		End:    "2011/04/04",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2011/04/30",
		End:    "2011/05/02",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2011/06/04",
		End:    "2011/06/06",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2011/09/10",
		End:    "2011/09/12",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2011/10/01",
		End:    "2011/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
var holiday2010 = []OneCollection{
	{
		Start:  "2010/01/01",
		End:    "2010/01/03",
		ChName: ChHolidays[NewYearDay],
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2010/02/13",
		End:    "2010/02/19",
		ChName: ChHolidays[SpringFestivalDay],
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2010/04/03",
		End:    "2010/04/05",
		ChName: ChHolidays[TombSweepingDay],
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2010/05/01",
		End:    "2010/05/03",
		ChName: ChHolidays[LaborDay],
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2010/06/14",
		End:    "2010/06/16",
		ChName: ChHolidays[DragonBoatFestivalDay],
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2010/09/22",
		End:    "2010/09/24",
		ChName: ChHolidays[MidAutumnFestivalDay],
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2010/10/01",
		End:    "2010/10/07",
		ChName: ChHolidays[NationalDay],
		EnName: EnHolidays[NationalDay],
	},
}
