package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func ValidationStringLength(fieldName, value string, min, max int) error {
	l := len(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be between %d and %d characters", fieldName, min, max)
	}

	return nil
}

func ValidationRegex(fieldName, value string, re *regexp.Regexp) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s must match %s", fieldName, re)
	}

	return nil
}

func ValidationPositiveInt(fieldName, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a number", fieldName)
	}

	if v < 0 {

		return 0, fmt.Errorf("%s must be positive", fieldName)
	}

	return v, nil
}

func ValidationUuid(fieldName, value string) (*uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return &uid, fmt.Errorf("%s must be a valid Uuid", fieldName)
	}

	return &uid, nil
}

func ValidationIntList(fieldName, value string, allowed map[string]bool) error {
	if !allowed[value] {
		return fmt.Errorf("%s must be one of: %v", fieldName, keys(allowed))
	}

	return nil
}

func keys(m map[string]bool) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}

func HandleValidationErrors(err error) gin.H {
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := make(map[string]string)

		for _, e := range validationError {
			log.Printf("%+v", e.Tag())
			switch e.Tag() {
			case "gt":
				errs[e.Field()] = e.Field() + " phai lon hon gia tri toi thieu"
			case "uuid":

				errs[e.Field()] = e.Field() + " khong dung dinh dang"
			}

		}
		return gin.H{"error": errs}
	}

	return gin.H{"error": "Yeu cau khong hop le " + err.Error()}
}
