package core

import (
	"encoding/json"
	"fmt"
	"os"
	//"strings"
)

type input struct {
	Right []struct {
		TimeSign []int `json:"ts"`
		Notes    []struct {
			NoteName string `json:"n"`
			Finger   int    `json:"f"`
		} `json:"notes"`
	} `json:"Right"`
	Left []struct {
		TimeSign []int `json:"ts"`
		Notes    []struct {
			NoteName string `json:"n"`
			Finger   int    `json:"f"`
		} `json:"notes"`
	} `json:"Left"`
}

func (p *PVF) FillFinger() {

	file, err := os.ReadFile("assets/input.json")
	if err != nil {
		fmt.Println("Error read user JSON:", err)
		return
	}

	var userInputData input
	err = json.Unmarshal(file, &userInputData)
	if err != nil {
		fmt.Println("Error unmarshaling user JSON:", err)
		return
	}
	//fmt.Println(userInputData)

	for i := range p.TracksV2.Right {
		if i < len(userInputData.Right) { // 确保不越界
			for j := range p.TracksV2.Right[i].Notes {
				if j < len(userInputData.Right[i].Notes) {
					p.TracksV2.Right[i].Notes[j].Finger = userInputData.Right[i].Notes[j].Finger
				}
			}
		}
	}

	for i := range p.TracksV2.Left {
		if i < len(userInputData.Left) {
			for j := range p.TracksV2.Left[i].Notes {
				if j < len(userInputData.Left[i].Notes) {
					p.TracksV2.Left[i].Notes[j].Finger = userInputData.Left[i].Notes[j].Finger
				}
			}
		}
	}
	fmt.Println(p)
	tmp, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
	err = os.WriteFile("output.json", tmp, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return // Exit if there's an error
	}
}
