# Go tutorial
We tought that the best way to learn go is to go into the documentation and check out what is the most important things to retrieve from it. To help us in this matter we found a spectacular tutorial called tour go. This is the website: https://go.dev/tour/list. This tour gives all the basic concepts that we need to move forward in go, so as the exercises and nedded resources to move forward.

# Notes

## Welcome
go run (name of the program).go , runs the program

go build (name of the program).go , compilees the program

to execute then you use ./(name of the exec generated)

## Packages
- Program always start into a package called main
- When we import a package, by convention the name of the package is always the last element of the import path
```
Ex: import ("math/rand")
rand.Intn(10)
```

## Imports
- We can either do imports like this:
 ```
 import "fmt"
 import "math"
 ```
- Or like this:
 ```
 import(
    "fmt",
    "math
 )
 ```
Note that with package fmt we have access to printf from C and also println from java for example

## Exported names
 In go, in order to some component be exported, it needs to start with capital letter.
```
Ex: 
math.Pi
fmt.Println
``` 
All the components need to be capital letter in order to be exported. If there is something we do not want to be exported but to become internal, we create that component with lower-case.
```
func test(){...}
func Testa(){
    test();
}
We can use Testa but we cannot use test
```

## Functions
- The type declaration of return or type of the variables comes always in the end with the exception of pointers
```
x int;
func add(x int,y int) int
*p
```
- We can instance the variables all at once like this:
```
x,y int;
```

## Multiple Results
- In go we can return multiple results like so:
  ```
  func (x,y string) (string,string){
    return y,x
  }
  func main(){
    a,b:= swap("Hello","World")
  }
  ```

## Named return values
- In go, we can return simply variables that we already declared in the return statement
  ```
  func test(x,y int) (z,v int){
    z = x/2
    v = y+x
    return
  }
  func main(){
    fmt.Println(test(17))
  }
  ``` 
- This is named "naked" return, it is good practise to only use this in short functions

## Variables
- We use "var" when we need to declare explicity a certain var like so:
 ```
 var count int
 ```
- Cases where we do not use the var:
 ```
 count := ..
 ```
- Note in the exercise that when we do not specify the value of bool, it turns out to be "false"

## Variables with initializers
- We can initialize the variables with and without type at the same time like so:
  ```
  var i,j int = 1,2
  var c,python,java = true, false ,"no!"
  ```

## Short variable declarations
- We can declare variables like so:
  ```
   k := 3
  ```

## Basic Types
- bool
- string
- int
- int8
- int16
- int32
- int64
- byte (same as int8)
- rune (same as int32)
- float32
- float64
- complex64
- complex128
  
## Zero Values
- When you dont initialize a certain variable, it is given a default value
  ```
    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.
  ```

## Type conversions
- We can convert numbers like so:
  ```
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)
  ```

## Type Inference
- We can infere the type as so:
  ```
    i := 42           // int
    f := 3.142        // float64
    g := 0.867 + 0.5i // complex128
  ```

## Constants
- Constants are like variables
- Can be anything
- Cannot be declared with :=

## Numeric Constants
- High precision values
- An untyped constant takes the type needed by its context

## Final notes of basic
">>n" operator removes n zeros from the binary representation of a certain number

"<<n" operator adds n zeros from the binary representation of a certain number

## For
- Normal for is like this
  ```
  for i:=0 ; i<10; i++ {
    sum+=i
  }
  ```

## For continued
- Another way to make a for
  ```
  for ; sum < 1000; {
    sum += sum
  }
  ```

## While in go
- The while in go uses for
```
for sum < 1000{
    sum += sum
}
```

## Forever Loop 
```
for{

}
```

## If Statement
- The if statement does not need parentheses
```
if x {}
```

## If with short statement
- We can make a if in a short way but the value will only be accessible inside of the if
  ```
  if v:= x*x; v < y {
    return v
  }
  ```
- This "v" can only be accessed inside of the {}

## If and else
```
if v:=x*x; v < 10{
    return v
}else{
    //I CAN USE V here
    return lim
}
// I cannot use v here
```

## Switch
```
switch(x){
    case 1:
    fmt.Println("lul")
    default:
    fmt.Println("Hallo")
}
```

## Switch with no condition
```
switch{
    case x < 12:
    ...
    case x > 12:
    ...
    default:
    ...
}
```

## Defer
- This is a keyword that would a certain task until the neighboor task get executed
- If multiple defers are stacked, they will run in order. The oldest will be the last to run.

## Pointers
- Pointers hold the memory address of a value
  ```
  var p *int
  ```
- This is how we pass the address to the pointer
  ```
  i := 42 
  p = &i
  ```
- This is how we get the value:
  ```
  fmt.Println(*p) // PRINT THE VALUE
  *p=21 //SET THE VALUE
  ```

## Structs
```
type Person struct{
    name string,
    familyName string
}
```

## Structs Fields
```
type Vertex struct{
    X int
    Y int
}
func main(){
    v:=Vertex{1,2}
    v.X=4
}
```

## Pointers to Structs
The proccess is the same as other variable
```
v := Vertex{1, 2}
p := &v
```

## Struct Literals
```
type Vertex struct{
    Name string
    Age int
    Email string
}
```