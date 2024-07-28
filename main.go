package main

import (
	"memCache/util"
)

func main() {
	util.GetValueSize(1)
	util.GetValueSize(false)
	util.GetValueSize("你好你好你好你好你好你好")
	util.GetValueSize(map[string]string{
		"name": "fwf",
	})
	util.GetValueSize(map[string]map[string]string{
		"aaa": map[string]string{
			"bbb": "ccc",
		},
	})
}
