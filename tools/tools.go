//go:build tools
// +build tools

// Tools Pattern as followed by the resource below
// https://www.jvt.me/posts/2024/09/30/go-tools-module/

package main

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
