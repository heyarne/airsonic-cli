package scan

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"airsonic-cli/config"
	"airsonic-cli/request"
	"airsonic-cli/utils"
)

// Scan define the json subsonic response structure for scan calls
type Scan struct {
	SubsonicResponse struct {
		Status     string `json:"status"`
		Version    string `json:"version"`
		ScanStatus struct {
			Scanning bool `json:"scanning"`
			Count    int  `json:"count"`
		} `json:"scanStatus"`
	} `json:"subsonic-response"`
}

// StartScanAction call a scan on the server's API
func StartScanAction(conf *config.Config) error {
	if config.IsVerbose(conf) {
		utils.InfoMsg("Start scan send to " + config.GetServer(conf))
	}
	var data = request.Get(conf, "/rest/startScan", "")

	if request.CheckResponse(conf, data) {
		if config.IsVerbose(conf) {
			utils.InfoMsg("Scan started successfully")
		}
		return nil
	}
	return errors.New("error while calling startScan enpoint")
}

// ScanStatusAction call a scan status on the server's API
func ScanStatusAction(conf *config.Config) error {
	if config.IsVerbose(conf) {
		utils.InfoMsg("Get scan status from " + config.GetServer(conf))
	}
	var data = request.Get(conf, "/rest/getScanStatus", "")

	if request.CheckResponse(conf, data) {
		scan := Scan{}
		jsonErr := json.Unmarshal(data, &scan)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		utils.InfoMsg("Scanning => " + strconv.FormatBool(scan.SubsonicResponse.ScanStatus.Scanning))
		utils.InfoMsg("Count    => " + strconv.Itoa(scan.SubsonicResponse.ScanStatus.Count))
		return nil
	}
	return errors.New("error while calling getScanStatus enpoint")
}
