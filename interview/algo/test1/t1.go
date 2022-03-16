package test1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

// 3    1.01<=ele<=3  maxOneDay=3
// [1.01, 1.01, 1.01, 1.4, 2.4]
func findMinimumDays(durations []float32) int32 {
	// Write your code here
	arr := make([]float64, len(durations))
	for i, val := range durations {
		arr[i] = float64(val)
	}
	sort.Float64s(arr)
	i, j := 0, 0
	lenNum := len(arr)
	maxHour := float64(3)
	dayNum := 0
	for i < lenNum {
		j = i + 1
		curAdd := arr[i] + arr[j]
		if curAdd <= maxHour {
			i += 2
			dayNum += 1
		} else {
			i += 1
			dayNum += 2
		}
	}

	return int32(dayNum)
}

func getPhoneNumbers(country string, phoneNumber string) string {
	resp, err := queryCodes(country)
	if err != nil {
		fmt.Printf("[err] %v \n", err)
		return phoneNumber
	}
	if len(resp.Data) > 0 {
		callingCodes := resp.Data[0].CallingCodes
		if len(callingCodes) > 0 {
			return fmt.Sprintf("+%s %s", callingCodes[0], phoneNumber)
		}
	}

	return phoneNumber
}

type CodeResp struct {
	Page int                  `json:"page"`
	Data []CodeRespStructData `json:"data"`
}

type CodeRespStructData struct {
	Name         string   `json:"name"`
	Capital      string   `json:"capital"`
	CallingCodes []string `json:"callingCodes"`
}

func queryCodes(countryName string) (*CodeResp, error) {
	countryName = url.QueryEscape(countryName)
	r, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/countries?name=%s", countryName))
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	resp := CodeResp{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

func canBeEqualized(firstStrings []string, secondStrings []string) []string {
	// Write your code here
	res := make([]string, len(firstStrings))
	for i, s1 := range firstStrings {
		map1 := make(map[rune]struct{}, 0)
		map2 := make(map[rune]struct{}, 0)
		for _, c := range s1 {
			map1[c] = struct{}{}
		}
		s2 := secondStrings[i]
		for _, c := range s2 {
			map2[c] = struct{}{}
		}
		isOk := "YES"
		diffCount := 0
		for c1, _ := range map1 {
			if _, ok := map2[c1]; !ok {
				diffCount += 1
			}
			if diffCount > 1 {
				isOk = "NO"
				break
			}
		}
		res = append(res, isOk)
	}

	return res
}
