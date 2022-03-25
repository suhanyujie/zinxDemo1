package jsonCode

import (
	"encoding/json"
	"reflect"
	"testing"
	"zinx_demo1/interview/other/jsonCode/jsonx"
)

func TestEncodeAndDecode1(t *testing.T) {
	// encode
	arr := []int64{21192121, 21192122, 21192123}
	resByte, _ := json.Marshal(arr)
	resStr := string(resByte)
	if resStr != "[21192121,21192122,21192123]" {
		t.Error("Marshal err001")
		return
	}
	// decode
	arr2 := make([]int64, 0)
	if err := json.Unmarshal([]byte(resStr), &arr2); err != nil {
		t.Error("Unmarshal error001")
		return
	}
	if !reflect.DeepEqual(arr2, arr) {
		t.Error("Unmarshal error002")
		return
	}

	t.Log("--end--")
}

func TestEncode1(t *testing.T) {
	arr := []int64{21192121, 21192122, 21192123}
	json2, err := jsonx.ToJson(arr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json2)
}

func TestJson1(t *testing.T) {
	var err error
	json1 := `["21192121", "21192122", "21192123", "21192124"]`
	arr := make([]int64, 0)
	err = jsonx.FromJson(json1, &arr)
	if err != nil {
		t.Error(err)
		return
	}
	json2, err := jsonx.ToJson(arr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json2)
}
