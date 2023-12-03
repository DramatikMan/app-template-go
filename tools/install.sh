#!/bin/sh
set -e

cd tools

go install \
	"github.com/cweill/gotests/gotests" \
	"github.com/fatih/gomodifytags" \
	"github.com/go-delve/delve/cmd/dlv" \
	"github.com/haya14busa/goplay/cmd/goplay" \
	"github.com/josharian/impl" \
	"golang.org/x/tools/gopls" \
	"honnef.co/go/tools/cmd/staticcheck"
