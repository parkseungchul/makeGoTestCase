package make

import (
	"os"
	"fmt"
	"strings"
	"errors"
	"go/parser"
	"regexp"
	"go/ast"
	"go/token"
)

const (
	excluded = "[-]"
	comments ="// "
	divide = "=>"
	returnVariable ="abcdefghijklmnopqrstuvwxyz"
)

func Maker2(input string) {
	//info := &Info{inputName:`testing/make/testdata/testCase.go`}
	info := &Info{inputName:input}
	err := envCheck(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	funInfos := ParserFun(info)


	fileConent := "package "+ info.packageName +"\n\n"
	fileConent = fileConent + `import "testing"`+ "\n\n"

	for _, funInfo := range funInfos {
		fileConent = fileConent + funInfo.makeTestCode()+"\n\n"
	}

	fmt.Println(fileConent)
	makeFile(fileConent,  info.testPath)

}

func makeFile(contents, testPath string)(err error){
	file1, _ :=os.Create(testPath)
	defer file1.Close()
	fmt.Fprint(file1, contents)
	fmt.Print("Success Create FileName:" ,testPath)

	return
}


type Info struct {
	inputName string
	filePath string
	testPath string
	packageName string
}

func (i Info)Print(){
	fmt.Println(i.inputName)
	fmt.Println(i.filePath)
	fmt.Print(i.testPath)
}

type FieldInfo struct {
	fType string
	fName string
}

type FuncInfo struct{
	fType string   // receive R, function F
	name string
	receive FieldInfo
	parameters []FieldInfo
	results  []FieldInfo
	comments string
	isValue bool  // exist in, out variable
	paraValue string
	resultValue string
	resultValues[] string

}

// 파라미터 변수를 리턴합니다.
func getReturnVariable(cnt int)(result string){
	if cnt == 0 {
		return
	}
	for i:=0; i<cnt; i++ {
		if i == cnt - 1 {
			result = result + fmt.Sprintf("%c",returnVariable[i])
		}else {
			result = result + fmt.Sprintf("%c, ",returnVariable[i])
		}
	}
	return result + " := "
}


func (f FuncInfo)makeTestCode()(result string){
	// 주석 만들기
	if f.comments != "" {
		f.comments = comments + f.comments
		f.comments = strings.Replace(f.comments, "\n", "\n"+comments, -1)
		result = f.comments + "\n"
	}

	if !f.isValue {
		result = result + comments
	}

   // 리턴 값 만들기
	result = result + getReturnVariable(len(f.results))

	if f.fType == "F" {
		result = result +  f.name+"( "+expectBlace(f.paraValue) +" )"
	} else {
		a := strings.Replace(f.paraValue, "}.{", "}."+f.name+"(", -1)
		i := strings.LastIndex(a, "}")
		result = result +  a[:i]+ expectBlace2(f.paraValue) +" )"
	}
	result = result + getResultCheck(f)

	result = strings.Replace(result,"\n", "\n"+getSpace(1 ,""), -1)
	result = getSpace(1, "")+ result


	result = "func Test_"+f.name+"(t *testing.T){\n"+ result
	result = result +"\n}"


	return
}

func getResultCheck(f FuncInfo)(result string){



	result = "\n"
	if f.resultValues == nil || len(f.resultValues)  == 0 {
		return
	}
	result = result + "if !("
	for i := 0; i < len(f.resultValues);  i++ {
		if i == 0 {
			result = result + fmt.Sprintf("%c",returnVariable[i]) + " == " +  f.resultValues[i]
		} else {
			result = result + " && " +fmt.Sprintf("%c",returnVariable[i]) + " == " +  f.resultValues[i]
		}
	}

	result = result + " ){\n"
	result = result + getSpace(1, "")+`t.Error("Error `+f.name+`")`+ "\n}"
	return
}


// 일반함수 블레이스 벗겨내기
func expectBlace(str string)(result string){
	expression := `[^{][a-zA-Z0-9\.\,\s\"\{\}]*[^}]`
	re := regexp.MustCompile(expression)
	results := re.FindAllString(str,1)
	if len(results) == 1 {
		return results[0]
	} else{
		return ""
	}
}

// 일반함수 블레이스 벗겨내기
func expectBlace2(str string)(result string){
	post := strings.LastIndex(str, "}")
	pre := strings.Index(str[:post],"}.{")
	return str[pre+len(divide)+1:post]

}


func getNameType(fieldInfo []FieldInfo, onlyType bool)(result string) {
	if fieldInfo == nil ||  len(fieldInfo) == 0 {
		return
	}else {
		for i := 0; i < len(fieldInfo); i++ {
			if ( len(fieldInfo) - 1 ) == i {  // check last
				if onlyType {
					result = result + fieldInfo[i].fType
				}else {
					result = result + fieldInfo[i].fName +" "+ fieldInfo[i].fType
				}
			}else{
				if onlyType {
					result = result + fieldInfo[i].fType + ","
				}else {
					result = result + fieldInfo[i].fName +" "+ fieldInfo[i].fType + ","
				}
			}
		}
	}
	return
}



// 1. Dose it exist file?
// 2. Is it golang file?
// 3. what does it include package name?
func envCheck(info *Info)(err error){
	// 1. Dose it exist file?
	filePath := os.Getenv("GOPATH")+"/src/"+ info.inputName
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return
	}
	// 2. Is it golang file?
	if !strings.HasSuffix(info.inputName, ".go") {
		err = errors.New("this is not golang file")
		return
	}
	// 3. what does it include package name?
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return
	}
	if astFile.Name == nil {
		err = errors.New("package name not found")
		return
	}
	if strings.Compare(strings.Trim(astFile.Name.Name," "),"main") == 0 {
		err = errors.New("package name dose not allow main")
		return
	}
	info.packageName = strings.Trim(astFile.Name.Name," ")


	//  analysis File path
	expression := `(\/?[a-zA-Z0-9\.]+)`
	re := regexp.MustCompile(expression)
	names := re.FindAllString(filePath, -1)

	// get file name
	fileName := names[len(names)-1]
	testName := strings.Replace(fileName,".go", "_temp_test.go", -1)
	testPath := strings.Replace(filePath, fileName, testName, -1)

	info.filePath = filePath
	info.testPath = testPath

	return
}

