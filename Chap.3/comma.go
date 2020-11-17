package main

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:-3]) + ',' + s[-3:]
}
