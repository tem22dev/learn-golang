package validation

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidation(v *validator.Validate) error {
	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	errSlugRegex := v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})
	if errSlugRegex != nil {
		return errSlugRegex
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
