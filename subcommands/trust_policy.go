package subcommands

import "flag"

// TrustPolicySubcommandParsedArgs holds all the bits of data that are
// needed for the trust policy subcommand to work properly.
type TrustPolicySubcommandParsedArgs struct {
	Arn, AccountID, ExternalID string
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
