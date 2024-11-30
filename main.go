package main

import "fmt"

func main() {
	fmt.Println("Start processing tasks...")

	//region //для случая обработки заданий одного типа (без использования интерфейсов)
	//// Инициализируем слайс из 20 задач для выполнения
	//tasks := make([]Task, 20)
	//// присваиваем всем задачам ID
	//for i := 0; i < len(tasks); i++ {
	//	tasks[i] = Task{ID: i + 1}
	//}
	//endregion

	//region Для случая с различными типами задач (с использованием интерфейсов)
	tasks := []Task{
		&EmailTask{Email: "email1@codeheim.io", Subject: "test", Message: "test 1"},
		&ImageProcessingTask{ImageUrl: "/images/sample1.jpg"},
		&EmailTask{Email: "email2@codeheim.io", Subject: "test", Message: "test 2"},
		&ImageProcessingTask{ImageUrl: "/images/sample2.jpg"},
		&EmailTask{Email: "email3@codeheim.io", Subject: "test", Message: "test 3"},
		&ImageProcessingTask{ImageUrl: "/images/sample3.jpg"},
		&EmailTask{Email: "email4@codeheim.io", Subject: "test", Message: "test 4"},
		&ImageProcessingTask{ImageUrl: "/images/sample4.jpg"},
		&EmailTask{Email: "email5@codeheim.io", Subject: "test", Message: "test 5"},
		&ImageProcessingTask{ImageUrl: "/images/sample5.jpg"},
	}
	//endregion

	// Создаем пул воркеров с заданиями и устанавливаем максимальное количество выполняемых заданий 5
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	// запускаем пул воркеров для обработки заданий
	wp.Run()
	fmt.Println("All tasks have been processed")
}
