package service

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

func ValidateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.New("invalid UUID format")
	}
	return nil
}

func ValidateSlug(slug string) error {
	match, err := regexp.MatchString("^[a-z0-9]+(?:-[a-z0-9]+)*$", slug)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("invalid slug format")
	}
	return nil
}
