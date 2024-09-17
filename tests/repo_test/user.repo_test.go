package repo_test

import (
	"websocket/internal/models"
	"testing"

	"github.com/bxcodec/faker/v4"
)

func TestAddUser(t *testing.T) {
	// Create a new user object
	user := generateUser()

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.Add(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailExists(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{name: "Test with correct email", email: user.Email, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := u.EmailExists(tt.email)
			if got != tt.want {
				t.Errorf("userSql.EmailExists() got = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestPhoneExists(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{name: "Test with correct phone", phone: user.PhoneNumber, want: true},
		{name: "Test with wrong phone", phone: faker.Phonenumber(), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := u.PhoneExists(tt.phone)
			if got != tt.want {
				t.Errorf("userSql.PhoneExists() got = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestGetByEmail(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{name: "Test with correct email", email: user.Email, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.GetByEmail(tt.email)
			if (err != nil) != tt.want {
				t.Errorf("userSql.GetByEmail() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name string
		id   string
		want bool
	}{
		{name: "Test with correct id", id: user.Id, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.GetById(tt.id)
			if (err != nil) != tt.want {
				t.Errorf("userSql.GetById() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		_ = createAndAddUser(nil)
	}

	tests := []struct {
		name string
		want bool
	}{
		{name: "Test with correct details", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.GetAll()
			if (err != nil) != tt.want {
				t.Errorf("userSql.GetAll() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
