package config

import (
	"fmt"
	//"github.com/go-yaml/yaml"
	"os"
	"path/filepath"
)

type DBConfig struct {

}

func NewConfig() error {
	configPath,err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err !=nil {
		return err
	}
	////读取配置文件
	//ioutil.ReadFile()

	//yaml.Unmarshal()
	fmt.Println(configPath)
	fmt.Println(os.Args[0])
	return nil
}
