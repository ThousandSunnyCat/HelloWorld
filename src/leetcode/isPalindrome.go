package leetcode

import (
	"regexp"
	"strings"
)

// 125
func IsPalindrome(s string) bool {
	p, q := 0, len(s)-1
	for p < q {
		i, j := s[p], s[q]
		if isUpper(i) {
			if isNum(j) {
				return false
			}
			if isUpper(j) {
				if i == j {
					p++
					q--
					continue
				} else {
					return false
				}
			}
			if isLower(j) {
				if i == j - 32 {
					p++
					q--
					continue
				} else {
					return false
				}
			}
			q--
			continue
		}
		if isLower(i) {
			if isNum(j) {
				return false
			}
			if isLower(j) {
				if i == j {
					p++
					q--
					continue
				} else {
					return false
				}
			}
			if isUpper(j) {
				if i == j + 32 {
					p++
					q--
					continue
				} else {
					return false
				}
			}
			q--
			continue
		}
		if isNum(i) {
			if isNum(j) {
				if i == j {
					p++
					q--
					continue
				} else {
					return false
				}
			}
			if isUpper(j) {
				return false
			}
			if isLower(j) {
				return false
			}
			q--
			continue
		}
		p++
		continue
	}
	return true
}

func isUpper(k byte) bool {
	return k > 64 && k < 91
}

func isLower(k byte) bool {
	return k > 96 && k < 123
}

func isNum(k byte) bool {
	return k > 47 && k < 58
}

func IsPalindrome2(s string) bool {
	reg, _ := regexp.Compile(`[^a-zA-Z0-9]`)
	s = strings.ToLower(s)
	s2 := reg.ReplaceAll([]byte(s), []byte(""))
	p, q := 0, len(s2)-1
	for p < q {
		if s2[p] != s2[q] {
			return false
		}
		p++
		q--
	}
	return true
}

func IsPalindrome3(s string) bool {
	p, q := 0, len(s)-1
	for p < q {
		var i, j byte
		if i = toNewByte(s[p]); i == 0 {
			p++
			continue
		}
		if j = toNewByte(s[q]); j == 0 {
			q--
			continue
		}

		if i != j {
			return false
		}

		p++
		q--
	}

	return true
}

func toNewByte(i byte) byte {
	if isUpper(i) {
		return i - 'A' + 11
	} else if isLower(i) {
		return i - 'a' + 11
	} else if isNum(i) {
		return i - '0' + 1
	}

	return 0
}