package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"unicode"
)

func ValidateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.New("invalid id")
	}

	return nil
}

func ValidateSlug(slug string) error {
	for _, char := range slug {
		if unicode.IsDigit(char) {
			return fmt.Errorf("slug shouldn't contain digits")
		}
	}

	specialCharRegex := regexp.MustCompile(`[^a-zA-Zа-яА-Я ]`)
	if specialCharRegex.MatchString(slug) {
		return fmt.Errorf("slug shouldn't contain any special symbols")
	}

	return nil
}
