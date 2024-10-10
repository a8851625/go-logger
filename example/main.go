package main

import (
	"gitea.heheda.vip/root/go-logger"
)

func main() {
	// 使用预设的 logger
	logger.Debug.Println("This is a debug message")
	logger.Info.Println("This is an info message")
	logger.Warning.Println("This is a warning message")
	logger.Error.Println("This is an error message")

	// 使用格式化输出
	logger.Info.Printf("Formatted message: %s", "Hello, World!")

	// 更改输出位置（例如，输出到文件）
	// f, _ := os.OpenFile("logfile.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// logger.Error.SetOutput(f)
	// logger.Error.Println("This error message will be written to the file")

	// 获取 logger 级别
	level := logger.Debug.GetLevel()
	logger.Info.Printf("Debug logger level: %v", level)

	// 注意：下面的代码会导致程序退出
	// logger.Error.Fatalln("This is a fatal error message")
}
