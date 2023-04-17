package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func genKeys(win fyne.Window, privKeyWidget *widget.Label, pubKeyWidget *widget.Label, secretWidget *widget.Label) {
	privKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		dialog.ShowError(err, win)
		return
	} else {
		pubKeyWidget.SetText(privKey.PublicKey().String())
		privKeyWidget.SetText(privKey.String())
	}

	secret, err := wgtypes.GenerateKey()
	if err != nil {
		dialog.ShowError(err, win)
		return
	} else {
		secretWidget.SetText(secret.String())
	}

}

func copyToClipboard(cp fyne.Clipboard, wid *widget.Label) {
	cp.SetContent(wid.Text)
}

func makeContent(win fyne.Window) fyne.CanvasObject {
	privKeyLable := widget.NewLabel("None")
	pubKeyLable := widget.NewLabel("None")
	secretLable := widget.NewLabel("None")

	priKeyCopy := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		copyToClipboard(win.Clipboard(), privKeyLable)
	})
	pubKeyCopy := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		copyToClipboard(win.Clipboard(), pubKeyLable)
	})
	secretCopy := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		copyToClipboard(win.Clipboard(), secretLable)
	})

	genKeys(win, privKeyLable, pubKeyLable, secretLable)

	return container.NewVBox(
		widget.NewLabel("Privat Key:"),
		container.NewHBox(
			priKeyCopy, privKeyLable),
		widget.NewLabel("Public Key:"),
		container.NewHBox(
			pubKeyCopy, pubKeyLable),
		widget.NewLabel("Pre-Shared Key:"),
		container.NewHBox(
			secretCopy, secretLable),
		widget.NewSeparator(),
		widget.NewButton(
			"Re-Generate Keys",
			func() { genKeys(win, privKeyLable, pubKeyLable, secretLable) }))
}

func main() {
	a := app.New()
	w := a.NewWindow("Wireguard Key Generator")

	w.SetContent(makeContent(w))
	w.Resize(fyne.NewSize(480, 300))
	w.SetFixedSize(true)
	w.SetMaster()
	w.ShowAndRun()
}
