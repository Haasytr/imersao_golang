package http

import (
	"encoding/json"
	"net/http"

	"github.com/Haastyr/imersao_golang/internal/events/usecase"
)

type EventsHandler struct {
	listEventsUseCase *usecase.ListEventsUseCase
	getEventUseCase   *usecase.GetEventUseCase
	buyTicketUseCase  *usecase.BuyTicketUseCase
	listSpotsUseCase  *usecase.ListSpotsUseCase
}

func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase: listEventsUseCase,
		getEventUseCase:   getEventUseCase,
		listSpotsUseCase:  listSpotsUseCase,
		buyTicketUseCase:  buyTicketsUseCase,
	}
}

func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.buyTicketUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
