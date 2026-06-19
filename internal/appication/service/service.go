package service

import (
	"io"
	"log"
	"net/http"
)

type Service struct {
	data []byte
}

func NewService() IService {
	s := &Service{}
	s.LoadCache()
	return s
}

func (s *Service) LoadCache() error {
	log.Println("getting cache data")

	url := "https://api.victorsanmartin.com/holidays/en.json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return err
	}

	s.data = body
	return nil
}

func (s *Service) Get() []byte {
	return s.data
}
