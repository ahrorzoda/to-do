package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Создаем канал для приема сигналов
	sigs := make(chan os.Signal, 1)

	// Передаем сигналы интересующего нас типа (syscall.SIGINT) в канал
	signal.Notify(sigs, syscall.SIGINT)

	// Ожидаем сигнал
	fmt.Println("Ожидание сигнала Ctrl+C (SIGINT). Нажмите Ctrl+C для завершения...")

	// Блокируем выполнение программы до получения сигнала
	<-sigs

	// Здесь можно добавить ваш код обработки сигнала Ctrl+C
	fmt.Println("Получен сигнал Ctrl+C. Завершение программы.")
}
