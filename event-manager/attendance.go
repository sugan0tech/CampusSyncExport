package eventmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const URI = "https://possible-flamingo-large.ngrok-free.app"

type AcamedicDate struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}

func GetAcademicDates() []AcamedicDate {
	var emptyBody []byte
	response, err := http.Post(URI+"/attendance/get-dates", "none", bytes.NewBuffer(emptyBody))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var dates []AcamedicDate

	fmt.Print(response)
	fmt.Println(body)
	json.Unmarshal(body, &dates)

	return dates
}
