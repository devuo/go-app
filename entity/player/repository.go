package player

import "github.com/devuo/go-app/service"

type Repository struct{}

func (r Repository) Fetch() ([]Player, error) {
	var players []Player
	_, err := service.DB.Select("*").From("Player").Load(&players)
	return players, err
}
