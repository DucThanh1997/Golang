package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
	"time"
	"os"
	"bufio"
	"io/ioutil"
	"encoding/base64"
	"strconv"
	"sync"
)
const baseURL = "http://192.168.30.150:5000/api/v1/"

func postURL(url string, jsonData RecordPC) (string, error) {
	var jsonStr, err = json.Marshal(jsonData)
	if err != nil {
		fmt.Println("err marshal: ", err.Error())
		return "0", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwcm9kdWN0IjoiIiwiZW1haWwiOiJ0aGFuaDE5MTk5N0BnbWFpbC5jb20iLCJleHAiOjE2MDg3MTk2Njl9.tEb-YS75wLR4Jl6e24lpxcXWxYQl0YAqP-YMj79XH7c")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("err when send request to people counting " + err.Error())
		return "0", err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	return resp.Status, nil
}

type RecordPC struct {
	CameraID		string		`json:"cameraID" bson:"cameraID"`
	TimeStamp		string		`json:"timeStamp" bson:"timeStamp"`
	Image			string		`json:"image" bson:"image"`
	Direction		int 		`json:"direction" bson:"direction"`
	Count			int			`json:"count" bson:"count"`
}

func main() {
	start := time.Now()
    defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()

	f, _ := os.Open("a.png")

    // Read entire JPG into byte slice.
    reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

    // Encode as base64.
    encoded := base64.StdEncoding.EncodeToString(content)
	
	wg := sync.WaitGroup{}

	var data = RecordPC {
		CameraID: "a7bc3de8-a2f2-4f37-943d-e15e05426bb4",
		TimeStamp: strconv.Itoa(int(time.Now().Unix())),
		Image: encoded,
		Direction: 2,
		Count: 1,
	}

	for i:= 0; i <= 10000; i++ {
		wg.Add(1)
		go func(data RecordPC) {
            statusCode, _ := postURL(baseURL + "people-counting/record-pc", data)
			fmt.Println("statusCode: ", statusCode)
            wg.Done()
        }(data)
	}
	
	wg.Wait()
	


}