package leet_04

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cast"
)

var _, yearSize = utf8.DecodeRuneInString("年")
var _, monSize = utf8.DecodeRuneInString("月")

func Format(str string) (string, error) {
	str = strings.TrimSpace(str)
	idx := strings.Index(str, "年")
	if idx == -1 {
		return "", errors.New("not found")
	}
	year := str[:idx]
	remain := str[idx:]
	idx = strings.Index(remain, "月")
	if idx == -1 {
		return "", errors.New("not found")
	}
	mon := remain[yearSize:idx]
	remain = remain[idx:]
	idx = strings.Index(remain, "日")
	if idx == -1 {
		return "", errors.New("not found")
	}
	days := remain[monSize:idx]
	return fmt.Sprintf("%s-%02d-%02d", year, cast.ToInt(mon), cast.ToInt(days)), nil
}

func Format2(str string) (string, error) {
	str = strings.TrimSpace(str)
	str = strings.TrimRight(str, "日")
	str = strings.ReplaceAll(str, "年", "-")
	str = strings.ReplaceAll(str, "月", "-")
	ss := strings.Split(str, "-")

	if len(ss) != 3 {
		return "", errors.New("not found")
	}
	year, mons, days := ss[0], cast.ToInt(ss[1]), cast.ToInt(ss[2])

	return fmt.Sprintf("%s-%02d-%02d", year, mons, days), nil
}
