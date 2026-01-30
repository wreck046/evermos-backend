package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProvinces(c *gin.Context) {
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch provinces"})
		return
	}
	defer resp.Body.Close()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(200, data)
}

func GetCities(c *gin.Context) {
	provinceID := c.Query("province_id")
	if provinceID == "" {
		c.JSON(400, gin.H{"error": "province_id is required"})
		return
	}

	url := "https://www.emsifa.com/api-wilayah-indonesia/api/regencies/" + provinceID + ".json"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch cities"})
		return
	}
	defer resp.Body.Close()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(200, data)
}

