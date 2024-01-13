package pkg

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

func DecodeBody(c *gin.Context, storeVar any) error {
	err := json.NewDecoder(c.Request.Body).Decode(&storeVar)
	if HasError(err) {
		return err
	}

	return Validate(storeVar)
}

func GetUrlVars(c *gin.Context, param string) (string, error) {
	urlParam := c.Param(param)

	if urlParam == STRING_EMPTY {
		return STRING_EMPTY, errors.New("url var not found")
	}

	return urlParam, nil
}

func GetUrlParams(c *gin.Context, param string) (string, error) {
	urlParam := c.Query(param)

	if urlParam == STRING_EMPTY {
		return STRING_EMPTY, errors.New("url param not found")
	}

	return urlParam, nil
}
