package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type NodeChallengeHandler struct {
}

func NewNodeChallengeHandler() *NodeChallengeHandler {
	return &NodeChallengeHandler{}
}

func (h *NodeChallengeHandler) NodeChallenge(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("utils/hard.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	data := make([][]int, 0)
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	for i := len(data) - 2; i >= 0; i-- {
		for j := 0; j < len(data[i]); j++ {
			data[i][j] += max(data[i+1][j], data[i+1][j+1])
		}
	}
	sum := data[0][0]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"sum": sum})
}
