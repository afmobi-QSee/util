# 校验参数valid.go

校验参数方法ValidParams
例子如下：

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
		H  *TestStruct1 `InStruct:"true"`
	}

	type TestStruct1 struct {
		I string `NotNull:"true"`
		J int `Min:"5"`
	}

该方法基于struct的tag属性，利用反射功能自动校验每个参数，一些常用tag设计如下

	NotNull   参数不为空,只支持string, true为校验,其他值或没有为不校验
	Min,Max   数字不大于或不小于某个值,支持int,uint,float,string类型, 后面的值为其限定范围
	SizeMin,SizeMax 字符长度不大于或不小于某个值,只支持string, 后面的值为其限定范围
	Pattern   用正则表达式匹配参数值,支持string,int类型,后面的值为校验的正则表达式
	InStruct  表示该参数为嵌套的struct,且需要校验里面的值,true为需要校验,其他值不用校验
