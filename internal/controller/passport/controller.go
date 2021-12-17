package passportctrl

import (
	"github.com/rs/zerolog/log"
	"store_server/internal/usecase/passport"
)

// May exist several of controllers

type Controller struct {
	SaveUseCase passport.SavePassportUseCase
}

func New(s passport.SavePassportUseCase) *Controller {
	return &Controller{SaveUseCase: s}
}

func (c *Controller) SavePassport(p passport.PassportModel) {
	err := c.SaveUseCase.Save(p)
	if err != nil {
		log.Err(err)
	}
}

func (c *Controller) LoadPassport(id int) *passport.PassportModel {
	return nil
}
