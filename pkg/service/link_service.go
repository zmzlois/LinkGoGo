package service

import (
	"context"

	"github.com/a-h/templ"
	"github.com/zmzlois/LinkGoGo/ent"
	"github.com/zmzlois/LinkGoGo/ent/links"
	"github.com/zmzlois/LinkGoGo/ent/users"
	"github.com/zmzlois/LinkGoGo/pkg/model"
)

type LinkService struct {
	db *ent.Client
}

func NewLinkService(db *ent.Client) *LinkService {
	return &LinkService{db: db}
}

func (b *LinkService) GetLinks(ctx context.Context, externalId string) ([]model.NewLinkInput, error) {

	var userLinks = []model.NewLinkInput{}

	task, err := b.db.Links.
		Query().
		Where(links.HasUserWith(users.ExternalID(externalId))).
		Order(ent.Asc(links.FieldOrder)).
		All(ctx)

	if err != nil {
		return userLinks, err
	}

	if len(task) == 0 {
		return userLinks, nil
	}

	for _, link := range task {

		safeURL := templ.SafeURL(link.URL)

		userLinks = append(userLinks, model.NewLinkInput{
			Title:       link.Title,
			Description: link.Description,
			Url:         safeURL,
			Order:       link.Order,
		})
		if link.URL == "" {
			continue
		}
	}

	return userLinks, nil
}
