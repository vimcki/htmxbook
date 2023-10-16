package repo

import (
	"strings"

	"github.com/vimcki/htmxbook/internal/model"
)

type Repo struct {
	contacts []model.Contact
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
