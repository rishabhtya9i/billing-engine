package main

import (
    "github.com/gin-gonic/gin"
    "billing-engine/handlers"
)

func main() {
    r := gin.Default()

    r.POST("/loan", handlers.CreateLoan)
    r.POST("/loan/:id/payment", handlers.MakePayment)
    r.GET("/loan/:id/outstanding", handlers.GetOutstanding)
    r.GET("/loan/:id/delinquent", handlers.CheckDelinquent)

    r.Run(":8080")
}