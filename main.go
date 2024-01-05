package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type FileInfo struct {
	File   string `json:"file"`   // File name
	Line   int    `json:"line"`   // Line number where error_reporting was found
	Column int    `json:"column"` // Position of the character where error_reporting was found
	Value  string `json:"value"`  // Value of error_reporting
}

type Json struct {
	File   string `json:"file"`   // File name
	Line   string `json:"line"`   // Line number where error_reporting was found
	Column string `json:"column"` // Position of the character where error_reporting was found
	Value  string `json:"value"`  // Display of PHP's predefined error level and whether it's enabled/disabled
}

// PHP error levels and their values
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
	// Check os.Args to ensure arguments are present
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go '[JSON data]'")
		os.Exit(1)
	}

	// Receive the first argument as JSON data
	jsonData := os.Args[1]

	// Initialize the FileInfo struct
	var fileInfo FileInfo

	// Unmarshal the JSON data into the FileInfo struct
	err := json.Unmarshal([]byte(jsonData), &fileInfo)
	if err != nil {
		// Panic if an error occurs during unmarshalling
		fmt.Println(err)
		return
	}

	// Define regular expression pattern
	pattern := "[0-9]+"
	// Compile the regular expression pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Extract the numeric part from the string
	s := re.FindString(fileInfo.Value)
	if s == "" {
		return
	}
	// Calculate the corresponding level from the set integer value
	i, _ := strconv.Atoi(s)
	errorLevels := getErrorLevels(i)

	j := Json{
		File:   fileInfo.File,
		Line:   strconv.Itoa(fileInfo.Line),
		Column: strconv.Itoa(fileInfo.Column),
		Value:  errorLevels,
	}
	// Encode to JSON
	result, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output the JSON data
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
		result string
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
	// Trim the trailing comma and space
	result = strings.TrimSuffix(result, ", ")

	return result
}
