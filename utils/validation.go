package utils

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
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
			root := strings.Split(e.Namespace(), ".")[0]
			rawPath := strings.TrimPrefix(e.Namespace(), root+".")
			parts := strings.Split(rawPath, ".")

			log.Printf("parts: %s", parts)
			for i, part := range parts {
				if strings.Contains(part, "[") {
					idx := strings.Index(part, "[")
					base := pascalToSnakeCase(part[:idx]) // 0 to before [
					index := part[idx:]
					parts[i] = base + index
					log.Printf("idx: %v", idx)
					log.Printf("base: %v", base)
					log.Printf("index: %v", index)
				} else {
					parts[i] = pascalToSnakeCase(part)
				}
			}

			fieldPath := strings.Join(parts, ".")

			log.Printf("rawPath: %s", rawPath)
			log.Printf("fieldPath: %s", fieldPath)
			switch e.Tag() {
			case "gt":
				errs[fieldPath] = fmt.Sprintf("%s phải lớn hơn %s", fieldPath, e.Param())
			case "lt":
				errs[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn %s", fieldPath, e.Param())
			case "gte":
				errs[fieldPath] = fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", fieldPath, e.Param())
			case "lte":
				errs[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", fieldPath, e.Param())
			case "uuid":
				errs[fieldPath] = fmt.Sprintf("%s phải là UUID hợp lệ", fieldPath)
			case "slug":
				errs[fieldPath] = fmt.Sprintf("%s chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm", fieldPath)
			case "min":
				errs[fieldPath] = fmt.Sprintf("%s phải lớn hơn %s ký tự", fieldPath, e.Param())
			case "max":
				errs[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn %s ký tự", fieldPath, e.Param())
			case "min_int":
				errs[fieldPath] = fmt.Sprintf("%s phải có giá trị lớn hơn %s", fieldPath, e.Param())
			case "max_int":
				errs[fieldPath] = fmt.Sprintf("%s phải có giá trị nhỏ hơn %s", fieldPath, e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				errs[fieldPath] = fmt.Sprintf("%s phải là một trong các giá trị: %s", fieldPath, allowedValues)
			case "required":
				errs[fieldPath] = fmt.Sprintf("%s là bắt buộc", fieldPath)
			case "search":
				errs[fieldPath] = fmt.Sprintf("%s chỉ được chứa chữ thường, in hoa, số và khoảng trắng", fieldPath)
			case "email":
				errs[fieldPath] = fmt.Sprintf("%s phải đúng định dạng là email", fieldPath)
			case "datetime":
				errs[fieldPath] = fmt.Sprintf("%s phải đúng định dạng YYYY-MM-DD", fieldPath)
			case "file_ext":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				errs[fieldPath] = fmt.Sprintf("%s chỉ cho phép những file có extension: %s", fieldPath, allowedValues)
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

	errMinInt := v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minVal
	})
	if errMinInt != nil {
		return errMinInt
	}

	errMaxInt := v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= maxVal
	})
	if errMaxInt != nil {
		return errMaxInt
	}

	errFileExt := v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()
		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}
		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExt {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}

		return false

	})
	if errFileExt != nil {
		return errFileExt
	}

	return nil
}
