package testdata

import "fmt"

// sum1 function
// {1,2}=>{3,"12"}
func sum1(a int, b int)( int, string) {
	return a + b, fmt.Sprintf("%d%d",a, b)
}

// sum2 function
// {1,2}=>{ TestStruct{sum:3} }
func sum2(a int, b int)TestStruct{
	return TestStruct{a+b}
}

// sum3 function
// {1,2}=>{}
func sum3(a int, b int){

}

// receiveFun1 method
// TestStruct{1}.{}=>{1}
func (t TestStruct)receiveFun1()int{
	return t.sum
}

// receive function
// TestStruct{1}.{1}=>{2}
func (t TestStruct)receiveFun2(a int)int{
	return t.sum + a
}

// receive function
// TestStruct{1}.{1,1}=>{3}
func (t TestStruct)receiveFun3(a int, b int)int{
	return t.sum + a + b
}

// [-] This symbol, except for creating a test case
func subtract1(a int, b int)(int){
	return a - b
}

// {2,2}=>{0}
// {4,4}=>{0}
func subtract2(a int, b int)(int){
	return a - b
}

type TestStruct struct {
	sum int
}