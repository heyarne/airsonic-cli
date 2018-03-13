package ping

import (
	"errors"

	"airsonic-cli/config"
	"airsonic-cli/request"
	"airsonic-cli/utils"
)

type Ping struct {
	SubsonicResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
	} `json:"subsonic-response"`
}

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
	return errors.New("PingAction => Exiting...")
}
