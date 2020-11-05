package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestParam(c *gin.Context) {
	Ids := c.Param("declare_ids")
	for _,id := range Ids {
		fmt.Println()
	}
	c.JSON(200, gin.H{"message": Ids, "status": 200})
}
