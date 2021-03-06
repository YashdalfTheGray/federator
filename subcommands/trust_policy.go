package subcommands

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/YashdalfTheGray/federator/models"
)

// TrustPolicySubcommandParsedArgs holds all the bits of data that are
// needed for the trust policy subcommand to work properly.
type TrustPolicySubcommandParsedArgs struct {
	Arn, AccountID, ExternalID string
	OutputJSON                 bool
}

// TrustPolicySubcommand holds the parsed args, when populated,
// as well as internal state that is needed to make this work.
type TrustPolicySubcommand struct {
	Parsed     TrustPolicySubcommandParsedArgs
	subcommand *flag.FlagSet
}

func newTrustPolicySubcommandParsedArgs() TrustPolicySubcommandParsedArgs {
	return TrustPolicySubcommandParsedArgs{}
}

// NewTrustPolicySubcommand creates an empty container for all the
// data that will be set up by calling .Setup and wil be populated by
// calling .Parse.
func NewTrustPolicySubcommand() TrustPolicySubcommand {
	return TrustPolicySubcommand{
		Parsed: newTrustPolicySubcommandParsedArgs(),
	}
}

// Setup will setup the subcommand with flags and descriptions.
func (cmd *TrustPolicySubcommand) Setup() {
	cmd.subcommand = flag.NewFlagSet("trust-policy", flag.ExitOnError)

	cmd.subcommand.StringVar(
		&cmd.Parsed.Arn,
		"arn",
		"",
		"the ARN of an IAM object to add to the trust policy, either this or an account ID must be provided",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.AccountID,
		"account-id",
		"",
		"the account ID to add to the trust policy, either this or an arn must be provided",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.ExternalID,
		"external-id",
		"",
		"the external ID that can optionally be provided to be added to the trust policy",
	)
	cmd.subcommand.BoolVar(
		&cmd.Parsed.OutputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)
}

// Parse will parse the flags, according to the arguments setup in .Setup
func (cmd TrustPolicySubcommand) Parse(args []string) error {
	return cmd.subcommand.Parse(args)
}

// SetOutput is a mirror of flag.FlagSet.SetOutput
func (cmd TrustPolicySubcommand) SetOutput(output io.Writer) {
	cmd.subcommand.SetOutput(output)
}

// PrintDefaults is a mirror of flag.FlagSet.PrintDefaults
func (cmd TrustPolicySubcommand) PrintDefaults() {
	cmd.subcommand.PrintDefaults()
}

// TrustPolicyString prints out a trust policy built from the given parameters
// in JSON form
func (cmd TrustPolicySubcommand) TrustPolicyString() string {
	if os.Getenv("CI_MODE") == "true" {
		if cmd.Parsed.OutputJSON {
			fmt.Println("<Running in quiet mode because of CI but would print JSON>")
		} else {
			fmt.Println("<Running in quiet mode because of CI>")
		}
		return ""
	}

	var trustPolicy models.TrustPolicy

	if cmd.Parsed.Arn != "" {
		trustPolicy = models.NewTrustPolicy(cmd.Parsed.Arn, cmd.Parsed.ExternalID)

	} else if cmd.Parsed.AccountID != "" {
		trustPolicy = models.NewTrustPolicy(cmd.Parsed.AccountID, cmd.Parsed.ExternalID)
	}

	jsonString, _ := trustPolicy.ToJSONString()
	return jsonString
}
