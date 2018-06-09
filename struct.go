package main

import "time"

type IotButtonEvent struct {
	DeviceInfo struct {
		DeviceID      string  `json:"deviceId"`
		Type          string  `json:"type"`
		RemainingLife float64 `json:"remainingLife"`
		Attributes    struct {
			ProjectRegion      string `json:"projectRegion"`
			ProjectName        string `json:"projectName"`
			PlacementName      string `json:"placementName"`
			DeviceTemplateName string `json:"deviceTemplateName"`
		} `json:"attributes"`
	} `json:"deviceInfo"`
	DeviceEvent struct {
		ButtonClicked struct {
			ClickType    string    `json:"clickType"`
			ReportedTime time.Time `json:"reportedTime"`
		} `json:"buttonClicked"`
	} `json:"deviceEvent"`
	PlacementInfo struct {
		ProjectName   string `json:"projectName"`
		PlacementName string `json:"placementName"`
		Attributes    struct {
			Key string `json:"key"`
		} `json:"attributes"`
		Devices struct {
			SampleRequest string `json:"Sample-Request"`
		} `json:"devices"`
	} `json:"placementInfo"`
}