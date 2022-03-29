package config

import "context"

const (
	EmTypeNone   = 0
	EmTypeBase64 = 1
	EmTypeAES256 = 2
)

type Config interface {
	LoadConfigFromFile(ctx context.Context, filename string) error
	SaveConfigToFile(ctx context.Context, filename string) error
	SetValue(ctx context.Context, key string, value string, emType int) error
	GetValue(ctx context.Context, value string) (string, error)
}

type cfg struct {
	filename string
	records  []rec
}

func NewConfig(filename string) *cfg {
	return &cfg{filename: filename}
}

type rec struct {
	key           string
	value         string
	isEncrypted   bool
	encryptMethod int
}

func (r *rec) LoadConfigFromFile(ctx context.Context, filename string) error {
	//TODO implement me
	panic("implement me")
}

func (r *rec) SaveConfigToFile(ctx context.Context, filename string) error {
	//TODO implement me
	panic("implement me")
}

func (r *rec) SetValue(ctx context.Context, key string, value string, emType int) error {
	//TODO implement me
	panic("implement me")
}

func (r *rec) GetValue(ctx context.Context, value string) (string, error) {
	//TODO implement me
	panic("implement me")
}
