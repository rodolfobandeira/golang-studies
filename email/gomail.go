package main

import (
    "gopkg.in/gomail.v2"
    "fmt"
)

func main() {
    m := gomail.NewMessage()
    m.SetHeader("From", "pub@example.org")
    m.SetHeader("To", "foo@example.org")
    m.SetHeader("Subject", "my subject")
    m.SetBody("text/plain", "This is the body of the message.")

    d := gomail.NewPlainDialer("smtp.example.org", 587, "foo@example.org", "password")

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

    fmt.Println("Sent")
}

