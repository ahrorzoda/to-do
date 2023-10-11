package loggers

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	// Создаем файлы логирования (если они не существуют)
	if _, err := os.Stat("../pkg/logger/file/info.log"); os.IsNotExist(err) {
		os.Create("info.log")
	}
	if _, err := os.Stat("../pkg/logger/file/error.log"); os.IsNotExist(err) {
		os.Create("error.log")
	}

	// Открываем файлы логирования для записи
	infoFile, _ := os.OpenFile("../pkg/logger/file/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errorFile, _ := os.OpenFile("../pkg/logger/file/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// Устанавливаем вывод логов в соответствующие файлы
	Info = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
