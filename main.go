package main

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Views interface {
    Load() error
    Render(io.Writer, string, interface{}, ...string) error
}

type Count struct {
	Count int
}

type Contact struct {
	Email string
	Name string
	Phone string
}

func newContact(name, email, phone string) Contact {
	return Contact{
		Email: email,
		Name: name,
		Phone: phone,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@gmail.com", "123456789"),
			newContact("Doe", "dj@gmail.com", "987654321"),
			newContact("Jane", "janeD@gmail.com", "123456789"),
			newContact("Doe", "doeJ@gmail.com", "987654321"),
		},
	}
}


func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	data := newData()

	app := fiber.New(fiber.Config{
        Views: engine,
    })

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", data)
	})

	app.Post("/contacts", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		phone := c.FormValue("phone")

		data.Contacts = append(data.Contacts, newContact(name, email, phone))
		return c.Render("display", data)
	});


	app.Listen(":6942")
}
