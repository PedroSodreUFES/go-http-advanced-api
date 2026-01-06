package api

import (
	"main/internal/services"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type Api struct {
	AuctionLobby   services.AuctionLobby
	BidsService    services.BidsService
	ProductService services.ProductService
	Router         *chi.Mux
	Sessions       *scs.SessionManager
	UserService    services.UserService
	WsUpgrader     websocket.Upgrader
}
