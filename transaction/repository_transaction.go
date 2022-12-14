package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignIDRepository(campaignID int) ([]Transaction, error)
	GetByUserIDTransactionRepository(userID int) ([]Transaction, error)
	GetByIDTransactionRepository(ID int) (Transaction, error)
	SaveTransactionRepository(transaction Transaction) (Transaction, error)
	UpdateTransactionRepository(transaction Transaction) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignIDRepository(campaignID int) ([]Transaction, error) {
	var transaction []Transaction

	//ORDER mengurutkan data

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

func (r *repository) GetByUserIDTransactionRepository(userID int) ([]Transaction, error) {
	var transactionS []Transaction

	// CARA MENGHUBUNGKAN RELASI DATABASE YG TIDAK PUNYA RELASI LANGSUNG
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id=?", userID).Order("id desc").Find(&transactionS).Error

	if err != nil {
		return transactionS, err
	}

	return transactionS, nil
}

func (r *repository) GetByIDTransactionRepository(ID int) (Transaction, error) {
	var transaction Transaction

	// CARA MENGHUBUNGKAN RELASI DATABASE YG TIDAK PUNYA RELASI LANGSUNG
	err := r.db.Where("id=?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) SaveTransactionRepository(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) UpdateTransactionRepository(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
