package service

import (
	"project-name/internal/forms"
	"project-name/internal/models"
	repo "project-name/internal/repository"
	"project-name/internal/se"
)

type UserService interface {
	Create(req *forms.Create) (*models.User, *se.ServiceError)
	Login(req *forms.Login) (*models.Auth, *se.ServiceError)
}

type userSrv struct {
	userRepo     repo.UserRepo
	authRepo     repo.AuthRepo
	validatorSrv ValidationService
	cryptoSrv    CryptoService
	authSrv      AuthService
	emailSrv     EmailService
}

func (u *userSrv) Create(req *forms.Create) (*models.User, *se.ServiceError) {
	err := u.validatorSrv.Validate(req)
	if err != nil {
		return nil, se.Validating(err)
	}

	if ok, err := u.userRepo.EmailExists(req.Email); ok {
		return nil, se.ConflictOrInternal(err, "user already exists")
	}

	password, err := u.cryptoSrv.HashPassword(req.Password)
	if err != nil {
		return nil, se.Internal(err, "could not hash password")
	}

	var user models.User

	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Password = password

	usr, err := u.userRepo.Add(&user)
	if err != nil {
		return nil, se.Internal(err)
	}

	return usr, nil
}

func (u *userSrv) Login(req *forms.Login) (*models.Auth, *se.ServiceError) {
	err := u.validatorSrv.Validate(req)
	if err != nil {
		return nil, se.Validating(err)
	}

	user, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "user does not exist")
	}

	ok := u.cryptoSrv.ComparePassword(user.Password, req.Password)
	if !ok {
		return nil, se.BadRequest("password does not match")
	}

	auth := new(models.Auth)

	auth.AccessToken, auth.RefreshToken, err = u.authSrv.Create(user.Id, user.Email)
	if err != nil {
		return nil, se.Internal(err, "Error when creating token")
	}

	// Create auth row
	ath, err := u.authRepo.Add(auth)
	if err != nil {
		return nil, se.Internal(err, "Error when adding/updating user token")
	}

	return ath, nil
}

func NewUserSrv(repo repo.UserRepo, authRepo repo.AuthRepo, validator ValidationService, cryptoSrv CryptoService, authSrv AuthService, emailSrv EmailService) UserService {
	return &userSrv{userRepo: repo, authRepo: authRepo, validatorSrv: validator, cryptoSrv: cryptoSrv, authSrv: authSrv, emailSrv: emailSrv}
}
