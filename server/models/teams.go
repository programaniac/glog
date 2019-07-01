package models

import (
	"community-inviter/server/database"
	"community-inviter/server/slack"
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var _teams_collection string = "teams"

type Team struct {
	ID             string              `json:"_id" bson:"_id"`
	Name           string              `json:"name" bson:"name"`
	Domain         string              `json:"domain" bson:"domain"`
	EmailDomain    string              `json:"email_domain" bson:"email_domain"`
	Icon           slack.TeamInfoIcons `json:"icon" bson:"icon"`
	EnterpriseID   string              `json:"enterprise_id" bson:"enterprise_id"`
	EnterpriseName string              `json:"enterprise_name" bson:"enterprise_name"`
	AccessToken    string              `json:"access_token" bson:"access_token"`
	Subscription   bool                `json:"subscription" bson:"subscription"`
	SubscriptionID string              `json:"subscription_id" bson:"subscription_id"`
	RegisterDate   time.Time           `json:"register_date" bson:"register_date"`
}

func GetTeam(id string) (*Team, error) {
	collection, err := database.GetCollection(_teams_collection)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	_team := &Team{}
	collection.Find(bson.M{"_id": id}).One(_team)

	return _team, nil
}

func UpdateTeam(id, name, domain, email_domain, access_token string, icons slack.TeamInfoIcons) (*Team, error) {
	collection, err := database.GetCollection(_teams_collection)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	_team := &Team{}
	err = collection.Find(bson.M{"_id": id}).One(_team)
	if err != nil {
		return nil, errors.New("User not found. Please try again.")
	}

	_team.Name = name
	_team.Domain = domain
	_team.EmailDomain = email_domain
	_team.AccessToken = access_token
	_team.Icon = icons

	err = collection.Update(bson.M{"_id": id}, bson.M{"$set": _team})
	if err != nil {
		return nil, errors.New("Update operation unsuccessfull. Please try again...")
	}
	return _team, nil
}

func CreateTeam(id, name, domain, email_domain, enterprise_id, enterprise_name, access_token string, icons slack.TeamInfoIcons) (*Team, error) {
	collection, err := database.GetCollection(_teams_collection)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	_team := &Team{}
	_team.ID = id
	_team.Name = name
	_team.Domain = domain
	_team.EmailDomain = email_domain
	_team.EnterpriseID = enterprise_id
	_team.EnterpriseName = enterprise_name
	_team.Icon = icons
	_team.AccessToken = access_token
	_team.Subscription = false
	_team.SubscriptionID = ""
	_team.RegisterDate = time.Now()

	err = collection.Insert(_team)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	return _team, nil
}
