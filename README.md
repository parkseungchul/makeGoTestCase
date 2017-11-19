# This is to make simply test-case template in golang

## This program wants only file-name on $GOPATH

### git clone  https://github.com/parkseungchul/makeGoTestCase.git

### and then look at this [makeTest_test.go](src/testing/make/makeTestCase2_test.go)

### Execute this.  it will make test-case template about input golang file


### Detail function

#### 1. if you want to excepts test-case then you must include "[-]" in comment.

before 
// sum1 function
// {1,2}=>{3,"12"}
func sum1(a int, b int)( int, string) {
	return a + b, fmt.Sprintf("%s%s",a, b)
}

after 
func Test_sum1(t *testing.T){
    // sum1 function
    // {1,2}=>{3,"12"}
    a, b := sum1( 1,2 )
    if a == 3 && b == "12" {
    }
}