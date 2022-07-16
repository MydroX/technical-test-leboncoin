package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeBody struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func Home(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body HomeBody

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": fmt.Sprintf("error parsing body : %s", err)})
		log.Printf("home error: %v", err)
		return
	}

	if body.Limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "limit must be greater than 0"})
		return
	}
	if body.Str1 == "" || body.Str2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "str1 and str2 must be non-empty"})
		return
	}

	var res string

	for i := 1; i <= body.Limit; i++ {
		if body.Int1 == 0 {
			if i%body.Int2 == 0 {
				if i == 1 {
					res = fmt.Sprintf("%s%s", body.Str1, body.Str2)
				}
				res = fmt.Sprintf("%s,%s%s", res, body.Str1, body.Str2)
			} else {
				if i == 1 {
					res = body.Str1
				}
				res = fmt.Sprintf("%s,%s", res, body.Str1)
			}
			continue
		} else if body.Int2 == 0 {
			if i%body.Int1 == 0 {
				if i == 1 {
					res = fmt.Sprintf("%s%s", body.Str1, body.Str2)
				}
				res = fmt.Sprintf("%s,%s%s", res, body.Str1, body.Str2)
			} else {
				if i == 1 {
					res = body.Str2
				}
				res = fmt.Sprintf("%s,%s", res, body.Str2)
			}
			continue
		}

		if i == 1 {
			if i%body.Int1 == 0 && i%body.Int2 == 0 {
				res = fmt.Sprintf("%s,%s%s", res, body.Str1, body.Str2)
			} else if i%body.Int1 == 0 {
				res = body.Str1
			} else if i%body.Int2 == 0 {
				res = body.Str2
			} else {
				res = fmt.Sprintf("%d", i)
			}
			continue
		}

		if i%body.Int1 == 0 && i%body.Int2 == 0 {
			res = fmt.Sprintf("%s,%s%s", res, body.Str1, body.Str2)
		} else if i%body.Int1 == 0 {
			res = fmt.Sprintf("%s,%s", res, body.Str1)
		} else if i%body.Int2 == 0 {
			res = fmt.Sprintf("%s,%s", res, body.Str2)
		} else {
			res = fmt.Sprintf("%s,%d", res, i)
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "result": res})
}
