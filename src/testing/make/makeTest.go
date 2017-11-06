package make

import (
	"os"
	"fmt"
	"go/token"
	"go/parser"
	"go/ast"
	"strings"
	"errors"
)

const (
	preComment string = "// "
	divsion string =", "
	funDivsion string = "\n\n"
	expectKey ="[-]"
)

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
}

func Maker(inputName string)(err error){

	contents, pathName, fileName, err := makeTest(inputName)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = makeFile(contents, pathName, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func makeFile(contents, pathName, fileName string)(err error){
	newFileName := strings.Replace(pathName + fileName,".go", "_temp_test.go", 1)

	file1, _ :=os.Create(newFileName)
	defer file1.Close()
	fmt.Fprint(file1, contents)
	fmt.Print("Success Create FileName:" ,newFileName)

	return
}

// isName, isType
func (f FieldInfo)Gender(isName bool, isType bool)string{
	if isType && isName {
		if f.fName == "" {
			return f.fType
		}
		return f.fName + " " + f.fType
	}else if isType && !isName{
		return f.fType
	}else if !isType && isName{
		return f.fName
	}else {
		return ""
	}
}

func (f FuncInfo)Gender()(parser string){
	testCaseName := ""
	funcName :=""
	if f.fType == "F" {
		testCaseName = "_"+f.name
		funcName = f.name
	}else{
		testCaseName = "_"+f.receive.fType+"_"+f.name
		funcName =f.receive.fType+"{}."+ f.name
	}

parser =
`func Test`+ testCaseName+`(t *testing.T){
#genComment#
#genBody#
}`

	parser = strings.Replace(parser, "#genComment#", f.comments,1)

	genBody :=  getSpace(1) + preComment + funcName +"(#genParameter#) "+preComment +"returns: #genResult#"
    genParameter := ""
	if f.parameters != nil {
		for i:=0; i<len(f.parameters); i++{
			if  i == 0 {
				genParameter = f.parameters[i].Gender(true, true)
			}else {
				genParameter = genParameter + divsion + f.parameters[i].Gender(true, true)
			}
		}
	}
	genBody = strings.Replace(genBody, "#genParameter#",genParameter,1)

	genResult := ""
	if f.results != nil {
		for i:=0; i<len(f.results); i++{
			if  i == 0 {
				genResult = f.results[i].Gender(false, true)
			}else {
				genResult = genResult + divsion + f.results[i].Gender(false, true)
			}
		}
	}

	genBody = strings.Replace(genBody, "#genResult#",genResult,1)

	parser = strings.Replace(parser, "#genBody#", genBody, 1)
	return
}


func getSpace(level int)(result string){
	for i :=0;  i < level; i++ {
		for j := 0; j <4; j++ {
			result = result + " "
		}
	}
	return
}


func makeTest(inutName string)(contents, pathName, fileName string, err error){
	packageName, pathName, fileName , err := envCheck(inutName)
	if err != nil {
		return
	}

	fset := token.NewFileSet()
	nodes, err := parser.ParseFile(fset, pathName + fileName, nil, parser.ParseComments)
	if err != nil {
		return
	}

   funcInfos := make([]FuncInfo, 0)

	for _, node := range nodes.Decls {
		switch n := node.(type) {
		case *ast.FuncDecl:

			funcInfo := FuncInfo{}

			if (n.Recv != nil) {
				funcInfo.fType = "R"
				for _ ,recv := range n.Recv.List {
					//fmt.Println("recv",recv.Names ,recv.Type )
					for _, value := range recv.Names{
						funcInfo.receive = FieldInfo{fmt.Sprintf("%s",recv.Type),value.Name}
					}
				}
			}else{
				funcInfo.fType = "F"
			}

			funcInfo.name = n.Name.Name
			if  n.Doc.Text() != "" {
				comentsConvert := n.Doc.Text()[0:strings.LastIndex(n.Doc.Text(), "\n")]
				comentsConvert = getSpace(1) + preComment + comentsConvert
				comentsConvert = strings.Replace(comentsConvert, "\n","\n"+getSpace(1)+preComment, -1)
				funcInfo.comments = comentsConvert
			}else{
				funcInfo.comments = getSpace(1) + preComment
			}

			if strings.ContainsAny(funcInfo.comments, expectKey) {
				continue
			}


			if len(n.Type.Params.List) != 0 {
				for _, getPara := range n.Type.Params.List {
					//fmt.Println("parameters" ,getPara.Names,getPara.Names[0].Name, getPara.Type)
					for _, value := range getPara.Names {
						//fmt.Println("parameters" , name, getPara.Type)
						if funcInfo.parameters == nil {
							funcInfo.parameters = make([]FieldInfo, 0)
							funcInfo.parameters = append(funcInfo.parameters, FieldInfo{fmt.Sprintf("%s",getPara.Type),value.Name})
						}else{
							funcInfo.parameters = append(funcInfo.parameters, FieldInfo{fmt.Sprintf("%s", getPara.Type),value.Name})
						}
					}
				}
			}
			if n.Type.Results != nil  {
				for _,getResult :=  range n.Type.Results.List {
					//fmt.Println("returns", getResult.Type, getResult.Names, len(getResult.Names))
					if len(getResult.Names) != 0 {
						for _, value := range getResult.Names {
							if funcInfo.results == nil {
								funcInfo.results = make([]FieldInfo, 0)
								funcInfo.results = append(funcInfo.results, FieldInfo{fmt.Sprintf("%s",getResult.Type),value.Name})
							}else{
								funcInfo.results = append(funcInfo.results, FieldInfo{fmt.Sprintf("%s",getResult.Type),value.Name})
							}
						}
					}else{
						if funcInfo.results == nil {
							funcInfo.results = make([]FieldInfo, 0)
							funcInfo.results = append(funcInfo.results, FieldInfo{fmt.Sprintf("%s",getResult.Type),""})
						}else{
							funcInfo.results = append(funcInfo.results, FieldInfo{fmt.Sprintf("%s",getResult.Type),""})
						}
					}
				}
			}
			funcInfos = append(funcInfos, funcInfo)
		default:
		}
	}


	body := ""
	for index, funcInfo := range funcInfos {
		if (index == len(funcInfos)-1){
			body = body + funcInfo.Gender()
		}else{
			body = body + funcInfo.Gender() + funDivsion
		}
	}

contents =
`package #packageName#

import "testing"

#body#`

	contents = strings.Replace(contents,"#packageName#",packageName, 1)
	contents = strings.Replace(contents,"#body#",body, 1)

	return
}

//check exist a file
func envCheck(inputName string)(packageName, pathName, fileName string, err error){
	inputName = os.Getenv("GOPATH")+"/src/"+ inputName

	// is checked to exist file
	if _, err = os.Stat(inputName); os.IsNotExist(err) {
		return
	}

	// is checked golang file
	if !strings.HasSuffix(inputName, ".go") {
		err = errors.New("this is not golang file")
		return
	 }

	 // is checked that golang file include package is main
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, inputName, nil, parser.PackageClauseOnly)
	if err != nil {
		return
	}
	if astFile.Name == nil {
		return "","","", fmt.Errorf("no package name found")
	}else {
		if strings.Compare(strings.Trim(astFile.Name.Name," "),"main") == 0 {
			return  "","","",fmt.Errorf("main package is not makes test-case ")
		}
		packageName = strings.Trim(astFile.Name.Name," ")
	}

	names := strings.Split(inputName, "/")
	fileName = names[len(names)-1]
	pathName = strings.Replace(inputName, fileName, "", 1)
	return packageName, pathName, fileName , err
}


func (f FuncInfo)Print(){
	fmt.Println("===========")
	fmt.Print(f.fType)
	fmt.Println(f.name)

	if f.parameters != nil {
		fmt.Println("parameters", len(f.parameters))
		for _, value := range f.parameters {
			fmt.Println(value)
		}
	}

	if f.results != nil {
		fmt.Println("results", len(f.results))
		for _, value := range f.results {
			fmt.Println(value)
		}
	}

	if f.comments != "" {
		fmt.Println("comments")
		fmt.Print(f.comments)
	}
	fmt.Println("\n===========")
}
