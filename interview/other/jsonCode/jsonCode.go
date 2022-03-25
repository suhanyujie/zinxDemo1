package jsonCode

/**
## json 的序列化和反序列化问题
在很久之前，在开发公司业务的时候，遇到一个 json 序列化的坑。虽然这个问题现在看来很简单，但是当时困扰了一段时间。

在 Go 中我们可以通过标准库中的 Marshal 和 Unmarshal 对数据进行序列化和反序列化：

```go
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
```

当然，因为标准库中 `encoding/json` 包的速度慢，可以替换成 `github.com/json-iterator/go` 进行 json 的序列化和反序列化：

```go
import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"testing"
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
```

毫无意外的，这段测试运行通过。但是在公司内部使用反序列化时，结果却不一致：

```go
arr := []int64{21192121, 21192122, 21192123}
byte1, err := json2.Marshal(arr)
if err != nil {
	t.Error(err)
	return
}
t.Log(string(byte1)) // 输出：["21192121","21192122","21192123"]
```

假设 json2 是我们自己定义的包。调用内部的 Marshal 函数后，输出结构 `["21192121","21192122","21192123"]`。因为这个值会被存入到数据库中，每次查看数据库，都看到这样的数据，我总以为是前端传过来的参数没做校验，导致后端就直接存入到数据库中，当修改完后，依然有数据是这一种字符串数组的形式，我甚至以为又是其他的地方没有做好入参转换。在那几天，我一点都没有怀疑是公司的内部包的问题（一个如此常用的包，怎么会有问题？）

不久后，我尝试在自己本地写 test，这才怀疑到是公司的公共包的问题。为了进一步验证我的猜想，我这才写了同样的 test case，分别在本地，以及使用公司的包执行，这才发现了它们直接的区别，于是进入公共包中查看源码，发现了问题所在 —— 自定义 int64 类型的编码问题。

通过排查，主要是这一段注册扩展导致的问题：

```go
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

json.RegisterExtension(&int64Extension{})
```

而如果去除 int64 的扩展注册，则执行正常：

```go
arr := []int64{21192121, 21192122, 21192123}
json2, err := json2.Marshal(arr)
if err != nil {
	t.Error(err)
	return
}
t.Log(string(json2)) // 输出 [21192121,21192122,21192123]
```

## 总结
回想这段经历，感觉自己定位问题还是太草率了，仅凭个人经验和直觉，导致错误的定位。而如果重来一次，我觉得可以用如下的思路来排查：

* 遇到问题时，可以从接口入参开始测试（如果你知道问题所在那就另当别论了），通过 debug 的方式确定问题所在。
* 如果一个地方你觉得不容易出错，也需要怀疑一下，通过写 test case 进行排除，缩小排查范围。
* 多学习一些库的源码，标准库源码。最好可以自己复刻，这样理解会更加深刻！

*/
