package util

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/ACking-you/byte_douyin_project/config"
	"log"
)

// 可以更改
var (
	globalMutex        sync.RWMutex
	defaultVideoSuffix = ".mp4"
	defaultImageSuffix = ".jpg"
)

type Video2Image struct {
	inputPath  string
	outputPath string
	startTime  string
	keepTime   string
	filter     string
	frameCount int
	debug      bool
}

func NewVideo2Image() *Video2Image {
	return &Video2Image{} // 每次返回新实例，避免共享状态
}

func GetDefaultVideoSuffix() string {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	return defaultVideoSuffix
}

func GetDefaultImageSuffix() string {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	return defaultImageSuffix
}

// ChangeVideoDefaultSuffix 全局配置修改方法（线程安全）
func ChangeVideoDefaultSuffix(suffix string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	defaultVideoSuffix = normalizeExtension(suffix)
}

func ChangeImageDefaultSuffix(suffix string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	defaultImageSuffix = normalizeExtension(suffix)
}

func normalizeExtension(ext string) string {
	if ext == "" {
		return ext
	}
	if ext[0] != '.' {
		return "." + ext
	}
	return ext
}

// SetInputPath 方法链式调用（非共享实例，无需加锁）
func (v *Video2Image) SetInputPath(path string) *Video2Image {
	v.inputPath = path
	return v
}

func (v *Video2Image) SetOutputPath(path string) *Video2Image {
	v.outputPath = path
	return v
}

func (v *Video2Image) SetTimeOptions(start, duration string) *Video2Image {
	v.startTime = start
	v.keepTime = duration
	return v
}

func (v *Video2Image) SetFilter(filter string) *Video2Image {
	v.filter = filter
	return v
}

func (v *Video2Image) SetFrameCount(count int) *Video2Image {
	v.frameCount = count
	return v
}

func (v *Video2Image) SetDebug(debug bool) *Video2Image {
	v.debug = debug
	return v
}

func (v *Video2Image) buildArgs() ([]string, error) {
	if v.inputPath == "" || v.outputPath == "" {
		return nil, errors.New("input and output path must be specified")
	}

	args := []string{
		"-i", filepath.ToSlash(v.inputPath),
		"-f", "image2",
	}

	if v.filter != "" {
		args = append(args, "-vf", v.filter)
	}
	if v.startTime != "" {
		args = append(args, "-ss", v.startTime)
	}
	if v.keepTime != "" {
		args = append(args, "-t", v.keepTime)
	}
	if v.frameCount > 0 {
		args = append(args, "-frames:v", strconv.Itoa(v.frameCount))
	}

	args = append(args, "-y", filepath.ToSlash(v.outputPath))
	return args, nil
}

func (v *Video2Image) Execute() error {
	args, err := v.buildArgs()
	if err != nil {
		return fmt.Errorf("参数构建失败: %w", err)
	}

	ffmpegPath := filepath.FromSlash(config.Global.FfmpegPath)

	cmd := exec.Command(ffmpegPath, args...)
	if v.debug {
		log.Printf("执行命令: %q", cmd.String())
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg执行失败: %w (命令: %q)", err, cmd.String())
	}
	return nil
}
