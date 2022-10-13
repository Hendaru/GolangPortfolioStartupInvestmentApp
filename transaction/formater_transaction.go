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
