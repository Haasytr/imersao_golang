package main

import (
	"database/sql"
	"net/http"

	"github.com/Haastyr/imersao_golang/internal/events/infra/repository"
	"github.com/Haastyr/imersao_golang/internal/events/infra/service"
	"github.com/Haastyr/imersao_golang/internal/events/usecase"

	httpHandler "github.com/Haastyr/imersao_golang/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}
	db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}
	partnerBaseURLs := map[int]string{
		1: "http://localhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		getEventUseCase,
		buyTicketsUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

}
