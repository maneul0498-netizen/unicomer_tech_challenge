package service

type IService interface {
	Get() []byte
	LoadCache() error
}
