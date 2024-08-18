package transaction

import (
	"crowdfunding/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transactions, error)
	GetTransactionByUserID(userID int) ([]Transactions, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transactions, error) {

	//get campaign by id
	//check campaign.userid != user.id yang melakukan request

	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return []Transactions{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transactions{}, errors.New("Unauthorized")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTransactionByUserID(userID int) ([]Transactions, error) {
	transactions, err := s.repository.GetByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
