package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("eu-west-1")},
		Profile: "jumpbox",
	})
	if err != nil {
		fmt.Println(err)
	}

	selectedarn := os.Args[1]

	svc2 := sts.New(sess)
	params := &sts.AssumeRoleInput{
		RoleArn:         aws.String(selectedarn),
		RoleSessionName: aws.String("gosume"),
		DurationSeconds: aws.Int64(3600),
	}
	resp, err := svc2.AssumeRole(params)
	if err != nil {
		fmt.Println(err)
	}
	var creds credentials.Value
	creds.AccessKeyID = *resp.Credentials.AccessKeyId
	creds.SecretAccessKey = *resp.Credentials.SecretAccessKey
	creds.SessionToken = *resp.Credentials.SessionToken
	exportstr := "\nexport AWS_ACCESS_KEY_ID=" + creds.AccessKeyID + " && export AWS_SECRET_ACCESS_KEY=" + creds.SecretAccessKey + " && export AWS_SESSION_TOKEN=" + creds.SessionToken + " && export AWS_SECURITY_TOKEN=" + creds.SessionToken + " && export AWS_DEFAULT_REGION=eu-west-1\n"
	fmt.Println("🠯🠯Copy below command and execute in shell to set AWS permissions as per chosen role 🠯🠯")
	fmt.Println(exportstr)
}
