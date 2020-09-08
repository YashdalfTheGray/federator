package subcommands

import (
	"flag"
	"io"

	"github.com/YashdalfTheGray/federator/constants"
)

// LinkSubcommandParsedArgs holds all the bits of data that are
// needed for the link subcommand to work properly.
type LinkSubcommandParsedArgs struct {
	RoleArn, ExternalID, Region, IssuerURL, DestinationURL string
	OutputJSON                                             bool
}

// LinkSubcommand holds the parsed args, when populated as well as internal
// state that is needed to make this work.
type LinkSubcommand struct {
	Parsed     LinkSubcommandParsedArgs
	subcommand *flag.FlagSet
}

func newLinkSubcommandParsedArgs() LinkSubcommandParsedArgs {
	return LinkSubcommandParsedArgs{}
}

// NewLinkSubcommand creates an empty container for all the
// data that will be set up by calling .Setup and wil be populated by
// calling .Parse.
func NewLinkSubcommand() LinkSubcommand {
	return LinkSubcommand{
		Parsed: newLinkSubcommandParsedArgs(),
	}
}

// Setup will setup the subcommand with flags and descriptions.
func (cmd LinkSubcommand) Setup() {
	linkCmd := flag.NewFlagSet("link", flag.ExitOnError)
	linkCmd.StringVar(
		&cmd.Parsed.RoleArn,
		"role-arn",
		"",
		"the role arn to assume for federating with AWS",
	)
	linkCmd.StringVar(
		&cmd.Parsed.ExternalID,
		"external-id",
		"",
		"the external ID that can optionally be provided if the assume role requires it",
	)
	linkCmd.StringVar(
		&cmd.Parsed.Region,
		"region",
		"",
		"the region to make the call against, will be read from the CLI config if omitted",
	)
	linkCmd.StringVar(
		&cmd.Parsed.IssuerURL,
		"issuer",
		constants.DefaultIssuer,
		"the link where the user will be taken when the session has expired",
	)
	linkCmd.StringVar(
		&cmd.Parsed.DestinationURL,
		"destination",
		constants.DefaultDestination,
		"the link that the user will be redirected to after login",
	)
	linkCmd.BoolVar(
		&cmd.Parsed.OutputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)
}

// Parse will parse the flags, according to the arguments setup in .Setup
func (cmd LinkSubcommand) Parse(args []string) error {
	return cmd.subcommand.Parse(args)
}

// SetOutput is a mirror of flag.FlagSet.SetOutput
func (cmd LinkSubcommand) SetOutput(output io.Writer) {
	cmd.subcommand.SetOutput(output)
}

// PrintDefaults is a mirror of flag.FlagSet.PrintDefaults
func (cmd LinkSubcommand) PrintDefaults() {
	cmd.subcommand.PrintDefaults()
}
