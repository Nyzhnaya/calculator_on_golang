package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Operand1 int
	Action   string
	Operand2 int
}

func main() {
	fmt.Println("Please enter 4 tasks: ")
	var a, c, d, f, g, i, j, l int
	var b, e, h, k string
	_, err := fmt.Scanf("%d %s %d %d %s %d %d %s %d %d %s %d", &a, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k, &l)
	if err != nil {
		fmt.Println(err)
	}

	startTime := time.Now()

	task1 := Task{
		Operand1: a,
		Action:   b,
		Operand2: c,
	}

	task2 := Task{
		Operand1: d,
		Action:   e,
		Operand2: f,
	}

	task3 := Task{
		Operand1: g,
		Action:   h,
		Operand2: i,
	}

	task4 := Task{
		Operand1: j,
		Action:   k,
		Operand2: l,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		task1.Answer(task1)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		task2.Answer(task2)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		task3.Answer(task3)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		task4.Answer(task4)
		wg.Done()
	}()

	wg.Wait()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Finished execution. Time: %v\n", elapsedTime)
}

func (t *Task) Answer(task Task) {

	var result int

	switch task.Action {
	case "+":
		result = (task.Operand1 + task.Operand2)
		fmt.Printf("%d + %d = %d\n", task.Operand1, task.Operand2, result)
	case "-":
		result = task.Operand1 - task.Operand2
		fmt.Printf("%d - %d = %d\n", task.Operand1, task.Operand2, result)
	case "*":
		result = task.Operand1 * task.Operand2
		fmt.Printf("%d * %d = %d\n", task.Operand1, task.Operand2, result)
	case "/":
		result = task.Operand1 / task.Operand2
		fmt.Printf("%d / %d = %d\n", task.Operand1, task.Operand2, result)
	default:
		fmt.Println("The data is not correct. Please enter task: ")
	}
}
