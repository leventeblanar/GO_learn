package main

import (
	"fmt"
	"errors"
)

type Book struct {
	title		string
	author		string
	isbn		string
	available	bool
}

type Library struct {
	name	string
	books 	[]Book
}

func (l *Library) AddBook(book Book)  {
	l.books = append(l.books, book)
	fmt.Printf("%s hozzáadva a könyvtár leltárhoz\n", book.title)
}

func (l *Library) BorrowBook(isbn string) error {
	for i := range l.books {
		if l.books[i].isbn == isbn {
			if l.books[i].available {
				l.books[i].available = false
				return nil
			} else {
				return errors.New("a könyv már ki van kölcsönözve")
			}
		} 
	}
	return errors.New("könyv nem található")
}

func (l *Library) ReturnBook(isbn string) error {
	for i := range l.books {
		if l.books[i].isbn == isbn {
			if l.books[i].available {
				return errors.New("a könyv ki se volt kölcsönözve")
			} else {
				l.books[i].available = true
			}
		}
	}
	return errors.New("nincs a könyvtárban ilyen könyv")
}

func (l Library) ListAvailableBooks() {
	fmt.Println("\nElérhető könyvek:")
	for _, book := range l.books {
		if book.available {
					fmt.Printf("- %s (%s) - ISBN: %s\n", book.title, book.author, book.isbn)

		}
	}
}



func main() {
	szegediKönyvtár := Library{
		name: "Szegedi Könyvtár",
		books: []Book{},
	}

	szegediKönyvtár.AddBook(Book{"A Gyűrűk Ura", "J.R.R. Tolkien", "978-963-11-8171-4", true})
	szegediKönyvtár.AddBook(Book{"1984", "George Orwell", "978-963-13-5582-4", true})
	szegediKönyvtár.AddBook(Book{"Harry Potter és a Bölcsek Köve", "J.K. Rowling", "978-963-07-8491-6", true})
	szegediKönyvtár.AddBook(Book{"A kis herceg", "Antoine de Saint-Exupéry", "978-963-08-0123-7", true})

	szegediKönyvtár.ListAvailableBooks()

	err := szegediKönyvtár.BorrowBook("978-963-13-5582-4")
	if err != nil {
		fmt.Println("Hiba:", err)
	} else {
		fmt.Println("sikeres kölcsönzés!")
	}
}