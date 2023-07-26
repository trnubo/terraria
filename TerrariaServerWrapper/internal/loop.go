package internal

import (
	"bytes"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// startInputLoop begins a goroutine that continuously forwards
// os.stdin to the server's stdin pipe
func (server *Server) startInputLoop(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				logrus.Infof("Exiting read loop")
				return
			default:
				buf := make([]byte, 1024)
				n, _ := os.Stdin.Read(buf)
				if n == 0 {
					continue
				}
				buf = bytes.Trim(buf, "\x00")
				server.Stdin.Write(buf)
			}
		}
	}()
}

// sigtermHandler starts a goroutine which waits for a SIGTERM and
// safely shuts down the server when it receives one
func (server *Server) startSigtermHandler(ctx context.Context) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-ctx.Done():
			return
		case <-sigChan:
			logrus.Infof("Received signal, shutting down safely...")
			server.Shutdown()
			return
		}
	}()
}
