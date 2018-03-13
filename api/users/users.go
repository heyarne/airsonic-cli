package users

// FIXME: Add folders a user has access to (needs extra request)

import (
	"encoding/json"

	"airsonic-cli/config"
	"airsonic-cli/request"
	"airsonic-cli/utils"
)

type User struct {
	Name              string `json:"username"`
	Email             string `json:"email"`
	ScrobblingEnabled bool   `json:"srobblingEnabled"`
	AdminRole         bool   `json:"adminRole"`
	SettingsRole      bool   `json:"settingsRole"`
	UploadRole        bool   `json:"uploadRole"`
	DownloadRole      bool   `json:"downloadRole"`
	PlaylistRole      bool   `json:"playlistRole"`
	CoverArtRole      bool   `json:"coverArtRole"`
	CommentRole       bool   `json:"commentRole"`
	PodcastRole       bool   `json:"podcastRole"`
	StreamRole        bool   `json:"streamRole"`
	JukeboxRole       bool   `json:"jukeboxRole"`
	ShareRole         bool   `json:"shareRole"`
}

type GetUserResponse struct {
	SubsonicResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
		User    User   `json:"user"`
	} `json:"subsonic-response"`
}

type ListUsersResponse struct {
	SubsonicResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
		Users   struct {
			User []User `json:"user"`
		} `json:"users"`
	} `json:"subsonic-response"`
}

func GetUserAction(conf *config.Config, name string) error {
	if config.IsVerbose(conf) {
		utils.InfoMsg("Get user status for " + name + " from " + config.GetServer(conf))
	}
	data := request.Get(conf, "/rest/getUser", "username="+name)
	if request.CheckResponse(conf, data) {
		res := GetUserResponse{}
		jsonErr := json.Unmarshal(data, &res)
		if jsonErr != nil {
			return jsonErr
		}
		user := res.SubsonicResponse.User
		utils.PrettyPrint(user)
	}
	return nil
}

func ListUsersAction(c *config.Config) error {
	if config.IsVerbose(c) {
		utils.InfoMsg("Listing users on " + config.GetServer(c))
	}
	data := request.Get(c, "/rest/getUsers", "")
	if request.CheckResponse(c, data) {
		res := ListUsersResponse{}
		jsonErr := json.Unmarshal(data, &res)
		if jsonErr != nil {
			return jsonErr
		}
		users := res.SubsonicResponse.Users.User
		for _, u := range users {
			utils.PrettyPrint(u)
		}
	}
	return nil
}

// FIXME: Implement this
func CreateUserAction(c *config.Config) error {
	if config.IsVerbose(c) {
		utils.InfoMsg("Creating new user on " + config.GetServer(c))
	}
	// data := request.Get(c)
	return nil
}
