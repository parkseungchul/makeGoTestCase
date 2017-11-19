# This is to make simply test-case template in golang

# Before code
![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/as-is.PNG?raw=true)


# auto make test code 
![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/to-be.PNG?raw=true)


## This program wants only file-name on $GOPATH

### git clone  https://github.com/parkseungchul/makeGoTestCase.git

### and then look at this [makeTestCase2_test.go](src/testing/make/makeTestCase2_test.go)

### Execute this.  it will make test-case template about input golang file

## 고랭 파일을 스캔하여 테스트 코드 파일을 만들어주는 프로그램입니다.

### 테스트 코드는 다음과 같으며 $GOPATH 이후의 패키지 경로와 파일명을 입력값으로 필요합니다.
 
### [makeTestCase2_test.go](src/testing/make/makeTestCase2_test.go)

### 다음과 같은 조건으로 테스트 코드를 생성

#### 1. 메인 파일은 테스트 코드를 생성할 수 없습니다.

#### 2. 입력 파일 값의 같은 디렉토리에 ${file_name}_temp.go 파일로 생성

#### 3. 주석에 아래와 같은 규칙을 넣게 되면 인 아웃을 자동으로 맵핑하여 테스트 코드를 만들어 줍니다.

##### 규칙:{input}=>{output} 

##### 일반함수는 {1,2}=>{3,"12"}

##### 리시버 함수는 TestStruct{1}.{}=>{1}

##### 리턴 값이 없을 경우는 {}=>{}

##### 반드시 => 있어야 정상적으로 파싱을 할 수 있습니다.

#### 4. 주석안에 [-] 있다면 테스트 케이스를 만들지 않습니다.









