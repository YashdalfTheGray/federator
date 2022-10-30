package subcommands

import (
	"flag"
	"io"
	"log"

	"github.com/YashdalfTheGray/federator/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// CredsSubcommandParsedArgs holds all the bits of data that are
// needed for the link subcommand to work properly.
type CredsSubcommandParsedArgs struct {
	RoleArn, ExternalID, Region, Profile string
	OutputJSON, OutputAwsCli             bool
}

// CredsSubcommand holds the parsed args, when populated as well as internal
// state that is needed to make this work.
type CredsSubcommand struct {
	Parsed     CredsSubcommandParsedArgs
	subcommand *flag.FlagSet
}

func newCredsSubcommandParsedArgs() CredsSubcommandParsedArgs {
	return CredsSubcommandParsedArgs{}
}

// NewCredsSubcommand creates an empty container for all the
// data that will be set up by calling .Setup and wil be populated by
// calling .Parse.
func NewCredsSubcommand() CredsSubcommand {
	return CredsSubcommand{
		Parsed: newCredsSubcommandParsedArgs(),
	}
}

// Setup will setup the subcommand with flags and descriptions.
func (cmd *CredsSubcommand) Setup() {
	cmd.subcommand = flag.NewFlagSet("creds", flag.ExitOnError)

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
		&cmd.Parsed.Profile,
		"profile",
		"",
		"the aws credentials profile to use when making the assume role call",
	)
	cmd.subcommand.BoolVar(
		&cmd.Parsed.OutputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)
	cmd.subcommand.BoolVar(
		&cmd.Parsed.OutputAwsCli,
		"awscli",
		false,
		"output results in JSON format suitable for use with credentials_process",
	)
}

// Validate runs some general validations on the arguments
func (cmd CredsSubcommand) Validate() {
	if cmd.Parsed.Region != "" && !utils.IsRegionValid(cmd.Parsed.Region) {
		log.Fatalln("invalid value passed in for the --region flag")
	}

	if cmd.Parsed.RoleArn == "" {
		log.Fatalln("the --role-arn flag is required for this subcommand")
	}
}

// GetAWSConfig gets the right AWS config based on whether the
// region is passed in or read from the CLI configuration
func (cmd CredsSubcommand) GetAWSConfig() aws.Config {
	opts := []func(*config.LoadOptions) error{}
	if cmd.Parsed.Region != "" {
		opts = append(opts, config.WithRegion(cmd.Parsed.Region))
	}
	if cmd.Parsed.Profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(cmd.Parsed.Profile))
	}
	return utils.GetAWSConfigOpts(opts...)
}

// Parse will parse the flags, according to the arguments setup in .Setup
func (cmd CredsSubcommand) Parse(args []string) error {
	return cmd.subcommand.Parse(args)
}

// SetOutput is a mirror of flag.FlagSet.SetOutput
func (cmd CredsSubcommand) SetOutput(output io.Writer) {
	cmd.subcommand.SetOutput(output)
}

// PrintDefaults is a mirror of flag.FlagSet.PrintDefaults
func (cmd CredsSubcommand) PrintDefaults() {
	cmd.subcommand.PrintDefaults()
}
