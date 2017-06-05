package smooch_test

import (
	"log"
	"testing"

	smooch "github.com/Nyarum/go-smooch"
)

func TestListWebhooks(t *testing.T) {
	sm := smooch.NewSmooch()
	sm.Auth("", "")

	listWebhooks, err := sm.ListWebhooks()
	if err != nil {
		t.Error(err)
	}

	log.Println(listWebhooks)
}

func TestSubscribeToWebhook(t *testing.T) {
	sm := smooch.NewSmooch()
	sm.Auth("", "")

	err := sm.SubscribeToWebhook("/hook", ":8080")
	if err != nil {
		t.Error(err)
	}
}
