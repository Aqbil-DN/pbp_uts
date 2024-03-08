package model

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Game struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MaxPlayers int    `json:"max_players"`
}

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	IDGame   int    `json:"game_id"`
}

type Participant struct {
	ID        int `json:"id"`
	IDRoom    int `json:"room_id"`
	IDAccount int `json:"account_id"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetAllRoomsResponseDataRoom struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
}

type GetAllRoomsResponseData struct {
	Rooms []GetAllRoomsResponseDataRoom `json:"rooms"`
}

type GetAllRoomsResponse struct {
	Status int                     `json:"status"`
	Data   GetAllRoomsResponseData `json:"data"`
}

type GetDetailRoomResponseParticipant struct {
	ID        int    `json:"id"`
	AccountID int    `json:"account_id"`
	Username  string `json:"username"`
}

type GetDetailRoomResponseDataRoom struct {
	ID           int                                `json:"id"`
	RoomName     string                             `json:"room_name"`
	Participants []GetDetailRoomResponseParticipant `json:"participants"`
}

type GetDetailRoomResponseData struct {
	Room GetDetailRoomResponseDataRoom `json:"room"`
}

type GetDetailRoomResponse struct {
	Status int                       `json:"status"`
	Data   GetDetailRoomResponseData `json:"data"`
}
