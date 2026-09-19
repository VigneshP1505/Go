package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type Session struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Duration string `json:"duration"`
}

const dataFile = "codeclock.json"

func main() {
	fmt.Println("================")
	fmt.Println("Code clock")
	fmt.Println("================")

	fmt.Print("What task are you starting?:")
	reader := bufio.NewReader(os.Stdin)
	taskName, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("error encountered")
	}

	var startTime time.Time = time.Now()
	fmt.Printf("\nClocked IN at %s\n", startTime.Format("15:04:05"))

	fmt.Println("Press [Enter] when you are finished...")

	_, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalln("error encountere")
	}

	var endTime time.Time = time.Now()
	exactDuration := endTime.Sub(startTime)
	duration := math.Round(exactDuration.Seconds())
	fmt.Printf("Clocked OUT! Time Spent:%d\n", duration)

	newSession := &Session{
		Name:     taskName[:len(taskName)-1],
		Date:     time.Now().Format("2006-01-02"),
		Duration: strconv.Itoa(int(duration)),
	}

	saveSession(*newSession)

}

func saveSession(newSession Session) error {
	var sessions []Session
	data, err := os.ReadFile("codeclock.json")

	if len(data) > 0 {
		err = json.Unmarshal(data, &sessions)
		if err != nil {
			fmt.Println("error encountered unmarshaling json")
		}
	}

	sessions = append(sessions, newSession)
	data, err = json.MarshalIndent(sessions, "", " ")
	fmt.Println("data")
	if err != nil {
		fmt.Println("error encountered marshaling data")
	}

	return os.WriteFile("codeclock.json", data, 0644)

}
