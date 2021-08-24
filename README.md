# ASCII-Live

It is a demo project based upon [ANSI-art](https://github.com/EtoDemerzel0427/ANSI-art). ASCII-Live is created with ReactJS and Golang.
It is used to transform your live video captured by a webcam into an ASCII-form to display in command line:

https://user-images.githubusercontent.com/32998901/130675904-b090e856-af42-4c07-b732-84ca9127cb14.mp4

## How to run?
I assume you already have an environment that could run ReactJS and Golang.
If not, please refer to:
* [Create-React-App](https://github.com/facebook/create-react-app) for React env.
* [Golang.org](https://golang.org/) for Golang.

Once you have your environment ready, first open one terminal window:
```shell
$ cd ascii-live
$ go run main.go
```
Now you have the server running. There will not be any output till now, because we haven't
open the webcam via a browser.

Then open another terminal window (and cd to the same folder):
```shell
$ npm start
```

You will witness a browser running on http://localhost:3000. This webpage will open your webcam,
and display the video live inside the browser. In the same time, you can see the ascii-live showing in the first
terminal window.
## How it works?

ASCII-Live has Four major dependencies:
* [React-Webcam](https://github.com/mozmorris/react-webcam).
* [Gorilla Websocket](https://github.com/gorilla/websocket).
* [tcell](https://github.com/gdamore/tcell).
* [ANSI-art](https://github.com/EtoDemerzel0427/ANSI-art).

React-Webcam provides a convenient interface for us to utilize the web camera from a browser,
the image data it collects is transferred to the server via a WebSocket. 

WebSocket is an upgraded HTTP protocol
which is more suitable for long-term consistent communication, compared to the notorious long poll of HTTP. This functionality
is offered by Gorilla WebSocket. 

Then, as in my previous project, the image data is processed by ANSI-art to transform into ASCII
text, which can be directly output to terminal. 

While it is what we did in ANSI-art, this time we utilize tcell for better terminal visualization.

## Acknowledgement

This project is inspired by [ascii-fluid](https://github.com/esimov/ascii-fluid). What is worth noting is, during
studying his code, I spot a memory bug in this project, and found the same codebase was reused in many projects of the author's, 
including a 3k-star one. I think this story tells us writing WebAssembly in Golang is a bad idea 😂.