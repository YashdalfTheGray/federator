package subcommands

import (
	"flag"

	"github.com/YashdalfTheGray/federator/constants"
)

// LinkSubcommandParsedArgs holds all the bits of data that are
// needed for the link subcommand to work properly.
type LinkSubcommandParsedArgs struct {
	RoleArn, ExternalID, Region, IssuerURL, DestinationURL string
	OutputJSON                                             bool
	subcommand                                             *flag.FlagSet
}

// NewLinkSubcommandParsedArgs creates an empty container for all the
// data that will be set up by calling .Setup and wil be populated by
// calling .Parse.
func NewLinkSubcommandParsedArgs() LinkSubcommandParsedArgs {
	return LinkSubcommandParsedArgs{}
}

// Setup will setup the subcommand with flags and descriptions.
func (lpa LinkSubcommandParsedArgs) Setup() {
	linkCmd := flag.NewFlagSet("link", flag.ExitOnError)
	linkCmd.StringVar(
		&lpa.RoleArn,
		"role-arn",
		"",
		"the role arn to assume for federating with AWS",
	)
	linkCmd.StringVar(
		&lpa.ExternalID,
		"external-id",
		"",
		"the external ID that can optionally be provided if the assume role requires it",
	)
	linkCmd.StringVar(
		&lpa.Region,
		"region",
		"",
		"the region to make the call against, will be read from the CLI config if omitted",
	)
	linkCmd.StringVar(
		&lpa.IssuerURL,
		"issuer",
		constants.DefaultIssuer,
		"the link where the user will be taken when the session has expired",
	)
	linkCmd.StringVar(
		&lpa.DestinationURL,
		"destination",
		constants.DefaultDestination,
		"the link that the user will be redirected to after login",
	)
	linkCmd.BoolVar(
		&lpa.OutputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)
}

// Parse will parse the flags, according to the arguments setup in .Setup
func (lpa LinkSubcommandParsedArgs) Parse(args []string) error {
	return lpa.subcommand.Parse(args)
}
