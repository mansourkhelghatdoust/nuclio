/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package command

import (
	"github.com/nuclio/nuclio/pkg/functionconfig"
	"github.com/nuclio/nuclio/pkg/platform"

	"github.com/nuclio/errors"
	"github.com/spf13/cobra"
)

type deleteCommandeer struct {
	cmd            *cobra.Command
	rootCommandeer *RootCommandeer
}

func newDeleteCommandeer(rootCommandeer *RootCommandeer) *deleteCommandeer {
	commandeer := &deleteCommandeer{
		rootCommandeer: rootCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete resources",
	}

	deleteFunctionCommand := newDeleteFunctionCommandeer(commandeer).cmd
	deleteProjectCommand := newDeleteProjectCommandeer(commandeer).cmd
	deleteFunctionEventCommand := newDeleteFunctionEventCommandeer(commandeer).cmd

	cmd.AddCommand(
		deleteFunctionCommand,
		deleteProjectCommand,
		deleteFunctionEventCommand,
	)

	commandeer.cmd = cmd

	return commandeer
}

type deleteFunctionCommandeer struct {
	*deleteCommandeer
	functionConfig functionconfig.Config
}

func newDeleteFunctionCommandeer(deleteCommandeer *deleteCommandeer) *deleteFunctionCommandeer {
	commandeer := &deleteFunctionCommandeer{
		deleteCommandeer: deleteCommandeer,
		functionConfig:   *functionconfig.NewConfig(),
	}

	cmd := &cobra.Command{
		Use:     "functions [name[:version]]",
		Aliases: []string{"fu", "fn", "function"},
		Short:   "(or function) Delete functions",
		RunE: func(cmd *cobra.Command, args []string) error {

			// if we got positional arguments
			if len(args) != 1 {
				return errors.New("Function delete requires an identifier")
			}

			// initialize root
			if err := deleteCommandeer.rootCommandeer.initialize(); err != nil {
				return errors.Wrap(err, "Failed to initialize root")
			}

			commandeer.functionConfig.Meta.Name = args[0]
			commandeer.functionConfig.Meta.Namespace = deleteCommandeer.rootCommandeer.namespace

			return deleteCommandeer.rootCommandeer.platform.DeleteFunction(&platform.DeleteFunctionOptions{
				FunctionConfig: commandeer.functionConfig,
			})
		},
	}

	commandeer.cmd = cmd

	return commandeer
}

type deleteProjectCommandeer struct {
	*deleteCommandeer
	projectMeta platform.ProjectMeta
}

func newDeleteProjectCommandeer(deleteCommandeer *deleteCommandeer) *deleteProjectCommandeer {
	commandeer := &deleteProjectCommandeer{
		deleteCommandeer: deleteCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "projects name",
		Aliases: []string{"proj", "prj", "project"},
		Short:   "(or project) Delete projects",
		RunE: func(cmd *cobra.Command, args []string) error {

			// if we got positional arguments
			if len(args) != 1 {
				return errors.New("Project delete requires an identifier")
			}

			// initialize root
			if err := deleteCommandeer.rootCommandeer.initialize(); err != nil {
				return errors.Wrap(err, "Failed to initialize root")
			}

			commandeer.projectMeta.Name = args[0]
			commandeer.projectMeta.Namespace = deleteCommandeer.rootCommandeer.namespace

			return deleteCommandeer.rootCommandeer.platform.DeleteProject(&platform.DeleteProjectOptions{
				Meta: commandeer.projectMeta,
			})
		},
	}

	commandeer.cmd = cmd

	return commandeer
}

type deleteFunctionEventCommandeer struct {
	*deleteCommandeer
	functionEventMeta platform.FunctionEventMeta
}

func newDeleteFunctionEventCommandeer(deleteCommandeer *deleteCommandeer) *deleteFunctionEventCommandeer {
	commandeer := &deleteFunctionEventCommandeer{
		deleteCommandeer: deleteCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "functionevents name",
		Aliases: []string{"fe", "functionevent"},
		Short:   "(or functionevent) Delete function event",
		RunE: func(cmd *cobra.Command, args []string) error {

			// if we got positional arguments
			if len(args) != 1 {
				return errors.New("Function event delete requires an identifier")
			}

			// initialize root
			if err := deleteCommandeer.rootCommandeer.initialize(); err != nil {
				return errors.Wrap(err, "Failed to initialize root")
			}

			commandeer.functionEventMeta.Name = args[0]
			commandeer.functionEventMeta.Namespace = deleteCommandeer.rootCommandeer.namespace

			return deleteCommandeer.rootCommandeer.platform.DeleteFunctionEvent(&platform.DeleteFunctionEventOptions{
				Meta: commandeer.functionEventMeta,
			})
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
