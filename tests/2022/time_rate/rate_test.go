package time_rate

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRate1(t *testing.T) {
	ctx := context.Background()
	r := rate.NewLimiter(2, 10)
	isOk := r.AllowN(time.Now(), 12)
	t.Log(isOk)
	rs := r.ReserveN(time.Now(), 10)
	if !rs.OK() {
		t.Error("err1")
		return
	}
	t.Log(rs.Delay())
	time.Sleep(rs.Delay())
	ctx1, _ := context.WithTimeout(ctx, time.Second*4)
	for i := 0; i < 12; i++ {
		if err := r.Wait(ctx1); err != nil {
			t.Error(err)
			return
		}
		go func(i int) {
			fmt.Printf("[TestRate1] %d \n", i)
		}(i)
		//time.Sleep(time.Second)
	}
}
