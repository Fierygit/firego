/*
 * @Author: Firefly
 * @Date: 2020-10-15 14:45:14
 * @Descripttion:
 * @LastEditTime: 2020-10-16 12:07:32
 */

package proxy

import (
	"firego/src/common/util"
	"flag"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var proxyConfig = make(map[string]string)

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/proxy")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	for k, v := range viper.AllSettings() {
		proxyConfig[k] = strconv.Itoa(v.(int))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	logrus.Info("proxy for " + r.URL.Path)
	prifix := "http://127.0.0.1:"
	truePort := ""

	for k, v := range proxyConfig {
		i := 0
		for ; i < util.Min(len(k), len(r.URL.Path)); i++ {
			if k[i] != r.URL.Path[i] {
				break
			}
		}
		if i == len(k) {
			truePort = v
			break
		}
	}

	if truePort == "" {
		if proxyConfig["default"] == "" {
			io.WriteString(w, "error\n")
			return
		}
		truePort = proxyConfig["default"]
	}

	url, err := url.Parse(prifix + truePort)
	if err != nil {
		io.WriteString(w, url.Path+"error\n")
		log.Println(err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

// Run start the proxy
func Run() {

	initConfig()
	var port int
	flag.IntVar(&port, "proxyPort", 80, "the proxy port")
	flag.Parse()
	logrus.Info("start to proxy 80 port" + " : " + strconv.Itoa(port))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	os.Exit(0)

}
