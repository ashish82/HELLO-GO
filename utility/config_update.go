package utility

import (
	"encoding/json"
	"fmt"
)

func GetUpdatedJson() string {
	var updatedResult string
	jsonObj := getJsonObj()
	var result map[string]interface{}
	errorObj := json.Unmarshal([]byte(jsonObj), &result)

	if errorObj == nil {
		updateMileStone(result)
	}

	finalResult, errorMar := json.Marshal(result)
	if errorMar == nil {
		updatedResult = string(finalResult)
	}

	return updatedResult
}

func updateMileStone(result map[string]interface{}) map[string]interface{} {
	fmt.Println("in update Milestone function")
	if result != nil {
		dataObj := result["referralRewardProgram"].(map[string]interface{})["data"].(map[string]interface{})["friend"].(map[string]interface{})["milestones"]
		dataSliceObj := dataObj.([]interface{})
		for _, result := range dataSliceObj {
			mileStoneObj := result.(map[string]interface{})
			eventName := mileStoneObj["eventName"]
			if eventName == "string" {
				mileStoneObj["eventName"] = "Ashish"
			}
		}
	}
	return result
}

func getJsonObj() string {
	return "{\"referralRewardProgram\":{\"data\":{\"bannerImage\":\"string\",\"friend\":{\"cta\":\"string\",\"header\":\"string\",\"milestones\":[{\"eventName\":\"string\",\"rewardAmount\":0,\"text\":\"string\"}],\"rewardType\":\"string\",\"totalRewardAmt\":0},\"referrer\":{\"cta\":\"string\",\"header\":\"string\",\"milestones\":[{\"eventName\":\"string\",\"rewardAmount\":0,\"text\":\"string\"}],\"rewardType\":\"string\",\"totalRewardAmt\":0},\"shareMessage\":\"string\",\"title\":\"string\",\"referralCode\":\"!@#!@$!@$!@$\",\"earnings\":{\"title\":\"Hurray\",\"subtitle\":\"You have earned\",\"rewardAmount\":\"Currency with ruppe symbol\",\"cta\":{\"text\":\"\",\"deeplink\":\"\"}}}}}"
}
