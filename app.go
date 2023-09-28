package main

import (
	"context"
	"fmt"
	"hosteditor/backend/services"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx       context.Context
	hosts_ctx string
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置

	hosts, err := services.NewHosts()

	if err != nil {
		log.Fatal(err)
	}

	hosts_ctx, err := hosts.ReadHosts()

	if err != nil {
		log.Fatal(err)
	}

	a.hosts_ctx = hosts_ctx

	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {

	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Exit?",
		Message: "Are you sure you want to exit?",
	})
	if err != nil {
		return false
	}
	return dialog != "Yes"

}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
}

func (a *App) GetCtx() context.Context {
	return a.ctx
}

func (a *App) Alert(title string, content string) {
	dialog, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   title,
		Message: content,
	})
	print(dialog)
}

func (a *App) Confirm(title string, content string) string {
	dialog, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         title,
		Message:       content,
		DefaultButton: "No",
	})
	print(dialog)
	return dialog
}

func (a *App) RequestUrl(url string) string {
	// 发起 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		a.Alert("Error", fmt.Sprintf("Request failed: %s", err))
		return ""
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		a.Alert("Error", fmt.Sprintf("Failed to read response content: %s", err))
		return ""
	}

	// 打印响应内容
	fmt.Println(string(body))

	return string(body)
}
