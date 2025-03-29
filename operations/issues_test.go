package operations

import (
	"strings"
	"testing"
)

func TestCreateIssueOptionsValidate(t *testing.T) {
	tests := []struct {
		name          string
		options       CreateIssueOptions
		wantErr       bool
		errorContains string
	}{
		{
			name: "missing title",
			options: CreateIssueOptions{
				Owner: "owner123",
				Repo:  "valid-repo",
				Title: "",
				Body:  "This is a test issue",
			},
			wantErr:       true,
			errorContains: "title is required",
		},
		{
			name: "invalid owner",
			options: CreateIssueOptions{
				Owner: "invalid owner",
				Repo:  "valid-repo",
				Title: "Test Issue",
				Body:  "This is a test issue",
			},
			wantErr:       true,
			errorContains: "owner name",
		},
		{
			name: "invalid repo",
			options: CreateIssueOptions{
				Owner: "owner123",
				Repo:  "invalid repo",
				Title: "Test Issue",
				Body:  "This is a test issue",
			},
			wantErr:       true,
			errorContains: "repository name",
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

func TestGetIssueOptionsValidate(t *testing.T) {
	tests := []struct {
		name          string
		options       GetIssueOptions
		wantErr       bool
		errorContains string
	}{
		{
			name: "invalid number",
			options: GetIssueOptions{
				Owner:  "owner123",
				Repo:   "valid-repo",
				Number: 0,
			},
			wantErr:       true,
			errorContains: "issue number",
		},
		{
			name: "negative number",
			options: GetIssueOptions{
				Owner:  "owner123",
				Repo:   "valid-repo",
				Number: -1,
			},
			wantErr:       true,
			errorContains: "issue number",
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

func TestListIssuesOptionsValidate(t *testing.T) {
	tests := []struct {
		name          string
		options       ListIssuesOptions
		wantErr       bool
		errorContains string
	}{
		{
			name: "invalid state",
			options: ListIssuesOptions{
				Owner: "owner123",
				Repo:  "valid-repo",
				State: "invalid",
			},
			wantErr:       true,
			errorContains: "state must be",
		},
		{
			name: "invalid sort",
			options: ListIssuesOptions{
				Owner: "owner123",
				Repo:  "valid-repo",
				Sort:  "invalid",
			},
			wantErr:       true,
			errorContains: "sort must be",
		},
		{
			name: "invalid direction",
			options: ListIssuesOptions{
				Owner:     "owner123",
				Repo:      "valid-repo",
				Direction: "invalid",
			},
			wantErr:       true,
			errorContains: "direction must be",
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
