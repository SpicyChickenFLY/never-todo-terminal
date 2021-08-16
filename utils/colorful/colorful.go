package colorful

import "fmt"

// MARK make terminal realize this is a color controller
const MARK = 0x1B

// Mode Map(Can be combined)
// code meaning
// --------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见
const (
	ModeDefault   = 0
	ModeHighLight = 1
	ModeLine      = 4
	ModeFlash     = 5
	ModeReWhite   = 6
	ModeHidden    = 7
)

//global variable
var modeMap = map[string]int{
	"default":   ModeDefault,
	"highlight": ModeHighLight,
	"line":      ModeLine,
	"flash":     ModeFlash,
	"rewrite":   ModeReWhite,
	"hidden":    ModeHidden,
}

// Color Map 8/16
// f   b   meaning
// --------------
// 39  49  default
// 30  40  black
// 31  41  red
// 32  42  green
// 33  43  yellow
// 34  44  blue
// 35  45  magenta
// 36  46  cyan
// 37  47  light grey
// 90  100 dark grey
// 91  101 light red
// 92  102 light green
// 93  103 light yellow
// 94  104 light blue
// 95  105 light magenta
// 96  106 light cyan
// 97  107 white
const (
	frontBlack = iota + 30
	frontRed
	frontGreen
	frontYellow
	frontBlue
	frontPurple
	frontCyan
	frontWhite
	frontDefault = 39
)
const (
	frontLightBlack = iota + 90
	frontLightRed
	frontLightGreen
	frontLightYellow
	frontLightBlue
	frontLightPurple
	frontLightCyan
	frontLightWhite
)
const (
	backBlack = iota + 40
	backRed
	backGreen
	backYellow
	backBlue
	backPurple
	backCyan
	backWhite
	backDefault = 49
)
const (
	backLightBlack = iota + 100
	backLightRed
	backLightGreen
	backLightYellow
	backLightBlue
	backLightPurple
	backLightCyan
	backLightWhite
)

var frontMap8_16 = map[string]int{
	"":        frontDefault,
	"crimson": frontRed,
	"red":     frontLightRed,
	"yellow":  frontLightYellow,
	"brown":   frontYellow,
	"green":   frontGreen,
	"cyan":    frontLightCyan,
	"sky":     frontLightBlue,
	"blue":    frontBlue,
	"magenta": frontLightPurple,
	"purple":  frontPurple,
	"grey":    frontWhite,
	"black":   frontLightBlack,
}

var backMap8_16 = map[string]int{
	"":        backDefault,
	"crimson": backRed,
	"red":     backLightRed,
	"yellow":  backLightYellow,
	"brown":   backYellow,
	"green":   backGreen,
	"cyan":    backLightCyan,
	"sky":     backLightBlue,
	"blue":    backBlue,
	"magenta": backLightPurple,
	"purple":  backPurple,
	"grey":    backWhite,
	"black":   backLightBlack,
}

// Color Map 88/256
// There are about 256 kind of colors so we just list color we used
const (
	Crimson = 88
	Yellow  = 11
	Brown   = 94
	Green   = 2
	Cyan    = 14
	Sky     = 4
	Blue    = 19
	Magenta = 13
	Purple  = 5
	Grey    = 249
	Black   = 239
)

// RenderStr is a func for render string with specified color and style
// str:  		your output string
// modeStr: 	your mode string(default/highlight/line/flash/rewrite/hidden)
// bColorStr:	your back color string(black/red/yellow/green/blue/cyan/purple/white)
// fColorStr: 	your front color string(same as back color)
func RenderStr(str, modeStr, bColorStr, fColorStr string) string {
	startMark := GetStartMark(modeStr, bColorStr, fColorStr)
	endMark := GetEndMark()
	return fmt.Sprintf("%s%s%s", startMark, str, endMark)
}

// GetStartMark return start mark with color controller
func GetStartMark(modeStr, bColorStr, fColorStr string) string {
	mode, ok := modeMap[modeStr]
	if !ok {
		mode = ModeDefault
	}

	bColor, ok := backMap8_16[bColorStr]
	if !ok {
		bColor = backDefault
	}
	bColorStr = fmt.Sprintf("%d", bColor)

	fColor, ok := frontMap8_16[fColorStr]
	if !ok {
		fColor = frontDefault
	}
	fColorStr = fmt.Sprintf("%d", fColor)

	return fmt.Sprintf("%c[%d;%s;%sm", MARK, mode, bColorStr, fColorStr)

}

// GetEndMark return end mark
func GetEndMark() string {
	return fmt.Sprintf("%c[0m", MARK)
}
