package repositories

import (
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/alireza-fa/ghofle/internal/db/models"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
)

type UserRepository struct {
	rdbms rdbms.RDBMS
	log   logger.Logger
}

func NewUserRepository(cfg *config.Config) *UserRepository {
	r, err := rdbms.New(cfg.Postgres)
	if err != nil {
		panic(err)
	}
	return &UserRepository{
		rdbms: r,
		log:   logger.NewLogger(cfg.Logger),
	}
}

const QueryGetUserByEmail = `
SELECT id, username, email, hash_password FROM
email=$1;`

func (repository *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	in := []interface{}{email}
	user := models.User{}
	out := []interface{}{&user.Id, &user.Username, &user.Email, &user.HashPassword}

	if err := repository.rdbms.QueryRaw(QueryGetUserByEmail, in, out); err != nil {
		repository.log.Error(logger.RDBMS, logger.Select, "Error find user by email", nil)
		return nil, err
	}

	return &user, nil
}

const QueryCheckUserExistByUsername = `
SELECT 1 FROM users
WHERE username=$1
LIMIT 1;`

func (repository *UserRepository) CheckUserExistByUsername(username string) (bool, error) {
	in := []interface{}{username}
	var id uint64
	out := []interface{}{&id}

	if err := repository.rdbms.QueryRaw(QueryCheckUserExistByUsername, in, out); err != nil {
		repository.log.Error(logger.RDBMS, logger.Select, "Error while select user", nil)
		return false, err
	}

	return true, nil
}

const QueryCheckUserExistByEmail = `
SELECT 1 FROM users
WHERE email=$1
LIMIT 1;`

func (repository *UserRepository) CheckUserExistByEmail(email string) (bool, error) {
	in := []interface{}{email}
	var id uint64
	out := []interface{}{&id}

	if err := repository.rdbms.QueryRaw(QueryCheckUserExistByEmail, in, out); err != nil {
		repository.log.Error(logger.RDBMS, logger.Select, "Error while select user", nil)
		return false, err
	}

	return true, nil
}

const QueryCheckUserExist = `
SELECT id FROM users
WHERE username=$1 OR email=$2`

func (repository *UserRepository) CheckUserExist(user *models.User) (bool, error) {
	in := []interface{}{user.Username, user.Email}
	out := []interface{}{&user.Id}

	if err := repository.rdbms.QueryRaw(QueryCheckUserExist, in, out); err != nil {
		repository.log.Error(logger.RDBMS, logger.Select, "Error while select user", nil)
		return false, err
	}

	return true, nil
}

const QueryCreateUser = `
INSERT INTO users (username, email, hash_password) VALUES($1, $2, $3)
RETURNING id;`

func (repository *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	in := []interface{}{user.Username, user.Email, user.HashPassword}
	out := []interface{}{&user.Id}

	if err := repository.rdbms.QueryRaw(QueryCreateUser, in, out); err != nil {
		repository.log.Error(logger.RDBMS, logger.Insert, "Error inserting user", nil)
		return nil, err
	}

	return user, nil
}
