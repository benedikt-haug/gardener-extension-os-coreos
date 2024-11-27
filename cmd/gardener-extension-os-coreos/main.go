// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/gardener/gardener/pkg/logger"
	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/gardener/gardener-extension-os-coreos/cmd/gardener-extension-os-coreos/app"
)

func main() {
	runtimelog.SetLogger(logger.MustNewZapLogger(logger.InfoLevel, logger.FormatJSON))
	cmd := app.NewControllerCommand()

	if err := cmd.ExecuteContext(signals.SetupSignalHandler()); err != nil {
		runtimelog.Log.Error(err, "Error executing the main controller command")
		os.Exit(1)
	}
}
