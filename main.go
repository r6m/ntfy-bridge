package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-chi/render"
	"github.com/r6m/httpx"
	"github.com/r6m/ntfy-bridge/provider"
)

var (
	reTopic = regexp.MustCompile(`(https?://.*?)/([-a-zA-Z0-9()@:%_\+.~#?&=]+)$`)
	port    = flag.String("port", "8080", "http port")

	ntfyUrl = "https://ntfy.sh"
)

func main() {

	http.HandleFunc("GET /grafana/{topic...}", handleGrafana)

	fmt.Println("starting server on http://localhost:" + *port)
	http.ListenAndServe(":"+*port, nil)
}

func prepareTopic(r *http.Request) (string, error) {
	topic := r.PathValue("topic")

	if strings.HasPrefix(topic, "http") {
		return "", fmt.Errorf("invalid topic: %v", topic)
	}

	endpoint := r.Header.Get("x-endpoint")
	if endpoint == "" {
		endpoint = ntfyUrl
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	u = u.JoinPath(topic)

	topic = u.String()

	return topic, nil
}

func handleGrafana(w http.ResponseWriter, r *http.Request) {
	topic, err := prepareTopic(r)
	if err != nil {
		httpx.BadRequest(w, r, err.Error())
		return
	}
	fmt.Println("topic is", topic)

	in := new(provider.GrafanaPayload)
	if err := render.Decode(r, in); err != nil {
		httpx.BadRequest(w, r, "invalid payload")
		return
	}

	payload := in.ToNotification(topic)
	if err := sendNotification(topic, payload); err != nil {
		httpx.BadRequest(w, r, fmt.Sprintf("can't send notification: %v", err))
		return
	}

	httpx.Respond(w, r, "notification sent")
}

func sendNotification(topic string, payload provider.Notification) error {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return fmt.Errorf("can't encode payload: %v", err)
	}
	req, err := http.NewRequest("POST", topic, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("can't send req: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ntfy status code: %d", resp.StatusCode)
	}

	return nil
}
