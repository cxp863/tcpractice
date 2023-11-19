package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"net"
	"time"
)

var shortConnParam struct {
	Endpoint string
}

func shortConnClient(cmd *cobra.Command, args []string) {
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
		close(done)
	}()

	t := time.NewTicker(10 * time.Millisecond)

	for {
		select {
		case <-done:
			return
		case <-t.C:
			nonPersistentConn()
		}
	}
}

func nonPersistentConn() {
	conn, err := net.Dial("tcp", echoServerParam.Endpoint)
	if err != nil {
		logs.Fatal("dial tcp conn failed", err)
	}

	sendBuffer := []byte("hello")
	if _, err := conn.Write(sendBuffer); err != nil {
		logs.Error("write to conn failed", err)
	}

	readBuffer := make([]byte, 1536)
	if count, err := conn.Read(readBuffer); err == io.EOF {
		logs.Info("read eof")
	} else if err != nil {
		logs.Error("read from conn failed", err)
	} else {
		logs.Infof("read %d bytes", count)
	}

	if err := conn.Close(); err != nil {
		logs.Error("close conn failed", err)
	}
}

var shortConnCmd = &cobra.Command{
	Use: "short-client",
	Run: shortConnClient,
}

func init() {
	shortConnCmd.Flags().StringVar(&shortConnParam.Endpoint, "endpoint", "localhost:1222", "endpoint to connect")

	rootCmd.AddCommand(shortConnCmd)
}
