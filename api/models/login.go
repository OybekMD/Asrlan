package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	EmailData struct {
		Code string
	}

	ResetData struct {
		DeepLink string
	}

	Signup struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Language int64  `json:"language"`
		Level    int64  `json:"level"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Username   string `json:"username"`
		Bio        string `json:"bio"`
		BirthDay   string `json:"birth_day"`
		Email      string `json:"email"`
		Avatar     string `json:"avatar"`
		Coint      int64  `json:"coint"`
		Score      int64  `json:"score"`
		LanguageId int64  `json:"language_id"`
		LevelId    int64  `json:"level_id"`
		CreatedAt  string `json:"created_at"`
		Access     string `json:"access_token"`
		Refresh    string `json:"refresh_token"`
	}

	Forgot struct {
		Email string `json:"email"`
	}

	Reset struct {
		Otp         string `json:"otp"`
		Email       string `json:"email"`
		Newpassword string `json:"new_password"`
	}

	VerifyRequest struct {
		Otp string `json:"otp"`
	}

	VerifyResponse struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		LanguageId int64  `json:"language_id"`
		LevelId    int64  `json:"level_id"`
		Access     string `json:"access_token"`
		Refresh    string `json:"refresh_token"`
		CreatedAt  string `json:"created_at"`
	}

	UserInfo struct {
		ID        string `json:"id"`
		Login     uint64 `json:"login"`
		Role      string `json:"role"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)

func (rm *Signup) ValidateEmail() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Email, validation.Required, is.Email),
	)
}

func (rm *Signup) ValidatePassword() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(
			&rm.Password,
			validation.Required,
			validation.Length(8, 30),
			validation.Match(regexp.MustCompile("[a-z]|[1-9]")),
		),
	)
}

func (rm *Reset) ValidatePassword() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(
			&rm.Newpassword,
			validation.Required,
			validation.Length(8, 30),
			validation.Match(regexp.MustCompile("[a-z]|[1-9]")),
		),
	)
}
