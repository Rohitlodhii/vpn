package protocol

type PacketType uint8

const (
	Ping PacketType = 1
	Pong PacketType = 2
)

type Packet struct {
	Type    PacketType
	Message []byte
}
