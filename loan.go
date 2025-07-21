package models

import (
    "errors"
)

const (
    TotalWeeks        = 50
    PrincipalAmount   = 5000000
    InterestAmount    = 500000
    TotalAmount       = PrincipalAmount + InterestAmount
    WeeklyInstallment = TotalAmount / TotalWeeks
)

type Loan struct {
    ID             string `json:"id"`
    WeeklySchedule []bool `json:"weekly_schedule"`
    CurrentWeek    int    `json:"current_week"`
}

func NewLoan(id string) *Loan {
    return &Loan{
        ID:             id,
        WeeklySchedule: make([]bool, TotalWeeks),
        CurrentWeek:    0,
    }
}

func (l *Loan) MakePayment(amount int) error {
    if l.CurrentWeek >= TotalWeeks {
        return errors.New("loan already completed")
    }
    if amount != WeeklyInstallment {
        return errors.New("only exact payment allowed")
    }
    l.WeeklySchedule[l.CurrentWeek] = true
    l.CurrentWeek++
    return nil
}

func (l *Loan) GetOutstanding() int {
    payments := 0
    for _, paid := range l.WeeklySchedule {
        if paid {
            payments++
        }
    }
    return (TotalWeeks - payments) * WeeklyInstallment
}

func (l *Loan) IsDelinquent() bool {
    missed := 0
    for i := 0; i < l.CurrentWeek; i++ {
        if !l.WeeklySchedule[i] {
            missed++
            if missed == 2 {
                return true
            }
        } else {
            missed = 0
        }
    }
    return false
}