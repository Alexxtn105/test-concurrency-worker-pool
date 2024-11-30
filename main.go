package main

import "fmt"

func main() {
	fmt.Println("Start processing tasks...")

	// Инициализируем слайс из 20 задач для выполнения
	tasks := make([]Task, 20)
	// присваиваем всем задачам ID

	for i := 0; i < len(tasks); i++ {
		tasks[i] = Task{ID: i + 1}
	}

	// Создаем пул воркеров с заданиями и устанавливаем максимальное количество выполняемых заданий 5
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	// запускаем пул воркеров для обработки заданий
	wp.Run()
	fmt.Println("All tasks have been processed")
}
