package main

import (
	"context"
	"net/http"
	"time"
	"log"
	"golang.org/x/time/rate"
	 pb "lubdub.nl/aruba-iot/src/proto"
	"google.golang.org/protobuf/proto"
	"io"
	"nhooyr.io/websocket"
)

// echoServer is the WebSocket echo server implementation.
// It ensures the client speaks the echo subprotocol and
// only allows one message every 100ms with a 10 message burst.
type echoServer struct {
	// logf controls where logs are sent.
	logf func(f string, v ...interface{})
}

func (s echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"194.171.98.6"},
		CompressionMode: websocket.CompressionContextTakeover,
	})
	if err != nil {
		s.logf("%v", err)
		return
	}
	defer c.CloseNow()

	l := rate.NewLimiter(rate.Every(time.Millisecond*10), 10)
	for {
		err = echo(r.Context(), c, l)
		if err != nil {
			s.logf("failed to echo with %v: %v", r.RemoteAddr, err)
			return
		}
	}
}

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func echo(ctx context.Context, c *websocket.Conn, l *rate.Limiter) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	err := l.Wait(ctx)
	if err != nil {
		return err
	}
	iot_blob := &pb.Telemetry{}
	typ, r, err := c.Reader(ctx)
	b, err := io.ReadAll(r)
	if err := proto.Unmarshal(b,iot_blob); err != nil {
		log.Fatalln("Failed to parse mssage", err)
	}
	log.Println(iot_blob)
	if err != nil {
		return err
	}
	log.Println(typ)
	log.Println(r)

	return err
}
