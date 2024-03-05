package services

import (
	"context"
)

func (s *Service) Hello(_ context.Context) map[string]string {
	return map[string]string{
		"page_title":    "Выбор игровых команд",
		"random_team_1": "Одутловатые",
		"random_team_2": "Отрафиты",
	}
}
