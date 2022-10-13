package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignIDRepository(campaignID int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignIDRepository(campaignID int) ([]Transaction, error) {
	var transaction []Transaction

	//ORDER mengurutkan data

	err := r.db.Preload("User").Where("campaigns_id = ?", campaignID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}
