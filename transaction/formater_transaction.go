package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactionS(transactionS []Transaction) []CampaignTransactionFormatter {
	if len(transactionS) == 0 {
		return []CampaignTransactionFormatter{}
	}
	var transactionSFormatter []CampaignTransactionFormatter

	for _, t := range transactionS {
		formatter := FormatCampaignTransaction(t)
		transactionSFormatter = append(transactionSFormatter, formatter)
	}
	return transactionSFormatter
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter
	return formatter
}

func FormatUserTransactionS(transactionS []Transaction) []UserTransactionFormatter {
	if len(transactionS) == 0 {
		return []UserTransactionFormatter{}
	}
	var transactionSFormatter []UserTransactionFormatter

	for _, t := range transactionS {
		formatter := FormatUserTransaction(t)
		transactionSFormatter = append(transactionSFormatter, formatter)
	}
	return transactionSFormatter
}
