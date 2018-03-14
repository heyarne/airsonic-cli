package ping

import (
	"errors"

	"airsonic-cli/config"
	"airsonic-cli/request"
	"airsonic-cli/utils"
)

// Ping define the json subsonic response structure for ping calls
type Ping struct {
	SubsonicResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
	} `json:"subsonic-response"`
}

// PingAction run a ping against the server's API
func PingAction(conf *config.Config) error {
	if config.IsVerbose(conf) {
		utils.InfoMsg("Ping send to " + config.GetServer(conf))
	}
	var data = request.Get(conf, "/rest/ping", "")

	if request.CheckResponse(conf, data) {
		if config.IsVerbose(conf) {
			utils.InfoMsg("Pong")
		}
		return nil
	}
	return errors.New("error while calling ping enpoint")
}
