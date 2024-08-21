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
			newContact("John", "aeou", "123456789"),
			newContact("Doe", "dj@gmail.com", "987654321"),
			newContact("Jane", "janeD@gmail.com", "123456789"),
			newContact("Doe", "doeJ@gmail.com", "987654321"),
		},
	}
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	page := newPage()

	app := fiber.New(fiber.Config{
        Views: engine,
    })

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", page)
	})

	app.Post("/contacts", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		phone := c.FormValue("phone")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Values["phone"] = phone
			formData.Errors["email"] = "Email already exists"
			c.Status(fiber.StatusUnprocessableEntity);
			return c.Render("form", formData)
		}

		page.Data.Contacts = append(page.Data.Contacts, newContact(name, email, phone))
		return c.Render("display", page.Data)
	});


	app.Listen(":6942")
}
