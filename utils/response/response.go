package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
)

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
		"data":    data,
	})
}

func CreatedResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	var ve validation.Errors
	if errors.As(err, &ve) {
		// mapping field => error message
		errorMap := make(map[string]string)
		for field, fe := range ve {
			errorMap[field] = fe.Error()
		}
		c.JSON(statusCode, gin.H{
			"code":    statusCode,
			"errors":  errorMap,
			"message": "Validation Error",
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": err.Error(),
	})
}
