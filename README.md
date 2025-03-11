Golang is a open source programming langugae created and supported by the Google and also maintain by the google.

Go is a natively compiled language that converts the codebase into the machine level language.

Mainly used in the cloud computing and devops and in its projects.

Go's execution speed is more similar to java and C# and use much less memory compare to the java and C#.

Java and C++ takes hours to compile the code and If there are any bugs then If we find and fix in the seconds then also it takes hours to compile and build the new production versions.

Go takes second to compile and build it.

Docker and kubernetes are written in go language.

Major companies are using Golang like Google, twitch, uber, Medium, soundCloud, dailymotion.

Go is platform independent language like Java and due to that Go is very much faster than other language.

Go is a strongly typed and statically typed language and this is very good thing for developers. And most of the JavaScript developers are moving from JS to TS for statically typed language.	

In statically typed language it gives the bug or error at compile time not at the production time.

Go is Garbage Collected language and It has automated memory management techniques.

Go doesn't have JVM. Go build the binary of each and every program and also includes the runtime on each and every binary program. It's like with binary and little bit extra code has been added to it called as a run time which manage the automated memory management and garbage collection.

# Multi-threading
Now in days in any software, we can do multiple tasks at once like in google drive we can download, upload, and delete at once and automatically we are redirected to same UI this is called multi-threading

# Some challenges of multi-threading
If users are booking a ticket at same time and If we need a single ticket single person then we need the concurrency rather than multi-threading
Some of the languages has multi-threading like go, Java has but the codebase is become very complex, and there are more number of human errors and there for Go has built in 2007 and became open-source in 2009.

# Use cases
For performant Applications
Running on scaled distributed systems

# Characterstics of Go
The developer of go is attempt to combine both
1. Simple and readable syntax like python
2. efficiency and safety like lower-level language like C++
3. Go is used in backend side or on server side like in microservices, web applications, database services
4. It is very fast build time, start up and run 
5. requires fewer resources to run the codebase
6. It is compiled language and compiled into single binary or in machine code very quickly

# Compilation:

-> Computers only understands the binary codes like 0's and 1's and we writting the code into the human readable format. In order to make it understandable by computer, we need to convert this into the binary format and this conversion is known as the compilation or compiling the code from human readable format to binary format.

# History:

-> It was created by engineers at Google.

-> When the go language was created the undue complexity was required to use languages like c++ and python and engineers was frusted from that.

-> Go was created to combine
	- The ease of programming of an interpreted, dynamically typed language (such as python)
	- With the efficiency and safety of a statically typed, compiled language (such as c++)
	- It also simed to be modern, with support for networked and multicore computing.

# create Module for projects in Go
- go mod init <module path>
- create new module
- It initialized a go.mod file 

# Every thing is organized into the packages in go
# The main function is the entry point in go program and only have one main function in every go file
# Go packages is like container and It gives you various functionality that we can directly use into our code
# There are 25 keywords in go

# First Go program:

	package main
	import "fmt"
	func main() {
    	fmt.Println("hello world")
	}

1. package main: 
   - packages are Go's way of organizing and reusing code.
   - Every go program start with the package declaration.
   - The syntax is very simple and It is "package <name of the package>".
   - When we build reusable pieces of code, we will develop a code as shared library.
   - but if we want to develop the executable programs, then we have to use the main package in order to make it as a executable program.

2. import "fmt":
	- we use the import keyword to include the code of other packages to our code.
	- It is a short form of "format".
	- And it implements formatting for input and output.
	- We import it because we need to print the sentence to console.
	- Syntax is `import "<package_name>" `

3. Function declarations:
   - main function is a special type of function and it is the entry point of the go file means every go program starts executing from main function.
   - every go program will contains the main package and main function in it code.
   - function always declare using the "func" keyword.

4. fmt.Println("Hello World"): 
	- syntax is first use the "fmt" then specify the name of the function that is present in the package.
	- the dot operator is called as a member access operator.

# Data type in Go:
	- data type is a classification made on the kind of data.
	- Go has String, Number which is furthur divide into the Integer and Float, Boolean, Arrays & Slices in which slices represents a flexible array-like data type but they also provide more control over memort allocation, Map is a collection of key-value pairs.
# Why we need it?    
	- categorize a set of related values
	- describe the operations that can be done on them
	- define the way the data is stored in the memory

# Static and Dynamic typed language: 

	- Statically typed language throws an error when types are used incorrectly such as c++, java.
	- Dynamically typed language not throws error when we write the code such as javascript


# Variable declarations:

