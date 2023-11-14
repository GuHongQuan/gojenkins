package gojenkins

import (
	"context"
	"fmt"
	"strings"
)

type BlueOceanStage struct {
	DisplayName string      `json:"displayName"`
	Id          string      `json:"id"`
	Input       interface{} `json:"input"`
	Result      string      `json:"result"`
	State       string      `json:"state"`
	Type        string      `json:"type"`
	FirstParent interface{} `json:"firstParent"`
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
