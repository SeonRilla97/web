package repository

import "web/config"

type RootRepository struct {
	config *config.Config
}

func NewRootRepository(config *config.Config) (*RootRepository, error) {
	r := &RootRepository{
		config: config,
	}
	return r, nil
}
