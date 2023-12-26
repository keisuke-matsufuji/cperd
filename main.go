package main

import (
	"fmt"
	"sort"
)

// PHPのエラーレベルとその値
var errorLevels = map[int]string{
    1:     "E_ERROR",
    2:     "E_WARNING",
    4:     "E_PARSE",
    8:     "E_NOTICE",
    16:    "E_CORE_ERROR",
    32:    "E_CORE_WARNING",
    64:    "E_COMPILE_ERROR",
    128:   "E_COMPILE_WARNING",
    256:   "E_USER_ERROR",
    512:   "E_USER_WARNING",
    1024:  "E_USER_NOTICE",
    2048:  "E_STRICT",
    4096:  "E_RECOVERABLE_ERROR",
    8192:  "E_DEPRECATED",
    16384: "E_USER_DEPRECATED",
}

func main() {
	// ここに`error_reporting`の整数値を設定する
	var errorReportingValue int = 6135 // 通常はE_ALLの値, あるいは異なる値で試すことができます。

	// 設定された整数値から対応するレベルを算出する
	activeErrorLevels := getErrorLevels(errorReportingValue)

	fmt.Println("Active error levels:")
	for _, level := range activeErrorLevels {
		fmt.Println(level)
	}
}

func getErrorLevels(value int) []string {
	result := []string{}
	var flagStr string

	// create a slice of keys
    var keys []int
    for k := range errorLevels {
        keys = append(keys, k)
    }

	// sort the keys
    sort.Ints(keys)

	for _, key := range keys {
		// using bitmasks to check if the level is active
		if value&key != 0 {
			flagStr = "+"
		} else {
			flagStr = "-"
		}
		str := fmt.Sprintf("%s: %s:%v", flagStr, errorLevels[key], key)

		result = append(result, str)
	}
	return result
}
