package provider

type Action struct {
	Action string `json:"action"`
	Label  string `json:"label"`
	Url    string `json:"url"`
	Clear  bool   `json:"clear"`
}

type Notification struct {
	Message string   `json:"message"`
	Topic   string   `json:"topic"`
	Title   string   `json:"title"`
	Actions []Action `json:"actions"`
}
