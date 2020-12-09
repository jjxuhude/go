package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	ResX, ResY float32
}
type Battery struct {
	Capacity int
}

func getJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchId bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			4000,
		},
		HasTouchId: true,
	}

	jsonData, _ := json.Marshal(raw)
	return jsonData

}

func main() {
	jsonData := getJsonData()
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	screenAndTouch := struct {
		Screen
		HasTouchId bool
	}{}

	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Printf("%v\n", screenAndTouch)
}
