package utils

import (
	"math"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func ContainStr(sentence, word string) bool {
	sentence = strings.ToLower(sentence)
	word = strings.ToLower(word)
	reg := regexp.MustCompile("\\s+")
	sentence = reg.ReplaceAllString(sentence, "")
	word = reg.ReplaceAllString(word, "")
	return strings.Contains(sentence, word)
}

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

func min(x int, y int, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}

func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func LessInAbs(a, b int) bool {
	return Abs(a) < Abs(b)
}

func LessInID(a, b int) bool {
	return LessInAbs(a, b)
}

func LessInTime(a, b time.Time) bool {
	return a.Before(b)
}

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
