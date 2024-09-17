package service_test

import (
	"websocket/internal/forms"
	"testing"
)

func TestAddUser(t *testing.T) {
	// Create a new user object
	user := generateUserForm()

	tests := []struct {
		name    string
		user    *forms.Create
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Create(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	// Create a new user object
	userForm := generateUserForm()
	user := createAndRegisterUser(userForm)

	tests := []struct {
		name    string
		user    *forms.Login
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: &forms.Login{Email: user.Email, Password: userForm.Password}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Login(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: user.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	// Create a new user object

	for i := 0; i < 10; i++ {
		_ = createAndRegisterUser(nil)
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
