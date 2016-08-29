package indicator

import "github.com/devuo/go-app/service"

type Repository struct{}

func (r Repository) Fetch() ([]Indicator, error) {
	var indicators []Indicator
	_, err := service.DB.Select("*").From("Indicator").Load(&indicators)
	return indicators, err
}
