package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

// GetBitcoinPrice отримує актуальний курс Bitcoin і повертає його у вигляді рядка.
func GetBitcoinPrice() string {
	// Отримання актуального курсу Bitcoin з CoinGecko API
	price, err := http.Get("https://api.coingecko.com/api/v3/coins/bitcoin")
	if err != nil {
		return fmt.Errorf("failed to get current price: %v", err).Error()
	}
	defer price.Body.Close()

	var data models.CoinGeckoResponse
	err = json.NewDecoder(price.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err).Error()
	}

	jsonData, err := json.Marshal(data.MarketData.CurrentPrice.UAH)
	if err != nil {
		return fmt.Errorf("failed to serialize JSON: %v", err).Error()
	}

	return string(jsonData)
}

// GetCurrentPrice є обробником Gin, який викликає функцію GetBitcoinPrice та надсилає результат у відповідь API.
func GetCurrentPrice(c *gin.Context) {
	// Отримання актуального курсу Bitcoin
	price := GetBitcoinPrice()
	if price == "" {
		// Обробка помилки, якщо не вдалося отримати курс Bitcoin
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current price"})
		return
	}

	// Відправка успішного результату з актуальним курсом Bitcoin
	c.JSON(http.StatusOK, gin.H{
		"Повертається актуальний курс BTC до UAH": price,
	})
}
