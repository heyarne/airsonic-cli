package request

import (
  "log"
  "time"
  "net/url"
  "net/http"
  "io/ioutil"

  "airsonic-cli/utils"
  "airsonic-cli/config"
)

func Get(conf *config.Config, endpoint string, payload string) []byte {
  	u, err := url.Parse(config.GetServer(conf) + endpoint + "?"+ loadCommonPayload(conf) + payload)
  	if err != nil {
  		log.Fatal(err)
  	}
    q := u.Query()
    q.Set("u", config.GetUsername(conf))
    q.Set("t", config.GetToken(conf))
    q.Set("s", config.GetSalt(conf))
    q.Set("c", config.GetAppName(conf))
    q.Set("v", config.GetAPIVersion(conf))
    q.Set("f", config.GetAPIFormat(conf))
    u.RawQuery = q.Encode()
    if config.IsVerbose(conf) { utils.InfoMsg("GET " + u.String()) }
  	req, reqErr := http.NewRequest(http.MethodGet, u.String(), nil)
    if reqErr != nil {
      log.Fatal(reqErr)
    }
    httpClient := http.Client{
    	Timeout: time.Second * 2,
    }
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
