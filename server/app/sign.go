package App

import (
	"community-inviter/server/helpers"
	"community-inviter/server/models"
	"community-inviter/server/slack"
	"errors"

	"github.com/kataras/iris"
)

func SignHandler(ctx iris.Context, token string) error {
	_team_info_response, _team_info_error := slack.GetTeamInfo(token)
	if _team_info_error != nil {
		return _team_info_error
	}

	if _team_info_response.Ok == false {
		if _team_info_response.Error != "" {
			return errors.New(_team_info_response.Error)
		} else {
			return errors.New("Something happened on slack side while getting team information. Please try again.")
		}
	}

	_team_info, _team_error := models.GetTeam(_team_info_response.Team.ID)
	if _team_error != nil {
		return errors.New(_team_error.Error())
	}

	if _team_info.AccessToken == "" {
		return _up(ctx, *_team_info_response, token)
	} else {
		return _in(ctx, *_team_info_response, token)
	}
}

func _in(ctx iris.Context, team_info slack.TeamInfoResponse, access_token string) error {
	team, err := models.UpdateTeam(team_info.Team.ID, team_info.Team.Name, team_info.Team.Domain, team_info.Team.EmailDomain, access_token, team_info.Team.Icon)
	if err != nil {
		return err
	}

	helpers.SetSession(ctx, *team)

	return nil
}

func _up(ctx iris.Context, team_info slack.TeamInfoResponse, access_token string) error {
	team, err := models.CreateTeam(team_info.Team.ID, team_info.Team.Name, team_info.Team.Domain, team_info.Team.EmailDomain, team_info.Team.EnterpriseID, team_info.Team.EnterpriseName, access_token, team_info.Team.Icon)
	if err != nil {
		return err
	}

	helpers.SetSession(ctx, *team)
	return nil
}
