package utils

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"runtime"
	"time"
)

type ServerInfo struct {
	Os  `json:"os"`
	Cpu `json:"cpu"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

func (s *ServerInfo) InitOS() {
	s.GOOS = runtime.GOOS
	s.NumCPU = runtime.NumCPU()
	s.Compiler = runtime.Compiler
	s.GoVersion = runtime.Version()
	s.NumGoroutine = runtime.NumGoroutine()
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
