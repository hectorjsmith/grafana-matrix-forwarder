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

var appVersion string

func main() {
	if appVersion == "" {
		appVersion = "dev"
	}

	ctx, _ := listenForInterrupt()

	cfg.Parse()
	if cfg.VersionMode {
		printAppVersion()
	} else {
		err := run(ctx, cfg.UserId, cfg.UserPassword, cfg.HomeserverUrl)
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
	fmt.Println(appVersion)
}

func run(ctx context.Context, userId, userPassword, homeserverUrl string) error {
	client, err := matrix.CreateClient(userId, userPassword, homeserverUrl)
	if err != nil {
		return err
	}
	return server.Start(ctx, client)
}
