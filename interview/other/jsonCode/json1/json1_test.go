package json1

import (
	"reflect"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"github.com/modern-go/reflect2"

	jsoniter "github.com/json-iterator/go"
)

func TestJson1(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

func TestJsonForTime(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	s1 := `["2022-03-24 01:11:11"]`
	timeArr := make([]time.Time, 0)
	json.Unmarshal([]byte(s1), &timeArr)
	t.Log(timeArr)
}

const AppDateFormat = "2006-01-02"
const AppTimeFormat = "2006-01-02 15:04:05"
const BlankString = ""

var loc, _ = time.LoadLocation("Local")

var json1 = jsoniter.ConfigCompatibleWithStandardLibrary

type TimeDecoder struct{}
type TimeEncoder struct {
	precision time.Duration
}

func (codec *TimeEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.IsZero()
}

func (codec *TimeEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	if !ts.IsZero() {
		timestamp := ts.Unix()
		tm := time.Unix(timestamp, 0)
		format := tm.Format(AppTimeFormat)
		stream.WriteString(format)
	} else {
		mayBlank, _ := time.Parse(AppTimeFormat, BlankString)
		stream.WriteString(mayBlank.Format(AppTimeFormat))
	}
}

func init() {
	jsoniter.RegisterTypeDecoder("time.Time", &TimeDecoder{})
	jsoniter.RegisterTypeEncoder("time.Time", &TimeEncoder{})
	//jsoniter.RegisterTypeEncoder("int64", &int64codec{})
	json2.RegisterExtension(&int64Extension{})
}

func (td *TimeDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()
	mayBlank, _ := time.Parse(AppTimeFormat, str)
	now, err := time.ParseInLocation(AppTimeFormat, str, loc)

	if err != nil {
		*((*time.Time)(ptr)) = time.Unix(0, 0)
	} else if mayBlank.IsZero() {
		*((*time.Time)(ptr)) = mayBlank
	} else {
		*((*time.Time)(ptr)) = now
	}
}

func TestJsonForTime2(t *testing.T) {
	s1 := `["2022-03-24 01:11:11"]`
	timeArr := make([]time.Time, 0)
	json.Unmarshal([]byte(s1), &timeArr)
	t.Log(timeArr)
}

type int64codec struct{}

func (codec *int64codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*int64)(ptr)) == 0
}

func (codec *int64codec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteString(strconv.FormatInt(*((*int64)(ptr)), 10))
}

func (codec *int64codec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	data := iter.Read()
	if in, ok := data.(string); ok {
		v, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			*((*int64)(ptr)) = 0
		} else {
			*((*int64)(ptr)) = v
		}
	} else if in, ok := data.(int64); ok {
		*((*int64)(ptr)) = in
	} else {
		v, err := strconv.ParseInt(iter.ReadString(), 10, 64)
		if err != nil {
			*((*int64)(ptr)) = 0
		} else {
			*((*int64)(ptr)) = v
		}
	}
}

type int64Extension struct{}

func (extension *int64Extension) CreateEncoder(typ reflect2.Type) jsoniter.ValEncoder {
	if typ == reflect2.TypeOfPtr((*int64)(nil)).Elem() {
		return &int64codec{}
	}
	return nil
}
func (extension *int64Extension) CreateDecoder(typ reflect2.Type) jsoniter.ValDecoder {
	if typ == reflect2.TypeOfPtr((*int64)(nil)).Elem() {
		return &int64codec{}
	}
	return nil
}

// UpdateStructDescriptor No-op
func (extension *int64Extension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (extension *int64Extension) CreateMapKeyDecoder(typ reflect2.Type) jsoniter.ValDecoder {
	if typ == reflect2.TypeOfPtr((*int64)(nil)).Elem() {
		return &int64codec{}
	}
	return nil
}

// CreateMapKeyEncoder No-op
func (extension *int64Extension) CreateMapKeyEncoder(typ reflect2.Type) jsoniter.ValEncoder {
	if typ == reflect2.TypeOfPtr((*int64)(nil)).Elem() {
		return &int64codec{}
	}
	return nil
}

// DecorateDecoder No-op
func (extension *int64Extension) DecorateDecoder(typ reflect2.Type, decoder jsoniter.ValDecoder) jsoniter.ValDecoder {
	return decoder
}

// DecorateEncoder No-op
func (extension *int64Extension) DecorateEncoder(typ reflect2.Type, encoder jsoniter.ValEncoder) jsoniter.ValEncoder {
	return encoder
}

var json2 = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

func TestJsonForInt64(t *testing.T) {
	arr := []int64{21192121, 21192122, 21192123}
	byte1, err := json2.Marshal(arr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(byte1))
}
