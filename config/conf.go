package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

type Server struct {
	IP   string
	Port int
}

type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

type Config struct {
	DB     Mysql `toml:"mysql"`
	RDB    Redis `toml:"redis"`
	Server `toml:"server"`
	Path   `toml:"path"`
}

var Global Config

func ensurePathValid() {
	var err error
	if _, err = os.Stat(Global.StaticSourcePath); os.IsNotExist(err) {
		if err = os.Mkdir(Global.StaticSourcePath, 0755); err != nil {
			log.Fatalf("mkdir error:path %s", Global.StaticSourcePath)
		}
	}
	if _, err = os.Stat(Global.FfmpegPath); os.IsNotExist(err) {
		if _, err = exec.Command("ffmpeg", "-version").Output(); err != nil {
			log.Fatalf("ffmpeg not valid %s", Global.FfmpegPath)
		} else {
			Global.FfmpegPath = "ffmpeg"
		}
	} else {
		Global.FfmpegPath, err = filepath.Abs(Global.FfmpegPath)
		if err != nil {
			log.Fatalln("get abs path failed:", Global.FfmpegPath)
		}
	}
	//把资源路径转化为绝对路径，防止调用ffmpeg命令失效
	Global.StaticSourcePath, err = filepath.Abs(Global.StaticSourcePath)
	if err != nil {
		log.Fatalln("get abs path failed:", Global.StaticSourcePath)
	}
}

//包初始化加载时候会调用的函数
func init() {
	if _, err := toml.DecodeFile("./config/config.toml", &Global); err != nil {
		panic(err)
	}
	//去除左右的空格
	strings.Trim(Global.Server.IP, " ")
	strings.Trim(Global.RDB.IP, " ")
	strings.Trim(Global.DB.Host, " ")
	//保证路径正常
	ensurePathValid()
}

// DBConnectString 填充得到数据库连接字符串
func DBConnectString() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Global.DB.Username, Global.DB.Password, Global.DB.Host, Global.DB.Port, Global.DB.Database,
		Global.DB.Charset, Global.DB.ParseTime, Global.DB.Loc)
	log.Println(arg)
	return arg
}
