package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/aws/session"
)

func informAndExit(message string, code int) {
	fmt.Println(message)
	os.Exit(code)
}

func getAWSSession(options session.Options) *session.Session {
	return session.Must(session.NewSessionWithOptions(options))
}

func getSessionName(roleArn string) (string, error) {
	user, _ := os.LookupEnv("USER")
	roleRegex := regexp.MustCompile("arn:aws:iam::[0-9]{12}:role/([a-zA-Z0-9-]+)")
	match := roleRegex.FindAllStringSubmatch(roleArn, -1)

	if !roleRegex.MatchString(roleArn) {
		return "", errors.New("Invalid Role ARN")
	}

	return fmt.Sprintf("federator-%s-%s", user, match[0][1]), nil
}

func main() {
	var roleArn string
	var loadEnv bool

	linkCmd := flag.NewFlagSet("link", flag.ExitOnError)
	linkCmd.StringVar(&roleArn, "role-arn", "", "the role arn to assume for federating with AWS")

	credsCmd := flag.NewFlagSet("creds", flag.ExitOnError)
	credsCmd.StringVar(&roleArn, "role-arn", "", "the role arn to assume for federating with AWS")
	credsCmd.BoolVar(&loadEnv, "load", false, "include this flag to load the creds into the environment")

	if len(os.Args) < 2 {
		informAndExit("This executable needs a subcommand and options to work. Use -h for help.", 1)
	}

	switch os.Args[1] {
	case "link":
		fmt.Println("Using AWS STS to get a federated console signin link...")
		linkCmd.Parse(os.Args[2:])
		if roleArn == "" {
			informAndExit("the --role-arn flag is required for this subcommand", 1)
		}
		break
	case "creds":
		fmt.Println("Using AWS STS to get temporary credentials")
		credsCmd.Parse(os.Args[2:])
		if roleArn == "" {
			informAndExit("the --role-arn flag is required for this subcommand", 1)
		}
		break
	case "-h", "--help":
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
	default:
		informAndExit(fmt.Sprintf("Invalid subcommand, %s. Valid options are link, creds", os.Args[1]), 1)
	}
}
