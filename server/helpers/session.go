package helpers

import (
	"community-inviter/server/models"
	"encoding/json"
	"errors"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var sess = sessions.New(sessions.Config{
	// Cookie string, the session's client cookie name, for example: "mysessionid"
	//
	// Defaults to "irissessionid"
	Cookie: "lahmacun",
	// it's time.Duration, from the time cookie is created, how long it can be alive?
	// 0 means no expire.
	// -1 means expire when browser closes
	// or set a value, like 2 hours:
	Expires: time.Hour * 24,
})

func SetSession(ctx iris.Context, team models.Team) {
	s := sess.Start(ctx)
	usr, _ := json.Marshal(team)
	s.Set("team", string(usr))
}

func GetSession(ctx iris.Context) (*models.Team, error) {
	usr := sess.Start(ctx).GetString("user")
	if usr == "" {
		return nil, errors.New("User not found !")
	}
	user := &models.Team{}
	err := json.Unmarshal([]byte(usr), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckUserPremium(ctx iris.Context) bool {
	usr := sess.Start(ctx).GetString("user")
	if usr == "" {
		return false
	}
	user := &models.Team{}
	err := json.Unmarshal([]byte(usr), user)
	if err != nil {
		return false
	}

	return user.Subscription
}
