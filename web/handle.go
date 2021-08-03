package web

import (
	"encoding/json"
	"errors"
	"fmt"
	models "github/estelasouza/api-star-wars/models/discussion"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DiscussionRepository interface {
	Create(models.Discussion) error
	List() ([]models.Discussion, error)
}
type Controller struct {
	nextId uint
	store  map[uint]models.Discussion
	repo   DiscussionRepository
}

func NewController(repo DiscussionRepository) *Controller {
	return &Controller{
		nextId: 1,
		store:  make(map[uint]models.Discussion),

		repo: repo,
	}
}

func (c *Controller) HandleCreateDiscussion(w http.ResponseWriter, r *http.Request) {

	d := models.Discussion{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	defer r.Body.Close()

	err = c.repo.Create(d)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) HandleListDiscussions(w http.ResponseWriter, r *http.Request) {

	people, err := c.repo.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(people)
	if err != nil {
		return
	}

}

func (c *Controller) HandleGetDiscussion(w http.ResponseWriter, r *http.Request) {
	id, err := c.parseDiscussionId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	d, found := c.store[id]

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
func (c *Controller) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := c.parseDiscussionId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, found := c.store[id]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(c.store, id)
	w.WriteHeader(http.StatusNoContent)
}
func (c *Controller) parseDiscussionId(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	DiscussionIdRaw := vars["id"]
	fmt.Println(DiscussionIdRaw)
	// strconv.ParseUint(DiscussionIdRaw, 10, 32)
	DiscussionId, err := strconv.Atoi(DiscussionIdRaw)

	if err != nil {
		return 0, err
	}

	if DiscussionId <= 0 {
		return 0, errors.New("id must be positive")
	}

	id := uint(DiscussionId)

	return id, nil
}

func (c *Controller) HandleUpdateDiscussion(w http.ResponseWriter, r *http.Request) {
	id, err := c.parseDiscussionId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, found := c.store[id]

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	d := models.Discussion{}
	err = json.NewDecoder(r.Body).Decode(&d)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	d.Id = id
	c.store[id] = d

	w.WriteHeader(http.StatusNoContent)
}
