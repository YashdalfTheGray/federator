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
