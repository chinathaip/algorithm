package main

func getMinDeletions(s string) int32 {
	var count int32
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			count++
		}
	}
	return int32(len(s)) - count
}
