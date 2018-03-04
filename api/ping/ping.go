package ping

import (
  "log"
  "strconv"
  "encoding/json"
  
  "airsonic-cli/utils"
  "airsonic-cli/request"
  "airsonic-cli/config"
)

type Ping struct {
	SubsonicResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
		Error   struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	} `json:"subsonic-response"`
}

func PingAction(conf *config.Config) error {
  if config.IsVerbose(conf) { utils.InfoMsg("Ping send to " + config.GetServer(conf)) }
  ping := Ping{}
  jsonErr := json.Unmarshal(request.Get(conf, "/rest/ping", ""), &ping)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }
  if ping.SubsonicResponse.Status == "failed" {
    utils.ErrorMsg("Response : " + ping.SubsonicResponse.Status)
    utils.ErrorMsg(strconv.Itoa(ping.SubsonicResponse.Error.Code) + " - " + ping.SubsonicResponse.Error.Message)
  } else {
    if config.IsVerbose(conf) {
      utils.InfoMsg("Response : " + ping.SubsonicResponse.Status)
    }
  }
  return nil
}
