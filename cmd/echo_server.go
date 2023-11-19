package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"net"
	"os"
)

var echoServerParam struct {
	Endpoint string
}

func echoServer(cmd *cobra.Command, args []string) {
	listener, err := net.Listen("tcp", echoServerParam.Endpoint)
	if err != nil {
		logs.Fatal(err)
	}
	logs.Infof("Listering on %v", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			logs.Fatal(err)
		}
		logs.Infof("Accepted connection to %v from %v", conn.LocalAddr(), conn.RemoteAddr())
		go worker(conn)
	}
}

func worker(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logs.Error("close conn failed")
		}
	}(conn)

	b := make([]byte, 1536)
	for {
		size, err := conn.Read(b)
		if err == io.EOF {
			logs.Info("read EOF")
			break
		}
		if err != nil {
			logs.Fatal("read from conn error, ", err)
			os.Exit(-1)
		}

		if size, err = conn.Write(b[0:size]); err != nil {
			logs.Fatal("write to conn error, ", err)
			os.Exit(-1)
		} else {
			logs.Infof("echo %d bytes", size)
		}
	}
}

var echoServerCmd = &cobra.Command{
	Use: "echo-server",
	Run: echoServer,
}

func init() {
	echoServerCmd.Flags().StringVar(&echoServerParam.Endpoint, "endpoint", "0.0.0.0:1222", "endpoint to listen")

	rootCmd.AddCommand(echoServerCmd)
}
