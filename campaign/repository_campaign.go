package campaign

import "gorm.io/gorm"

type Repository interface{
	FindAllCampaignRepository() ([]Campaign, error)
	FindByUserIDCampaignRepository(UserID int) ([]Campaign, error)
	FindByIDCampaignRepository(ID int)(Campaign, error)
	SaveCampaignRepository(campaign Campaign) (Campaign, error)
	UpdateCampaignRepository(campaign Campaign) (Campaign , error)
	CreateImageRepository(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimaryRepository(campaignID int)(bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllCampaignRepository() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByUserIDCampaignRepository(userID int)([]Campaign, error){
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByIDCampaignRepository(ID int) (Campaign, error){
	var campaign Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) SaveCampaignRepository(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) UpdateCampaignRepository(campaign Campaign) (Campaign, error) {
	err :=r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) CreateImageRepository(campaignImage CampaignImage) (CampaignImage, error){
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimaryRepository(campaignID int)(bool, error){
	// UPDATE campaign_images SET is_primary = false WHARE campaign_id=1
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?",campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
