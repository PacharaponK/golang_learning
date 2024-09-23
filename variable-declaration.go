package main

import ("fmt")

const PI float32 = 3.14

func main() {  
    var name string = "Bom"  //type is string  
    var name2 string         //type is string  
    var name3 = "Earth"      //type is inferred  
    name2 = "Pupu"  
    age := 20                //type is inferred  
    
    fmt.Println(name)  
    fmt.Println(name2)  
    fmt.Println(name3)  
    fmt.Println(age)  
    fmt.Println(PI)
}
