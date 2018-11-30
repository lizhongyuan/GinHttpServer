package main


type MyInterface interface{
	Print()
}

func TestFunc(x MyInterface) {}

type MyStruct struct {}

func (me MyStruct) Print() {}

func main() {
	var me MyStruct
	TestFunc(me)
}


type HybridInterface interface {
	Show()
}

type HybridStruct struct {

}
