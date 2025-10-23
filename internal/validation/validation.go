package validation

import (
	"errors"
	"fmt"
	"learn-golang/internal/utils"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	err := RegisterCustomValidation(v)
	if err != nil {
		return err
	}

	return nil
}

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
					base := utils.PascalToSnakeCase(part[:idx]) // 0 to before [
					index := part[idx:]
					parts[i] = base + index
					log.Printf("idx: %v", idx)
					log.Printf("base: %v", base)
					log.Printf("index: %v", index)
				} else {
					parts[i] = utils.PascalToSnakeCase(part)
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
			case "email_advanced":
				errs[fieldPath] = fmt.Sprintf("%s này nằm trong danh sách cấm", fieldPath)
			case "datetime":
				errs[fieldPath] = fmt.Sprintf("%s phải đúng định dạng YYYY-MM-DD", fieldPath)
			case "password_strong":
				errs[fieldPath] = fmt.Sprintf("%s phải ít nhất 8 ký tự bao gồm (chữ thường, chữ in hoa, số và ký tự đặc biệt)", fieldPath)
			case "file_ext":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				errs[fieldPath] = fmt.Sprintf("%s chỉ cho phép những file có extension: %s", fieldPath, allowedValues)
			}

		}
		return gin.H{"error": errs}
	}

	return gin.H{
		"error":  "Yêu cầu không hợp lệ",
		"detail": err.Error(),
	}
}
