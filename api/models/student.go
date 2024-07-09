package models

import (
	"regexp"
	"strings"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	StudentHome struct {
		TagName string   `json:"tagname"`
		Groups  []string `json:"groups"`
	}

	StudentInfo struct {
		Student  Student `json:"student"`
		Password string  `json:"password"`
		Login    string  `json:"login"`
	}

	Student struct {
		Id          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Age         uint8  `json:"age"`
		ClassNumber uint8  `json:"class_number"`
		PhoneNumber string `json:"phone_number"`
		Gender      string `json:"gender"`
		Email       string `json:"email"`
	}
)

// This method validate student
func (s *Student) Validate() error {
	s.Email = strings.TrimSpace(s.Email)
	s.PhoneNumber = strings.TrimSpace(s.PhoneNumber)
	s.Gender = strings.ToLower(s.Gender)
	return validation.ValidateStruct(
		s,
		validation.Field(
			&s.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&s.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(`^\+998([- ])?(90|91|93|94|95|98|99|33|50|97|71)([- ])?(\d{3})([- ])?(\d{2})([- ])?(\d{2})$`)),
		),
		validation.Field(
			&s.FirstName,
			validation.Required,
			validation.NilOrNotEmpty,
		),
		validation.Field(
			&s.LastName,
			validation.Required,
			validation.NilOrNotEmpty,
		),
		validation.Field(
			&s.ClassNumber,
			validation.Required,
			validation.NilOrNotEmpty,
		),
		validation.Field(
			&s.Gender,
			validation.Required,
			validation.In("male", "female"),
		),
	)
}

// This method validate student info
func (si *StudentInfo) Validate() error {
	si.Password = strings.TrimSpace(si.Password)
	return validation.ValidateStruct(
		si,
		validation.Field(
			&si.Password,
			validation.Required,
			validation.Length(8, 30),
		),
		validation.Field(
			&si.Student,
			validation.Required,
			validation.NilOrNotEmpty,
		),
	)
}
