package factory_method

import (
    "errors"
    "fmt"
)

/**
 * @author Rancho
 * @date 2021/12/26
 */

const (
    Cash      = 1
    DebitCard = 2
)

type Payer interface {
    Pay(amount float64) string
}

func GetPayer(m int) (Payer, error) {
    switch m {
    case Cash:
        return new(CashPM), nil
    case DebitCard:
        return new(DebitCardPM), nil
    default:
        return nil, errors.New(fmt.Sprintf("payment method %d not recognized", m))
    }
}

type CashPM struct {
}

func (c *CashPM) Pay(amount float64) string {
    return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

type DebitCardPM struct {
}

func (d *DebitCardPM) Pay(amount float64) string {
    return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}
