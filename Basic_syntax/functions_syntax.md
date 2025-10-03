package basicsyntax

```go
//  function szignatúra
func name(params...) returnTypes...
1. func -> kulcsszó
2. name -> a függvény neve (pl. sum, CountTo)
3. paraméterek -> név típus páros, pl. (n int, s strng)
4. return típus(ok) -> lehet 0, 1 vagy több


// Paraméterek
func greet(name string, age int)
name string = a name cáltozó string
age int = az age változó egy int

//  több azonos típusú, rövidíthetett:
func add(a, b int) int


//  Visszatérési értékek
- egy érték
func square (x int) int

- több érték
func divide (a, b int) (int, error)

nincs visszatérés
- func sayHello(name string)


//  Generics
//  Amikor []int-et vagy map[string]int-et látsz azok típusparaméterek
//          - []int = slice of int
//          - map[string]int = kulcs string, érték int

példa:
func sumSlice(nums []int) int


//   Speciális forma -> függvény, ami függvényt ad vissza
func countTo(n int) func(yield func(int) bool)
// func countTo(n int) -> egy függvény, ami n int-et kap
// visszaad: func(yield func(int) bool) -> egy másik függvényt, ami szintén paraméterként vár egy függvényt (yield), és az yield maga egy func(int bool)
// Tehát: outer function: gyár - inner function: iterátor

péld:
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}
//  multiplier paramétere: factor int
//  visszatérési értéke: egy func(int) int
```