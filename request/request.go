package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"airsonic-cli/config"
	"airsonic-cli/utils"
)

type Response struct {
	SubsonicResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
		Error   struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	} `json:"subsonic-response"`
}

func Get(conf *config.Config, endpoint string, payload string) []byte {
	u, parErr := url.Parse(config.GetServer(conf) + endpoint + "?" + loadCommonPayload(conf) + payload)
	if parErr != nil {
		log.Fatal(parErr)
	}

	q := u.Query()
	q.Set("u", config.GetUsername(conf))
	q.Set("t", config.GetToken(conf))
	q.Set("s", config.GetSalt(conf))
	q.Set("c", config.GetAppName(conf))
	q.Set("v", config.GetAPIVersion(conf))
	q.Set("f", config.GetAPIFormat(conf))
	u.RawQuery = q.Encode()

	if config.IsVerbose(conf) {
		utils.InfoMsg("GET " + u.String())
	}

	req, reqErr := http.NewRequest(http.MethodGet, u.String(), nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	httpClient := http.Client{Timeout: time.Second * 2}
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func CheckResponse(conf *config.Config, data []byte) bool {
	if config.IsVerbose(conf) {
		utils.InfoMsg("Checking response")
	}
	response := Response{}
	jsonErr := json.Unmarshal(data, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	if response.SubsonicResponse.Status == "failed" {
		utils.ErrorMsg("Response : " + response.SubsonicResponse.Status)
		utils.ErrorMsg(strconv.Itoa(response.SubsonicResponse.Error.Code) + " - " + response.SubsonicResponse.Error.Message)
		return false
	}
	if config.IsVerbose(conf) {
		utils.InfoMsg("Response : " + response.SubsonicResponse.Status)
	}
	return true
}

func loadCommonPayload(conf *config.Config) string {
	u := new(url.URL)
	q := u.Query()
	q.Set("u", config.GetUsername(conf))
	q.Set("t", config.GetToken(conf))
	q.Set("s", config.GetSalt(conf))
	q.Set("c", config.GetAppName(conf))
	q.Set("v", config.GetAPIVersion(conf))
	q.Set("f", config.GetAPIFormat(conf))
	return q.Encode()
}
