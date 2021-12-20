package holidays

var holiday2022 []OneCollection

func init() {
	holiday2022 = []OneCollection{
		{
			Start:  "2022/01/01",
			End:    "2022/01/03",
			ChName: ChHolidays[NewYearDay],
			EnName: EnHolidays[NewYearDay],
		},
		{
			Start:  "2022/01/31",
			End:    "2022/02/06",
			ChName: ChHolidays[SpringFestivalDay],
			EnName: EnHolidays[SpringFestivalDay],
		},
		{
			Start:  "2022/04/03",
			End:    "2022/04/05",
			ChName: ChHolidays[TombSweepingDay],
			EnName: EnHolidays[TombSweepingDay],
		},
		{
			Start:  "2022/04/30",
			End:    "2022/05/04",
			ChName: ChHolidays[LaborDay],
			EnName: EnHolidays[LaborDay],
		},
		{
			Start:  "2022/06/03",
			End:    "2022/06/05",
			ChName: ChHolidays[DragonBoatFestivalDay],
			EnName: EnHolidays[DragonBoatFestivalDay],
		},
		{
			Start:  "2022/09/10",
			End:    "2022/09/12",
			ChName: ChHolidays[MidAutumnFestivalDay],
			EnName: EnHolidays[MidAutumnFestivalDay],
		},
		{
			Start:  "2022/10/01",
			End:    "2022/10/07",
			ChName: ChHolidays[NationalDay],
			EnName: EnHolidays[NationalDay],
		},
	}
}
