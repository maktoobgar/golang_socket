package dto

type Room struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ConnectionsLength int    `json:"connection_length"`
}
