package validators

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateTaskInput(title, description, status string) []ValidationError {
	var errs []ValidationError

	// Title: required, min 3 chars, max 100
	if strings.TrimSpace(title) == "" {
		errs = append(errs, ValidationError{"title", "Title is required."})
	} else if len(title) < 3 {
		errs = append(errs, ValidationError{"title", "Title must be at least 3 characters."})
	} else if len(title) > 100 {
		errs = append(errs, ValidationError{"title", "Title must be less than 100 characters."})
	}

	// Description: optional, max 300
	if len(description) > 300 {
		errs = append(errs, ValidationError{"description", "Description must be less than 300 characters."})
	}

	// Status: required, must be one of allowed values
	allowedStatuses := []string{"pending", "completed", "archived"}
	validStatus := false
	for _, s := range allowedStatuses {
		if status == s {
			validStatus = true
			break
		}
	}
	if !validStatus {
		errs = append(errs, ValidationError{"status", fmt.Sprintf("Status must be one of: %v.", allowedStatuses)})
	}

	return errs
}

