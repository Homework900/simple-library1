package library

import (
    "fmt"
    "github.com/Homework900/simple-library1/domain"
)

type Library struct {
    Books        []*domain.Book
    Readers      []*domain.Reader
    lastBookID   int
    lastReaderID int
}

func New() *Library {
    return &Library{Books: []*domain.Book{}, Readers: []*domain.Reader{}}
}

func (lib *Library) AddBook(title, author string, year int) *domain.Book {
    lib.lastBookID++
    b := &domain.Book{ID: lib.lastBookID, Title: title, Author: author, Year: year}
    lib.Books = append(lib.Books, b)
    return b
}

func (lib *Library) AddReader(firstName, lastName string) *domain.Reader {
    lib.lastReaderID++
    r := &domain.Reader{ID: lib.lastReaderID, FirstName: firstName, LastName: lastName, IsActive: true}
    lib.Readers = append(lib.Readers, r)
    return r
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
    for _, b := range lib.Books {
        if b.ID == id {
            return b, nil
        }
    }
    return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) IssueBookToReader(bookID, readerID int) error {
    book, err := lib.FindBookByID(bookID)
    if err != nil {
        return err
    }
    var reader *domain.Reader
    for _, r := range lib.Readers {
        if r.ID == readerID {
            reader = r
            break
        }
    }
    if reader == nil {
        return fmt.Errorf("читатель с ID %d не найден", readerID)
    }
    return book.IssueBook(reader)
}

func (lib *Library) ReturnBook(bookID int) error {
    book, err := lib.FindBookByID(bookID)
    if err != nil {
        return err
    }
    return book.ReturnBook()
}

func (lib *Library) GetAllBooks() []*domain.Book { return lib.Books }

func (lib *Library) UpdateIDs() {
    for _, b := range lib.Books {
        if b.ID > lib.lastBookID {
            lib.lastBookID = b.ID
        }
    }
    for _, r := range lib.Readers {
        if r.ID > lib.lastReaderID {
            lib.lastReaderID = r.ID
        }
    }
}