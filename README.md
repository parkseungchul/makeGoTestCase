# 고랭으로 만든 코드를 테스트 코드로 만들어 주는 프로그램입니다. 

## 테스트 코드는 아래와 같으며 입력값은 $GOPATH 이후에 패키지를 파일명만 필요합니다.

![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/img/test.PNG?raw=true)


## 대상 파일 함수 [testCase.go](src/testing/make/testdata/testCaset.go)
![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/img/as-is.PNG?raw=true)

## 자동으로 생성 된 테스트 파일의 함수 [testCase_temp_test.go](src/testing/make/testdata/testCase_temp_test.go)
![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/img/to-be.PNG?raw=true)


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









