# wggen-ui
This program is a small, simplistic tool, to generate a private, public and shared secret.
For that, it's using the [wgtypes go packet](golang.zx2c4.com/wireguard/wgctrl/wgtypes) to generate the actual Keys, and [fyne](https://fyne.io/) as the GUI.

# Build
For this program, you need to run the following commands:
~~~
go get golang.zx2c4.com/wireguard/wgctrl/wgtypes@latest
go get fyne.io/fyne/v2@latest
go install fyne.io/fyne/v2/cmd/fyne@latest
go install github.com/fyne-io/fyne-cross@latest
~~~

To build, I use fyne-cross. With it, you can build an Executable for Windows, macOS and Linux on a single system.  For it to work, you must also install docker.
## Windows: 
~~~
fyne-cross windows -app-id wggen-ui
~~~
## macOS: 
~~~
fyne-cross darwin -app-id wggen-ui
~~~
## Linux: 
~~~
fyne-cross linux -app-id wggen-ui
~~~


# Legal
Please Note that the [wgtypes go packet](golang.zx2c4.com/wireguard/wgctrl/wgtypes) stands under the MIT Licence.

MIT License
Copyright (C) 2018-2022 Matt Layher

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
