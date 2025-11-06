package main

import (
	"fmt"
	"errors"
)

type Book struct {
	id			int
	title		string
	author		string
	BorrowedBy	string
	available 	bool
}

type BookTracker struct {
	books		[]Book
	nextID		int
}


// AddBook(title, author string) - új könyv hozzáadása (auto-increment ID)
func (bt *BookTracker) AddBook(title, author string) {
	book := Book {
		id: bt.nextID,
		title: title,
		author: author,
		available: true,
	}

	bt.books = append(bt.books, book)
	bt.nextID++

	fmt.Printf("Sikeresen hozzáadva: %s - %s\n", book.author, book.title)
}
// BorrowBook(id int, borrowerName string) error - kölcsönzés (error ha nincs meg vagy már ki van kölcsönözve)
func (bt *BookTracker) BorrowBook(id int, borrowerName string) error {
	for i := range bt.books {
		if id == bt.books[i].id {
			if !bt.books[i].available {
				return errors.New("a könyv jelenleg nem elérhető")
			} else {
			bt.books[i].BorrowedBy = borrowerName
			bt.books[i].available = false
			fmt.Printf("A könyv kikölcsönzésre került: könyvid: (%d) %s - %s (kölcsönző: %s) elérhető: %v\n", bt.books[i].id, bt.books[i].author, bt.books[i].title, borrowerName, bt.books[i].available)
			return nil
			}
		}
	}
	return errors.New("nincs ilyen idval rendelkező könyv")
}
// ReturnBook(id int) error - visszahozás (error ha nincs meg vagy nem volt kikölcsönözve)
func (bt *BookTracker) ReturnBook(id int) error {
	for i := range bt.books {
		if id == bt.books[i].id {
			if !bt.books[i].available {
				bt.books[i].available = true
				bt.books[i].BorrowedBy = ""
				fmt.Printf("Sikeres elérhetőre állítás: %s - %s\n", bt.books[i].author, bt.books[i].title)
				return nil
			} else {
				return errors.New("a könyv még nem volt kikölcsönözve")
			}
		}
	}
	return errors.New("nincs ilyen idval rendelkező könyv készleten")
}

// GetAvailableBooks() []Book - elérhető könyvek
func (bt *BookTracker) GetAvailableBooks() []Book {
	fmt.Println("Elérhető könyvek:")
	var availableBooks []Book
	for i, book := range bt.books {
		if bt.books[i].available {
		availableBooks = append(availableBooks, book)
		fmt.Printf("(id: %d) %s - %s\n", book.id, book.author, book.title)
		}
	}
	return availableBooks
}
// GetBorrowedBooks() []Book - kikölcsönzött könyvek
func (bt *BookTracker) GetBorrowedBooks() []Book {
	fmt.Println("Nem Elérhető könyvek:")
	var notavailableBooks []Book
	for i, book := range bt.books {
		if !bt.books[i].available {
		notavailableBooks = append(notavailableBooks, book)
		fmt.Printf("(id: %d) %s - %s\n", book.id, book.author, book.title)
		}
	}
	return notavailableBooks
}
// GetBooksByAuthor(author string) []Book - adott szerző összes könyve
func (bt *BookTracker) GetBooksByAuthor(author string) []Book {
	var authorBooks []Book
	fmt.Printf("%s szerzőhöz tartozó könyv.\n", author)
	for i, book := range bt.books {
		if bt.books[i].author == author {
			authorBooks = append(authorBooks, bt.books[i])
			fmt.Printf("(id: %d) %s - %s \n", book.id, book.author, book.title)
		}
	}
	return authorBooks
}
// ListAll() - minden könyv kilistázása (jelezve ki nála van)
func (bt *BookTracker) ListAll() {
	for i := range bt.books {
		fmt.Printf("(id: %d) %s - %s Elérhető: %v, kölcsönző:%s\n", bt.books[i].id, bt.books[i].author, bt.books[i].title, bt.books[i].available, bt.books[i].BorrowedBy)
	}
}

// GetBookByID(id int) (Book, error) - ID alapján lekérés
func (bt *BookTracker) GetBookByID(id int) (Book, error) {
	for i, book := range bt.books {
		if bt.books[i].id == id {
			fmt.Printf("(id: %d) %s - %s\n", bt.books[i].id, bt.books[i].author, bt.books[i].title)
			return book, nil
		}
	}
	return Book{}, errors.New("nem áll rendelkezésre könyv ezzel az id-val")
}

func main() {
	booktracker1 := BookTracker{
		nextID: 1,
	}

	booktracker1.AddBook("Tesztkönyv", "Teszt author")
	booktracker1.AddBook("Go for Dummies", "Vidéki Béla")
	booktracker1.AddBook("Futás lábak nélkül", "Csiga József")
	booktracker1.AddBook("Ház semmiből", "Csiga József")

	if err := booktracker1.BorrowBook(1, "Lajos"); err != nil {
		fmt.Println("Hiba a kölcsönzés során: ", err)
	}
	
	if err := booktracker1.BorrowBook(2, "Öcsike"); err != nil {
		fmt.Println("Hiba a kölcsönzés során: ", err)
	}

	if err := booktracker1.ReturnBook(1); err != nil {
		fmt.Println("Hiba a visszaszolgáltatás során: ", err)
	}

	if err := booktracker1.ReturnBook(3); err != nil {
	fmt.Println("Hiba a visszaszolgáltatás során: ", err)
	}

	booktracker1.GetAvailableBooks()
	booktracker1.GetBorrowedBooks()
	booktracker1.GetBooksByAuthor("Csiga József")
	booktracker1.ListAll()

	booktracker1.GetBookByID(2)
}