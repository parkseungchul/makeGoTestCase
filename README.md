## 고랭 소스코드를 분석하여 테스트 코드 파일을 만들어주는 프로그램

#####  주요 문법: ast,  정규화 표현을 사용 

##### 테스트 코드는 아래와 같으며 입력 파라미터는 $GOPATH 이후에 패키지를 포함한 파일명만 필요

![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/img/test.PNG?raw=true)


### 대상 파일 함수 [testCase.go](src/testing/make/testdata/testCaset.go)
![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/img/as-is.PNG?raw=true)

### 자동으로 생성 된 테스트 파일의 함수 [testCase_temp_test.go](src/testing/make/testdata/testCase_temp_test.go)
![screenshot](https://github.com/parkseungchul/makeGoTestCase/blob/master/src/testing/make/testdata/img/to-be.PNG?raw=true)


### 다음과 같은 조건으로 테스트 코드를 생성

#### 1. 메인 파일은 테스트 코드를 생성할 수 없음

#### 2. 입력 파일 값의 같은 디렉토리에 ${file_name}_temp.go 파일로 생성

#### 3. 주석에 아래와 같은 규칙을 넣게 되면 인 아웃을 자동으로 맵핑하여 테스트 코드를 생성

##### 규칙:{input}=>{output} 

##### 일반함수는 {1,2}=>{3,"12"}

##### 리시버 함수는 TestStruct{1}.{}=>{1}

##### 리턴 값이 없을 경우는 {}=>{}

##### 반드시 => 있어야 정상적으로 파싱

#### 4. 주석안에 [-] 있다면 테스트 케이스를 만들지 않음

##### -release note---------------------------------------------

##### 일부 인아웃 표현식 정상적으로 파싱 안되는 오류 수정 2017 11/23

##### 복수개의 인아웃 정보가 있을 경우 테스트 케이스 복수로 생성  2017 11/23

##### ----------------------------------------------









