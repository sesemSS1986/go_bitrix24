package crm

type ExternalchannelConnector struct {
	TypeId             string `json:"TYPE_ID"`
	OriginatorId       string `json:"ORIGINATOR_ID"`
	Name               string `json:"NAME"`
	AppId              string `json:"APP_ID"`
	ChannelId          string `json:"CHANNEL_ID"`
	ExternalServerHost string `json:"EXTERNAL_SERVER_HOST"`
}
