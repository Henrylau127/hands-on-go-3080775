// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	// add a book to the library
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	libraries := library{}
	jb := author{name: "James Baldwin"}


	// add 2 books to the library by the same author
	book1 := book{
		title: "book1",
		author: jb,
	}
	book2 := book {
		title: "book2",
		author: jb,
	}

	libraries.addBook(book1)
	libraries.addBook(book2)

	// add 1 book to the library by a different author
		book3 := book {
		title: "book3",
		author: jb,
	}

	libraries.addBook(book3)

	// dump the library with spew
	spew.Dump(libraries)

	// lookup books by known author in the library
	res := libraries.lookupByAuthor(jb.name)

	// print out the first book's title and its author's name
	if len(res) != 0 {
		b := res[0]
		fmt.Println(b.title, "by", b.author.name)
	}
}
