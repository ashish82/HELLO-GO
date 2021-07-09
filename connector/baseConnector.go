package connector

import (
	"HELLO-GO/http_config"
	"HELLO-GO/model/response"
	logger2 "HELLO-GO/utility/logger"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

var logger = logger2.GetLogger()

type BaseConnector struct {
	POST interface{}
}

func HttpPost(url string, apiName string, headers map[string]string, request *interface{}, response *interface{}) {

	bytesRepresentation, err := json.Marshal(request)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))

	for ky, val := range headers {
		req.Header.Set(ky, val)
	}

	httpClientObj := http_config.HTTPClientMap[apiName]
	resp, err := httpClientObj.Do(req)
	if err != nil {
		logger.Errorf("Error while getting api response for %s error %s", apiName, err.Error())
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logger.Errorf("error in body close for api %s  error %s", apiName, err.Error())
			}
		}(resp.Body)
		body, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(body, response)
	}
}

func HttpGet(url string, apiName string, headers map[string]string, response *response.Comments) {
	req, _ := http.NewRequest("GET", url, nil)
	for ky, val := range headers {
		req.Header.Set(ky, val)
	}
	httpClientObj := http_config.HTTPClientMap[apiName]
	resp, err := httpClientObj.Do(req)
	if err != nil {
		logger.Errorf("Error while getting api response for %s error %s", apiName, err.Error())
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logger.Errorf("error in body close for api %s  error %s", apiName, err.Error())
			}
		}(resp.Body)
		body, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(body, response)
	}
}
