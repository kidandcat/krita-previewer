package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/sys/windows/registry"
)

func main() {
	if len(os.Args) > 1 {
		path := strings.Join(os.Args[1:len(os.Args)], " ")

		fmt.Println("unzipping")
		os.RemoveAll(os.TempDir() + "/krita_preview")
		unzip(path, os.TempDir()+"/krita_preview")

		toOpen := filepath.Join(os.TempDir(), "/krita_preview/mergedimage.png")
		fmt.Println("opening " + toOpen)
		open.Run(toOpen)
	} else {
		/*
			INSTALL
		*/
		// All files
		k, _, err := registry.CreateKey(registry.CLASSES_ROOT, `*\shell\Krita preview`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			fmt.Println(err)
		}

		if err := k.SetExpandStringValue("", "Krita preview"); err != nil {
			fmt.Println(err)
		}
		if err := k.SetExpandStringValue("Icon", os.Args[0]); err != nil {
			fmt.Println(err)
		}
		if err := k.Close(); err != nil {
			fmt.Println(err)
		}

		// All files command
		command, _, err := registry.CreateKey(registry.CLASSES_ROOT, `*\shell\Krita preview\command`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			fmt.Println(err)
		}
		if err := command.SetExpandStringValue("", os.Args[0]+" %1"); err != nil {
			fmt.Println(err)
		}
		if err := command.Close(); err != nil {
			fmt.Println(err)
		}
	}

	beeep.Notify("Krita preview installed", "Run it again if you move its location!", "")

	// Uncomment for debugging
	// fmt.Scanln()
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		if f.Name != "mergedimage.png" {
			continue
		}
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
