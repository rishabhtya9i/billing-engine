package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "billing-engine/models"
)

var loans = make(map[string]*models.Loan)

func CreateLoan(c *gin.Context) {
    id := uuid.New().String()
    loans[id] = models.NewLoan(id)
    c.JSON(http.StatusCreated, gin.H{"message": "Loan created", "id": id})
}

func MakePayment(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Amount int `json:"amount"`
    }

    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
        return
    }

    loan, exists := loans[id]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
        return
    }

    err := loan.MakePayment(body.Amount)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Payment accepted"})
}

func GetOutstanding(c *gin.Context) {
    id := c.Param("id")
    loan, exists := loans[id]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"outstanding": loan.GetOutstanding()})
}

func CheckDelinquent(c *gin.Context) {
    id := c.Param("id")
    loan, exists := loans[id]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"is_delinquent": loan.IsDelinquent()})
}