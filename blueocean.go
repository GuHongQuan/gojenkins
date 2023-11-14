package gojenkins

import (
	"context"
	"fmt"
	"strings"
)

type BlueOceanStage struct {
	Class string `json:"_class"`
	Links struct {
		Self struct {
			Class string `json:"_class"`
			Href  string `json:"href"`
		} `json:"self"`
		Actions struct {
			Class string `json:"_class"`
			Href  string `json:"href"`
		} `json:"actions"`
		Steps struct {
			Class string `json:"_class"`
			Href  string `json:"href"`
		} `json:"steps"`
	} `json:"_links"`
	Actions            []interface{} `json:"actions"`
	DisplayDescription interface{}   `json:"displayDescription"`
	DisplayName        string        `json:"displayName"`
	DurationInMillis   int           `json:"durationInMillis"`
	Id                 string        `json:"id"`
	Input              interface{}   `json:"input"`
	Result             string        `json:"result"`
	StartTime          string        `json:"startTime"`
	State              string        `json:"state"`
	Type               string        `json:"type"`
	CauseOfBlockage    interface{}   `json:"causeOfBlockage"`
	Edges              []struct {
		Class string `json:"_class"`
		Id    string `json:"id"`
		Type  string `json:"type"`
	} `json:"edges"`
	FirstParent interface{} `json:"firstParent"`
	Restartable bool        `json:"restartable"`
}

func (j *Job) GetBlueOceanRunStages(ctx context.Context, number int) ([]*BlueOceanStage, error) {
	href := "/blue/rest/organizations/jenkins/pipelines/%s/runs/%d/nodes/"

	href = fmt.Sprintf(href, strings.Join(strings.Split(j.Base, "/job/"), "/"), number)
	result := make([]*BlueOceanStage, 0)
	_, err := j.Jenkins.Requester.Get(ctx, href, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
