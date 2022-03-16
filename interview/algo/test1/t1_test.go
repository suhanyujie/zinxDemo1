package test1

import "testing"

func TestFindMinimumDays(t *testing.T) {
	durations := []float32{1.01, 1.01, 1.01, 1.4, 2.4}
	res := findMinimumDays(durations)
	t.Log(res)
}

func TestQueryCode(t *testing.T) {
	// res := getPhoneNumbers("Puerto Rico", "564593986")
	res := getPhoneNumbers("Puerto Rico", "564593986")
	t.Log(res)
}

func TestCanBeEqualized(t *testing.T) {
	sArr1 := []string{"aaa", "abbc"}
	sArr2 := []string{"bbb", "ccc"}
	res := canBeEqualized(sArr1, sArr2)
	t.Log(res)
}
