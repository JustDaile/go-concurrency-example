package main

import (
  "log"
  "fmt"
  "time"
  "sync"
)

type Task struct {
	Name string
}

func (t Task) Run() {
	log.Printf("Started task %s\n", t.Name)
	time.Sleep(time.Second * 5)
	log.Printf("%s finished\n", t.Name)
	wg.Done()
}

const tasksToAdd = 5
var wg = sync.WaitGroup{}

func main(){
	tasks := []Task{}
	for i := 0; i < tasksToAdd; i++ {
		tasks = append(tasks, Task{
			Name: fmt.Sprintf("Tasks_%v",i),
		})
	}

	for _,t := range tasks {
		wg.Add(1)
		go t.Run()
	}
	wg.Wait()
	log.Println("all tasks done")
}
