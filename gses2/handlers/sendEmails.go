package handlers

import (
	"bufio"
	"fmt"
	"os"

	mail "github.com/go-mail/gomail"

	"github.com/gin-gonic/gin"
)

// SendEmailsHandler є обробником Gin, який відправляє електронні листи з оновленнями про ціну Bitcoin.
func SendEmailsHandler(c *gin.Context) {
	subject := "Bitcoin Price Update"

	price := GetBitcoinPrice()
	fmt.Println("Current Bitcoin price in UAH:", price)

	body := fmt.Sprintf("The current Bitcoin price in UAH is %s", price)

	file, err := os.Open(subscriptionFile)
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()

	// Створення нового читача файлу
	reader := bufio.NewScanner(file)

	// Читання рядків з файлу та відправка електронних листів
	for reader.Scan() {
		line := reader.Text()
		SendEmail(line, subject, body)
	}

	if err := reader.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
		return
	}
	fmt.Println("Email sent successfully!")
}

// SendEmail відправляє електронний лист з заданими параметрами.
func SendEmail(to, subject, body string) error {
	from := "karysel55@gmail.com"
	password := "uxhejpcuwpgtonuo"
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	email := mail.NewMessage()
	email.SetHeader("From", from)
	email.SetHeader("To", to)
	email.SetHeader("Subject", subject)
	email.SetBody("text/plain", body)

	dialer := mail.NewDialer(smtpHost, smtpPort, from, password)

	if err := dialer.DialAndSend(email); err != nil {
		return err
	}

	return nil
}
