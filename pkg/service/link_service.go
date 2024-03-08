package service

import (
	"context"
	"fmt"

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

func (b *LinkService) GetLinks(ctx context.Context) ([]model.NewLinkInput, error) {

	var userLinks = []model.NewLinkInput{}

	externalId := ctx.Value("user_id").(string)

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
			Id:          link.ID.String(),
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

func (b *LinkService) AddLink(ctx context.Context, externalId string, title string, url string) (*ent.Links, error) {

	links, err := b.GetLinks(ctx)

	if err != nil {
		return nil, fmt.Errorf("link_service.AddLink.links: %w", err)
	}

	linkLength := len(links)

	task1, err := b.GetUser(ctx)

	if err != nil {
		return nil, fmt.Errorf("link_service.AddLink.task1: %w", err)
	}

	task2, err := b.db.Links.
		Create().
		SetTitle(title).
		SetURL(url).
		SetUser(task1).
		SetOrder(linkLength + 1).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("link_service.AddLink.task2: %w", err)
	}

	return task2, nil
}

func (b *LinkService) GetUser(ctx context.Context) (*ent.Users, error) {
	externalId := ctx.Value("user_id").(string)

	var task1 *ent.Users

	task1, err := b.db.Users.Query().Where(users.ExternalID(externalId)).Only(ctx)

	if err != nil {
		return task1, fmt.Errorf("link_service.GetUser.task1: %w", err)
	}

	return task1, nil
}
