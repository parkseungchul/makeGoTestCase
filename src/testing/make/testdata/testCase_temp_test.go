package testdata

import "testing"

func Test_sum1(t *testing.T){
    // sum1 function
    // {1,2}=>{3,"12"}
    a, b := sum1( 1,2 )
    if !(a == 3 && b == "12" ){
        t.Error("Error sum1")
    }
    
}

func Test_sum2(t *testing.T){
    // sum2 function
    // {1,2}=>{ TestStruct{sum:3} }
    a := sum2( 1,2 )
    if !(a == TestStruct{sum:3} ){
        t.Error("Error sum2")
    }
}

func Test_sum3(t *testing.T){
    // sum3 function
    // {1,2}=>{}
    sum3( 1,2 )
    
}

func Test_receiveFun1(t *testing.T){
    // receiveFun1 method
    // TestStruct{1}.{}=>{1}
    a := TestStruct{1}.receiveFun1( )
    if !(a == 1 ){
        t.Error("Error receiveFun1")
    }
    
}

func Test_receiveFun2(t *testing.T){
    // receive function
    // TestStruct{1}.{1}=>{2}
    a := TestStruct{1}.receiveFun2(1 )
    if !(a == 2 ){
        t.Error("Error receiveFun2")
    }
    
}

func Test_receiveFun3(t *testing.T){
    // receive function
    // TestStruct{1}.{1,1}=>{3}
    a := TestStruct{1}.receiveFun3(1,1 )
    if !(a == 3 ){
        t.Error("Error receiveFun3")
    }
}

func Test_subtract2(t *testing.T){
    // {2,2}=>{0}
    // {4,4}=>{0}
    a := subtract2( 2,2 )
    if !(a == 0 ){
        t.Error("Error subtract2")
    }
    a = subtract2( 4,4 )
    if !(a == 0 ){
        t.Error("Error subtract2")
    }
}

