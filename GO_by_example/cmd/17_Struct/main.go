package main

import ("fmt")

//  Struct saját típus: több mező (field) tárolására alkalmas
//  Olyan, mint egy adatcsomag -> pl. egy könyv vagy egy felhasználó adatai
//  		- Mezők = név + típus
// 			- Struct maga is egy típus -> példányt hozol létre belőle
//  		- Értékek alapból bemásolódnak (copy), nem pointerként mennek át


type Book struct {
	Title		string
	Author		string
	Year		int
}

type Person struct {
	Name		string
	Age			int
}


func main() {
	b1 := Book{"The Hobbit", "tolkien", 1937}
	b2 := Book{Title: "1984", Author: "Orwell", Year: 1949}
	b3 := Book{Title: "Go in Action"} // hiányzó mezők = alapérték (Year = 0)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	// mezők elérése
	fmt.Println("Title:", b1.Title)
	fmt.Println("Author:", b1.Author)

	// mező módosítása
	b1.Year = 1938
	fmt.Println("Updated year:", b1.Year)


	// Mivel a struct az csak egy sablon, séma, típusdefiníció így tárolnunk kell valamiben aminek megadjuk a structot mint minta

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age:40},
	}

	for _, p := range people{
		fmt.Println(p.Name, "is", p.Age, "years old.") 
	}

	//  Létezik olyan megoldás is, hogy "Map of struct" ha pl. név szerint akarjuk gyorsan elérni

	peopleMap := map[string]Person{
		"alice": 	{Name: "Alice", Age: 30},
		"bob": 		{Name: "Bob", Age: 25},
		"charlie": 	{Name: "Charlie", Age:40},
	}

	fmt.Println(peopleMap["bob"].Age)
}
