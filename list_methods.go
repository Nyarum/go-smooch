package smooch

type ListWebhooks struct {
	Webhooks []struct {
		ID       string   `json:"_id"`
		Triggers []string `json:"triggers"`
		Secret   string   `json:"secret"`
		Target   string   `json:"target"`
	} `json:"webhooks"`
}

func (s *Smooch) ListWebhooks() (*ListWebhooks, error) {
	resp, err := s.request("webhooks")
	if err != nil {
		return nil, err
	}

	listWebhooks := &ListWebhooks{}

	err = resp.ToJSON(listWebhooks)
	if err != nil {
		return nil, err
	}

	return listWebhooks, nil
}
