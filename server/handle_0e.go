package server

import (
	"log"

	"github.com/minero/minero/proto/packet"
	"github.com/minero/minero/server/player"
)

// Handle0E handles incoming requests of packet 0x0E: PlayerAction
func Handle0E(server *Server, sender *player.Player) {
	pkt := new(packet.PlayerAction)
	pkt.ReadFrom(sender.Conn)

	log.Printf("PlayerAction: %+v", pkt)
}
