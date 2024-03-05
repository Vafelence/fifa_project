package services

import (
	"context"

	"github.com/pkg/errors"
)

type Player struct {
	ID   string
	Name string
}

func (p Player) IsNew() bool {
	return p.ID == ""
}

func (s *Service) PlayerManagement(ctx context.Context, player Player) error {
	if player.IsNew() {
		err := s.storage.AddPlayer(ctx, player.Name)
		if err != nil {
			return errors.Wrapf(err, "can't add to storage player data %s: %v", err, player)
		}
	} else {
		err := s.storage.UpdatePlayer(ctx, player.ID, player.Name)
		if err != nil {
			return errors.Wrapf(err, "can't update to storage player data %s: %v", err, player)
		}
	}
	return nil
}

func (p *Service) Delete(ctx context.Context, playerID string) error {
	return nil
}
