package db

import (
	"errors"
	"fmt"

	"github.com/S4mkiel/sso/domain/entity"
	user_errors "github.com/S4mkiel/sso/domain/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserRepository struct {
	orm *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{orm: db}
}

func (db UserRepository) Create(u *entity.User) (*entity.User, error) {
	err := db.orm.Create(&u).Error
	if err != nil {
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			switch pgError.Code {
			case "23505":
				if db.orm.Where(entity.User{Email: u.Email}).Take(&entity.User{}).Error == nil {
					return nil, user_errors.ErrDuplicatedEmail
				}
				if db.orm.Where(entity.User{ExternalID: u.ExternalID}).Take(&entity.User{}).Error == nil {
					return nil, user_errors.ErrDuplicatedExternalUID
				}
				if db.orm.Where(entity.User{Username: u.Username}).Take(&entity.User{}).Error == nil {
					return nil, user_errors.ErrDuplicatedUsername
				}
			}
		}

		return nil, err
	}
	return u, nil
}

func (db UserRepository) FindByID(uid string) (*entity.User, error) {
	var user entity.User
	result := db.orm.Where(fmt.Sprintf("id = '%v'", uid)).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}

	return &user, nil
}
