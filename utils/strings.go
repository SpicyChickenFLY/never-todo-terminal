package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// ContainStr search for sub string
func ContainStr(sentence, word string) bool {
	sentence = strings.ToLower(sentence)
	word = strings.ToLower(word)
	reg := regexp.MustCompile("\\s+")
	sentence = reg.ReplaceAllString(sentence, "")
	word = reg.ReplaceAllString(word, "")
	return strings.Contains(sentence, word)
}

// MinDistance calculate Edit-Distance for strings
func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// ContainChinese return result boolean
func ContainChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) ||
			(regexp.MustCompile(
				"[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

func getPrefixTable(search string) []int {
	ptLength := len(search)
	pt := make([]int, ptLength)
	pt[0] = 0
	len := 0
	i := 1
	for i < ptLength {
		if search[i] == search[len] {
			len++
			pt[i] = len
			i++
		} else {
			if len > 1 {
				len = pt[len-1]
			} else {
				pt[i] = len
				i++
			}
		}
	}
	return pt
}
func shiftPT(pt []int) {
	len := len(pt)
	for i := len - 1; i > 0; i-- {
		pt[i] = pt[i-1]
	}
	pt[0] = -1
}

func searchWithPT(text string, pattern string) (idxRec []int) {
	idxRec = []int{}
	pt := getPrefixTable(pattern)
	shiftPT(pt)
	// fmt.Printf("pt=%v", pt)
	M := len(text)
	N := len(pattern)
	if N > M {
		return
	}
	// i 追踪text的位置 ， j 追踪pattern的位置
	txtIdx, patIdx := 0, 0
	found := 0
	for txtIdx < M {
		if patIdx == N-1 && text[txtIdx] == pattern[patIdx] {
			// fmt.Printf("found pattern \"%s\" at index %d\n", pattern, txtIdx-patIdx)
			idxRec = append(idxRec, txtIdx-patIdx)
			found++
			patIdx = pt[patIdx]
		}
		if text[txtIdx] == pattern[patIdx] {
			txtIdx++
			patIdx++
		} else {
			patIdx = pt[patIdx]
			if patIdx == -1 {
				txtIdx++
				patIdx++
			}
		}
	}
	// fmt.Printf("find %d pattern in the text \n", found)
	// return found > 0
	return
}

func chinese2Unicode(str string) string {

	if len(str) > 0 {
		str = strconv.QuoteToASCII(str)
		str = str[1 : len(str)-1]
	}
	return str
}

func unicode2Chinese(str string) (string, error) {
	// fmt.Println("str:", str)
	ascii, err := strconv.ParseInt(str, 16, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%c", ascii), nil
}

// EncodeCmd transfer normal string into unicode string
func EncodeCmd(cmd string) string {
	return chinese2Unicode(cmd)
}

// DecodeCmd transfer unicode string into formal string
func DecodeCmd(cmd string) (result string, err error) {
	result = ""
	idxRec := searchWithPT(cmd, "\\u")
	if len(idxRec) == 0 {
		return cmd, nil
	}
	// fmt.Println("idxRec:", idxRec)
	if idxRec[0] != 0 {
		result += cmd[0:idxRec[0]]
	}
	for i := range idxRec {
		if i != 0 && idxRec[i]-idxRec[i-1] > 6 {
			result += cmd[idxRec[i-1]+6 : idxRec[i]-1]
		}
		word, err := unicode2Chinese(cmd[idxRec[i]+2 : idxRec[i]+6])
		if err != nil {
			return "", err
		}
		result += word
	}
	if len(idxRec) > 0 && idxRec[len(idxRec)-1]+6 != len(cmd) {
		result += cmd[idxRec[len(idxRec)-1]+6:]
	}
	return result, nil
}

// mustParseUint parses the given expression as an int or returns an error.
func mustParseUint(expr string) (uint, error) {
	num, err := strconv.Atoi(expr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int from %s: %s", expr, err)
	}
	if num < 0 {
		return 0, fmt.Errorf("negative number (%d) not allowed: %s", num, expr)
	}

	return uint(num), nil
}
