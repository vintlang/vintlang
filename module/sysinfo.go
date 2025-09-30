package module

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/vintlang/vintlang/object"
)

var SysInfoFunctions = map[string]object.ModuleFunction{}

func init() {
	SysInfoFunctions["os"] = getOS
	SysInfoFunctions["arch"] = getArch
	SysInfoFunctions["memInfo"] = getMemInfo
	SysInfoFunctions["cpuInfo"] = getCPUInfo
	SysInfoFunctions["diskInfo"] = getDiskInfo
	SysInfoFunctions["netInfo"] = getNetInfo
}

func getOS(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"sysinfo", "os",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`sysinfo.os() -> "linux"`,
		)
	}
	return &object.String{Value: runtime.GOOS}
}

func getArch(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"sysinfo", "arch",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`sysinfo.arch() -> "amd64"`,
		)
	}
	return &object.String{Value: runtime.GOARCH}
}

func getMemInfo(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"sysinfo", "memInfo",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`sysinfo.memInfo() -> {"total": "8GB", "available": "4GB", "used": "4GB", "free": "4GB", "percent": "50.0"}`,
		)
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get memory info: %v", err)}
	}

	memInfoMap := make(map[object.HashKey]object.DictPair)

	// Add memory information to the map
	memInfoMap[(&object.String{Value: "total"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "total"},
		Value: &object.String{Value: fmt.Sprintf("%.2f GB", float64(memStat.Total)/1024/1024/1024)},
	}
	memInfoMap[(&object.String{Value: "available"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "available"},
		Value: &object.String{Value: fmt.Sprintf("%.2f GB", float64(memStat.Available)/1024/1024/1024)},
	}
	memInfoMap[(&object.String{Value: "used"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "used"},
		Value: &object.String{Value: fmt.Sprintf("%.2f GB", float64(memStat.Used)/1024/1024/1024)},
	}
	memInfoMap[(&object.String{Value: "free"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "free"},
		Value: &object.String{Value: fmt.Sprintf("%.2f GB", float64(memStat.Free)/1024/1024/1024)},
	}
	memInfoMap[(&object.String{Value: "percent"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "percent"},
		Value: &object.String{Value: fmt.Sprintf("%.2f", memStat.UsedPercent)},
	}

	return &object.Dict{Pairs: memInfoMap}
}

func getCPUInfo(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"sysinfo", "cpuInfo",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`sysinfo.cpuInfo() -> {"cores": "4", "model": "Intel Core i7", "usage": "25.5"}`,
		)
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get CPU info: %v", err)}
	}

	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get CPU usage: %v", err)}
	}

	cpuInfoMap := make(map[object.HashKey]object.DictPair)

	if len(cpuInfo) > 0 {
		cpuInfoMap[(&object.String{Value: "cores"}).HashKey()] = object.DictPair{
			Key:   &object.String{Value: "cores"},
			Value: &object.String{Value: strconv.Itoa(int(cpuInfo[0].Cores))},
		}
		cpuInfoMap[(&object.String{Value: "model"}).HashKey()] = object.DictPair{
			Key:   &object.String{Value: "model"},
			Value: &object.String{Value: cpuInfo[0].ModelName},
		}
		cpuInfoMap[(&object.String{Value: "frequency"}).HashKey()] = object.DictPair{
			Key:   &object.String{Value: "frequency"},
			Value: &object.String{Value: fmt.Sprintf("%.2f MHz", cpuInfo[0].Mhz)},
		}
	}

	if len(cpuPercent) > 0 {
		cpuInfoMap[(&object.String{Value: "usage"}).HashKey()] = object.DictPair{
			Key:   &object.String{Value: "usage"},
			Value: &object.String{Value: fmt.Sprintf("%.2f", cpuPercent[0])},
		}
	}

	return &object.Dict{Pairs: cpuInfoMap}
}

func getDiskInfo(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"sysinfo", "diskInfo",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`sysinfo.diskInfo() -> {"total": "100GB", "used": "50GB", "free": "50GB", "percent": "50.0"}`,
		)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get disk info: %v", err)}
	}

	diskInfoMap := make(map[object.HashKey]object.DictPair)

	diskInfoMap[(&object.String{Value: "total"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "total"},
		Value: &object.String{Value: fmt.Sprintf("%.2f GB", float64(diskStat.Total)/1024/1024/1024)},
	}
	diskInfoMap[(&object.String{Value: "used"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "used"},
		Value: &object.String{Value: fmt.Sprintf("%.2f GB", float64(diskStat.Used)/1024/1024/1024)},
	}
	diskInfoMap[(&object.String{Value: "free"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "free"},
		Value: &object.String{Value: fmt.Sprintf("%.2f GB", float64(diskStat.Free)/1024/1024/1024)},
	}
	diskInfoMap[(&object.String{Value: "percent"}).HashKey()] = object.DictPair{
		Key:   &object.String{Value: "percent"},
		Value: &object.String{Value: fmt.Sprintf("%.2f", diskStat.UsedPercent)},
	}

	return &object.Dict{Pairs: diskInfoMap}
}

func getNetInfo(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"sysinfo", "netInfo",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`sysinfo.netInfo() -> [{"name": "eth0", "addrs": ["192.168.1.100"]}]`,
		)
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get network info: %v", err)}
	}

	var networkInterfaces []object.VintObject

	for _, iface := range interfaces {
		if len(iface.Addrs) > 0 {
			ifaceMap := make(map[object.HashKey]object.DictPair)

			ifaceMap[(&object.String{Value: "name"}).HashKey()] = object.DictPair{
				Key:   &object.String{Value: "name"},
				Value: &object.String{Value: iface.Name},
			}

			var addrs []object.VintObject
			for _, addr := range iface.Addrs {
				addrs = append(addrs, &object.String{Value: addr.Addr})
			}

			ifaceMap[(&object.String{Value: "addrs"}).HashKey()] = object.DictPair{
				Key:   &object.String{Value: "addrs"},
				Value: &object.Array{Elements: addrs},
			}

			networkInterfaces = append(networkInterfaces, &object.Dict{Pairs: ifaceMap})
		}
	}

	return &object.Array{Elements: networkInterfaces}
}
