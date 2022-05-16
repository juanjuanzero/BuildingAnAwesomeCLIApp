# Building an Awesome CLI App in Go

- Link to the resource: https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/
- Well i was wanting to build a CLI application, then it pointed me to cobra. :shrug: so i guess i'll just build a cli using that. Here are my notes for this.
- Here is cobra's github page: https://github.com/spf13/cobra/blob/master/user_guide.md

# Getting started

- Go to the docs page and read up on how to get started, you'll need to get the cobra downloaded so that it can be used.
- You will also need to get by initilizing a go mod file using `go mod init`, the clean up using `go mod tidy`.

# Getting Familiar

There is an example CLI that is available in the user_guide that you can, lets take a look at that.

```Go
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var echoTimes int

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}
    //add a flag to the times command.
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"} //the root command is the entry point
	rootCmd.AddCommand(cmdPrint, cmdEcho) //top-level commands are wired to the root
	cmdEcho.AddCommand(cmdTimes) //make the times command, a sub-command.
	rootCmd.Execute()
}

```

Here i've added some comments. The root command like the entry point for the application. Here we create a few commands to our CLI application. The `print` and `echo` commands are top-level commands, whereas `times` is a sub-command of `echo`. We've also introduced flags to the `times` command to modify its behavior.

As an example if you build the application using `go build`, you can execute commands with the application. Here is an example:

```
$ ./BuildingAnAwesomeCliApp echo times "hello" -t 5
//Output:
//Echo: hello
//Echo: hello
//Echo: hello
//Echo: hello
//Echo: hello
//Echo: hello
```

Notice how we have the `echo times` part. This is because the `times` command is a sub-command of the `echo` command. We also pass in the `-t` flag with an int value to tell how many times to repeat the value.

# What am I going to build?

That is the question, we'll try to build a to-do list application with a CLI. Here is the functionality that we can expect from the app:

- Add tasks
- List the tasks
- Remove tasks
- Mark as task as done
- Add a priority to a task
- Update a task's details, priority and/or due date

# Getting Started

So first well go ahead and clean up the main.go file and create a struct that describes our task. Our first priority would just be to add a task and list it out to the console. But before we do that we'll actually rearrange our file. You see the [library's user guide](https://github.com/spf13/cobra/blob/master/user_guide.md) suggests a specific structure. So we'll just follow that for now.

We'll create a godoit directory and move our main.go file there. In the src directory we'll create a cmd directory. This is where we will place all of our commands for the cli.

Now we will create root.go, which will be the entry point for our application. Here is what root.go looks like

```Go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godoit",
	Short: "this is my first cli-app so please be gentle",
	Long:  "a todo list application build for the love of learning",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

```

Next we are going to clean up our main.go file. Here is what it looks like now.

```Go
package main

import (
	"github.com/juanjuanzero/BuildingAnAwesomeCliApp/godoit/cmd"
)

func main() {
	cmd.Execute()
}
```

Since this will function as a todo list application we'll need a place to keep our to-dos somewhere. For now we'll just read from and write to a json file.
