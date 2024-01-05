package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type FileInfo struct {
	File   string `json:"file"`   // ファイル名
	Line   int    `json:"line"`   // error_reportingが見つかった行数
	Column int    `json:"column"` // error_reportingが見つかった文字の場所
	Value  string `json:"value"`  // error_reportingの値
}

type Json struct {
	File   string `json:"file"`   // ファイル名
	Line   string `json:"line"`   // error_reportingが見つかった行数
	Column string `json:"column"` // error_reportingが見つかった文字の場所
	Value  string `json:"value"`  // PHPの定義済みエラーレベルと有効/無効の表示
}

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
	// os.Argsをチェックして引数があることを確認します。
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go '[JSON data]'")
		os.Exit(1)
	}

	// 第1引数をJSONデータとして受け取ります。
	jsonData := os.Args[1]

	// FileInfoの構造体を初期化します。
	var fileInfo FileInfo

	// JSONデータをFileInfo構造体にアンマーシャルします。
	err := json.Unmarshal([]byte(jsonData), &fileInfo)
	if err != nil {
		// アンマーシャル中にエラーが発生した場合はpanicします。
		panic(err)
	}

	// 正規表現パターンを定義
	pattern := "[0-9]+"
	// 正規表現パターンをコンパイル
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	// 文字列から数字部分を抽出
	s := re.FindString(fileInfo.Value)
	if s == "" {
		return
	}
	// 設定された整数値から対応するレベルを算出する
	i, _ := strconv.Atoi(s)
	errorLevels := getErrorLevels(i)

	j := Json{
		File:   fileInfo.File,
		Line:   strconv.Itoa(fileInfo.Line),	
		Column: strconv.Itoa(fileInfo.Column),
		Value:  errorLevels,
	}
	// JSONにエンコード
	result, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
		return
	}

	// JSONデータを出力
	fmt.Println(string(result))
}

// func getErrorLevels(value int) string {

// 	var (
// 		result  string
// 		flagStr string
// 	)

// 	// create a slice of keys
// 	var keys []int
// 	for k := range errorLevels {
// 		keys = append(keys, k)
// 	}

// 	// sort the keys
// 	sort.Ints(keys)

// 	for _, key := range keys {
// 		// using bitmasks to check if the level is active
// 		if value&key != 0 {
// 			flagStr = "+"
// 		} else {
// 			flagStr = "-"
// 		}
// 		str := fmt.Sprintf("%s: %s:%v", flagStr, errorLevels[key], key)

// 		result += str + "\n"
// 	}
// 	return result
// }

func getErrorLevels(value int) string {

	var (
		result  string
	)

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
			result += fmt.Sprintf("%s, ", errorLevels[key])
		}
	}
	return result
}
