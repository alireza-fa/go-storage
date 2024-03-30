package repositories

import (
	"github.com/alireza-fa/ghofle/internal/db/models"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
)

type UserRepository struct {
	rdbms rdbms.RDBMS
}

const QueryGetUserByEmail = `
SELECT id, username, email, hash_password FROM
email=$1;`

func (repository *UserRepository) GetUserByEmail(email string) (*models.User, error) {
}
