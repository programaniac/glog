package slack

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var _client_id string = "84219785124.141983912389"
var _client_secret string = "c3a4cb490a35ef2c4544386456a84521"

// https://slack.com/oauth/authorize?client_id=84219785124.141983912389&scope=admin,client

type AccessTokenResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
	Error       string `json:"error,omitempty"`
}

func GetAccessToken(code string) (*AccessTokenResponse, error) {
	uri := "https://slack.com/api/oauth.access?client_id=" + _client_id + "&client_secret=" + _client_secret + "&code=" + code
	resp, err := http.Get(uri)
	if err != nil {
		return nil, errors.New("Something happened while connectiong to slack. Please try again.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	data := &AccessTokenResponse{}

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.New("Unsupported response from slack. Please try again.")
	}
	return data, nil
}

type TeamInfoResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
	Team  struct {
		ID             string        `json:"id"`
		Name           string        `json:"name"`
		Domain         string        `json:"domain"`
		EmailDomain    string        `json:"email_domain"`
		Icon           TeamInfoIcons `json:"icon"`
		EnterpriseID   string        `json:"enterprise_id"`
		EnterpriseName string        `json:"enterprise_name"`
	} `json:"team,omitempty"`
}

type TeamInfoIcons struct {
	Image34      string `json:"image_34" bson:"image_34"`
	Image44      string `json:"image_44" bson:"image_44"`
	Image68      string `json:"image_68" bson:"image_68"`
	Image88      string `json:"image_88" bson:"image_88"`
	Image102     string `json:"image_102" bson:"image_102"`
	Image132     string `json:"image_132" bson:"image_132"`
	ImageDefault bool   `json:"image_default" bson:"image_default"`
}

func GetTeamInfo(access_token string) (*TeamInfoResponse, error) {
	uri := "https://slack.com/api/team.info?token=" + access_token
	resp, err := http.Get(uri)
	if err != nil {
		return nil, errors.New("Something happened while connectiong to slack. Please try again.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	data := &TeamInfoResponse{}

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.New("Unsupported response from slack. Please try again.")
	}
	return data, nil
}
