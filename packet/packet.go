package packet

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Packet struct {
	Magic string
	Type  int32
	Size  int32
	Data  string
}

func NewPacket(magicValue string, typeValue int32, sizeValue int32, dataValue string) *Packet {

	return &Packet{Magic: magicValue, Type: typeValue, Size: sizeValue, Data: dataValue}
}

func BuildPacketFromConnection(conn net.Conn) *Packet {

	magicByte := make([]byte, HEADER_ITEM_LENGTH)
	typeByte := make([]byte, HEADER_ITEM_LENGTH)
	sizeByte := make([]byte, HEADER_ITEM_LENGTH)

	_, _ = conn.Read(magicByte)
	_, _ = conn.Read(typeByte)
	_, _ = conn.Read(sizeByte)

	var typeValue, sizeValue int32
	var magicValue, dataValue string

	magicValue = string(magicByte)
	binary.Read(bytes.NewBuffer(typeByte), binary.BigEndian, &typeValue)
	binary.Read(bytes.NewBuffer(sizeByte), binary.BigEndian, &sizeValue)

	dataByte := make([]byte, sizeValue)
	_, _ = conn.Read(dataByte)

	dataValue = string(dataByte)

	return NewPacket(magicValue, typeValue, sizeValue, dataValue)
}

// Encode a Packet to byte slice
func (this *Packet) Encode() []byte {

	// length of data
	dataLength := len(this.Data)

	// legth of header
	headerLength := HEADER_ITEM_LENGTH * HEADER_ITEM_COUNT

	// length of packet
	packetLength := dataLength + headerLength

	buffer := make([]byte, packetLength)

	copy(buffer[:HEADER_ITEM_LENGTH], MAGIC_REQ)
	binary.BigEndian.PutUint32(buffer[HEADER_ITEM_LENGTH:HEADER_ITEM_LENGTH*2], uint32(this.Type))
	binary.BigEndian.PutUint32(buffer[HEADER_ITEM_LENGTH*2:headerLength], uint32(dataLength))
	copy(buffer[headerLength:], []byte(this.Data))

	return buffer
}
