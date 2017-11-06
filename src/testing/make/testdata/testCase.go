package testdata

type TestStruct struct {
	str1 string
}

// sum1 function
func sum1(a int, b int)( int, string) {
	return a + b, ""
}

// receive function
func (t TestStruct)receiveFun()string{
	return t.str1
}

// [-] subtract two numbers
func subtract1(a int, b int)(int){
	return a - b
}