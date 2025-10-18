package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) gin.H {
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := make(map[string]string)

		for _, e := range validationError {
			log.Printf("%+v", e.Tag())
			switch e.Tag() {
			case "gt":
				errs[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu"
			case "lt":
				errs[e.Field()] = e.Field() + " phải nhỏ hơn giá trị tối thiểu"
			case "gte":
				errs[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng giá trị tối thiểu"
			case "lte":
				errs[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng giá trị tối thiểu"
			case "uuid":
				errs[e.Field()] = e.Field() + " khong dung dinh dang"
			case "slug":
				errs[e.Field()] = e.Field() + " chứa chữ thường, số hoặc dấu gạch ngang"
			case "min":
				errs[e.Field()] = fmt.Sprintf("%s phải lớn hơn %s ký tự", e.Field(), e.Param())
			case "max":
				errs[e.Field()] = fmt.Sprintf("%s phải nhỏ hơn %s ký tự", e.Field(), e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				errs[e.Field()] = fmt.Sprintf("%s phải là một trong các giá trị: %s", e.Field(), allowedValues)
			case "required":
				errs[e.Field()] = e.Field() + " là bắt buộc"
			case "search":
				errs[e.Field()] = e.Field() + " chỉ được chứa chữ thường, in hoa, số và khoảng trắng"
			}

		}
		return gin.H{"error": errs}
	}

	return gin.H{"error": "Yeu cau khong hop le " + err.Error()}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("Failed to get validator engine")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	err := v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})
	if err != nil {
		return err
	}

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	errSearchRegex := v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})
	if errSearchRegex != nil {
		return errSearchRegex
	}

	return nil
}
