Bhojpur RTMP - Server Engine
====

The Real-Time Messaging Protocol server

Run a simple media streaming server

	package main

	import "github.com/bhojpur/audio/pkg/rtmp"

	func main() {
		rtmp.SimpleServer()
	}

Use avconv to publish the media stream
	
	avconv -re -i a.mp4 -c:a copy -c:v copy -f flv rtmp://localhost/myapp/1

Use avplay to play media stream

	avplay rtmp://localhost/myapp/1
