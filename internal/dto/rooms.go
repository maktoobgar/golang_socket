package dto

type Room struct {
	Name              string `json:"name"`
	ConnectionsLength int    `json:"connection_length"`
}
