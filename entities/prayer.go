package entities

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"time"
)

type Prayer struct {
	ID        string `valid:"uuid"`
	Talked    string `valid:"required"`
	Date      time.Time `valid:"required"`
	CreatedAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:"-"`
}

func NewPrayer(talked string, date time.Time) (*Prayer, error) {
	prayer := &Prayer{
		Talked:    talked,
		Date:      date,
	}

	prayer.prepare()

	err := prayer.Validate()
	if err != nil {
		return nil, err
	}

	return prayer, nil
}

func (p *Prayer) prepare()  {
	p.ID = uuid.NewString()
	p.CreatedAt = time.Now()
}

func (p *Prayer) Validate() error {

	_,err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
