package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

func main() {
	if len(os.Args) > 1 {
		/*
			DELETE
		*/
		fmt.Println("Deleting")
		err := os.RemoveAll(os.Args[1])
		if err != nil {
			fmt.Println(err)
			fmt.Scanln()
		}
	} else {
		/*
			INSTALL
		*/
		// All files
		k, _, err := registry.CreateKey(registry.CLASSES_ROOT, `*\shell\Eliminar permanente`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			panic(err)
		}

		if err := k.SetExpandStringValue("", "Eliminar permanente"); err != nil {
			panic(err)
		}
		if err := k.SetExpandStringValue("Icon", os.Args[0]); err != nil {
			panic(err)
		}
		if err := k.Close(); err != nil {
			panic(err)
		}

		// All files command
		command, _, err := registry.CreateKey(registry.CLASSES_ROOT, `*\shell\Eliminar permanente\command`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			panic(err)
		}
		if err := command.SetExpandStringValue("", os.Args[0]+" %1"); err != nil {
			panic(err)
		}
		if err := command.Close(); err != nil {
			panic(err)
		}

		// All folders
		k2, _, err := registry.CreateKey(registry.CLASSES_ROOT, `Directory\shell\Eliminar permanente`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			panic(err)
		}

		if err := k2.SetExpandStringValue("", "Eliminar permanente"); err != nil {
			panic(err)
		}
		if err := k2.SetExpandStringValue("Icon", os.Args[0]); err != nil {
			panic(err)
		}
		if err := k2.Close(); err != nil {
			panic(err)
		}

		// All folders command
		command2, _, err := registry.CreateKey(registry.CLASSES_ROOT, `Directory\shell\Eliminar permanente\command`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			panic(err)
		}
		if err := command2.SetExpandStringValue("", os.Args[0]+" %1"); err != nil {
			panic(err)
		}
		if err := command2.Close(); err != nil {
			panic(err)
		}
	}
}
