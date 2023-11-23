package entity

import "errors"

type Order struct {
	ID         string  `json:"id,omitempty"`
	Price      float64 `json:"price,omitempty"`
	Tax        float64 `json:"tax,omitempty"`
	FinalPrice float64 `json:"final_price,omitempty"`
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}
	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if o.Tax <= 0 {
		return errors.New("tax must be greater than or equal to zero")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validate()
	if err != nil {
		return err
	}
	return nil
}
