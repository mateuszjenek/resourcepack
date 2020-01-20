package config

import "io"

type Options struct {
	Database struct {
		DataSourceName string
	}
}

var App *Options

func Read(reader io.Reader) error {
	return nil
}
