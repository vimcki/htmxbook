package repo

import (
	"strings"

	"github.com/vimcki/htmxbook/internal/model"
)

type Repo struct {
	contacts []model.Contact
	lastID   int
}

func New() *Repo {
	return &Repo{
		contacts: []model.Contact{
			{
				ID:    1,
				First: "alice",
				Last:  "smith",
				Phone: "1234",
				Email: "foo@gmail.com",
			},
			{
				ID:    2,
				First: "bob",
				Last:  "jones",
				Phone: "5678",
				Email: "bar@gmail.com",
			},
			{
				ID:    3,
				First: "charlie",
				Last:  "brown",
				Phone: "9012",
				Email: "baz@gmail.com",
			},
		},
		lastID: 3,
	}
}

func (r *Repo) Search(query string) []model.Contact {
	var contacts []model.Contact
	for _, contact := range r.contacts {
		if strings.Contains(contact.First, query) {
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

func (r *Repo) All() []model.Contact {
	return r.contacts
}

func (r *Repo) Save(contact model.Contact) {
	r.lastID++
	contact.ID = r.lastID
	r.contacts = append(r.contacts, contact)
}

func (r *Repo) Find(id int) (model.Contact, bool) {
	for _, contact := range r.contacts {
		if contact.ID == id {
			return contact, true
		}
	}
	return model.Contact{}, false
}

func (r *Repo) Update(id int, contact model.Contact) {
	for i, c := range r.contacts {
		if c.ID == id {
			r.contacts[i] = contact
			return
		}
	}
}

func (r *Repo) Delete(id int) {
	for i, c := range r.contacts {
		if c.ID == id {
			r.contacts = append(r.contacts[:i], r.contacts[i+1:]...)
			return
		}
	}
}
