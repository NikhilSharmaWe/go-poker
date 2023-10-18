package main

import "github.com/NikhilSharmaWe/go-poker/p2p"

func main() {
	// rand.Seed(time.Now().UnixNano())
	// d := deck.New()
	// fmt.Println(d)

	cfg := p2p.ServerConfig{
		ListenAddr: ":3000",
	}

	s := p2p.NewServer(cfg)

	s.Start()
}
