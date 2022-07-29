// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v3/backend"
	"github.com/pulumi/pulumi/pkg/v3/backend/filestate"
	"github.com/pulumi/pulumi/pkg/v3/backend/httpstate"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
)

func newLogoutCmd() *cobra.Command {
	var cloudURL string
	var localMode bool
	var all bool

	cmd := &cobra.Command{
		Use:   placeholder.Use,
		Short: placeholder.Short,
		Long:  placeholder.Long,
		Args:  cmdutil.MaximumNArgs(1),
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			// If a <cloud> was specified as an argument, use it.
			if len(args) > 0 {
				if cloudURL != "" || all {
					return errors.New("only one of --all, --cloud-url or argument URL may be specified, not both")
				}
				cloudURL = args[0]
			}

			// For local mode, store state by default in the user's home directory.
			if localMode {
				if cloudURL != "" {
					return errors.New("a URL may not be specified when --local mode is enabled")
				}
				cloudURL = "file://~"
			}

			if cloudURL == "" {
				var err error
				cloudURL, err = workspace.GetCurrentCloudURL()
				if err != nil {
					return fmt.Errorf("could not determine current cloud: %w", err)
				}
			}

			var be backend.Backend
			var err error
			if filestate.IsFileStateBackendURL(cloudURL) {
				fmt.Printf("Logged out of %s\n", cloudURL)
				return workspace.DeleteAccount(cloudURL)
			}

			be, err = httpstate.New(cmdutil.Diag(), cloudURL)
			if err != nil {
				return err
			}

			var logoutErr error
			if all {
				logoutErr = be.LogoutAll()
			} else {
				logoutErr = be.Logout()
			}
			fmt.Printf("Logged out of %s\n", be.URL())
			return logoutErr
		}),
	}

	cmd.PersistentFlags().BoolVar(&all, "all", false,
		"Logout of all backends")
	cmd.PersistentFlags().StringVarP(&cloudURL, "cloud-url", "c", "",
		"A cloud URL to log out of (defaults to current cloud)")
	cmd.PersistentFlags().BoolVarP(&localMode, "local", "l", false,
		"Log out of using local mode")

	return cmd
}
