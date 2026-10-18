package main

import (
    "fmt"
    "os"
    "github.com/Homework900/simple-library1/library"
    "github.com/Homework900/simple-library1/notifications"
    "github.com/Homework900/simple-library1/storage"
)

const dbFile = "books.json"

func main() {
    myLibrary, err := storage.LoadLibraryFromJSON(dbFile)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("Файл данных не найден, создана новая библиотека.")
            myLibrary = library.New()
        } else {
            fmt.Println("Ошибка загрузки:", err)
            return
        }
    } else {
        fmt.Println("Библиотека успешно загружена из файла.")
    }

    b1 := myLibrary.AddBook("Война и мир", "Лев Толстой", 1869)
    r1 := myLibrary.AddReader("Пётр", "Петров")

    if err := myLibrary.IssueBookToReader(b1.ID, r1.ID); err != nil {
        fmt.Println("Ошибка выдачи:", err)
    } else {
        r1.AssignBook(b1)
    }

    if err := myLibrary.ReturnBook(b1.ID); err != nil {
        fmt.Println("Ошибка возврата:", err)
    } else {
        fmt.Println("Книга возвращена:", b1)
    }

    if err := myLibrary.ReturnBook(b1.ID); err != nil {
        fmt.Println("Ожидаемая ошибка:", err)
    }
    notifiers := []notifications.Notifier{
        notifications.EmailNotifier{EmailAddress: "student@example.com"},
        notifications.SMSNotifier{PhoneNumber: "+79991234567"},
    }
    for _, n := range notifiers {
        n.Notify("Ваша книга просрочена!")
    }

    if err := storage.SaveLibraryToJSON(dbFile, myLibrary); err != nil {
        fmt.Println("Ошибка сохранения:", err)
    } else {
        fmt.Println("Данные сохранены в", dbFile)
    }
}