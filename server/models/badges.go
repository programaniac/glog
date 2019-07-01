package models

import (
	"community-inviter/server/database"
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var __badges_collection string = "badges"

type Badge struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	Title      string        `json:"title" bson:"title"`
	Icon       string        `json:"icon" bson:"icon"`
	Type       string        `json:"type" bson:"type"`
	Appearance struct {
		Position string `json:"position" bson:"position"`
		Colors   struct {
			Background       string `json:"background" bson:"background"`
			Text             string `json:"text" bson:"text"`
			ButtonBackground string `json:"button_background" bson:"button_background"`
			ButtonText       string `json:"button_text" bson:"button_text"`
		} `json:"colors" bson:"colors"`
		Font string `json:"font" bson:"font"`
	} `json:"appearance" bson:"appearance"`
	Content struct {
		MainText   string `json:"main_text" bson:"main_text"`
		ButtonText string `json:"button_text" bson:"button_text"`
	} `json:"content" bson:"content"`
	Success struct {
		MainText  string `json:"main_text" bson:"main_text"`
		Behaviour struct {
			State string `json:"state" bson:"state"`
			Time  int32  `json:"time" bson:"time"`
		} `json:"behaviour" bson:"behaviour"`
	} `json:"success" bson:"success"`
	Questions []struct {
		QuestionTitle   string `json:"question_title" bson:"question_title"`
		QuestionType    string `json:"question_type" bson:"question_type"`
		QuestionOptions []struct {
			Default bool   `json:"default" bson:"default"`
			Option  string `json:"option" bson:"option"`
		} `json:"question_options,omitempty" bson:"question_options,omitempty"`
		Required bool `json:"required" bson:"required"`
	} `json:"questions,omitempty" bson:"questions,omitempty"`
	Status      string    `json:"status" bson:"status"` // active, deactive, removed...
	CreatedDate time.Time `json:"created_date" bson:"created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date"`
}

func GetBadge(id string) (*Badge, error) {
	collection, err := database.GetCollection(__badges_collection)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	_badge := &Badge{}
	collection.Find(bson.M{"_id": id}).One(_badge)

	return _badge, nil
}

func CreateBadge(badge Badge) (*Badge, error) {
	collection, err := database.GetCollection(__badges_collection)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	badge.ID = bson.NewObjectId()
	badge.CreatedDate = time.Now()
	badge.UpdatedDate = time.Now()
	badge.Status = "active"

	err = collection.Insert(badge)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	return &badge, nil
}

func RemoveBadge(id string) error {
	collection, err := database.GetCollection(__badges_collection)
	if err != nil {
		return errors.New("Database connection error. Please try again or contact us.")
	}

	err = collection.Remove(bson.M{"_id": id})
	if err != nil {
		return errors.New("Database connection error. Please try again or contact us.")
	}

	return nil
}

func UpdateBadge(badge Badge) (*Badge, error) {
	collection, err := database.GetCollection(__badges_collection)
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	badge.UpdatedDate = time.Now()
	err = collection.Update(bson.M{"_id": badge.ID}, bson.M{"$set": badge})
	if err != nil {
		return nil, errors.New("Database connection error. Please try again or contact us.")
	}

	return &badge, nil
}
