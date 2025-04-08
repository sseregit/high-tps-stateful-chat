package service

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang-chat-controller/repository"
	"golang-chat-controller/types/table"
	"log"
)

type Service struct {
	repository    *repository.Repository
	AvgServerList map[string]bool
}

func NewService(repository *repository.Repository) *Service {
	s := &Service{repository: repository, AvgServerList: make(map[string]bool)}

	s.setServerInfo()

	if err := s.repository.Kafka.RegisterSubTopic("chat"); err != nil {
		panic(err)
	} else {
		go s.loopSubKafka()
	}

	return s
}

func (s *Service) loopSubKafka() {
	for {
		ev := s.repository.Kafka.Poll(100)

		switch event := ev.(type) {
		case *kafka.Message:
			type ServerInfoEvent struct {
				IP     string
				Status bool
			}

			var decoder ServerInfoEvent

			if err := json.Unmarshal(event.Value, &decoder); err != nil {
				log.Println("Failed To Decode Event", event.Value)
			} else {
				fmt.Println(decoder)
				s.AvgServerList[decoder.IP] = decoder.Status
			}

		case *kafka.Error:
			log.Println("Failed To Polling Event", event.Error())
		}
	}
}

func (s *Service) GetAvgServerList() []string {
	var res []string

	for ip, available := range s.AvgServerList {
		if available {
			res = append(res, ip)
		}
	}

	return res
}

func (s *Service) setServerInfo() {
	if serverList, err := s.GetAvailableServerList(); err != nil {
		panic(err)
	} else {
		for _, server := range serverList {
			s.AvgServerList[server.IP] = true
		}
	}
}

func (s *Service) GetAvailableServerList() ([]*table.Serverinfo, error) {
	return s.repository.GetAvailableServerList()
}
