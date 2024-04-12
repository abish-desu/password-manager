package validate

import (
	"errors"
)

func ValidateCreateCommand(args []string) (string, error) {
	if len(args) < 3 {
		return "", errors.New("key and value are missing")
	}
	if len(args) < 4 {
		return "", errors.New("value is missing")
	}
	if len(args) > 4 {
		return "", errors.New("invalid command")
	}
	return "Prompt valid ✅ ", nil
}
