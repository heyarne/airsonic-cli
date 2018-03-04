package scan

import (
  "log"
  "strconv"
	"errors"
  "encoding/json"

  "airsonic-cli/utils"
  "airsonic-cli/request"
  "airsonic-cli/config"
)

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

func StartScanAction(conf *config.Config) error {
  if config.IsVerbose(conf) { utils.InfoMsg("Start scan send to " + config.GetServer(conf)) }
  var data = request.Get(conf, "/rest/startScan", "")

  if request.CheckResponse(conf, data) {
    if config.IsVerbose(conf) { utils.InfoMsg("Scan started successfully") }
    return nil
  }
  return errors.New("StartScanAction => Exiting...")
}

func ScanStatusAction(conf *config.Config) error {
  if config.IsVerbose(conf) { utils.InfoMsg("Get scan status from " + config.GetServer(conf)) }
  var data = request.Get(conf, "/rest/getScanStatus", "")

  if request.CheckResponse(conf, data) {
    scan := Scan{}
    jsonErr := json.Unmarshal(data, &scan)
    if jsonErr != nil { log.Fatal(jsonErr) }
    utils.InfoMsg("Scanning => " + strconv.FormatBool(scan.SubsonicResponse.ScanStatus.Scanning))
    utils.InfoMsg("Count    => " + strconv.Itoa(scan.SubsonicResponse.ScanStatus.Count))
    return nil
  }
  return errors.New("ScanStatusAction => Exiting...")
}
