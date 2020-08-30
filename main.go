package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/YashdalfTheGray/federator/constants"
	"github.com/YashdalfTheGray/federator/utils"
)

func main() {
	var roleArn, issuerURL, destinationURL, externalID string
	var outputJSON bool

	linkCmd := flag.NewFlagSet("link", flag.ExitOnError)
	linkCmd.StringVar(
		&roleArn,
		"role-arn",
		"",
		"the role arn to assume for federating with AWS",
	)
	linkCmd.StringVar(
		&externalID,
		"external-id",
		"",
		"the external ID that can optionally be provided if the assume role requires it",
	)
	linkCmd.StringVar(
		&issuerURL,
		"issuer",
		constants.DefaultIssuer,
		"the link where the user will be taken when the session has expired",
	)
	linkCmd.StringVar(
		&destinationURL,
		"destination",
		constants.DefaultDestination,
		"the link that the user will be redirected to after login",
	)
	linkCmd.BoolVar(
		&outputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)

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
		linkCmd.Parse(os.Args[2:])

		if !outputJSON {
			fmt.Println("Using AWS STS to get a federated console signin link...")
			fmt.Print("\n")
		}

		if roleArn == "" {
			log.Fatalln("the --role-arn flag is required for this subcommand")
		}

		creds, credsErr := utils.AuthWithSTS(roleArn, externalID)
		if credsErr != nil {
			log.Fatalln(credsErr.Error())
		}

		signinTokenURL := utils.GetSigninTokenURL(creds)
		signinToken, signinErr := utils.GetSigninToken(signinTokenURL)
		if signinErr != nil {
			log.Fatalln(signinErr.Error())
		}
		loginURL := utils.GetLoginURL(signinToken, issuerURL, destinationURL)

		utils.PrintLoginURLDetails(creds, loginURL.String(), outputJSON)
		os.Exit(0)
		break
	case "creds":
		credsCmd.Parse(os.Args[2:])

		if !outputJSON {
			fmt.Println("Using AWS STS to get temporary credentials...")
			fmt.Print("\n")
		}

		if roleArn == "" {
			log.Fatalln("the --role-arn flag is required for this subcommand")
		}

		creds, credsErr := utils.AuthWithSTS(roleArn, externalID)
		if credsErr != nil {
			log.Fatalln(credsErr.Error())
		}

		utils.PrintCredsFromSTSOutput(creds, outputJSON)
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
