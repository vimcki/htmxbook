package repo

import (
	"strings"
	"time"

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
			{
				ID:    4,
				First: "dave",
				Last:  "smith",
				Phone: "3456",
				Email: "asdasd@daq.sd",
			},
			{
				ID:    5,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    6,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    7,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    8,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    9,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    10,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    11,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    12,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    13,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    14,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    15,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    16,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    17,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    18,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    19,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    20,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    21,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    22,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    23,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    24,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    25,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    26,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    27,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    28,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    29,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    30,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    31,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    32,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
			{
				ID:    33,
				First: "eve",
				Last:  "jones",
				Phone: "7890",
				Email: "dasokdaso@d.e",
			},
		},
		lastID: 33,
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

func (r *Repo) All(page int) []model.Contact {
	pageSize := 10
	start := (page - 1) * pageSize
	end := start + pageSize
	if end >= len(r.contacts) {
		end = len(r.contacts) - 1
	}
	return r.contacts[start:end]
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

func (r *Repo) Count() int {
	time.Sleep(1 * time.Second)
	return len(r.contacts)
}
