package operations

import (
	"strings"
	"testing"
)

func TestListCommitsOptionsValidate(t *testing.T) {
	tests := []struct {
		name          string
		options       ListCommitsOptions
		wantErr       bool
		errorContains string
	}{
		{
			name: "invalid owner",
			options: ListCommitsOptions{
				Owner:  "invalid owner",
				Repo:   "valid-repo",
				Branch: "main",
			},
			wantErr:       true,
			errorContains: "owner name",
		},
		{
			name: "invalid repo",
			options: ListCommitsOptions{
				Owner:  "validowner",
				Repo:   "invalid repo",
				Branch: "main",
			},
			wantErr:       true,
			errorContains: "repository name",
		},
		{
			name: "invalid branch",
			options: ListCommitsOptions{
				Owner:  "validowner",
				Repo:   "valid-repo",
				Branch: "invalid branch",
			},
			wantErr:       true,
			errorContains: "branch name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.options.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && tt.errorContains != "" {
				if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Validate() error = %v, should contain %v", err, tt.errorContains)
				}
			}
		})
	}
}
