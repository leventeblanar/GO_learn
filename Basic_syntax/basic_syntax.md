```
func main() {
	fmt.Println("Hello, Go!")
}
```

// Változók
```go
var x int = 20 			// teljes
var y = 20				// típus automatikus
z := 30					// rövid forma csak függvényen belül


// Alap típusok
var a int = 42
var b float64 = 3.14
var c string = "Go is cool"
var d bool = true


// Konstansok
const Pi = 3.14159
const Greeting = "Hello"


// For ciklus
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Végtelen ciklus

for {
    fmt.Println("looping forever")
    break
}
```


// If/else
```go
x := 10
if x > 5 {
    fmt.Println("x nagyobb mint 5")
} else {
    fmt.Println("x kisebb vagy egyenlő 5")
}
```


// Switch
```go
day := 3
switch day {
case 1:
	fmt.Println("Hétfő")
case 2,3:
	fmt.Println("Kedd vagy Szerda")
default:
	fm.tPrintln("Más nap")
}
```


// Függvények
```go
func add(a int, b int) int {
    return a + b
}

result := add(5, 3) // 8
```

// több érték visszaadása:
```go
func divide(a, b int) (int, int) {
	return a / b, a % b
}

q, r := divide(10 ,3)
```


// - Array és Slice
// Array (fix méret)
```go
var arr [3]int = [3]int{1, 2, 3}

//Slice(rugalmas)
nums := [int{1, 2, 3}]
nums = append(nums, 4) // [1 2 3 4]
```

// Iterálás
```go
for i, v := range nums {
	fmt.Println(i, v)
}
```


// MAP (dict)
```go
ages := map[string]int{
	"Alice": 25,
	"Bob": 30,
}

ages["Charlie"] = 40
delete(ages, "Bob")

for name, age := range ages {
	fmt.Println(name, "is", age)
}
```

// Struct
```go
type User struct[
	Name	string
	Age		int
]

u := User{Name: "Alice", Age: 25}
fmt.Println(u.Name, u.Age)
```

// Pointer
```go
x := 10
p := &x
fmt.Println(*p) // 10

*p = 20
fmt.Println(x)
```

// ERROR kezelés
```go
input := "123a"
num, err := strconv.Atoi(input)
if err != nil {
	fmt.Println("Hiba:", err)
} else {
	fmt.Println("Szám:", num)
}
```


// Package & Import
```go
fmt.Println(math.Sqrt(16))
```