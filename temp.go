package main

import (
	"github.com/mamemomonga/rpi-fan-control/configs"
	"log"
	"os/exec"
	"strconv"
	"time"
)

type TempC struct {
	stop       chan bool
	done       chan bool
	fan        *FanC
	cfg        *configs.Configs
	hysteresis int
}

func NewTempC(cfg *configs.Configs, fan *FanC) *TempC {
	t := new(TempC)
	t.cfg = cfg
	t.fan = fan
	t.stop = make(chan bool)
	t.done = make(chan bool)
	t.hysteresis = 0
	return t
}

func (t *TempC) Start() {
	go func() {
		defer func() {
			close(t.done)
		}()
		for {
			t.loop()
			select {
			case <-t.stop:
				log.Println("Stop TempC")
				return
			default:
			}
		}
	}()
}

func (t *TempC) Stop() {
	log.Println("Stopping TempC")
	close(t.stop)
	<-t.done
}

func (t *TempC) loop() {
	temp := t.readTemp()
	wait := t.fan.GetWait()

	switch {
	case t.hysteresis == 1:
		if temp <= float64(t.cfg.Configs.Hysteresis.High) {
			t.hysteresis = 0
		}
	case t.hysteresis == -1:
		if temp >= float64(t.cfg.Configs.Hysteresis.Low) {
			t.hysteresis = 0
		}
	default:
		for _, i := range t.cfg.Configs.FanControls {
			if temp >= float64(i.Temp) {
				switch i.High {
				case 255:
					t.hysteresis = 1
					t.fan.SetWait(255)
				case 0:
					t.hysteresis = -1
					t.fan.SetWait(0)
				default:
					t.fan.SetWait(i.High)
				}
				break
			}
		}
	}

	wait = t.fan.GetWait()
	switch wait {
	case 255:
		log.Printf("info: Temp: %2.1f C / Full \n", temp)

	case 0:
		log.Printf("info: Temp: %2.1f C / Stop \n", temp)

	default:
		log.Printf("info: Temp: %2.1f C / Wait: %d \n", temp, wait)
	}
	time.Sleep(time.Second * 5)
}

func (t *TempC) readTemp() float64 {
	out, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		log.Fatal(err)
	}
	buf := string(out)
	v, err := strconv.ParseFloat(buf[5:9], 32)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
