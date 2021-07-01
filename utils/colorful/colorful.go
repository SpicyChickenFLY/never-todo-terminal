package colorful

import "fmt"

// 前景 背景 颜色
// --------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
const (
	FrontBlack = iota + 30
	FrontRed
	FrontGreen
	FrontYellow
	FrontBlue
	FrontPurple
	FrontCyan
	FrontWhite
)
const (
	BackBlack = iota + 40
	BackRed
	BackGreen
	BackYellow
	BackBlue
	BackPurple
	BackCyan
	BackWhite
)

// 代码 意义
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

const MARK = 0x1B

//global variable
var modeMap = map[string]int{
	"default":   ModeDefault,
	"highlight": ModeHighLight,
	"line":      ModeLine,
	"flash":     ModeFlash,
	"rewrite":   ModeReWhite,
	"hidden":    ModeHidden,
}

var frontMap = map[string]int{
	"black":  FrontBlack,
	"red":    FrontRed,
	"green":  FrontGreen,
	"yellow": FrontYellow,
	"blue":   FrontBlue,
	"purple": FrontPurple,
	"cyan":   FrontCyan,
	"white":  FrontWhite,
}

var backMap = map[string]int{
	"black":  BackBlack,
	"red":    BackRed,
	"green":  BackGreen,
	"yellow": BackYellow,
	"blue":   BackBlue,
	"purple": BackPurple,
	"cyan":   BackCyan,
	"white":  BackWhite,
}

// RenderStr is a func for render string with specified color and style
// str:  		your output string
// modeStr: 	your mode string(default/highlight/line/flash/rewrite/hidden)
// bColorStr:	your back color string(black/red/yellow/green/blue/cyan/purple/white)
// fColorStr: 	your front color string(same as back color)
func RenderStr(str, modeStr, bColorStr, fColorStr string) string {
	mode, ok := modeMap[modeStr]
	if !ok {
		mode = ModeDefault
	}
	bColor, ok := backMap[bColorStr]
	if !ok {
		bColor = BackBlack
	}
	fColor, ok := frontMap[fColorStr]
	if !ok {
		fColor = FrontGreen
	}
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", MARK, mode, bColor, fColor, str, MARK)
}

func GetStartMark(modeStr, bColorStr, fColorStr string) string {
	mode, ok := modeMap[modeStr]
	if !ok {
		mode = ModeDefault
	}
	bColor, ok := backMap[bColorStr]
	if !ok {
		bColor = BackBlack
	}
	fColor, ok := frontMap[fColorStr]
	if !ok {
		fColor = FrontGreen
	}
	return fmt.Sprintf("%c[%d;%d;%dm", MARK, mode, bColor, fColor)
}

func GetEndMark() string {
	return fmt.Sprintf("%c[0m", MARK)
}