-> We have boolean, strings, numbers that furthur divide into the buckets and furthur in integers, unsighted integers, floats, complex numbers, and byte.

-> complex numbers are little funny and It is used to represent the concept of imaginary numbers.

-> integers are furthur divide into the int, int8, int16, int32, int64 and unsigned integers also.

-> Byte is also use a lot specially when you are marshaling a json object to be sent across a network connection or reading from a file or from disk.

var <variable_name> <data_type> = <value>

-> boolean value by default is false.

:= operator known as short assignment operator

-> instead of typing var empty string, we can use this operator for it like empty := "" and go understand that this is a string.

-> and := is a static type.

-> very often time variables declare like var <variable_name> <data_type> = <value> this.

-> we can also declare variables like this 

<first_variables>, <second_variable> := <value of first variable>, <value of second variable>

-> This above syntax is called Syntactic Sugar in Go and makes language sweeter for user but it doesn't work with constants

-> We can also convert numbers from one data_type to another like

<integer_variable_name> := float64(<value>)

-> This int8, int32, int64 8, 32, 64 define the bits not the bytes in go.

-> For the constant value just use the const key word before the variable name. And constant does not support the short declaration syntax.

-> Formatting string is as similar as in c language.

-> We can also declare ARRAY in go like 

<var keyword> <variable_name> = [size]<Type of data that we are going to store into the array>
var name = [2]string

-> If we have 50 size of an array but only 1 item is there in the array then 49 position will not be having any value and this is unsufficient for memory management.

-> We have to use the dynamic array like on every time if any new value come size will be increased and on deletion size will be decreased.

-> In go we can do it using SLICE.

-> Slice in go is a abstraction of an array

-> We can declare an slice with using this below syntax
var booking []int 

-> In go conditions like if-else and else if is same as other language but the brackets after if that is () are removed in go.

# -> Alternate way to write the if statements in go and that is this 

if <Initial_statements>; condition {

}

-> in which If variable that we are comparing that only used in that if statement then this syntax is useful.

-> like this 

if length := 3; length > 2 {
	fmt.Println("yes")
}

# We can declare the for loop like this 
1. for {
   }

2. for initialization; condition; increament {
 }

3. for variable 1, variable 2 := range condition variable {
 }

-> "_" this operator called dummy variable for all those variable that we are not going to use in the codebase but giving error like variable declared but not used because of Go is a statically typed language.

# We can declare function using func keyword

-> And pass the value to function like other programming language

-> return the value but need to specify the return type as we are doing for the varible like

func g() string {
	return "heelo"
}

# The variables that are using commonly in everytime, we can declare that variable as package level variables so that It can be used by everytime but there are some conditions to declare it

1. Package level variable cannot be declare using this := syntax
2. And after doing it no need to pass it in the argument It will automatically accessible by all functions.


# Array
var <name of the variable> = [define size]<data_type>{elements} 
or
<name of the variable> := [define size]<data_type>{elements}

# If we doesn't define size of the array then It will be called as slice

# Creating map
var <map_variable name> = make(<map keyword>[key datatype]value datatype)

# List of maps
var <variable_name> = make([]map[datatype]datatype, size) and It will be increased when we want

# struct
-> Sometime in project we need to have different datatypes and for that we have structure Datatype which is the user define datatype and we can define it using struct keyword

-> We can define mixed type of data into it.

-> It's like a lightweight class which doesn't support inheritance

-> We can define it using type keyword

type <name of the struct> struct {
	<user_define_types>
}

-> We can also define the list of maps of this structure like 
var <variable_name> = make([]structure_name, size)

# Sprintf, Sprint, Sprintln - Function using which we can store the output in variable

# In order to put some of the stop or we can say the time out then we can use the SLEEP function from time in-built package

-> In go using this kind of stop or time out function it stops the execution and after completion of it execute the other part of the code.

-> But in real-time this should be time consuming because It is like threading kind of thing

-> And in real world we need the concurrency kind of functionality

-> And for that just need to type the "GO" keyword in front of that code which we want has concurrent so that new user will come then they have there own goroutines

# goroutine

-> A goroutine is a lightweight thread managed by the Go runtime

-> Sometime what happens that by default, the main goroutine does not wait for other goroutine

-> For that we need to tell "main" thread that need to wait until other thread completes

-> We can do it using WaitGroup function that comes from sync package

-> It has three function under it
1. Add: add number of threads that main thread should wait for 
2. Wait: waits till WaitGroup counter become 0
3. Done: Which runs in main thread and It decreament the counter by 1
