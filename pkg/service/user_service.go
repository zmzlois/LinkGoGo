package service

import (
	"context"
	"time"

	"github.com/zmzlois/LinkGoGo/ent"
	"github.com/zmzlois/LinkGoGo/pkg/model"
)

type UserService struct {
	db *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		db: client,
	}
}

func (s UserService) CreateUser(payload *model.LoginUserInput, tokenPayload *model.TokenInput, tokenString string) (*ent.Users, *ent.Session, error) {
	user, err := s.db.Users.Create().
		SetExternalID(payload.Id).
		SetUsername(payload.Username).
		SetAvatar(payload.Avatar).
		SetGlobalName(payload.GlobalName).
		SetSlug(payload.Username).
		SetAccessToken(tokenPayload.AccessToken).
		SetRefreshToken(tokenPayload.RefreshToken).
		SetExpiresIn(tokenPayload.ExpiresIn).
		SetDeleted(false).
		Save(context.Background())
	if err != nil {
		return nil, nil, err
	}
	t := time.Unix(int64(tokenPayload.ExpiresIn), 0)

	session, err := s.db.Session.Create().
		SetSessionToken(tokenString).
		SetExpiresAt(t).SetDeleted(false).
		SetUser(user).
		Save(context.Background())

	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}
