package main

import (
	"flag"
	"github.com/mamemomonga/rpi-fan-control/configs"
	"log"
	"os"
	"os/signal"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	"syscall"
	//	"github.com/davecgh/go-spew/spew"
)

func main() {

	flg_config := flag.String("config","","Config File")
	flag.Parse()

	cfg := configs.New()
	cfg.Load(*flg_config)

	log.Println("info: Start")
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT)

	log.Println("info: Fan GPIO Pin:" + cfg.Configs.FanGPIOPin)
	fan := NewFanC(cfg, gpioreg.ByName(cfg.Configs.FanGPIOPin))
	temp := NewTempC(cfg, fan)

	fan.Start()
	temp.Start()

	<-quit
	log.Println("Terminate.")

	temp.Stop()
	fan.Stop()
}
