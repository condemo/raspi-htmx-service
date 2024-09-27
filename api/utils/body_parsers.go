package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/condemo/raspi-htmx-service/config"
)

func ConfigParser(r *http.Request, conf *config.UserConfig) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	t := r.FormValue("theme")
	tick, err := time.ParseDuration(r.FormValue("info-tick"))
	if err != nil {
		return err
	}

	conf.GeneralConf.CurrentTheme = config.Theme(t)
	conf.InfoConf.InfoTick = tick

	fmt.Printf("Final Struct -> %+v\n", conf)
	return nil
}
