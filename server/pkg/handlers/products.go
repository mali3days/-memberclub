package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mali3days/memberclub/pkg/data"
)

// Members is a http.Handler
type Members struct {
	l *log.Logger
}

// NewMembers creates a members handler with the given logger
func NewMembers(l *log.Logger) *Members {
	return &Members{l}
}

// getMembers returns the Members from the data store
func (m *Members) GetMembers(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Members")

	// fetch the Members from the datastore
	lp := data.GetMembers()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (m *Members) AddMember(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle POST Member")

	prod := r.Context().Value(KeyMember{}).(data.Member)
	data.AddMember(&prod)
}

type KeyMember struct{}

func (m Members) MiddlewareValidateMember(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Member{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			m.l.Println("[ERROR] deserializing member", err)
			http.Error(rw, "Error reading member", http.StatusBadRequest)
			return
		}

		// validate the member
		err = prod.Validate()
		if err != nil {
			m.l.Println("[ERROR] validating member", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating member: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the member to the context
		ctx := context.WithValue(r.Context(), KeyMember{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
