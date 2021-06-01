package io

import (
	"Glib/common"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	gnet "github.com/shirou/gopsutil/net"
	"net"
	"time"
)

//CpuInfo 获取 cpu 信息, 类似 cat /proc/cpuinfo
func CpuInfo() []string {
	cpuInfos, _ := cpu.Info()
	infos := []string{}
	for _, ci := range cpuInfos {
		infos = append(infos, ci.String())
	}
	return infos
}

//CpuLoad 获取Cpu在一段时间内的负载
func CpuLoad(durations []time.Duration) (map[time.Duration][]float64, error) {
	stats := map[time.Duration][]float64{}
	for _, duration := range durations {
		percent, _ := cpu.Percent(duration, false)
		fmt.Printf("cpu percent:%v\n", percent)
		stats[duration] = percent
	}
	return stats, nil
}

//MemInfo 获取内存使用情况
func MemInfo() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

//HostInfo 获取主机信息
func HostInfo() (*host.InfoStat, error) {
	return host.Info()
}

type DiskStat struct {
	Filesystem  string `json:"filesystem"`
	Size        string `json:"size"`
	Used        string `json:"used"`
	Avail       string `json:"avail"`
	UsedPercent string `json:"used_percent"`
	Mounted     string `json:"mounted"`
}

//DiskInfo 获取磁盘使用情况，类似 df
func DiskInfo() ([]DiskStat, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	var stats []DiskStat
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		stat := DiskStat{
			Filesystem:  part.Device,
			Size:        common.ByteFormat(diskInfo.Total),
			Used:        common.ByteFormat(diskInfo.Used),
			Avail:       common.ByteFormat(diskInfo.Free),
			UsedPercent: fmt.Sprintf("%.0f%%", diskInfo.UsedPercent),
			Mounted:     part.Mountpoint,
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

//NetworkCardInfo 获取网卡信息,类似 ifconfig eth0
func NetworkCardInfo() []string  {
	info, _ := gnet.IOCounters(true)
	infos := []string{}
	for index, v := range info {
		str := fmt.Sprintf("%v:%v send:%v recv:%v \n", index, v, v.BytesSent, v.BytesRecv)
		infos = append(infos, str)
	}
	return infos
}

//LocalIP 返回本机IP地址
func LocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

