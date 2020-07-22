package main

import (
	"fmt"
	"reflect"
)

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1] //get parameters
	if a.Kind() != b.Kind() {
		fmt.Println("Not same type")
		return nil
	}

	//Check type of a variable.
	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

//Make a dynamic function
func makeDynamicSumFunc(funcPtr interface{}) {
	funcPtrInfo := reflect.ValueOf(funcPtr).Elem()     //Pointer can get actual value information as a function of Elem().
	Value := reflect.MakeFunc(funcPtrInfo.Type(), sum) //Create a function info of sum().

	funcPtrInfo.Set(Value) //Set up function info and then connect a function.
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeDynamicSumFunc(&intSum)
	makeDynamicSumFunc(&floatSum)
	makeDynamicSumFunc(&stringSum)

	fmt.Println(intSum(100, 300))
	fmt.Println(floatSum(3.14, 3.15))
	fmt.Println(stringSum("hello ", "golang world!!"))
}
