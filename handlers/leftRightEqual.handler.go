package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist-service/models"
)

type LeftRightEqualHandler struct {
}

func NewLeftRightEqualHandler() *LeftRightEqualHandler {
	return &LeftRightEqualHandler{}
}

func (h *LeftRightEqualHandler) LeftRightEqual(w http.ResponseWriter, r *http.Request) {
	data := models.LeftRightEqualRequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	value := []int{}
	doString := []string{}
	num := 0
	for i := 0; i < len(data.Input); i++ { // 210122
		if i == 0 {
			if string(data.Input[i]) == "L" {
				num = 1
				value = append(value, 2)

			} else if string(data.Input[i]) == "R" {
				num = 2
				value = append(value, 1)
			} else if string(data.Input[i]) == "=" {
				value = append(value, 0)
			}
		} else {
			if string(data.Input[i]) == "L" {
				num--
			} else if string(data.Input[i]) == "R" {
				num++
			}
		}

		doString = append(doString, string(data.Input[i]))
		if num < 0 || num > 2 {
			value, num = h.Recheck(value, num, doString)
		} else {
			value = append(value, num)
		}
	}
	result := ""
	for _, num := range value {
		result += fmt.Sprintf("%d", num)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func (h *LeftRightEqualHandler) Recheck(value []int, num int, doString []string) (res []int, numRes int) {
	if num < 0 {
		num = 0
		for i := len(value) - 1; i >= 0; i-- {
			value[i] = value[i] + 1
			if i != 0 {
				if string(doString[i-1]) == "R" && value[i] > value[i-1] {
					value = append(value, num)
					return value, num
				} else if string(doString[i-1]) == "L" && value[i] < value[i-1] {
					value = append(value, num)
					return value, num
				} else if string(doString[i-1]) == "=" && value[i] == value[i-1] {
					value = append(value, num)
					return value, num
				}
			}
		}
	} else if num > 2 {
		num = 2
		for i := len(value) - 1; i >= 0; i-- {
			value[i] = value[i] - 1
			if i != 0 {
				if string(doString[i-1]) == "R" && value[i] > value[i-1] {
					value = append(value, num)
					return value, num
				} else if string(doString[i-1]) == "L" && value[i] < value[i-1] {
					value = append(value, num)
					return value, num
				} else if string(doString[i-1]) == "=" && value[i] == value[i-1] {
					value = append(value, num)
					return value, num
				}
			}
		}
	}
	value = append(value, num)
	return value, num
}
