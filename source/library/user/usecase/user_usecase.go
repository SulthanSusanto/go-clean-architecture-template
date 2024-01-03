package usecase

import (
	"context"
	"go-clean-architecture/source/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepo       entities.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(ur entities.UserRepository, to time.Duration) entities.UserUsecase {
	return &userUsecase{
		userRepo:       ur,
		contextTimeout: to,
	}
}

func (user *userUsecase) InsertOne(c context.Context, m *entities.User) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := user.userRepo.InsertOne(ctx, m)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (user *userUsecase) FindOne(c context.Context, id string) (*entities.User, error) {

	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	res, err := user.userRepo.FindOne(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (user *userUsecase) GetAllPagination(c context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]entities.User, int64, error) {

	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	res, count, err := user.userRepo.GetAllPagination(ctx, rp, p, filter, setsort)
	if err != nil {
		return res, count, err
	}

	return res, count, nil
}

func (user *userUsecase) UpdateOne(c context.Context, m *entities.User, id string) (*entities.User, error) {

	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	res, err := user.userRepo.UpdateOne(ctx, m, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
