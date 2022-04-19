package main

import (
	"fmt"
	"image/color"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/main-moonrain/analizador_de_sintaxis/src/mypackage"
)

func main() {
	a := app.New()
	w := a.NewWindow("Apertura de archivo")

	w.Resize(fyne.NewSize(500, 400))

	text := canvas.NewText("Seleccione un archivo", color.White)
	text.Alignment = fyne.TextAlignLeading
	text.TextStyle = fyne.TextStyle{Italic: true}

	btn := widget.NewButton("Open .txt files", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(r)
				result := fyne.NewStaticResource("name", data)
				errores, dataContent := mypackage.Analizador(data)
				entry := widget.NewMultiLineEntry()
				if errores == "" {
					errores = "No se encontraron errores"
				}
				entry.SetText(fmt.Sprintf("%s\n\n----ANÁLISIS---\n%s", string(result.StaticContent), errores))

				w := fyne.CurrentApp().NewWindow(
					"Texto")
				w.Resize(fyne.NewSize(500, 400))
				w.SetContent(container.NewScroll(entry))
				w.Show()

				table := widget.NewTable(
					func() (int, int) { return len(dataContent), len(dataContent[0]) },
					func() fyne.CanvasObject { return widget.NewLabel(("--------------------------------")) },
					func(i widget.TableCellID, obj fyne.CanvasObject) {
						obj.(*widget.Label).SetText(dataContent[i.Row][i.Col])
					},
				)

				w1 := fyne.CurrentApp().NewWindow("Tabla")
				w1.Resize(fyne.NewSize(500, 400))
				w1.SetContent(container.NewScroll(table))
				w1.Show()
			}, w)
		file_Dialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		file_Dialog.Show()
	})

	box2 := container.NewVBox(
		text,
		btn,
	)

	w.SetContent(box2)

	w.ShowAndRun()
	fmt.Println("fin")
}
