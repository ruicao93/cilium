// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/cilium/cilium-cli/sysdump"
)

var (
	sysdumpOptions = sysdump.Options{
		LargeSysdumpAbortTimeout: sysdump.DefaultLargeSysdumpAbortTimeout,
		LargeSysdumpThreshold:    sysdump.DefaultLargeSysdumpThreshold,
		Writer:                   os.Stdout,
	}
)

func newCmdSysdump(hooks sysdump.SysdumpHooks) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sysdump",
		Short: "Collects information required to troubleshoot issues with Cilium and Hubble",
		Long:  ``,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Honor --Namespace global flag in case it is set and --cilium-Namespace is not set
			if sysdumpOptions.CiliumNamespace == "" && cmd.Flags().Changed("Namespace") {
				sysdumpOptions.CiliumNamespace = Namespace
			}
			// Silence klog to avoid displaying "throttling" messages - those are expected.
			klog.SetOutput(io.Discard)
			// Collect the sysdump.
			collector, err := sysdump.NewCollector(K8sClient, sysdumpOptions, time.Now(), Version)
			if err != nil {
				return fmt.Errorf("failed to create sysdump collector: %w", err)
			}
			if err := hooks.AddSysdumpTasks(collector); err != nil {
				return fmt.Errorf("failed to add custom tasks: %w", err)
			}
			if err = collector.Run(); err != nil {
				return fmt.Errorf("failed to collect sysdump: %w", err)
			}
			return nil
		},
	}

	sysdump.InitSysdumpFlags(cmd, &sysdumpOptions, "", hooks)

	return cmd
}
