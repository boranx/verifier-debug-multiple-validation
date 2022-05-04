package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/credentials"
	ocmlog "github.com/openshift-online/ocm-sdk-go/logging"
	"github.com/openshift/osd-network-verifier/pkg/cloudclient"
)

func main() {
	subnetIDs := []string{"subnet-030455abded968aa2", "subnet-03146948aeaeb21de"}

	// build the credentials provider
	creds := credentials.NewStaticCredentialsProvider("XX", "XX", "")

	defaultTags := map[string]string{
		"osd-network-verifier": "owned",
		"red-hat-managed":      "true",
		"Name":                 "osd-network-verifier",
	}
	ctx := context.Background()
	logger, _ := ocmlog.NewStdLoggerBuilder().Debug(true).Build()

	// init a cloudclient
	cli, err := cloudclient.NewClient(ctx, logger, creds, "us-east-1", "t3.micro", defaultTags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, subnetID := range subnetIDs {
		// call the validation function
		out := cli.ValidateEgress(ctx, subnetID, "", "", 1*time.Second)
		if !out.IsSuccessful() {
			failures, exceptions, errors := out.Parse()
			var allErrors []error
			allErrors = append(allErrors, failures...)
			allErrors = append(allErrors, exceptions...)
			allErrors = append(allErrors, errors...)

			for _, v := range allErrors {
				fmt.Println(subnetID, "->", v.Error())
			}
		}
	}
}
