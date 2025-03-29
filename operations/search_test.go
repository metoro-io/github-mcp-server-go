package operations

import (
	"testing"
)

func TestSearchCodeOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		options SearchCodeOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: SearchCodeOptions{
				Query:   "test in:file language:go",
				Page:    1,
				PerPage: 30,
			},
			wantErr: false,
		},
		{
			name: "missing query",
			options: SearchCodeOptions{
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

func TestSearchIssuesOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		options SearchIssuesOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: SearchIssuesOptions{
				Query:   "test in:title state:open",
				Page:    1,
				PerPage: 30,
			},
			wantErr: false,
		},
		{
			name: "missing query",
			options: SearchIssuesOptions{
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

func TestSearchUsersOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		options SearchUsersOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: SearchUsersOptions{
				Query:   "test type:user",
				Page:    1,
				PerPage: 30,
			},
			wantErr: false,
		},
		{
			name: "missing query",
			options: SearchUsersOptions{
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
