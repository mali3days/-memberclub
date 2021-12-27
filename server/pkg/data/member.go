package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
)

type Member struct {
	Id               int       `json:"id"`
	Name             string    `json:"name" validate:"required"`
	Email            string    `json:"email" validate:"required,email,emailuniq"`
	RegistrationDate time.Time `json:"registrationDate"`
}

func (m *Member) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(m)
}

// Members is a collection of Member
type Members []*Member

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (m *Members) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

// GetMembers returns a list of members
func GetMembers() Members {
	return memberList
}

func AddMember(m *Member) {
	m.Id = getNextID()
	memberList = append(memberList, m)
}

func (m *Member) Validate() error {
	v := validator.New()

	v.RegisterValidation("emailuniq", validateEmailUniq)

	return v.Struct(m)
}

func validateEmailUniq(fl validator.FieldLevel) bool {
	for i := range memberList {
		if memberList[i].Email == fl.Field().String() {
			return false
		}
	}

	return true
}

func getNextID() int {
	lm := memberList[len(memberList)-1]
	return lm.Id + 1
}

var memberList = []*Member{
	{
		Id:               1,
		Name:             "Tom",
		Email:            "tomriddle@gmail.com",
		RegistrationDate: time.Now(),
	},
	{
		Id:               2,
		Name:             "Harry",
		Email:            "potter@gmail.com",
		RegistrationDate: time.Now(),
	},
}
