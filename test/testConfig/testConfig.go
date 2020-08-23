package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// 解析后面的名字一定要正确！！！
type Config struct {
	Labels Labels `yaml:"labels"`
}

type Labels struct{ // 转成数组不可以！！！
	Label1 Label `yaml:"label1"`
	Label2 Label `yaml:"label2"`
}

type Label struct{
	Label11 string `yaml:"label11"`
	Label12 string `yaml:"label12"`
}


//read yaml config
//注：path为yaml或yml文件的路径
func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		_ = yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}

//test yaml
func main() {


	conf, err := ReadYamlConfig("./test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf.Labels.Label1.Label11)

	byts, err := json.Marshal(conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byts))

}
