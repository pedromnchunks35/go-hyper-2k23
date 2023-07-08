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

## Arrays 
- To init a array we do something like [n]<type>, where n is the length and type is the type of the array
 ```
 var a [10]int
 ```
- Limited length

## Slices
- A slice is dynamically-sized
```
var s [] int = primes[1:4]
a[1:4]
```

## Slices are like references to arrays
- Slice does not store any data, it just describes a determined section
- Changing his elements only change the elements of his underlying array

## Slice literals
- THis is a literal array
  ```
  [3]bool{true,true,false}
  ```
- This is a literal slice
  ```
  []bool{true,false,true,false}
  ```

## Slice defaults
These slices expressions are equivalent:
```
a[0:10]
a[:10]
a[0:]
a[:]
```

## Slice length and capacity
- Slice has a length and a capacity
- if you use [:2], it will create a new slice with the first 2 members (length 2) 
- if you use [2:], it will create a new slice with the members that are not the first 2 members

## Nil slices
- The zero value of a slice is nil
- Nil slice has a length and capacity of 0

## Create Slice with make
- Slices can be created with a built-in make func
```
a:= make([]int,5)
```
- This results in a slice of length 5

```
b:=make([]int,0,5)
```
- This results in a slice of length 0 but capacity 5

## Slices of slices
- A slice can contain another slice

## Appending to a slice
We can append a slice to another
```
s = append(s,1)
```
We need to notice that if we crete a slice from a huge file for example.. if we do not cut that file into a piece and still mention it as a whole, we will pretty much not let the garbadge collection to take out the full mentioned file. We could take the part that interest us and mention that part exclusively, making the garbadge collector let out the entire data:
```
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}
```

## Range
- Range is a form of loop
- We can iterate all over the members, get that value and also the index with it
  ```
  for i,v := range pow {
    fmt.Printf("2**%d=%d\n",i,v)
  }
  ```
- You can skip the index or the value using "_" or ommiting it
  ```
  for _,v := range pow{

  }
  for i,_ := range pow{

  }
  for i := range pow{

  }
  ```

## Maps
- Maps keys to values
- Zero value of map is nil
  ```
  var m map[string]int
  ```

## Map Literals
- You can put a map literals like this:
  ```
  var m = map[string]Person{
    "Pedro": Person{
      Name: "Pedro",
      Age: 12,
    },
    "Martins": Person{
      Name: "Martins",
      Age: 1
    }
  }
  ```

## Mutating Maps
Insert or update map
```
m[key] = elem
```
Retrieve an element
```
elem=m[key]
```
Delete an element
```
delete(m,key)
```
Test that a key is present with a two-value assignment (if key is in m then ok=true, else ok=false. Also case the ok is false the elem is the zero value of the element type)
```
elem,ok=m[key]
```

## Function values
We can pass function like values to functions
```
func compute(fc func(float64,float64)) float64{
  fc(3,4)
}
```

## Function closures
- It is a function that calls a variable ouside his body

## Methods
- Go does not have classes
- You can define methods on types
```
type Person struct{
  Name: string
  Age: int
}

func (p Person) Abs() float64{
  return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main(){
  p:= Person(Name: "Pedro", Age: 1)
  fmt.Println(p.Abs())
}
```
- You can implement it in a normal way like this
  ```
  func Abs(p Person)float64{
    ...
  }
  ```

## Methods continued
- You can also aggregate methods to a variable like int for example
  ```
  type teste int
  func (t teste) roda(){
    ...
  }
  ```

## Pointer receivers
- With pointer receivers you can change the value of the given type you used to ivoke the function
  ```
  type Teste int
  func (t *Teste) changeValue(){
    t=Teste(4)
  }
  ```
- We can do the same thing using simply functions
 ```
 type Teste int
 func changeValye(t *Teste){}
 ```

## Methods and pointer indirection
- We can both use the receiver or create a function just for it, but in case we do the function we need to pass the pointer
```
RECEIVER:
v.calc(2)
Func way:
calc(&v,2)
```
- Note that for pass the pointer the function needs to take a pointer, otherwise it will throw a error
- Note that if you dont pass a pointer, it will create a copy of the variable you are passing
  
## Choosing a value or pointer receiver
- Normally it is better to use pointer receiver but not both at the same time

## Interfaces
- It is a method signature
- It describes which methods it is implementing
- To implement a interface in go we do as so:
 ```
 type test interface{
  printValue()
 }

 var t test

 t = someStructure // SOME STRUCTURE IMPLEMENTS t
 ```

## Interfaces implemented implicity
```
type I interface{
  M()
}
var i I = someStructure{someProp}
i.M()
```

## Interface values
- With interface we can assign different types to a given function
- When we call the function present in the interface, that function will run according to his perspective, because different structs, with different types can have the same signature of function but the function itself is different
```
type Person struct{
  Name string
}
type Albert struct{
  Age int
}
type someVal interface{
  GetProp()
}
func (p Person) GetProp(){
  return p.Name
}
func (a Albert) GetProp(){
  return a.Age
}
var x someVal = Person{Name: "Pedro"}
var y someVal = Albert{Age: 2}
x.GetProp()
y.GetProp()
``` 
- Note that x returns a string on GetProp()
- Note that y returns a int on GetProp()
- When creating the method of that interface, we should implement a way out case it is nil that prop
- Running a method on a nil interface causes a run time error
- An empty interface can be used for when we dont know the type of a certain variable

## Type assertions (interface)
- We can and we should make verifications of type of a certain value before doing operations over it
```
var i interface{} = "hello"
s,ok := i.(string) // CASE IT WENT OK = true, else false
``` 

## Type Switches (interface)
- We can create a switch according to the type
  ```
  switch v:= i.(type){
    case T:
    // here v has type T
    case S:
    // here v has type S
    default:
    // no match; here v has the same type as i
  }
  ``` 

## Stringers
- This is the .toString(), that we have in java
  ```
  type Person struct{
    Name string
  }
  func(p Person) String() string{
    return "THIS WILL BE THE OUTPUT WHATEVER THE NAME VALUE"
  }
  ```

## Errors
- This is how a error occur
  ```
  i,err := strconv.Atoi("42")
  if err != nil{
    ... //ERROR HANDLING HERE
  }
  //NO ERROR HANDLING
  ```
- Another way to handle errors
  ```
  type MyError struct{
    when time.time
    what string
  }
  func (e *MyError) Error() string{
    return fmt.Sprintf("at %v, %s",e.when,e.what)
  }
  func run() error{
    return &MyError{
      time.Now(),
      "it did not work"
    }
  }

  func main(){
    if err := run(); err != nil {
      fmt.Println(err)
    }
  }
  ```

## Reader
- It is something like a buffer, that we use to read a stream of data
- It is a slice of bytes, that we can specify the length and store content of a certain file.. We can loop all over a file using that slice of bytes and try to see if a file as something that we want

## Images
- We can define a image using interface image

## Type Parameters
- Go allows that we create a function that can take multiple types. But for that we need to use a state called comparable, so we can compare that multiple type
  ```
  func Index[T comparable](s[]T,x T) int{
    ...
  }
  ```

## Generic Types
- Go also implements generic types for a certain struct
  ```
  type List[T any] struct{
    ..
  }
  ```