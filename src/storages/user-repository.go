package storages

import (
	"fmt"

	"boiler_plate.com/src/configuration/environments"
	"boiler_plate.com/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IUserRepository interface {
	GetUser(username, password string) models.User
	CreateUser(username, password string) models.User
}

type userRepository struct {
	DbConnection *gorm.DB
}

func (u *userRepository) GetUser(username, password string) models.User {
	var user models.User

	u.DbConnection.Where("username = ? AND password = ?", username, password).Find(&user)
	u.DbConnection.Close()

	return user
}

func (u *userRepository) CreateUser(username, password string) models.User {
	user := models.User{Username: username, Password: password}

	u.DbConnection.Create(&user)
	u.DbConnection.Close()

	return user
}

func NewUserRepository() (IUserRepository, error) {
	env := environments.GetEnvironment()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		env.PostgresHost, env.PostgresPort, env.PostgresUser, env.PostgresDatabaseName, env.PostgresPassword)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return &userRepository{
		DbConnection: db,
	}, nil
}
