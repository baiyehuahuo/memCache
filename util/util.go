package util

import (
	"encoding/json"
	"log"
	"memCache/define"
	"strconv"
	"strings"
	"unicode"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

func ParseSize(size string) (parseByteSize int64, parseByteSizeStr string) {
	defer func() {
		if parseByteSize == 0 || len(parseByteSizeStr) == 0 {
			log.Println("解析size失败，返回默认值")
			parseByteSize = define.DefaultMemSize
			parseByteSizeStr = define.DefaultMemSizeStr
		}
	}()

	length := len(size)
	if length < 2 || unicode.ToUpper(rune(size[length-1])) != 'B' {
		return
	}

	var unit string
	// B 只有一个字母位，单独做下处理
	if unicode.IsDigit(rune(size[length-2])) {
		unit = "B"
	} else {
		unit = strings.ToUpper(size[length-2:])
	}

	byteNum, err := strconv.Atoi(size[:length-len(unit)])
	if err != nil {
		return
	}
	parseByteSize = int64(byteNum)
	size = strings.ToUpper(size)

	switch unit {
	case "B":
		return parseByteSize * B, size
	case "KB":
		return parseByteSize * KB, size
	case "MB":
		return parseByteSize * MB, size
	case "GB":
		return parseByteSize * GB, size
	case "TB":
		return parseByteSize * TB, size
	case "PB":
		return parseByteSize * PB, size
	case "EB":
		return parseByteSize * EB, size
	default:
	}
	return
}

func GetValueSize(val any) int64 {
	// todo make it better
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	log.Println("GetValueSize", val, size)
	return size
}
