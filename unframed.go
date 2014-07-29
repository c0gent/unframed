package unframed

import (
	"encoding/json"
	"github.com/nsan1129/unframed/log"
	"io/ioutil"
	"strconv"
	"os"
)

func Atoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Error("unframed.Atoi; string:", s, "int:", i, "][", err)
	}
	return
}

func Itoa(i int) (s string) {
	s = strconv.Itoa(i)
	return
}

type Config struct {
	Wd,
	DbType,
	ConnStr,
	ListenPort string
}

func WriteConfig(cfg *Config, cfgFile string) {

	file, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(cfgFile, file, 0644); err != nil {
		panic(err)
	}
}

func ReadConfig(cfgFile string) (cfg *Config) {
	file, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		panic("Opening Config Failed *** " + cfgFile + " *** \n" + err.Error())
	}

	cfg = new(Config)
	if err = json.Unmarshal(file, cfg); err != nil {
		panic("Parsing Config Failed *** " + cfgFile + " *** \n" + err.Error())
	}

	os.Chdir(cfg.Wd)
	return
}