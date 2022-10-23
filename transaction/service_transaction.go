package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"

	"strconv"

	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

type Service interface {
	GetTransactionByCampaignIDService(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserIDService(UserID int) ([]Transaction, error)
	CreatTransactionService(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionByCampaignIDService(input GetCampaignTransactionInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByIDCampaignRepository(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignIDRepository(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionByUserIDService(UserID int) ([]Transaction, error) {
	transactionS, err := s.repository.GetByUserIDTransactionRepository(UserID)
	if err != nil {
		return transactionS, err
	}
	return transactionS, nil
}

func (s *service) CreatTransactionService(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pendding"

	newTransaction, err := s.repository.SaveTransactionRepository(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.TransactionEntityPayment{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURLService(paymentTransaction, input.User)

	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL
	newTransaction, err = s.repository.UpdateTransactionRepository(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByIDTransactionRepository(transaction_id)
	if err != nil {
		return err
	}
	if input.PaymentType == "credit_cart" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" && input.TransactionStatus == "expire" && input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updateTransaction, err := s.repository.UpdateTransactionRepository(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindByIDCampaignRepository(updateTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updateTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updateTransaction.Amount

		_, err := s.campaignRepository.UpdateCampaignRepository(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
