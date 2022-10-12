package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface{
	GetCampaignsService(userID int) ([]Campaign, error)
	GetCampaignByIDService(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaignService(input CreateCampaignInput) (Campaign, error)
	UpdateCampaignService(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImageService(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

	if campaign.UserID != inputData.User.ID{
		return campaign, errors.New("Not an owner of the campaign")
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

func (s *service) SaveCampaignImageService(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error){
	campaign, err := s.repository.FindByIDCampaignRepository(input.CampaignID)

	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID{
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary{
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimaryRepository(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}
	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImageRepository(campaignImage)

	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}


