package subcommands

import (
	"flag"
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
