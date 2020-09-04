package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	"golang.org/x/sys/windows/registry"
)

func main() {
	if len(os.Args) > 1 {
		/*
			DELETE
		*/
		path := strings.Join(os.Args[1:len(os.Args)], " ")
		fmt.Println("Deleting", path)
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
			err := beeep.Alert("Error eliminando", err.Error(), "")
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		/*
			INSTALL
		*/
		// All files
		k, _, err := registry.CreateKey(registry.CLASSES_ROOT, `*\shell\Eliminar permanente`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			fmt.Println(err)
		}

		if err := k.SetExpandStringValue("", "Eliminar permanente"); err != nil {
			fmt.Println(err)
		}
		if err := k.SetExpandStringValue("Icon", os.Args[0]); err != nil {
			fmt.Println(err)
		}
		if err := k.Close(); err != nil {
			fmt.Println(err)
		}

		// All files command
		command, _, err := registry.CreateKey(registry.CLASSES_ROOT, `*\shell\Eliminar permanente\command`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			fmt.Println(err)
		}
		if err := command.SetExpandStringValue("", os.Args[0]+" %1"); err != nil {
			fmt.Println(err)
		}
		if err := command.Close(); err != nil {
			fmt.Println(err)
		}

		// All folders
		k2, _, err := registry.CreateKey(registry.CLASSES_ROOT, `Directory\shell\Eliminar permanente`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			fmt.Println(err)
		}

		if err := k2.SetExpandStringValue("", "Eliminar permanente"); err != nil {
			fmt.Println(err)
		}
		if err := k2.SetExpandStringValue("Icon", os.Args[0]); err != nil {
			fmt.Println(err)
		}
		if err := k2.Close(); err != nil {
			fmt.Println(err)
		}

		// All folders command
		command2, _, err := registry.CreateKey(registry.CLASSES_ROOT, `Directory\shell\Eliminar permanente\command`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			fmt.Println(err)
		}
		if err := command2.SetExpandStringValue("", os.Args[0]+" %1"); err != nil {
			fmt.Println(err)
		}
		if err := command2.Close(); err != nil {
			fmt.Println(err)
		}
	}

	// Uncomment for debugging
	// fmt.Scanln()
}
