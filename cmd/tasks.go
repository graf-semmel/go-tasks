package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Done        bool
	Description string
	Created     string
}

func GetAll() []Task {
	tasks := []Task{}
	records := ReadCSV()
	for i, record := range records {
		// header
		if i == 0 {
			continue
		}
		task := Task{
			Done:        record[1] == "true",
			Description: record[2],
			Created:     record[3],
		}
		if id, err := strconv.Atoi(record[0]); err != nil {
			log.Fatal("Error converting ID to int: ", err)
		} else {
			task.ID = id
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func GetPending() []Task {
	tasks := GetAll()
	openTasks := []Task{}
	for _, task := range tasks {
		if !task.Done {
			openTasks = append(openTasks, task)
		}
	}
	return openTasks
}

func Complete(id int) *Task {
	tasks := GetAll()
	var task *Task

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			task = &tasks[i]
			break
		}
	}

	if task == nil {
		log.Fatalf("Task with ID %d not found", id)
		return nil
	}

	WriteCSV(tasks)
	return task
}

func Add(desc string) {
	tasks := GetAll()

	task := Task{
		Description: desc,
		Done:        false,
		Created:     time.Now().Format(time.RFC3339),
	}

	if len(tasks) == 0 {
		task.ID = 1
	} else {
		// find max ID
		maxId := 0
		for _, t := range tasks {
			if t.ID > maxId {
				maxId = t.ID
			}
		}
		task.ID = maxId + 1
	}

	WriteCSV(append(tasks, task))
}

func Delete(id int) {
	tasks := GetAll()
	deleted := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			deleted = true
			break
		}
	}

	if !deleted {
		fmt.Printf("Task with ID %v not found\n", id)
		return
	}

	WriteCSV(tasks)
}
