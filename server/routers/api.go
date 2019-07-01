package routers

import (
	App "community-inviter/server/app"
	"community-inviter/server/helpers"
	"community-inviter/server/models"
	"community-inviter/server/slack"
	"net/url"

	"github.com/kataras/iris"
)

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func APIHandler(api iris.Party) {
	api.Get("/sign", func(ctx iris.Context) {
		err := ctx.URLParam("error")
		if err != "" {
			ctx.Redirect("/?error=" + url.QueryEscape("You have to give the access in order to use our apilication."))
			return
		}
		code := ctx.URLParam("code")
		if code == "" {
			ctx.Redirect("/?error=" + url.QueryEscape("Slack didn't return requeired parameters. Please try again."))
			return
		}

		access_info, access_err := slack.GetAccessToken(code)
		if access_err != nil {
			ctx.Redirect("/?error=" + url.QueryEscape(access_err.Error()))
			return
		}

		if access_info.Ok == false || access_info.AccessToken == "" {
			ctx.Redirect("/?error=" + url.QueryEscape("Data is not correct. Please try again."))
			return
		}

		sign_error := App.SignHandler(ctx, access_info.AccessToken)
		if sign_error != nil {
			ctx.Redirect("/?error=" + url.QueryEscape(sign_error.Error()))
		}

		ctx.Redirect("/dashboard")
	})
	api.PartyFunc("/badge", BadgeHandler)
	api.Post("/logout", func(ctx iris.Context) {})
}

func BadgeHandler(api iris.Party) {
	api.Get("/get/{id:string}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		badge, err := models.GetBadge(id)
		if err != nil {
			ctx.JSON(&BasicResponse{false, err.Error()})
			return
		}

		if badge.ID == "" {
			ctx.JSON(&BasicResponse{false, "Not Found"})
			return
		}

		ctx.JSON(badge)
		return
	})
	api.Post("/create", func(ctx iris.Context) {
		team, err := helpers.GetSession(ctx)
		if err != nil {
			ctx.JSON(&BasicResponse{false, "Can not find the session. Please login first."})
			return
		}

		if team.ID == "" {
			ctx.JSON(&BasicResponse{false, "Can not find the session. Please login first."})
			return
		}

		var badge models.Badge
		ctx.ReadJSON(&badge)

		App.BadgeCreateHandler(ctx, badge)
	})
	api.Post("/update/{id: string}", func(ctx iris.Context) {
		team, err := helpers.GetSession(ctx)
		if err != nil {
			ctx.JSON(&BasicResponse{false, "Can not find the session. Please login first."})
			return
		}

		if team.ID == "" {
			ctx.JSON(&BasicResponse{false, "Can not find the session. Please login first."})
			return
		}

		var badge models.Badge
		ctx.ReadJSON(&badge)

		// ! DO SOMETHING
		models.UpdateBadge(badge)
	})
	api.Get("/remove/{id:string}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		err := models.RemoveBadge(id)
		if err != nil {
			ctx.JSON(&BasicResponse{false, err.Error()})
			return
		}
		ctx.JSON(&BasicResponse{false, ""})
		return
	})
}

// global sign function
func signOP(username, pass string) models.Team {
	user, err := models.GetTeam(username)
	if err != nil {
		team, err = models.CreateTeam()

	}

	return user
}
