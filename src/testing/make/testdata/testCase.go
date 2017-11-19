package testdata

import "fmt"

type TestStruct struct {
	sum int
}

// sum1 function
// {1,2}=>{3,"12"}
func sum1(a int, b int)( int, string) {

	return a + b, fmt.Sprintf("%s%s",a, b)
}

func sum2(a int, b int)TestStruct{
	return TestStruct{a+b}
}


// receive function
// TestStruct{1}.{}=>{1}
func (t TestStruct)receiveFun()int{
	return t.sum
}

// [-] subtract two numbers - this expects to make a test-case
func subtract1(a int, b int)(int){
	return a - b
}

// {2,2}=>{0}
func subtract2(a int, b int)(int){
	return a - b
}