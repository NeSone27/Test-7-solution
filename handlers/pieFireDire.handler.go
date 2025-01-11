package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type PieFireDireHandler struct {
}

func NewPieFireDireHandler() *PieFireDireHandler {
	return &PieFireDireHandler{}
}

func (h *PieFireDireHandler) PieFireDire(w http.ResponseWriter, r *http.Request) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	text := string(body)
	text = strings.ToLower(text)
	text = regexp.MustCompile(`[^a-z\s-]`).ReplaceAllString(text, "")
	beefType := []string{
		"t-bone",
		"fatback",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
	}
	beefSum := make(map[string]int)
	for _, beef := range beefType {
		beefSum[beef] = strings.Count(text, beef)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]map[string]int{"beef": beefSum})
}
