package mailSender

import (
	"fmt"
	"gitlab.com/ast3am77/test-go/internal/models"
	"gopkg.in/gomail.v2"
)

type logger interface {
	DebugMsg(msg string)
	ErrorMsg(msg string, err error)
}

type MailSender struct {
	dealer *gomail.Dialer
	log    logger
}

func NewSender(cfg *models.Config, log logger) (*MailSender, error) {
	if cfg.EmailDealerConfig.Password == "" || cfg.EmailDealerConfig.Username == "" {
		err := fmt.Errorf("email auth login and password cannot be empty")
		return nil, err
	}
	dealer := gomail.NewDialer(cfg.EmailDealerConfig.Host, cfg.EmailDealerConfig.Port,
		cfg.EmailDealerConfig.Username, cfg.EmailDealerConfig.Password)
	s, err := dealer.Dial()
	if err != nil {
		return nil, err
	} else {
		defer s.Close()
	}

	result := MailSender{
		dealer: dealer,
		log:    log,
	}
	return &result, err
	//return nil, nil
}

func (ms *MailSender) SendVerificationEmail(worker *models.AddWorkersDTO, url string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", ms.dealer.Username)
	m.SetHeader("To", worker.Email)

	m.SetHeader("Subject", "invite link")
	m.SetBody("text/plain", url)

	if err := ms.dealer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
