package trapingRain

import "testing"

func TestTrap1(t *testing.T) {
	heightArr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(heightArr)
	expected := 6
	if res != expected {
		t.Error("res is error[1]")
		return
	}
	t.Log("--ok, end--")
}

func TestTrapDp(t *testing.T) {
	heightArr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trapDp(heightArr)
	expected := 6
	if res != expected {
		t.Error("res is error[1]")
		return
	}
	t.Log("--ok, end--")
}
