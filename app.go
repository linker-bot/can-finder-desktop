package main

import (
	"context"
	"encoding/json"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DeviceInfo struct {
	Name    string `json:"name"`
	IP      string `json:"ip"`
	MAC     string `json:"mac"`
	Model   string `json:"model"`
	Version string `json:"version"`
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.listenUDP()
}

func (a *App) listenUDP() {
	addr := net.UDPAddr{
		Port: 9999,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		runtime.LogError(a.ctx, "UDP listener error: "+err.Error())
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			runtime.LogError(a.ctx, "UDP read error: "+err.Error())
			continue
		}

		var device DeviceInfo
		if err := json.Unmarshal(buffer[:n], &device); err != nil {
			runtime.LogError(a.ctx, "JSON parse error: "+err.Error())
			continue
		}

		runtime.EventsEmit(a.ctx, "device-update", device)
	}
}
