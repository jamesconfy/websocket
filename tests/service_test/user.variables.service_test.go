package service_test

import (
	"project-name/internal/forms"
	"project-name/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateUserForm() *forms.Create {
	return &forms.Create{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Password:    faker.Password(),
	}
}

func createAndRegisterUser(user *forms.Create) *models.User {
	if user == nil {
		user = generateUserForm()
	}

	resultUser, err := userSrv.Create(user)
	if err != nil {
		panic(err)
	}

	return resultUser
}
