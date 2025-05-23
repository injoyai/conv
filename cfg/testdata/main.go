package main

import (
	"github.com/injoyai/conv/cfg"
	"log"
)

// -a=101 -aa=101
func main() {
	cfg.Init(
		cfg.WithFlag(&cfg.Flag{Name: "a"}, &cfg.Flag{Name: "aa"}),
		cfg.WithFile("./testdata/config.yaml"),
	)
	log.Println(cfg.GetInt("child[0].time"))
	log.Println(cfg.GetDuration("child[0].time"))
	log.Println(cfg.GetSecond("child[1].time"))
	log.Println(cfg.GetInt("a"))
	log.Println(cfg.GetInt("aa"))
	log.Println(cfg.GetInt("b", 103))
	log.Println(cfg.GetInt("text.a"))
	log.Println(cfg.GetString("text.b"))
}
