package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct {
}

var _ fyne.Theme = (*myTheme)(nil)

func (t *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (t *myTheme) Font(s fyne.TextStyle) fyne.Resource {
	return resourceGoNotoCurrentTtf
}

func (t *myTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func (t *myTheme) Icon(s fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(s)
}

func main() {
	a := app.New()
	if len(os.Args) > 1 {
		a.Settings().SetTheme(&myTheme{})
	}

	w := a.NewWindow("Hello World")

	data := make([][2]string, 0)

	addMoreData := func() {
		if len(data) < 50000 {
			for i := 0; i < 100; i++ {
				data = append(data, [2]string{randomPhrase(), randomPhrase()})
			}
		}
	}

	addMoreData()

	list := widget.NewList(
		func() int { return len(data) },
		func() fyne.CanvasObject {
			grid := container.NewGridWithColumns(2)
			l1 := widget.NewRichTextWithText("")
			l1.Wrapping = fyne.TextTruncate
			l2 := widget.NewRichTextWithText("")
			l2.Segments[0].(*widget.TextSegment).Style.TextStyle.Bold = true
			l2.Wrapping = fyne.TextTruncate
			grid.Add(l1)
			grid.Add(l2)
			num := widget.NewLabel("")
			c := container.NewBorder(nil, nil, num, nil, grid)

			return c
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			c := co.(*fyne.Container)
			num := c.Objects[1].(*widget.Label)
			num.SetText(strconv.Itoa(lii))
			grid := c.Objects[0].(*fyne.Container)
			row := data[lii]
			grid.Objects[0].(*widget.RichText).Segments[0].(*widget.TextSegment).Text = row[0]
			grid.Objects[1].(*widget.RichText).Segments[0].(*widget.TextSegment).Text = row[1]
			c.Refresh()
			if lii > len(data)-20 {
				addMoreData()
			}
		},
	)

	w.SetContent(list)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}

func randomPhrase() string {
	r := func() int {
		return rand.Intn(len(words))
	}
	return fmt.Sprintf("%s %s %s %s %s",
		words[r()], words[r()], words[r()], words[r()], words[r()])
}
