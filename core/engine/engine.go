package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
)

var (
	// Version 引擎版本号
	Version = "1.0.1"
	// EngineInfo 引擎信息
	EngineInfo = &struct {
		Version          *string
		StartTime        time.Time //启动时间
		EnableWaitStream *bool
		RingSize         *int
	}{
		&Version,
		time.Now(),
		&config.EnableWaitStream,
		&config.RingSize,
	}
	// config 配置信息
	config = &struct {
		EnableWaitStream bool
		EnableAudio      bool
		EnableVideo      bool
		RingSize         int
		PublishTimeout   time.Duration
	}{
		true,
		true,
		true,
		10,
		time.Minute,
	}
	// ConfigRaw 配置信息的原始数据
	ConfigRaw []byte

	// Gin框架控制引擎
	Gin *gin.Engine
)

func init() {
	//Gin初始化设置
	gin.SetMode(gin.ReleaseMode)
	Gin = gin.Default()
	Gin.Use(gzip.Gzip(gzip.DefaultCompression))
	Gin.Use(gin.Recovery())
	Gin.SetTrustedProxies(nil)
}

// Run启动Gin-Apps引擎
func Run(configFile string) (err error) {
	//Version
	Print(aurora.BgGreen(aurora.Black("Start Gin-Apps Engine @ v" + Version)))

	//停止脚本
	if runtime.GOOS == "windows" {
		ioutil.WriteFile("stop.bat", []byte(fmt.Sprintf("taskkill /pid %d  -t  -f", os.Getpid())), 0777)
	} else {
		ioutil.WriteFile("stop.sh", []byte(fmt.Sprintf("kill -9 %d", os.Getpid())), 0777)
	}

	//Config字典
	if ConfigRaw, err = ioutil.ReadFile(configFile); err != nil {
		Print(aurora.Red("read config file error:"), err)
		return
	}
	var cg map[string]interface{}
	if _, err = toml.Decode(string(ConfigRaw), &cg); err == nil {
		for name, config := range App {
			if cfg, ok := cg[name]; ok {
				b, _ := json.Marshal(cfg)
				if err = json.Unmarshal(b, config.Config); err != nil {
					log.Println(err)
					continue
				}
			} else if config.Config != nil {
				continue
			}
			if config.Run != nil {
				//执行run函数
				time.Sleep(300 * time.Microsecond)
				go config.Run()
			}
		}
	} else {
		Print(aurora.Red("decode config file error:"), err)
	}
	return
}
