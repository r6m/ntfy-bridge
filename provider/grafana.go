package provider

type GrafanaPayload struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []GrafanaAlert    `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	OrgID             int               `json:"orgId"`
	Title             string            `json:"title"`
	State             string            `json:"state"`
	Message           string            `json:"message"`
}

type GrafanaAlert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     string            `json:"startsAt"`
	EndsAt       string            `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
	SilenceURL   string            `json:"silenceURL"`
	DashboardURL string            `json:"dashboardURL"`
	PanelURL     string            `json:"panelURL"`
	Values       any               `json:"values"`
	ValueString  string            `json:"valueString"`
}

func (p *GrafanaPayload) ToNotification(topic string) Notification {
	firstAlert := p.Alerts[0]

	payload := Notification{
		Message: p.Message,
		Title:   p.Title,
		Topic:   topic,
		Actions: []Action{
			{
				Action: "view",
				Label:  "Open in Grafana",
				Url:    p.ExternalURL,
				Clear:  true,
			},
			{
				Action: "view",
				Label:  "Silence",
				Url:    firstAlert.SilenceURL,
				Clear:  false,
			},
		},
	}

	return payload
}
