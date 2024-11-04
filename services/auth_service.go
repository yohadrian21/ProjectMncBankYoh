package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"temp.project/transferbank/dtos"
	"temp.project/transferbank/models"
	"temp.project/transferbank/repositories"
)

type authService struct {
	db             *gorm.DB
	UserRepository repositories.UserRepository
}

type AuthService interface {
	Register(registerDto *dtos.RegisterDto) error
	Login(loginDto *dtos.LoginDto) (string, error) // Add the Login method
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{UserRepository: userRepo}
}
func (s *authService) Register(registerDto *dtos.RegisterDto) error {
	// Check if the user already exists
	var existingUser models.User
	if err := s.db.Where("username = ?", registerDto.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	// Create new user
	newUser := models.User{
		Username: registerDto.Username,
		Password: registerDto.Password, // Consider hashing the password before saving
		Balance:  0,                    // Set default balance
	}

	if err := s.db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(loginDto *dtos.LoginDto) (string, error) {
	var user models.User
	// Find user by username
	if err := s.db.Where("username = ?", loginDto.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid username or password")
		}
		return "", err // Return any other errors
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate token (for simplicity, we are just returning a mock token)
	token := "mockToken" // Replace this with actual token generation logic (e.g., JWT)
	log.Printf("User %s logged in successfully", user.Username)
	return token, nil
}

func generateToken(username string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   username,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your_secret_key")) // Replace with your secret key
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
