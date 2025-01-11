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

	sum := data[0][0]
	indexCal := 0
	for i := 1; i < len(data); i++ {
		vsum := 0
		if data[i][indexCal] > data[i][indexCal+1] {
			vsum = data[i][indexCal]
		} else {
			vsum = data[i][indexCal+1]
			indexCal = indexCal + 1
		}
		sum += vsum
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"sum": sum})
}
