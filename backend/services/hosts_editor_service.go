package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	R "runtime"
	"strings"
	"time"
)

// hostsService 结构体，用于处理hostsService文件操作
type hostsService struct {
	HostsPath string // Hosts文件的路径
}

// 获取Hosts文件路径
func GetHostsFilePath() string {
	var hostsFilePath string

	// 根据操作系统设置Hosts文件路径
	switch os := R.GOOS; os {
	case "windows":
		// Windows的Hosts文件路径
		hostsFilePath = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	case "linux":
		// Linux的Hosts文件路径
		hostsFilePath = "/etc/hosts"
	case "darwin":
		// macOS的Hosts文件路径
		hostsFilePath = "/etc/hosts"
	default:
		return ""
	}
	return hostsFilePath
}

// NewHosts 创建一个Hosts对象
func NewHosts() (*hostsService, error) {
	filePath := GetHostsFilePath()
	if filePath == "" {
		return nil, errors.New("Unsupported operating system")
	}

	return &hostsService{
		HostsPath: filePath,
	}, nil
}

// 读取Hosts文件的内容
func (h *hostsService) ReadHosts() (string, error) {
	content, err := ioutil.ReadFile(h.HostsPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 编辑Hosts文件
func (h *hostsService) EditHosts(newContent string) string {
	// 读取原文件内容
	originalContent, err := ioutil.ReadFile(h.HostsPath)
	if err != nil {
		return fmt.Sprintf("fail to read original file: %s", err)
	}

	// 获取当前时间并格式化为 yyyymmddhhmmss 格式
	currentTime := time.Now()
	timestamp := currentTime.Format("20060102150405")

	// 构建备份文件名，加上时间戳
	backupFileName := h.HostsPath + "." + timestamp + ".bak"

	// 备份原文件内容到备份文件
	err = ioutil.WriteFile(backupFileName, originalContent, 0644)
	if err != nil {
		return fmt.Sprintf("fail to backup original file: %s", err)
	}

	// 写入新内容到原文件
	err = ioutil.WriteFile(h.HostsPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Sprintf("fail to write new content: %s", err)
	}

	return "Ok"
}

func (h *hostsService) GetHostsBackupFilePaths() ([]string, error) {
	hostDirectory := strings.Replace(h.HostsPath, "hosts", "", 1)
	// 读取目录中的所有文件
	files, err := ioutil.ReadDir(hostDirectory)
	if err != nil {
		return nil, fmt.Errorf("fail to read directory: %s", err)
	}

	// 存储符合条件的备份文件路径的切片
	backupFilePaths := []string{}

	// 遍历目录中的文件
	for _, file := range files {
		// 检查文件名是否以 "hosts" 开头且以 ".bak" 结尾
		if strings.HasPrefix(file.Name(), "hosts") && strings.HasSuffix(file.Name(), ".bak") {
			// 构建备份文件的完整路径
			fullPath := filepath.Join(hostDirectory, file.Name())
			backupFilePaths = append(backupFilePaths, fullPath)
		}
	}

	return backupFilePaths, nil
}
