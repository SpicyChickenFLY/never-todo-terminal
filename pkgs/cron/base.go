package cron

var monthName = []string{
	"month",
	"January", "February", "March", "April",
	"May", "June", "July", "August",
	"September", "October", "November", "December",
}

var dayName = []string{
	"day", "Monday", "Tuesday", "Wednesday",
	"Thursday", "Friday", "Saturday", "Sunday",
}

const monthWith31Days = 1<<1 | 1<<3 | 1<<5 | 1<<7 | 1<<8 | 1<<10 | 1<<12

// field index
const (
	secType = iota
	minType
	hourType
	domType
	monthType
	dowType
	yearType
	fieldCount
)
