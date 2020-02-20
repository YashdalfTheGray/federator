package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/sts"

	"github.com/YashdalfTheGray/federator/utils"
)

const envAWSAccessKeyID = "AWS_ACCESS_KEY_ID"
const envAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
const envAWSSessionToken = "AWS_SESSION_TOKEN"
const federationEndpoint = "https://signin.aws.amazon.com/federation"
const defaultIssuer = "https://aws.amazon.com"
const defaultDestination = "https://console.aws.amazon.com/console/home"

func getSessionString(creds *sts.AssumeRoleOutput) string {
	session := struct {
		SessionID    string `json:"sessionId"`
		SessionKey   string `json:"sessionKey"`
		SessionToken string `json:"sessionToken"`
	}{
		SessionID:    *creds.Credentials.AccessKeyId,
		SessionKey:   *creds.Credentials.SecretAccessKey,
		SessionToken: *creds.Credentials.SessionToken,
	}

	sessionStr, err := json.Marshal(session)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(sessionStr)
}

func getSigninTokenURL(creds *sts.AssumeRoleOutput) url.URL {
	u, err := url.Parse(federationEndpoint)
	if err != nil {
		log.Fatalln(err.Error())
	}

	q := u.Query()
	q.Set("Action", "getSigninToken")
	q.Set("SessionDuration", "3600")
	q.Set("Session", getSessionString(creds))

	u.RawQuery = q.Encode()
	return *u
}

func getSigninToken(signinURL url.URL) string {
	var signinResponse struct {
		SigninToken string `json:"SigninToken"`
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, signinReqErr := client.Get(signinURL.String())
	if signinReqErr != nil {
		log.Fatalln(signinReqErr.Error())
	}
	defer resp.Body.Close()

	body, readBodyErr := ioutil.ReadAll(resp.Body)
	if readBodyErr != nil {
		log.Fatalln(readBodyErr.Error())
	}

	unmarshalErr := json.Unmarshal(body, &signinResponse)
	if unmarshalErr != nil {
		log.Fatalln(unmarshalErr.Error())
	}

	return signinResponse.SigninToken
}

func getLoginURL(signinToken string) url.URL {
	u, err := url.Parse(federationEndpoint)
	if err != nil {
		log.Fatalln(err.Error())
	}

	q := u.Query()
	q.Set("Action", "login")
	q.Set("Issuer", defaultIssuer)
	q.Set("Destination", defaultDestination)
	q.Set("SigninToken", signinToken)

	u.RawQuery = q.Encode()
	return *u
}

func printCredsFromOutput(out *sts.AssumeRoleOutput) {
	fmt.Println("Successfully authenticated with STS. Commands to use below.")
	fmt.Println(fmt.Sprintf("This session will expire at %s", out.Credentials.Expiration.Local().String()))
	fmt.Println(fmt.Sprintf("export %s=%s", envAWSAccessKeyID, *out.Credentials.AccessKeyId))
	fmt.Println(fmt.Sprintf("export %s=%s", envAWSSecretAccessKey, *out.Credentials.SecretAccessKey))
	fmt.Println(fmt.Sprintf("export %s=%s", envAWSSessionToken, *out.Credentials.SessionToken))
}

func printLoginURLDetails(out *sts.AssumeRoleOutput, loginURL string) {
	fmt.Println("Successfully authenticated with STS. Login URL below.")
	fmt.Println(fmt.Sprintf("This session will expire at %s", out.Credentials.Expiration.Local().String()))
	fmt.Println(loginURL)
}

func main() {
	var roleArn string

	linkCmd := flag.NewFlagSet("link", flag.ExitOnError)
	linkCmd.StringVar(&roleArn, "role-arn", "", "the role arn to assume for federating with AWS")

	credsCmd := flag.NewFlagSet("creds", flag.ExitOnError)
	credsCmd.StringVar(&roleArn, "role-arn", "", "the role arn to assume for federating with AWS")

	if len(os.Args) < 2 {
		log.Fatalln("This executable needs a subcommand and options to work. Use -h for help.")
	}

	switch os.Args[1] {
	case "link":
		fmt.Println("Using AWS STS to get a federated console signin link...")
		linkCmd.Parse(os.Args[2:])
		if roleArn == "" {
			log.Fatalln("the --role-arn flag is required for this subcommand")
		}

		creds, credsErr := utils.AuthWithSTS(roleArn)
		if credsErr != nil {
			log.Fatalln(credsErr.Error())
		}

		signinTokenURL := getSigninTokenURL(creds)
		signinToken := getSigninToken(signinTokenURL)
		loginURL := getLoginURL(signinToken)

		printLoginURLDetails(creds, loginURL.String())
		os.Exit(0)
		break
	case "creds":
		fmt.Println("Using AWS STS to get temporary credentials...")
		fmt.Print("\n")
		credsCmd.Parse(os.Args[2:])

		if roleArn == "" {
			log.Fatalln("the --role-arn flag is required for this subcommand")
		}

		creds, credsErr := utils.AuthWithSTS(roleArn)
		if credsErr != nil {
			log.Fatalln(credsErr.Error())
		}

		printCredsFromOutput(creds)
		os.Exit(0)
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
		log.Fatalln(fmt.Sprintf("Invalid subcommand, %s. Valid options are link, creds", os.Args[1]))
	}
}
