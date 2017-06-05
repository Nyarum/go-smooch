package smooch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type MessageReceive struct {
	Trigger string `json:"trigger"`
	App     struct {
		ID string `json:"_id"`
	} `json:"app"`
	Messages []struct {
		ID       string  `json:"_id"`
		Type     string  `json:"type"`
		Text     string  `json:"text"`
		Role     string  `json:"role"`
		AuthorID string  `json:"authorId"`
		Name     string  `json:"name"`
		Received float64 `json:"received"`
		Source   struct {
			Type string `json:"type"`
		} `json:"source"`
	} `json:"messages"`
	AppUser struct {
		ID         string `json:"_id"`
		UserID     string `json:"userId"`
		Properties struct {
		} `json:"properties"`
		SignedUpAt time.Time `json:"signedUpAt"`
		Clients    []struct {
			Active   bool      `json:"active"`
			ID       string    `json:"id"`
			LastSeen time.Time `json:"lastSeen"`
			Platform string    `json:"platform"`
		} `json:"clients"`
	} `json:"appUser"`
}

func (s *Smooch) SubscribeToWebhook(endpoint, host string) error {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("We got a GET method, youhuuu! :)")
		case "POST":
			payload, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("Error to get payload from webhook:", err)
				return
			}

			messageReceive := MessageReceive{}

			err = json.Unmarshal(payload, &messageReceive)
			if err != nil {
				log.Println("Error to unmarshal payload from webhook:", err)
				return
			}

			if messageReceive.Trigger != "api" {
				for _, message := range messageReceive.Messages {
					text := strings.ToLower(message.Text)

					if strings.Contains(text, "хочу меню") {
						responseMessage, err := s.SendMessage(messageReceive.AppUser.ID, `
						Меню:
						1. Пицца
						2. Салаты
						3. Бургеры
						`)
						if err != nil {
							log.Println("Error to send message from webhook:", err)
							return
						}

						log.Println("Receive result from send message:", responseMessage)
					} else if strings.Contains(text, "1") {
						responseMessage, err := s.SendMessage(messageReceive.AppUser.ID, `
						Из пицц у нас есть:
						- Натуральная
						- Без ГМО
						- Демо версия
						`)
						if err != nil {
							log.Println("Error to send message from webhook:", err)
							return
						}

						log.Println("Receive result from send message:", responseMessage)
					} else if strings.Contains(text, "2") {
						responseMessage, err := s.SendMessage(messageReceive.AppUser.ID, `
						В разработке...
						`)
						if err != nil {
							log.Println("Error to send message from webhook:", err)
							return
						}

						log.Println("Receive result from send message:", responseMessage)
					} else if strings.Contains(text, "3") {
						responseMessage, err := s.SendMessage(messageReceive.AppUser.ID, `
						В разработке...
						`)
						if err != nil {
							log.Println("Error to send message from webhook:", err)
							return
						}

						log.Println("Receive result from send message:", responseMessage)
					}
				}
			}
		}
	})

	fmt.Println("Subscribed to webhook -", host+endpoint)

	return http.ListenAndServe(host, nil)
}
