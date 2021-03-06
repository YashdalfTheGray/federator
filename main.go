package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/YashdalfTheGray/federator/constants"
	"github.com/YashdalfTheGray/federator/subcommands"
	"github.com/YashdalfTheGray/federator/utils"
)

func main() {
	var config aws.Config

	linkCmd := subcommands.NewLinkSubcommand()
	linkCmd.Setup()

	credsCmd := subcommands.NewCredsSubcommand()
	credsCmd.Setup()

	trustPolicyCmd := subcommands.NewTrustPolicySubcommand()
	trustPolicyCmd.Setup()

	if len(os.Args) < 2 {
		log.Fatalln("This executable needs a subcommand and options to work. Use -h for help.")
	}

	switch os.Args[1] {
	case "link":
		linkCmd.Parse(os.Args[2:])
		linkCmd.Validate()

		config = linkCmd.GetAWSConfig()

		if !linkCmd.Parsed.OutputJSON {
			fmt.Printf(
				"Using AWS STS in region %s to get a federated console signin link...\n",
				config.Region,
			)
			fmt.Print("\n")
		}

		creds, credsErr := utils.AuthWithSTS(
			linkCmd.Parsed.RoleArn,
			linkCmd.Parsed.ExternalID,
			config,
		)
		if credsErr != nil {
			log.Fatalln(credsErr.Error())
		}

		signinTokenURL := utils.GetSigninTokenURL(creds)
		signinToken, signinErr := utils.GetSigninToken(signinTokenURL)
		if signinErr != nil {
			log.Fatalln(signinErr.Error())
		}
		loginURL := utils.GetLoginURL(
			signinToken,
			linkCmd.Parsed.IssuerURL,
			linkCmd.Parsed.DestinationURL,
		)

		utils.PrintLoginURLDetails(creds, loginURL.String(), linkCmd.Parsed.OutputJSON)
		os.Exit(0)
		break
	case "creds":
		credsCmd.Parse(os.Args[2:])
		credsCmd.Validate()

		config = credsCmd.GetAWSConfig()

		if !credsCmd.Parsed.OutputJSON {
			fmt.Printf(
				"Using AWS STS in region %s to get temporary credentials...\n",
				config.Region,
			)
			fmt.Print("\n")
		}

		creds, credsErr := utils.AuthWithSTS(
			credsCmd.Parsed.RoleArn,
			credsCmd.Parsed.ExternalID,
			config,
		)
		if credsErr != nil {
			log.Fatalln(credsErr.Error())
		}

		utils.PrintCredsFromSTSResponse(creds, credsCmd.Parsed.OutputJSON)
		os.Exit(0)
		break
	case "trust-policy":
		trustPolicyCmd.Parse(os.Args[2:])
		if !trustPolicyCmd.Parsed.OutputJSON {
			fmt.Println("The trust policy with the provided details is below:")
		}
		fmt.Println(trustPolicyCmd.TrustPolicyString())
		os.Exit(0)
		break
	case "-h", "--help":
		fmt.Println(fmt.Sprintf("\nfederator %s", constants.Version))
		fmt.Println("\nUsage: federator <subcommand> <options>")
		fmt.Println("Options for the link subcommand")
		linkCmd.SetOutput(os.Stdout)
		linkCmd.PrintDefaults()
		fmt.Println("\nOptions for the creds subcommand")
		credsCmd.SetOutput(os.Stdout)
		credsCmd.PrintDefaults()
		fmt.Println("\nOptions for the trust-policy subcommand")
		trustPolicyCmd.SetOutput(os.Stdout)
		trustPolicyCmd.PrintDefaults()
		fmt.Println("\nUse `federator -h` or `federator --help` to display this help text.")
		os.Exit(0)
		break
	case "-v", "--version":
		fmt.Println(fmt.Sprintf("v%s", constants.Version))
	default:
		log.Fatalln(fmt.Sprintf("Invalid subcommand, %s. Valid options are link, creds", os.Args[1]))
	}
}
