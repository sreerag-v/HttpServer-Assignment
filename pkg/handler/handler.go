package handler

import (
	model "Inter/webhook/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func Worker(ch <-chan model.RequestData, wg *sync.WaitGroup) {
	for req := range ch {
		// Process the request and convert it
		responseData := convertData(req)

		// Set Content-Type
		fmt.Println("Processing:", responseData)

		// Send the converted data to the webhook
		sendToWebhook(responseData)

		wg.Done()
	}
}

func convertData(requestData model.RequestData) model.ResponseData {
	responseData := model.ResponseData{
		Event:           requestData.Ev,
		EventType:       requestData.Et,
		AppID:           requestData.Id,
		UserID:          requestData.Uid,
		MessageID:       requestData.Mid,
		PageTitle:       requestData.T,
		PageURL:         requestData.P,
		BrowserLanguage: requestData.L,
		ScreenSize:      requestData.Sc,
		Attributes:      make(map[string]model.Attribute),
		Traits:          make(map[string]model.Attribute),
	}

	// Process attributes
	processAttributes(requestData.Atrk1, requestData.Atrv1, requestData.Atrt1, responseData.Attributes)
	processAttributes(requestData.Atrk2, requestData.Atrv2, requestData.Atrt2, responseData.Attributes)

	// Process user attributes
	processAttributes(requestData.Uatrk1, requestData.Uatrv1, requestData.Uatrt1, responseData.Traits)
	processAttributes(requestData.Uatrk2, requestData.Uatrv2, requestData.Uatrt2, responseData.Traits)
	processAttributes(requestData.Uatrk3, requestData.Uatrv3, requestData.Uatrt3, responseData.Traits)

	return responseData
}

func processAttributes(key, value, typ string, attributeMap map[string]model.Attribute) {
	if key != "" {
		attributeMap[key] = model.Attribute{
			Value: value,
			Type:  typ,
		}
	}
}

func sendToWebhook(data model.ResponseData) {
	// Convert the result to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Send the JSON data to the webhook (replace with your actual webhook URL)
	_, err = http.Post("https://webhook.site/54903581-4b46-43b3-b1ca-123c6234276d", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending to webhook:", err)
		return
	}
}
