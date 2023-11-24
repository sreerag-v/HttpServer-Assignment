 package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"sync"
// )

// type RequestData struct {
// 	Ev     string `json:"ev"`
// 	Et     string `json:"et"`
// 	Id     string `json:"id"`
// 	Uid    string `json:"uid"`
// 	Mid    string `json:"mid"`
// 	T      string `json:"t"`
// 	P      string `json:"p"`
// 	L      string `json:"l"`
// 	Sc     string `json:"sc"`
// 	Atrk1  string `json:"atrk1"`
// 	Atrv1  string `json:"atrv1"`
// 	Atrt1  string `json:"atrt1"`
// 	Atrk2  string `json:"atrk2"`
// 	Atrv2  string `json:"atrv2"`
// 	Atrt2  string `json:"atrt2"`
// 	Uatrk1 string `json:"uatrk1"`
// 	Uatrv1 string `json:"uatrv1"`
// 	Uatrt1 string `json:"uatrt1"`
// 	Uatrk2 string `json:"uatrk2"`
// 	Uatrv2 string `json:"uatrv2"`
// 	Uatrt2 string `json:"uatrt2"`
// 	Uatrk3 string `json:"uatrk3"`
// 	Uatrv3 string `json:"uatrv3"`
// 	Uatrt3 string `json:"uatrt3"`
// }

// type ResponseData struct {
// 	Event           string               `json:"event"`
// 	EventType       string               `json:"event_type"`
// 	AppID           string               `json:"app_id"`
// 	UserID          string               `json:"user_id"`
// 	MessageID       string               `json:"message_id"`
// 	PageTitle       string               `json:"page_title"`
// 	PageURL         string               `json:"page_url"`
// 	BrowserLanguage string               `json:"browser_language"`
// 	ScreenSize      string               `json:"screen_size"`
// 	Attributes      map[string]Attribute `json:"attributes"`
// 	Traits          Traits               `json:"traits"`
// }

// type Attribute struct {
// 	Value string `json:"value"`
// 	Type  string `json:"type"`
// }

// type Traits struct {
// 	Name  Attribute `json:"name"`
// 	Email Attribute `json:"email"`
// 	Age   Attribute `json:"age"`
// }

// var wg sync.WaitGroup

// func main() {
// 	requestChannel := make(chan RequestData, 10) // Adjust buffer size based on your needs

// 	// Start worker
// 	go worker(requestChannel, &wg)

// 	// Start HTTP server
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != "POST" {
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 			return
// 		}

// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		var requestData RequestData
// 		err = json.Unmarshal(body, &requestData)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		// Send the request to the channel
// 		wg.Add(1)
// 		requestChannel <- requestData

// 		// Respond to the client
// 		fmt.Fprintf(w, "Request received and being processed.")
// 	})

// 	// Start the server
// 	fmt.Println("Starting server on port 8080...")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Close the channel when the server exits
// 	close(requestChannel)

// 	// Wait for the worker to finish processing
// 	wg.Wait()
// }

// func worker(ch <-chan RequestData, wg *sync.WaitGroup) {
// 	for req := range ch {
// 		// Process the request and convert it
// 		responseData := convertData(req)

// 		// Set Content-Type
// 		fmt.Println("Processing:", responseData)

// 		// Send the converted data to the webhook
// 		sendToWebhook(responseData)

// 		wg.Done()
// 	}
// }

// func convertData(requestData RequestData) ResponseData {
// 	responseData := ResponseData{
// 		Event:           requestData.Ev,
// 		EventType:       requestData.Et,
// 		AppID:           requestData.Id,
// 		UserID:          requestData.Uid,
// 		MessageID:       requestData.Mid,
// 		PageTitle:       requestData.T,
// 		PageURL:         requestData.P,
// 		BrowserLanguage: requestData.L,
// 		ScreenSize:      requestData.Sc,
// 		Attributes:      make(map[string]Attribute),
// 		Traits:  Traits{
// 			Name:  Attribute{Value: requestData.Uatrv1, Type: requestData.Uatrt1},
// 			Email: Attribute{Value: requestData.Uatrv2, Type: requestData.Uatrt2},
// 			Age:   Attribute{Value: requestData.Uatrv3, Type: requestData.Uatrt3},
// 		},
// 	}

// 	// Process attributes
// 	processAttributes(requestData.Atrk1, requestData.Atrv1, requestData.Atrt1, responseData.Attributes)
// 	processAttributes(requestData.Atrk2, requestData.Atrv2, requestData.Atrt2, responseData.Attributes)

// 	return responseData
// }

// func processAttributes(key, value, typ string, attributeMap map[string]Attribute) {
// 	if key != "" {
// 		attributeMap[key] = Attribute{
// 			Value: value,
// 			Type:  typ,
// 		}
// 	}
// }

// func sendToWebhook(data ResponseData) {
// 	// Convert the result to JSON
// 	jsonData, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error converting to JSON:", err)
// 		return
// 	}

// 	// Send the JSON data to the webhook (replace with your actual webhook URL)
// 	_, err = http.Post("https://webhook.site/54903581-4b46-43b3-b1ca-123c6234276d", "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		fmt.Println("Error sending to webhook:", err)
// 		return
// 	}
// }
