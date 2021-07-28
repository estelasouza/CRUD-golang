package web

import (
	"encoding/json"
	"io"
	"net/http"
)

type Controller struct {
	nextId uint
	store  map[uint]Person
}
type Person struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func NewController() *Controller {
	return &Controller{
		nextId: 1,
		store:  make(map[uint]Person),
	}
}

func (c *Controller) HandlePing(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "pong")
}

func (c *Controller) HandleCreatePerson(rw http.ResponseWriter, r *http.Request) {

	p := Person{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		io.WriteString(rw, err.Error())
		return
	}

	id := c.nextId
	p.Id = id
	c.store[id] = p
	c.nextId++

	rw.WriteHeader(http.StatusNoContent)
}

func (c *Controller) HandleListPeople(rw http.ResponseWriter, r *http.Request) {
	people := make([]Person, len(c.store))

	i := 0

	for _, value := range c.store {
		people[i] = value
		i++
	}

	rw.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(people)
	if err != nil {
		return
	}

}
