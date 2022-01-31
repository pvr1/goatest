// Code generated by goa v3.5.4, DO NOT EDIT.
//
// helloserver gRPC client CLI support package
//
// Command:
// $ goa gen github.com/pvr1/goatest/design

package cli

import (
	"flag"
	"fmt"
	"os"

	hellosvcc "github.com/pvr1/goatest/gen/grpc/hellosvc/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `hellosvc greet
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` hellosvc greet --message '{
      "name": "Iste ratione voluptatum."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		hellosvcFlags = flag.NewFlagSet("hellosvc", flag.ContinueOnError)

		hellosvcGreetFlags       = flag.NewFlagSet("greet", flag.ExitOnError)
		hellosvcGreetMessageFlag = hellosvcGreetFlags.String("message", "", "")
	)
	hellosvcFlags.Usage = hellosvcUsage
	hellosvcGreetFlags.Usage = hellosvcGreetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "hellosvc":
			svcf = hellosvcFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "hellosvc":
			switch epn {
			case "greet":
				epf = hellosvcGreetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "hellosvc":
			c := hellosvcc.NewClient(cc, opts...)
			switch epn {
			case "greet":
				endpoint = c.Greet()
				data, err = hellosvcc.BuildGreetPayload(*hellosvcGreetMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// hellosvcUsage displays the usage of the hellosvc command and its subcommands.
func hellosvcUsage() {
	fmt.Fprintf(os.Stderr, `The hello service
Usage:
    %[1]s [globalflags] hellosvc COMMAND [flags]

COMMAND:
    greet: Greet implements greet.

Additional help:
    %[1]s hellosvc COMMAND --help
`, os.Args[0])
}
func hellosvcGreetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] hellosvc greet -message JSON

Greet implements greet.
    -message JSON: 

Example:
    %[1]s hellosvc greet --message '{
      "name": "Iste ratione voluptatum."
   }'
`, os.Args[0])
}