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

type fieldConf struct {
	min, max uint
	units    []string
}

var fcs = []*fieldConf{
	{0, 59, []string{"second"}},
	{0, 59, []string{"minute"}},
	{0, 23, []string{"hour"}},
	{1, 31, []string{"day"}},
	{1, 12, monthName},
	{1, 7, dayName},
	{1, 9999, []string{"year"}},
}

// field value related
const (
	maskLen     = 64
	maskAll     = ^(uint64(0))
	maskOpt     = 1<<63 | 1<<62 | 1<<61 | 1<<60
	maskData    = maskAll ^ maskOpt
	maskOptL    = 1 << 62
	maskOptStar = 1 << 63
)

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
