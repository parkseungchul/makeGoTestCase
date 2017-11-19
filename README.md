# This is to make simply test-case template in golang

## This program wants only file-name on $GOPATH

### git clone  https://github.com/parkseungchul/makeGoTestCase.git

### and then look at this [makeTest_test.go](src/testing/make/makeTestCase2_test.go)

### Execute this.  it will make test-case template about input golang file


### Detail function

#### 1. if you want to excepts test-case then you must include "[-]" in comment.


고랭 파일을 스캔하여 테스트 코드 파일을 만들어주는 프로그램입니다. 
메인 파일은 테스트 코드를 만들수 없으며 대상 파일의 같은 디렉토리에  ${file}_temp_test.go 파일로 생성됩니다. 
테스트에 필요한 입력값은 GOPATH 이후에 파일 경로와 파일명입니다.
아래와 같이 사용하시면 됩니다.
Maker2("/testing/make/testdata/testCase.go")

그리고 
주석에 특정 규칙을 넣게 되면 인 아웃을 자동으로 맵핑하여 테스트 코드를 만들어 줍니다.
{input}=>{output} 
예를 들어 
{1,2}=>{3,"12"}

리시버 함수의 경우는
TestStruct{1}.{}=>{1}






