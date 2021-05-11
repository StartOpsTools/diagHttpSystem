package sniffer

import (
	"fmt"
	"github.com/StartOpsTools/diagHttpSystem/v1/pkg/convert"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

)

type Header struct {
	SourcePort    uint16
	DestPort      uint16
	SeqNumber     uint32
	AckNumber     uint32
	HeaderLength  uint8
	Reserved      uint8
	NS            uint8
	CWR           uint8
	ECE           uint8
	URG           uint8
	ACK           uint8
	PSH           uint8
	RST           uint8
	SYN           uint8
	FIN           uint8
	WindowsSize   uint8
	Checksum      uint8
	UrgentPointer uint8
	Options       uint8
	Padding       uint8
}

func TcpPacket(packet gopacket.Packet) {
	tpcLayer := packet.Layer(layers.LayerTypeTCP)
	tcpLayerPack, b := tpcLayer.(*layers.TCP)
	if b {
		Contents := tcpLayerPack.Contents

		contentsHex := convert.PayloadToHex(Contents)
		fmt.Println("contentsHex: ", contentsHex)
		
		Payload := tcpLayerPack.Payload
		
		fmt.Println("Contents: ", Contents)
		fmt.Println("Payload: ", Payload)
		//fmt.Println("payload: ", string(Payload))
		
		SrcPort := tcpLayerPack.SrcPort
		DstPort := tcpLayerPack.DstPort
		Seq := tcpLayerPack.Seq
		Ack := tcpLayerPack.Ack
		
		URG := tcpLayerPack.URG
		ACK := tcpLayerPack.ACK
		PSH := tcpLayerPack.PSH
		RST := tcpLayerPack.RST
		SYN := tcpLayerPack.SYN
		FIN := tcpLayerPack.FIN
		Options := tcpLayerPack.Options
		Padding := tcpLayerPack.Padding
		
		fmt.Println("SrcPort: ", SrcPort, ".    DstPort: ", DstPort, ".    SeqNumber: ", Seq, ".    Ack: ", Ack,
			".    URG: ", URG, ".    ACK: ", ACK, ".    PSH: ", PSH ,".    RST: ", RST, ".    SYN: ", SYN,
			".    FIN: ", FIN, ".    Options: ", Options, ".    Padding: ", Padding)
		

	}
}
