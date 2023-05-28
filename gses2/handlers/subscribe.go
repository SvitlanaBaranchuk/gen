package handlers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"../models"
	"github.com/gin-gonic/gin"
)

const subscriptionFile = "./storage/subscriptions.txt"

// CheckEmailExists перевіряє, чи існує електронна адреса в файлі підписок.
func CheckEmailExists(email string) (bool, error) {
	filePath := subscriptionFile

	// Відкриття файлу
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Читання файлу рядок за рядком
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Перевірка на співпадіння електронної адреси
		if strings.TrimSpace(scanner.Text()) == strings.TrimSpace(email) {
			return true, nil
		}
	}

	// Перевірка на помилку читання файлу
	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error while reading file: %v", err)
	}

	return false, nil
}

// SaveEmailToFile зберігає електронну адресу в файлі підписок.
func SaveEmailToFile(email string) error {
	file, err := os.OpenFile(subscriptionFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(email + "\n")
	if err != nil {
		return err
	}

	return nil
}

// SubscribeHandler є обробником Gin, який обробляє запит на підписку на електронну адресу.
func SubscribeHandler(c *gin.Context) {
	var email models.Email

	// Читання тіла запиту
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Присвоєння електронної адреси з тіла запиту
	email.Address = string(body)

	// Перевірка, чи електронна адреса вже існує
	exists, err := CheckEmailExists(email.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email existence"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Збереження електронної адреси в файлі
	err = SaveEmailToFile(email.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save email"})
		return
	}

	// Відправка успішної відповіді
	c.JSON(http.StatusOK, gin.H{"message": "Email subscribed successfully"})
}
