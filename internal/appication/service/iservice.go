package service

type IService interface {
	Get() (interface{}, error)
	LoadCache() (interface{}, error)
}
