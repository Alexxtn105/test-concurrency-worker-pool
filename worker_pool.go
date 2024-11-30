package main

import (
	"fmt"
	"sync"
	"time"
)

// Task Задача
type Task struct {
	ID int
}

// Process Метод, выполняемый задачей
func (t *Task) Process() {
	// симулируем длительную операцию
	fmt.Println("Processing task ", t.ID)
	time.Sleep(2 * time.Second)
}

// WorkerPool Пул воркеров
type WorkerPool struct {
	Tasks       []Task         // Список заданий
	concurrency int            // Количество одновременно выполняемых воркеров
	tasksChan   chan Task      // Канал, в котором задачи отправляются воркерам
	wg          sync.WaitGroup // WaitGroup для синхронизации завершения заданий, ожидает завершения заданий
}

// Метод, принимающий задание из канала taskChan и обрабатывающий его
func (wp *WorkerPool) worker() {
	// Принимаем задание из канала
	for task := range wp.tasksChan {
		// Вызываем метод Process задания для его обработки
		task.Process()

		// Сигнализируем в WaitGroup о завершении задания
		wp.wg.Done()
	}
}

// Run Метод инициализирует канал, устанавливает количество одновременно выполняемых задач (concurrency), создает горутины и отправляет задачи в канал
func (wp *WorkerPool) Run() {
	// инициализируем буферизованный канал длиной по количеству заданий
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// запускаем горутины воркеров (устанавливается в переменной concurrency)
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// добавляем счетчик WaitGroup по количеству заданий
	wp.wg.Add(len(wp.Tasks))

	// отправляем задания в канал
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}

	// закрываем канал после отправки ВСЕХ заданий для сигнализации о том, что более нельзя отправлять задачи
	close(wp.tasksChan)

	// ожидаем завершения работы всех задач
	wp.wg.Wait()
}
