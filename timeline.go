package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TimelineInput struct {
	Date string
}

type TimelineOutput struct {
	Result string
	ProductionData   []TimelineData
	DowntimeData     []TimelineData
	PowerOffData     []TimelineData
}

type TimelineData struct {
	Date  int64
	Value int
}

func getTimeLineData(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logInfo("MAIN", "Timeline function called")
	var data TimelineInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("MAIN", "Error parsing data: "+err.Error())
		var responseData TimelineOutput
		responseData.Result = "nok: " + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		logInfo("MAIN", "Parsing data ended")
		return
	}
	logInfo("MAIN", "Serving data from "+data.Date)

	var responseData TimelineOutput
	responseData.Result = "ok"
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	logInfo("MAIN", "Parsing data ended")
}
