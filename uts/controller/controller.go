package controller

import (
	h "aqbiluts/handler"
	"aqbiluts/model"
	m "aqbiluts/model"
	"aqbiluts/response"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := h.ConnectGormDBHandler()

	var rooms []m.Room

	err := db.Find(&rooms).Error
	if err != nil {
		response.PrintError(500, "Failed to retrieve rooms", w)
		return
	}

	var roomsResponse []model.GetAllRoomsResponseDataRoom
	for _, room := range rooms {
		roomResponse := model.GetAllRoomsResponseDataRoom{
			ID:       room.ID,
			RoomName: room.RoomName,
		}
		roomsResponse = append(roomsResponse, roomResponse)
	}

	response := model.GetAllRoomsResponse{
		Status: 200,
		Data: model.GetAllRoomsResponseData{
			Rooms: roomsResponse,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := h.ConnectGormDBHandler()
	err := r.ParseForm()
	if err != nil {
		response.PrintError(400, "Failed to parse form data", w)
		return
	}

	roomID := r.URL.Query().Get("id")

	var room m.Room
	err = db.Where("id = ?", roomID).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.PrintError(404, "Room not found", w)
		} else {
			response.PrintError(500, "Failed to retrieve room", w)
		}
		return
	}

	var participants []m.Participant
	err = db.Where("id_room = ?", roomID).Find(&participants).Error
	if err != nil {
		response.PrintError(500, "Failed to retrieve participants", w)
		return
	}

	var participantsResponse []model.GetDetailRoomResponseParticipant
	for _, participant := range participants {
		var account m.Account
		err := db.Where("id = ?", participant.IDAccount).First(&account).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response.PrintError(404, "Account not found", w)
			} else {
				response.PrintError(500, "Failed to retrieve account", w)
			}
			return
		}

		participantResponse := model.GetDetailRoomResponseParticipant{
			ID:        participant.ID,
			AccountID: participant.IDAccount,
			Username:  account.Username,
		}

		participantsResponse = append(participantsResponse, participantResponse)
	}
	response := model.GetDetailRoomResponse{
		Status: 200,
		Data: model.GetDetailRoomResponseData{
			Room: model.GetDetailRoomResponseDataRoom{
				ID:           room.ID,
				RoomName:     room.RoomName,
				Participants: participantsResponse,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
