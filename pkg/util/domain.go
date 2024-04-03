package util

import "strings"

func GetMainDomain(domain string) string {
	arr := strings.Split(domain, ".")
	if len(arr) < 2 {
		return ""
	}

	arr = arr[len(arr)-2:]
	var mainDomain string
	for idx, str := range arr {
		mainDomain += str
		if idx+1 != len(arr) {
			mainDomain += "."
		}
	}
	return mainDomain
}

func GetPrefix(domain string) string {
	arr := strings.Split(domain, ".")
	if len(arr) < 2 {
		return ""
	}

	arr = arr[:len(arr)-2]
	var prefix string
	for idx, str := range arr {
		prefix += str
		if idx+1 != len(arr) {
			prefix += "."
		}
	}
	return prefix
}
