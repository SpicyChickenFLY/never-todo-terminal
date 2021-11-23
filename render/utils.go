package render

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	tioWinSZ      = 0x5413
	tioCGWinSZOSX = 1074295912
)

// func lenOfTerminal() (int, error) {
// 	type window struct {
// 		Row    uint16
// 		Col    uint16
// 		Xpixel uint16
// 		Ypixel uint16
// 	}
// 	w := new(window)
// 	tio := tioWinSZ
// 	if runtime.GOOS == "darwin" {
// 		tio = tioCGWinSZOSX
// 	}
// 	res, _, err := syscall.Syscall(
// 		syscall.SYS_IOCTL,
// 		uintptr(syscall.Stdin),
// 		uintptr(tio),
// 		uintptr(unsafe.Pointer(w)),
// 	)
// 	if int(res) == -1 {
// 		return 0, err
// 	}
// 	return int(w.Col), nil
// }

func lenOfTerminal() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	sizeStr := strings.ReplaceAll(string(out), "\n", "")
	sizes := strings.Split(sizeStr, " ")
	width, err := strconv.Atoi(sizes[1])
	return width, nil
}

func lenOnScreen(str string) int {
	length := 0
	for _, r := range []rune(str) {
		rVal := int(r)
		length++
		if rVal >= 128 {
			length++
		}
	}
	return length
}

func splitStrOnScreen(str string, l int) []string {
	result := []string{}
	tempStr := ""
	tempLen := 0
	for _, r := range []rune(str) {
		rVal := int(r)
		rLen := 1
		if rVal >= 128 {
			rLen = 2
		}
		if tempLen+rLen >= l {
			result = append(result, tempStr)
			tempStr = string(r)
			tempLen = rLen
		} else {
			tempStr += string(r)
			tempLen += rLen
		}
	}
	return result
}

func fillSpace(l int) {
}
