package main

import (
	"flag"
	"fmt"
	todo "go-cap2/todoCli"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	task := flag.String("task", "", "Tarefa a ser inclusa no ToDo List")
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

	case *task != "":
		// ↓ adiciona nova tarefa
		l.Add(*task)

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
