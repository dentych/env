package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Tmux struct {
	printer *Printer
}

func (Tmux) Name() string {
	return "Tmux"
}

func (t Tmux) Install(conf Configuration) {
	t.printer.Print("Installing...")
	t.printer.Print("Adding .tmux.conf file to homedir")
	err := ioutil.WriteFile(t.filePath(), []byte(TmuxConf), 644)
	if err != nil {
		log.Fatalln("Tmux installation failed.", err)
	}
	t.printer.Print("Done!")
}

func (t Tmux) Uninstall(conf Configuration) {
	t.printer.Print("Uninstalling...")
	t.printer.Print("Removing .tmux.conf file to homedir")
	err := os.Remove(t.filePath())
	if err != nil {
		log.Fatalln("Tmux uninstallation error.", err)
	}
	t.printer.Print("Done!")
}

func (t *Tmux) SetPrinter(printer *Printer) {
	t.printer = printer
}

func (Tmux) filePath() string {
	return fmt.Sprintf("%s/%s", homedir, TmuxConfFileName)
}