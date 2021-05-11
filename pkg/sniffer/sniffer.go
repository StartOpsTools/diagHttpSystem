package sniffer

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func OpenLive(bpfFilter string) string {
	var result string
	
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Error("获取 devices 失败")
		return "获取 devices 失败"
	}
	
	//bpfFilter := fmt.Sprintf("tcp and port %d and dst host %s", port, host)
	fmt.Println("bpfFilter: ", bpfFilter)
	//bpfFilter := "tcp and port 80"
	
	fmt.Println("开始抓包时间: ", time.Now().Format("2006-01-02 15:04:05"))
	for _, device := range devices {
		//fmt.Println("device: ", device.Name)
		go PcapPackage(device, bpfFilter)
	}
	
	time.Sleep(60 * time.Second)
	fmt.Println("结束抓包时间: ", time.Now().Format("2006-01-02 15:04:05"))
	
	return result
}




func PcapPackage(device pcap.Interface, bpfFilter string) {
	
	handle, err := pcap.OpenLive(device.Name, snapshotLen, true, pcap.BlockForever)
	if err != nil {
		log.Error("err: ", err)
		//log.Fatal(err)
		return
	}
	
	defer handle.Close()
	
	if err := handle.SetBPFFilter(bpfFilter); err != nil {
		log.Error("set bpf filter failed. err: ", err)
		//log.Fatal("set bpf filter failed. err: ", err)
		return
	}
	
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		
		if packet.TransportLayer().LayerType() == layers.LayerTypeTCP {
			fmt.Println("TCP Transmission")
			TcpPacket(packet)
		}
		
	}
}