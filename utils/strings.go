package utils

import (
	s "strings"
)

func Capitalize(str string) (newStr string) {
	list := s.Split(str, "")
	list[0] = s.ToUpper(list[0])

	return s.Join(list, "")
}
