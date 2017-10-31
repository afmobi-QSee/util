/* 绝密 TOP SECRET, COPYRIGHT © AFMOBI GROUP */
package util

import (
	"testing"
	"fmt"
)

type TestStruct struct {
	A  string `NotNull:"true"`
	B  string `Min:"5"`
	B1 uint8 `Min:"5"`
	B2 float32 `Min:"5"`
	C  int `Min:"2" Max:"9"`
	D  string `SizeMin:"5"`
	E  string `SizeMin:"5" SizeMax:"10"`
	F  string `Pattern:"abc|def"`
	F1 string `Pattern:"[\u4e00-\u9fa5]"`
	H  TestStruct1 `InStruct:"true"`
	G []TestStruct1 `InStructArray:"true"`
}

type TestStruct1 struct {
	I string `NotNull:"true"`
	J int `Min:"5"`
}

//func (this *TestStruct)Customization()string{
//	if this.A == "" {
//		return "A is null"
//	}
//	return "ok"
//}

func newTestStruct() *TestStruct {
	wt := TestStruct1{I:"test", J:9}
	wt1 := TestStruct1{I:"xxxx", J:6}
	w := []TestStruct1{wt, wt1}
	return &TestStruct{H:wt,G:w}
}

func TestValidParams(t *testing.T) {
	testStruct := newTestStruct()

	result := ValidParams(testStruct)
	fmt.Println("result:", result)
	if result != "ok" {
		t.Error(result)
	}
	//test A NotNull
	testStruct.A = ""
	testError(t, testStruct, "NotNull")
	//test I innerNotNull
	testStruct = newTestStruct()
	testStruct.H.I = ""
	testError(t, testStruct, "innerNotNull")
	//test B Min string
	testStruct = newTestStruct()
	testStruct.B = "1"
	testError(t, testStruct, "Min string")
	//test B1 Min uint
	testStruct = newTestStruct()
	testStruct.B1 = 1
	testError(t, testStruct, "Min string")
	//test B2 Min uint
	testStruct = newTestStruct()
	testStruct.B2 = 1.1
	testError(t, testStruct, "Min uint")
	//test C Min Max
	testStruct = newTestStruct()
	testStruct.C = 20
	testError(t, testStruct, "Min Max")
	//test SizeMin
	testStruct = newTestStruct()
	testStruct.D = "1"
	testError(t, testStruct, "SizeMin")
	//test SizeMax
	testStruct = newTestStruct()
	testStruct.E = "123456789123"
	testError(t, testStruct, "SizeMax")
	//test Regexp
	testStruct = newTestStruct()
	testStruct.F = "email"
	testError(t, testStruct, "Regexp")
	testStruct = newTestStruct()
	testStruct.F1 = "1"
	testError(t, testStruct, "Regexp")
}

func testError(t *testing.T, testStruct *TestStruct, tag string) {
	result := ValidParams(testStruct)
	fmt.Println(result)
	if result == "ok" {
		t.Error(tag, "test failed")
	}
}
