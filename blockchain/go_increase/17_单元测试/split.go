package split

import "strings"

func Split(s, sep string) (result []string) {
	// Index 返回子串 sep 在字符串 s 中第一次出现的位置
	// 如果找不到，则返回 -1，如果 sep 为空，则返回 0
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)

	return
}
