package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 1
var filepath string

func main() {
	a := app.New()
	btn := 1
	_ = btn
	w := a.NewWindow("Text Editor")

	w.Resize(fyne.NewSize(400, 400))

	///content aline vertially

	input := widget.NewMultiLineEntry()   //NewMultiLineEntry creates a new entry that allows multiple lines
	input.SetPlaceHolder("Enter text...") //SetPlaceHolder sets the text that will be displayed if the entry is otherwise empty
	input.Move(fyne.NewPos(0, 0))         //Move the widget to a new position,
	input.Resize(fyne.NewSize(500, 500))

	new1 := fyne.NewMenuItem("New", func() { //NewMenuItem creates a new menu item from the passed label and action parameters.
		filepath = ""

		input.Text = ""
		input.Refresh()
	})
	save1 := fyne.NewMenuItem("Save", func() {
		if filepath != "" {
			f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0)
			//Flags to OpenFile wrapping those of the underlying system. Not all flags may be implemented on a given system.
			//
			if err != nil {
				//handle error
			}
			defer f.Close()
			//Close closes the File, rendering it unusable for I/O. On files that support SetDeadline, any pending I/O operations will be canceled and return immediately with an error.
			//f.Write([]byte(input.Text))
			f.WriteString(input.Text)
			//WriteString is like Write, but writes the contents of string s rather than a slice of bytes.
		} else {
			saveFileDialog := dialog.NewFileSave(
				func(r fyne.URIWriteCloser, _ error) {
					textData := []byte(input.Text)
					r.Write(textData)
					filepath = r.URI().Path()
					w.SetTitle(filepath)
				}, w)
			saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
			saveFileDialog.Show()
		}
	})
	saveAs1 := fyne.NewMenuItem("Save as..", func() {
		saveFileDialog := dialog.NewFileSave(
			func(r fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				r.Write(textData)
				filepath = r.URI().Path()
				w.SetTitle(filepath)
			}, w)
		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		saveFileDialog.Show()
	})
	open1 := fyne.NewMenuItem("Open", func() {
		openfileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(r)
				result := fyne.NewStaticResource("name", data)
				input.SetText(string(result.StaticContent))
				fmt.Println(result.StaticName + r.URI().Path())
				filepath = r.URI().Path()
				w.SetTitle(filepath)
			}, w)
		openfileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		openfileDialog.Show()
	})

	saveBtn := widget.NewButton("Save file", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			}, w)

		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		saveFileDialog.Show()
	})

	menuItem := fyne.NewMenu("File", new1, save1, saveAs1, open1)
	menux1 := fyne.NewMainMenu(menuItem)
	w.SetMainMenu(menux1)

	w.SetContent(
		container.NewVBox(input,
			container.NewHBox(saveBtn),
		),
	)

	w.ShowAndRun()
}
