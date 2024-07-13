package auth

import (
	"errors"
	"gobase/db"
	"gorm.io/gorm"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"gobase/config"
	"gobase/models/entities"
	"gobase/models/repos"
	"gobase/services/auth/dto"
	userDTO "gobase/services/users/dto"
	"gobase/utils"
)

var (
	JWTExpiration = 24 * time.Hour
)

type AuthService struct {
	userRepo  repos.UserRepository
	redisRepo repos.RedisRepository
}

// NewAuthService creates a new instance of AuthService.
//
// No parameters.
// Returns a pointer to AuthService.
func NewAuthService() *AuthService {
	userRepo := repos.NewUserRepository(db.GetDB())
	redisRepo := repos.NewRedisRepository()
	return &AuthService{userRepo: *userRepo, redisRepo: *redisRepo}
}

func (au *AuthService) Login(ctx *gin.Context, req dto.LoginRequest) (string, error) {
	user, err := au.userRepo.FindByUsername(req.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		errors.New("Incorrect username or password")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		errors.New("Incorrect username or password")
	}

	token, claims, err := utils.CreateJWTToken(user, config.GetConfig().Security.JWTSecret)
	if err != nil {
		return "", err
	}
	err = au.redisRepo.SetWithDuration(ctx, user.Username, token, JWTExpiration)
	if err != nil {
		return "", errors.New("Login unsuccessful")
	}
	ctx.Set("claims", claims)
	return token, nil
}

func (au *AuthService) Register(ctx *gin.Context, req dto.RegisterRequest) (userDTO.UserDTO, error) {
	if len(req.Password) < 8 {
		return userDTO.UserDTO{}, errors.New("Password must be at least 8 characters")
	}

	existingUser, err := au.userRepo.FindByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return userDTO.UserDTO{}, err
	}
	if existingUser != nil && existingUser.ID != 0 {
		return userDTO.UserDTO{}, errors.New("User already exists")
	}

	password, err := utils.BcryptHash(req.Password)
	if err != nil {
		return userDTO.UserDTO{}, errors.New("Failed to create password")
	}

	user := &entities.User{
		Username: req.Username,
		Email:    req.Email,
		Password: password,
	}

	err = au.userRepo.Create(user)
	if err != nil {
		return userDTO.UserDTO{}, err
	}
	return userDTO.ToUserDTO(*user), nil
}

func (au *AuthService) Logout(ctx *gin.Context) error {
	return nil
}
