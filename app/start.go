package app

import (
	"context"
	"math/rand"
	"time"
)

var onStart appHooks

func OnStart(name string, fn HookFunc) {
	onStart.Add(newHook(name, fn))
}

//------------------------------------------------------------------------------

var (
	onStop      appHooks
	onAfterStop appHooks
)

func OnStop(name string, fn HookFunc) {
	onStop.Add(newHook(name, fn))
}

func OnAfterStop(name string, fn HookFunc) {
	onAfterStop.Add(newHook(name, fn))
}

//------------------------------------------------------------------------------

func Start(ctx context.Context, service, envName string) error {
	cfg, err := ReadConfig(service, envName)
	if err != nil {
		return err
	}
	return StartConfig(ctx, cfg)
}

func StartConfig(ctx context.Context, cfg *AppConfig) error {
	rand.Seed(time.Now().UnixNano())

	TheApp = New(ctx, cfg)
	return onStart.Run(ctx, TheApp)
}

func Stop() {
	ctx := Context()
	_ = onStop.Run(ctx, TheApp)
	_ = onAfterStop.Run(ctx, TheApp)
	TheApp = nil
}
