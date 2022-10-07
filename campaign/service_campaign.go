package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface{
	GetCampaignsService(userID int) ([]Campaign, error)
	GetCampaignByIDService(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaignService(input CreateCampaignInput) (Campaign, error)
	UpdateCampaignService(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaignsService(userID int) ([]Campaign, error){
	if userID != 0 {
		campaigns, err := s.repository.FindByUserIDCampaignRepository(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAllCampaignRepository()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByIDService(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByIDCampaignRepository(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaignService(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.SaveCampaignRepository(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *service) UpdateCampaignService(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error){
	campaign, err := s.repository.FindByIDCampaignRepository(inputID.ID)

	if err != nil {
		return campaign, err
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.UpdateCampaignRepository(campaign)

	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}

