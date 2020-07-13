package income

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanmalv/go-crud-pod/internal/platform/database"
	"net/http"
)

//Income model details the data included on a basic income saving
type Income struct {
	Amount   *float64 `json:"amount"`
	Date     string   `json:"date"`
	Reason   string   `json:"reason"`
	Category string   `json:"category"`
}


func (operation Income)ValidateIncomeRequest(c *gin.Context) error {

	if operation.Amount == nil{
		c.JSON(http.StatusBadRequest, "Invalid Request")
		return fmt.Errorf("%v %v must be filled",database.IncomeCollection,"amount")
	}

	if *operation.Amount <= 0 {
		c.JSON(http.StatusBadRequest, "Invalid Request")
		return fmt.Errorf("%v %v must be greater than 0",database.IncomeCollection,"amount")
	}

	return nil
}