func ParserFun(info *Info)(funcInfos []FuncInfo){
	fset := token.NewFileSet()
	nodes, err := parser.ParseFile(fset, info.filePath, nil, parser.ParseComments)
	if err != nil {
		return
	}

	for _, node := range nodes.Decls {
		switch n := node.(type) {
		case *ast.FuncDecl:

			//  check excluded key word in comment
			// if it include excluded key word, it is skip
			// 주석문에서 테스트 케이스 만드는 스킵 대상인지 검사
			doc := ""
			if  n.Doc.Text() != "" {
				doc = n.Doc.Text()
				if !strings.Contains(doc, excluded) {
					re := regexp.MustCompile(`\n$`)
					doc = re.ReplaceAllString(doc, "")
				}else{
					continue
				}
			}
			funcInfo := FuncInfo{}
			funcInfo.comments = doc

			funType := ""
			if (n.Recv != nil) {
				funType = "R"
				for _ ,recv := range n.Recv.List {
					for _, value := range recv.Names{
						funcInfo.receive.fType =fmt.Sprintf("%s",recv.Type)
						funcInfo.receive.fName = value.Name
					}
				}
			}else {
				funType = "F"
			}
			funcInfo.fType = funType
			funcInfo.name = n.Name.Name


			lines := strings.Split(doc, "\n")
			for _, line := range lines {
				// 주석에서 인아웃 전문이 있는지 검사
				inOutExpress := `(([a-zA-Z0-9]+)?{[a-zA-Z0-9\.\,\s\"\{\}]*})=>({[a-zA-Z0-9\.\,\s\"\{\}]*})`
				isMatchLine, _:= regexp.MatchString(inOutExpress, line)
				if isMatchLine {
					funcInfo.isValue = true
					inOutRegexp := regexp.MustCompile(inOutExpress)
					ios := inOutRegexp.FindAllString(line, -1)

					// 정상적으로 값을 뽑아낼수 있었음. 여기서 => 분리해야 함
					inOutStr := ios[0]
					inOutIndex := strings.Index(inOutStr,divide)
					// fmt.Println("[PRE]",inOutStr[:inOutIndex],"[POST]",inOutStr[inOutIndex+len(divide):])
					funcInfo.paraValue = inOutStr[:inOutIndex]
					funcInfo.resultValue = inOutStr[inOutIndex+len(divide):]


					// 블래이스 안에 절대 분해 표현식
					dismemberExpress := `([a-zA-Z0-9]+{[a-zA-Z0-9\,\"\s]+})|("[a-zA-Z0-9\,\s]+")|(true)|(false)|([0-9\.]+)`
					dismemberRegexp := regexp.MustCompile(dismemberExpress)
					funcInfo.resultValues = dismemberRegexp.FindAllString(funcInfo.resultValue, -1)
					break
				}
			}

			paras := make([]FieldInfo, 0)
			if len(n.Type.Params.List) != 0 {
				for _, getPara := range n.Type.Params.List {
					for _, value := range getPara.Names {
						//fmt.Println(getPara.Type ,value.Name)
						paras = append(paras, FieldInfo{fType:fmt.Sprintf("%s",getPara.Type), fName:value.Name})
					}
				}
			}
			funcInfo.parameters = paras

			results := make([]FieldInfo, 0)
			if n.Type.Results != nil  {
				for _,getResult :=  range n.Type.Results.List {
					if len(getResult.Names) != 0 {
						for _, value := range getResult.Names {
							//fmt.Println(getResult.Type,value.Name)
							results = append(results, FieldInfo{fType:fmt.Sprintf("%s",getResult.Type), fName:value.Name})
						}
					}else{
						//fmt.Println(getResult.Type,"")
						results = append(results, FieldInfo{fType:fmt.Sprintf("%s",getResult.Type), fName:""})
					}
				}
			}
			funcInfo.results = results
			funcInfos = append(funcInfos, funcInfo )
		default:
		}
	}
	return
}


func getSpace(level int, msg string)(result string){
	for i :=0;  i < level; i++ {
		for j := 0; j <4; j++ {
			result = result + " "
		}
	}
	result = result + msg
	return
}