package configs

type C struct {
	FanGPIOPin       string        `yaml:"fan_gpio_pin"`
	FanControls      []CFanControl `yaml:"fan_control"`
	TempCheckSeconds int           `yaml:"temp_check_seconds"`
	LowMillis        int           `yaml:"low_millis"`
	Hysteresis       CHysteresis   `yaml:"hysteresis"`
}

type CFanControl struct {
	Temp int `yaml:"temp"`
	High int `yaml:"high"`
}

type CHysteresis struct {
	High int `yaml:"high"`
	Low  int `yaml:"low"`
}
