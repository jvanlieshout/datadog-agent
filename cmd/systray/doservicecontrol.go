// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.
// +build windows

package main

import (
	"github.com/DataDog/datadog-agent/cmd/agent/app"
	"golang.org/x/sys/windows/svc"
)

func onRestart() {
	if err := app.ControlService(svc.Stop, svc.Stopped); err == nil {
		err = app.StartService(nil, nil)
	}

}