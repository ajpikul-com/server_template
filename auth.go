package main

import (
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/ajpikul-com/uwho"
)

type stageFactory struct{}

func (f *stageFactory) New() uwho.ReqByCoord {
	return &stageState{}
}

type stageState struct {
	email       string
	session     uuid.NullUUID
	expired     bool
	failedLogin bool
	failedAuth  bool
}

func (s *stageState) InitState() error {
	if s.session.Valid {
		return uwho.ErrStateExists
	} else if s.email == "" {
		s.failedLogin = true
		return uwho.ErrNoCredential
	}
	s.session.UUID = uuid.New()
	s.session.Valid = true
	s.expired = false
	s.failedLogin = false
	s.failedAuth = false
	return nil
}

func (s *stageState) AcceptData(claims map[string]interface{}) bool {
	if !s.session.Valid { // Someone is just supplying their e-mail for the first time
		s.email = claims["email"].(string)
		if s.email != "ajpikul@gmail.com" {
			s.email = ""
			s.failedLogin = true
			return false
		}
	}
	if s.session.Valid {
		defaultLogger.Debug("It seems like user somehow managed to login while already in a valid state/session")
	}
	return true
}

func (s *stageState) DeleteState() {
	s.email = ""
	s.session.Valid = false
	s.expired = false
	s.failedLogin = false
	s.failedAuth = false
}

func (s *stageState) SessionToState(sessString string, expired bool) bool {
	if expired {
		s.DeleteState()
		s.expired = true
		return false
	}
	values := strings.Split(sessString, "&")
	if len(values) != 2 {
		s.DeleteState()
		return false
	}
	var err error
	s.session.UUID, err = uuid.Parse(values[0])
	if err != nil {
		s.DeleteState()
		return false
	}
	s.session.Valid = true
	s.email = values[1]
	s.expired = expired
	return true
}

func (s *stageState) StateToSession() string {
	return s.session.UUID.String() + "&" + s.email
}

func (s *stageState) AuthorizeUser(w http.ResponseWriter, r *http.Request) bool {
	// Check that all papers in order, if it is, renew things if you want
	if s.session.Valid && s.email == "ajpikul@gmail.com" {
		return true
	}
	s.failedAuth = true
	return false
}
