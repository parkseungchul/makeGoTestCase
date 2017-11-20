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
    // a := sum2(  )
    
}

func Test_receiveFun(t *testing.T){
    // receive function
    // TestStruct{1}.{}=>{1}
    a := TestStruct{1}.receiveFun( )
    if !(a == 1 ){
        t.Error("Error receiveFun")
    }
}

func Test_subtract2(t *testing.T){
    // {2,2}=>{0}
    a := subtract2( 2,2 )
    if !(a == 0 ){
        t.Error("Error subtract2")
    }
}

