package main

import (
	"github.com/mamemomonga/rpi-fan-control/configs"
	"log"
	"periph.io/x/periph/conn/gpio"
	"sync"
	"time"
)

type FanC struct {
	stop chan bool
	done chan bool
	w    int
	m    sync.Mutex
	fan  gpio.PinIO
	cfg  *configs.Configs
}

func NewFanC(cfg *configs.Configs, pin gpio.PinIO) *FanC {
	t := new(FanC)
	t.cfg = cfg
	t.fan = pin
	t.stop = make(chan bool)
	t.done = make(chan bool)
	return t
}

func (t *FanC) GetWait() int {
	t.m.Lock()
	v := t.w
	t.m.Unlock()
	return v
}

func (t *FanC) SetWait(v int) {
	t.m.Lock()
	t.w = v
	t.m.Unlock()
}

func (t *FanC) Start() {
	go func() {
		defer func() {
			close(t.done)
		}()
		for {
			t.loop()
			select {
			case <-t.stop:
				log.Println("Stop FanC")
				if err := t.fan.Out(gpio.Low); err != nil {
					log.Fatal(err)
				}
				return
			default:
			}
		}
	}()
}

func (t *FanC) Stop() {
	log.Println("Stopping FanC")
	close(t.stop)
	<-t.done
}

func (t *FanC) loop() {
	wa := t.GetWait()
	switch wa {
	case 0:
		if err := t.fan.Out(gpio.Low); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 500)
	case 255:
		if err := t.fan.Out(gpio.High); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 500)
	default:
		if err := t.fan.Out(gpio.High); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * time.Duration(wa))
		if err := t.fan.Out(gpio.Low); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * time.Duration(t.cfg.Configs.LowMillis))
	}
}
