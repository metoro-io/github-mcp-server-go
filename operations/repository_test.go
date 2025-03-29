package operations

import (
	"testing"
)

func TestSearchRepositoriesOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		options SearchRepositoriesOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: SearchRepositoriesOptions{
				Query:   "test",
				Page:    1,
				PerPage: 30,
			},
			wantErr: false,
		},
		{
			name: "missing query",
			options: SearchRepositoriesOptions{
				Query:   "",
				Page:    1,
				PerPage: 30,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.options.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateRepositoryOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		options CreateRepositoryOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: CreateRepositoryOptions{
				Name:        "test-repo",
				Description: "Test repository",
				Private:     true,
				AutoInit:    true,
			},
			wantErr: false,
		},
		{
			name: "missing name",
			options: CreateRepositoryOptions{
				Name:        "",
				Description: "Test repository",
				Private:     true,
				AutoInit:    true,
			},
			wantErr: true,
		},
		{
			name: "invalid name with spaces",
			options: CreateRepositoryOptions{
				Name:        "invalid repo name",
				Description: "Test repository",
				Private:     true,
				AutoInit:    true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.options.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
