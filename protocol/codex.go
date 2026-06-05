package protocol

import (
	"bytes"
	"encoding/binary"
)

func Encode(p Packet) ([]byte, error) {

	buf := new(bytes.Buffer)

	buf.WriteByte(byte(p.Type))

	err := binary.Write(
		buf,
		binary.BigEndian,
		uint16(len(p.Message)),
	)

	if err != nil {
		return nil, err
	}

	buf.Write([]byte(p.Message))

	return buf.Bytes(), nil
}

func Decode(data []byte) (Packet, error) {

	var p Packet

	p.Type = PacketType(data[0])

	msgLen := binary.BigEndian.Uint16(data[1:3])

	p.Message = data[3 : 3+msgLen]

	return p, nil
}
