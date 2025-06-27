package main

import (
	"context"
	"errors"

	"github.com/anthhub/forwarder"
)

func forwardWithError(config *Config) (err error) {
	if len(config.Options) == 0 {
		return errors.New("no target pods defined")
	}

	session, err := forwarder.WithForwarders(context.Background(), config.Options, config.KubeConfigPath)
	if err != nil {
		return err
	}
	defer session.Close()

	if _, err := session.Ready(); err != nil {
		return err
	}

	session.Wait()
	return nil
}
