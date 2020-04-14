package configs

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"regexp"
	//	"github.com/davecgh/go-spew/spew"
)

type Configs struct {
	Configs C
}

func New() (t *Configs) {
	t = new(Configs)
	t.Configs = C{}
	return t
}

func (t *Configs) Load(configFile string) error {
	if configFile == "" {
		log.Println()
		return errors.New("alert: 設定ファイルがありません")
	}

	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	s := regexp.MustCompile(`\r\n|\r|\n`).ReplaceAllString(string(buf), "\n")
	err = yaml.Unmarshal([]byte(s), &t.Configs)

	if err != nil {
		return err
	}

	return nil
}
