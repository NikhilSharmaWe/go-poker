package main

import (
	"github.com/NikhilSharmaWe/go-poker/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:    "GGPOKER V0.1-alpha",
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	remoteCfg := p2p.ServerConfig{
		Version:    "GGPOKER V0.1-alpha",
		ListenAddr: ":4000",
	}
	remoteServer := p2p.NewServer(remoteCfg)
	go remoteServer.Start()
	if err := remoteServer.Connect("localhost:3000"); err != nil {
		panic(err)
	}

	select {}
}
