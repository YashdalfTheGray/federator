package subcommands

import (
	"flag"
	"io"
	"log"

	"github.com/YashdalfTheGray/federator/constants"
	"github.com/YashdalfTheGray/federator/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
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
func (cmd *LinkSubcommand) Setup() {
	cmd.subcommand = flag.NewFlagSet("link", flag.ExitOnError)

	cmd.subcommand.StringVar(
		&cmd.Parsed.RoleArn,
		"role-arn",
		"",
		"the role arn to assume for federating with AWS",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.ExternalID,
		"external-id",
		"",
		"the external ID that can optionally be provided if the assume role requires it",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.Region,
		"region",
		"",
		"the region to make the call against, will be read from the CLI config if omitted",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.IssuerURL,
		"issuer",
		constants.DefaultIssuer,
		"the link where the user will be taken when the session has expired",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.DestinationURL,
		"destination",
		constants.DefaultDestination,
		"the link that the user will be redirected to after login",
	)
	cmd.subcommand.BoolVar(
		&cmd.Parsed.OutputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)
}

// Validate runs some general validations on the arguments
func (cmd LinkSubcommand) Validate() {
	if cmd.Parsed.Region != "" && !utils.IsRegionValid(cmd.Parsed.Region) {
		log.Fatalln("invalid value passed in for the --region flag")
	}

	if cmd.Parsed.RoleArn == "" {
		log.Fatalln("the --role-arn flag is required for this subcommand")
	}
}

// GetAWSConfig gets the right AWS config based on whether the
// region is passed in or read from the CLI configuration
func (cmd LinkSubcommand) GetAWSConfig() aws.Config {
	if cmd.Parsed.Region == "" {
		return utils.GetAWSConfig()
	}

	return utils.GetAWSConfigForRegion(cmd.Parsed.Region)
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
