/*
 * Copyright 2000-2020 JetBrains s.r.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"flag"
	"github.com/JetBrains/docker-credential-space/cli"
	"github.com/google/subcommands"
	"os"
)

const (
	dockerCredStoreGroup = "Docker credential store API"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(cli.NewStoreSubcommand(), dockerCredStoreGroup)
	subcommands.Register(cli.NewGetSubcommand(), dockerCredStoreGroup)
	subcommands.Register(cli.NewEraseSubcommand(), dockerCredStoreGroup)
	subcommands.Register(cli.NewListSubcommand(), dockerCredStoreGroup)

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
