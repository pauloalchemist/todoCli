package main

import (
	"flag"
	"fmt"
	"go-cap2/todoCli/todo"
	"os"
)

var todoFileName = ".todo.json"

func main() {

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	add := flag.Bool("add", false, "Tarefa a ser inclusa no ToDo Cli")
	list := flag.Bool("list", false, "Listar todas as tarefas")
	complete := flag.Int("complete", 0, "Alterar tarefa para concluida")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list: // lista a tarefas a fazer
		fmt.Print(l)

	case *complete > 0:
		// ↓ altera tarefa pata finalizada (complete)
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// ↓ salva no item na lista de tarefas
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		// ↓ adiciona nova tarefa
		t, err := todo.GetTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(t)

		// ↓ salva na lista
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		// ↓ se uma flag inválida for passada
		fmt.Fprintln(os.Stderr, "Opção inválida. Tente novamente.")
		os.Exit(1)
	}
}
