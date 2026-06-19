package service

import (
	"io"
	"log"
	"net/http"
)

type Service struct {
	data interface{}
}

func NewService() IService {
	s := &Service{}
	s.LoadCache()
	return s
}

func (s *Service) LoadCache() (interface{}, error) {
	log.Println("getting data")

	url := "https://api.victorsanmartin.com/feriados/en.json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(string(body))

	s.data = body
	return s.data, nil
}

func (s *Service) Get() (interface{}, error) {
	return s.data, nil
}
