package factory_method

import (
    "strings"
    "testing"
)

/**
 * @author Rancho
 * @date 2021/12/27
 */

func TestGetPaymentMethodCash(t *testing.T) {
    payment, err := GetPayer(Cash)
    if err != nil {
        t.Fatal("A payment method of type 'Cash' must exist")
    }

    msg := payment.Pay(10.30)
    if !strings.Contains(msg, "payed using cash") {
        t.Error("The cash payment method message wasn't correct")
    }
    t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
    payment, err := GetPayer(DebitCard)
    if err != nil {
        t.Fatal("A payment method of type 'DebitCard' must exist")
    }
    msg := payment.Pay(22.30)
    if !strings.Contains(msg, "payed using debit card") {
        t.Error("The debit card payment method message wasn't correct")
    }
    t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
    _, err := GetPayer(20)
    if err == nil {
        t.Error("A payment method with ID 20 must return an error")
    }
    t.Log("LOG:", err)
}
