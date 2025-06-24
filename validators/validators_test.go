package validators

import (
	"testing"
)

func TestValidateTaskInput(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		status      string
		wantErrs    int
	}{
		{
			name:        "Valid input",
			title:       "Write README",
			description: "Prepare the final README.md",
			status:      "pending",
			wantErrs:    0,
		},
		{
			name:        "Empty title",
			title:       "",
			description: "Some desc",
			status:      "completed",
			wantErrs:    1,
		},
		{
			name:        "Title too short",
			title:       "Hi",
			description: "Desc",
			status:      "pending",
			wantErrs:    1,
		},
		{
			name:        "Title too long",
			title:       string(make([]byte, 101)), // 101-char string
			description: "Desc",
			status:      "pending",
			wantErrs:    1,
		},
		{
			name:        "Description too long",
			title:       "Task title",
			description: string(make([]byte, 301)), // 301-char string
			status:      "pending",
			wantErrs:    1,
		},
		{
			name:        "Invalid status",
			title:       "Task title",
			description: "Valid desc",
			status:      "ongoing",
			wantErrs:    1,
		},
		{
			name:        "Multiple errors",
			title:       "",
			description: string(make([]byte, 301)),
			status:      "invalid",
			wantErrs:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := ValidateTaskInput(tt.title, tt.description, tt.status)
			if len(errs) != tt.wantErrs {
				t.Errorf("expected %d errors, got %d", tt.wantErrs, len(errs))
			}
		})
	}
}

