package handler_test

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

func generateLoginForm(user *forms.Create) *forms.Login {
	if user == nil {
		user = generateUserForm()

		_ = createAndRegisterUser(user)
	}

	return &forms.Login{
		Email:    user.Email,
		Password: user.Password,
	}
}

// func loginUserAndGenerateAuth(login *forms.Login) string {
// 	if login == nil {
// 		login = generateLoginForm(nil)
// 	}

// 	auth, err := userSrv.Login(login)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return fmt.Sprintf("Bearer %v", auth.AccessToken)
// }
