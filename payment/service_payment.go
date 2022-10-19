package payment

import (
	"bwastartup/user"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type service struct{}

type Service interface {
	GetPaymentURLService(transaction TransactionEntityPayment, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURLService(transaction TransactionEntityPayment, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-A9ymdOC3YZmNf_MMvfLCvdsv"
	midclient.ClientKey = "SB-Mid-client-NtoI5hBe1KBmnlSp"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil

}
