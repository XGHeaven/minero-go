package server

import (
	"log"

	"github.com/minero/minero/proto/packet"
	"github.com/minero/minero/server/player"
)

// Handle6C handles incoming requests of packet 0x6C: EnchantItem
func Handle6C(server *Server, sender *player.Player) {
	pkt := new(packet.EnchantItem)
	pkt.ReadFrom(sender.Conn)

	log.Printf("EnchantItem: %+v", pkt)
}
