package tools

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	sysinfo "github.com/condemo/raspi-htmx-service/services/common/genproto/services/sys_info"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var fanFilePath string

func newMemInfo() *sysinfo.MemInfo {
	m, err := mem.VirtualMemory()
	checkErr(err)

	return &sysinfo.MemInfo{
		MemPercent: fmt.Sprint(strconv.FormatFloat(m.UsedPercent, 'f', 2, 64), "%"),
		MemUsed:    fmt.Sprint(strconv.FormatUint(m.Used/1024/1024, 10), "mb"),
		MemTotal:   fmt.Sprint(strconv.FormatUint(m.Total/1024/1024, 10), "mb"),
	}
}

func newDiskInfo() *sysinfo.DiskInfo {
	var dl []*sysinfo.USBDrive

	di := &sysinfo.DiskInfo{}

	parts, err := disk.Partitions(false)

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		checkErr(err)

		// TODO: Mucho codigo repetido, mejorable
		if u.Path == "/" {
			di.RootUsed = fmt.Sprint(
				strconv.FormatUint(u.Used/1024/1024/1024, 10), "GB")
			di.RootTotal = fmt.Sprint(
				strconv.FormatUint(u.Total/1024/1024/1024, 10), "GB")
		} else if strings.HasPrefix(u.Path, "/mnt/") {
			n := strings.ToTitle(strings.TrimPrefix(u.Path, "/mnt/"))
			uDrive := &sysinfo.USBDrive{
				Name:  n,
				Used:  fmt.Sprint(strconv.FormatUint(u.Used/1024/1024/1024, 10), "GB"),
				Total: fmt.Sprint(strconv.FormatUint(u.Total/1024/1024/1024, 10), "GB"),
			}
			dl = append(dl, uDrive)
		} else if strings.HasPrefix(u.Path, "/media/condemopi/") {
			n := fmt.Sprint(
				"USB ", strings.ToTitle(strings.TrimPrefix(u.Path, "/media/condemopi/")))
			uDrive := &sysinfo.USBDrive{
				Name:  n,
				Used:  fmt.Sprint(strconv.FormatUint(u.Used/1024/1024/1024, 10), "GB"),
				Total: fmt.Sprint(strconv.FormatUint(u.Total/1024/1024/1024, 10), "GB"),
			}
			dl = append(dl, uDrive)
		}
	}

	di.USBDrives = dl

	checkErr(err)

	return di
}

func newCpuInfo() *sysinfo.CpuInfo {
	cpuPer, err := cpu.Percent(0, true)
	checkErr(err)

	c := &sysinfo.CpuInfo{}

	for _, cpu := range cpuPer {
		c.CoreInfoList = append(
			c.CoreInfoList,
			fmt.Sprintf("%s%%", strconv.FormatFloat(cpu, 'f', 2, 64)))
	}

	f, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	checkErr(err)
	defer f.Close()

	cb, err := io.ReadAll(f)
	checkErr(err)

	cs := strings.TrimSuffix(string(cb), "\n")

	cTemp, err := strconv.ParseInt(cs, 10, 64)
	checkErr(err)

	c.CpuTemp = fmt.Sprintf("%d°C", cTemp/1000)

	return c
}

func newFanInfo() *sysinfo.FanInfo {
	f := &sysinfo.FanInfo{}

	stateF, err := os.Open("/sys/devices/virtual/thermal/cooling_device0/cur_state")
	checkErr(err)
	defer stateF.Close()

	s, err := io.ReadAll(stateF)
	checkErr(err)

	if state, _ := strconv.ParseInt(
		strings.TrimSuffix(string(s), "\n"), 10, 8); state >= 1 {
		f.FanStatus = true
	}

	ff, err := os.Open(fanFilePath)
	checkErr(err)
	defer ff.Close()

	fb, err := io.ReadAll(ff)
	checkErr(err)

	fanSpeed := strings.TrimSuffix(string(fb), "\n")
	checkErr(err)

	f.FanSpeed = fanSpeed

	return f
}

func newNetInfo() *sysinfo.NetInfo {
	nf, err := os.Open("/home/condemopi/scripts/custom_output/net_speed")
	checkErr(err)

	nb, err := io.ReadAll(nf)
	checkErr(err)

	netSpeed := strings.TrimSuffix(string(nb), "\n")
	speedSlice := strings.Split(netSpeed, "|")

	// TODO: Cutre Fix
	if len(speedSlice) < 2 {
		speedSlice = []string{"0,0 kb", "0,0 kb"}
	}

	return &sysinfo.NetInfo{
		NetUp:   speedSlice[0],
		NetDown: speedSlice[1],
	}
}

func NewSysInfo() *sysinfo.SysInfo {
	return &sysinfo.SysInfo{
		DiskInfo: newDiskInfo(),
		MemInfo:  newMemInfo(),
		CpuInfo:  newCpuInfo(),
		FanInfo:  newFanInfo(),
		NetInfo:  newNetInfo(),
	}
}

// this init check where the fan1_input file is located
func init() {
	p := "/sys/devices/platform/cooling_fan/hwmon/"
	items, err := os.ReadDir(p)
	if err != nil {
		log.Fatalf("error leyendo Items %s", err)
	}

	for _, item := range items {
		if item.IsDir() {
			subItems, err := os.ReadDir(p + item.Name())
			if err != nil {
				log.Fatalf("error leyendo subItems %s", err)
			}
			for _, subItem := range subItems {
				if !subItem.IsDir() {
					if subItem.Name() == "fan1_input" {
						fanFilePath = p + item.Name() + "/" + subItem.Name()
					}
				}
			}
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("sys_info error: %v", err)
	}
}
