package utils

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"runtime"
	"time"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type ServerInfo struct {
	Os   `json:"os"`
	Cpu  `json:"cpu"`
	Ram  `json:"ram"`
	Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
	GoRoot       string `json:"goRoot"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	Used        int     `json:"used"`
	Total       int     `json:"total"`
	UsedPercent float64 `json:"usedPercent"`
}

type Disk struct {
	Used        int     `json:"used"`
	Total       int     `json:"total"`
	UsedPercent float64 `json:"usedPercent"`
}

func (s *ServerInfo) InitOS() {
	s.GOOS = runtime.GOOS
	s.NumCPU = runtime.NumCPU()
	s.Compiler = runtime.Compiler
	s.GoVersion = runtime.Version()
	s.NumGoroutine = runtime.NumGoroutine()
	s.GoRoot = runtime.GOROOT()
}

func (s *ServerInfo) InitCpu() error {
	percent, err := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	if err != nil {
		return err
	}
	s.Cpus = percent
	counts, err := cpu.Counts(false)
	if err != nil {
		return err
	}
	s.Cores = counts
	return nil
}

func (s *ServerInfo) InitRam() error {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return err
	}
	s.Ram.Used = int(memory.Used)
	s.Ram.Total = int(memory.Total)
	s.Ram.UsedPercent = memory.UsedPercent
	return nil
}

func (s *ServerInfo) InitDisk() error {
	usage, err := disk.Usage("/")
	if err != nil {
		return err
	}
	s.Disk.Used = int(usage.Used) / GB
	s.Disk.Total = int(usage.Total) / GB
	s.Disk.UsedPercent = usage.UsedPercent
	return nil
}
