package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/YashdalfTheGray/federator/constants"
	"github.com/YashdalfTheGray/federator/subcommands"
	"github.com/YashdalfTheGray/federator/utils"
)

func main() {
	var roleArn, issuerURL, destinationURL, externalID, region string
	var outputJSON bool
	var config aws.Config

	argsWithoutCommand := make([]string, 10)
	copy(argsWithoutCommand, os.Args[2:])

	linkCmd := subcommands.NewLinkSubcommand()
	linkCmd.Setup()

	credsCmd := flag.NewFlagSet("creds", flag.ExitOnError)
	credsCmd.StringVar(
		&roleArn,
		"role-arn",
		"",
		"the role arn to assume for federating with AWS",
	)
	credsCmd.StringVar(
		&externalID,
		"external-id",
		"",
		"the external ID that can optionally be provided if the assume role requires it",
	)
	credsCmd.StringVar(
		&region,
		"region",
		"",
		"the region to make the call against, will be read from the CLI config if omitted",
	)
	credsCmd.BoolVar(
		&outputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)

	if len(os.Args) < 2 {
		log.Fatalln("This executable needs a subcommand and options to work. Use -h for help.")
	}

	switch os.Args[1] {
	case "link":
		linkCmd.Parse(argsWithoutCommand)

		if linkCmd.Parsed.Region == "" {
			config = utils.GetAWSConfig()
		} else {
			config = utils.GetAWSConfigForRegion(region)
		}

		if !linkCmd.Parsed.OutputJSON {
			fmt.Printf(
				"Using AWS STS in region %s to get a federated console signin link...\n",
				config.Region,
			)
			fmt.Print("\n")
		}

		if linkCmd.Parsed.RoleArn == "" {
			log.Fatalln("the --role-arn flag is required for this subcommand")
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
		loginURL := utils.GetLoginURL(signinToken, issuerURL, destinationURL)

		utils.PrintLoginURLDetailsv2(creds, loginURL.String(), linkCmd.Parsed.OutputJSON)
		os.Exit(0)
		break
	case "creds":
		credsCmd.Parse(os.Args[2:])

		if region == "" {
			config = utils.GetAWSConfig()
		} else {
			config = utils.GetAWSConfigForRegion(region)
		}

		if !outputJSON {
			fmt.Printf(
				"Using AWS STS in region %s to get temporary credentials...\n",
				config.Region,
			)
			fmt.Print("\n")
		}

		if roleArn == "" {
			log.Fatalln("the --role-arn flag is required for this subcommand")
		}

		creds, credsErr := utils.AuthWithSTS(roleArn, externalID, config)
		if credsErr != nil {
			log.Fatalln(credsErr.Error())
		}

		utils.PrintCredsFromSTSResponse(creds, outputJSON)
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
		fmt.Println("\nUse `federator -h` or `federator --help` to display this help text.")
		os.Exit(0)
		break
	case "-v", "--version":
		fmt.Println(fmt.Sprintf("v%s", constants.Version))
	default:
		log.Fatalln(fmt.Sprintf("Invalid subcommand, %s. Valid options are link, creds", os.Args[1]))
	}
}
