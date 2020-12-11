package main

import (
	"context"
	"fmt"
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/matrix"
	"grafana-matrix-forwarder/server"
	"log"
	"os"
	"os/signal"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	ctx, _ := listenForInterrupt()

	appSettings := cfg.Parse()
	if appSettings.VersionMode {
		printAppVersion()
	} else {
		err := run(ctx, appSettings)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("done")
	}
}

func listenForInterrupt() (context.Context, context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		cancel()
	}()
	return ctx, cancel
}

func printAppVersion() {
	fmt.Println(version)
	fmt.Printf("    build date:  %s\r\n    commit hash: %s\r\n    built by:    %s\r\n", date, commit, builtBy)
}

func run(ctx context.Context, appSettings cfg.AppSettings) error {
	writeCloser, err := matrix.NewMatrixWriteCloser(appSettings.UserID, appSettings.UserPassword, appSettings.HomeserverURL)
	if err != nil {
		return err
	}
	return server.BuildServer(ctx, writeCloser, appSettings).Start()
}
