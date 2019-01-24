package plugin

/*
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/apache/trafficcontrol/grove/web"
	log "github.com/apache/trafficcontrol/lib/go-log"
	geoip2 "github.com/oschwald/geoip2-golang"
)

type GeoipConfig struct {
	Database string `json:"database"`
}

func init() {
	AddPlugin(10000, Funcs{onRequest: setGeoHeader, startup: geoipStartup, load: geoipLoad})
}

func setGeoHeader(icfg interface{}, d OnRequestData) bool {
	db := (*d.Context).(*geoip2.Reader)
	if d.R.Method == http.MethodOptions {
		return false
	}
	clientAddr, _ := web.GetClientIPPort(d.R)
	country, err := db.Country(net.ParseIP(clientAddr))
	if err == nil && country.Country.IsoCode != "" {
		d.R.Header.Set("X-Country", country.Country.IsoCode)
	}
	return false
}

func geoipStartup(icfg interface{}, d StartupData) {
	geoipConfig := icfg.(*GeoipConfig)
	db, err := geoip2.Open(geoipConfig.Database)
	if err != nil {
		log.Errorf("GeoIP startup error: %v", err.Error())
	} else {
		*d.Context = db
		log.Debugf("GeoIP startup success")
	}
}

func geoipLoad(b json.RawMessage) interface{} {
	geoipConfig := GeoipConfig{}
	err := json.Unmarshal(b, &geoipConfig)
	if err != nil {
		log.Errorln("Geoip config loading error: " + err.Error())
		return nil
	}
	log.Debugf("Geoip config loaded")
	return &geoipConfig
}
