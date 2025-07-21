package tests

import (
    "testing"
    "billing-engine/models"
)

func TestLoanFlow(t *testing.T) {
    loan := models.NewLoan("test-loan")

    err := loan.MakePayment(models.WeeklyInstallment)
    if err != nil {
        t.Fatal(err)
    }

    err = loan.MakePayment(models.WeeklyInstallment)
    if err != nil {
        t.Fatal(err)
    }

    loan.CurrentWeek += 2

    if !loan.IsDelinquent() {
        t.Error("Expected borrower to be delinquent")
    }

    if loan.GetOutstanding() != 5280000 {
        t.Errorf("Expected outstanding 5280000, got %d", loan.GetOutstanding())
    }
}