package utils

import "regexp"

func HasSpecialText(text string) bool {
	pattern := `[~!@#$%^&*()+=|{}':;',\\\[\].<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。，、？ ]`
	matcher := regexp.MustCompile(pattern)
	result := matcher.FindAllString(text, -1)
	return len(result) > 0
}
