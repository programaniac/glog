package App

import (
	"community-inviter/server/helpers"
	"community-inviter/server/models"
	"errors"

	"github.com/kataras/iris"
)

func BadgeCreateHandler(ctx iris.Context, badge models.Badge) (*models.Badge, error) {
	if BadgeCheckFields(ctx, badge) == false {
		return nil, errors.New("Premium features can not use without premium.")
	}

	return nil, nil
}

func BadgeCheckFields(ctx iris.Context, badge models.Badge) bool {
	premium := helpers.CheckUserPremium(ctx)
	if premium == false {
		if badge.Type == "question" {
			return false
		}
	}

	return true
}
