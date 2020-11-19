// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAuthorizationTokenInput struct {
	_ struct{} `type:"structure"`

	// A list of AWS account IDs that are associated with the registries for which
	// to get AuthorizationData objects. If you do not specify a registry, the default
	// registry is assumed.
	RegistryIds []string `locationName:"registryIds" min:"1" type:"list"`
}

// String returns the string representation
func (s GetAuthorizationTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAuthorizationTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAuthorizationTokenInput"}
	if s.RegistryIds != nil && len(s.RegistryIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegistryIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAuthorizationTokenOutput struct {
	_ struct{} `type:"structure"`

	// A list of authorization token data objects that correspond to the registryIds
	// values in the request.
	AuthorizationData []AuthorizationData `locationName:"authorizationData" type:"list"`
}

// String returns the string representation
func (s GetAuthorizationTokenOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAuthorizationToken = "GetAuthorizationToken"

// GetAuthorizationTokenRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Retrieves an authorization token. An authorization token represents your
// IAM authentication credentials and can be used to access any Amazon ECR registry
// that your IAM principal has access to. The authorization token is valid for
// 12 hours.
//
// The authorizationToken returned is a base64 encoded string that can be decoded
// and used in a docker login command to authenticate to a registry. The AWS
// CLI offers an get-login-password command that simplifies the login process.
// For more information, see Registry Authentication (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth)
// in the Amazon Elastic Container Registry User Guide.
//
//    // Example sending a request using GetAuthorizationTokenRequest.
//    req := client.GetAuthorizationTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/GetAuthorizationToken
func (c *Client) GetAuthorizationTokenRequest(input *GetAuthorizationTokenInput) GetAuthorizationTokenRequest {
	op := &aws.Operation{
		Name:       opGetAuthorizationToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAuthorizationTokenInput{}
	}

	req := c.newRequest(op, input, &GetAuthorizationTokenOutput{})

	return GetAuthorizationTokenRequest{Request: req, Input: input, Copy: c.GetAuthorizationTokenRequest}
}

// GetAuthorizationTokenRequest is the request type for the
// GetAuthorizationToken API operation.
type GetAuthorizationTokenRequest struct {
	*aws.Request
	Input *GetAuthorizationTokenInput
	Copy  func(*GetAuthorizationTokenInput) GetAuthorizationTokenRequest
}

// Send marshals and sends the GetAuthorizationToken API request.
func (r GetAuthorizationTokenRequest) Send(ctx context.Context) (*GetAuthorizationTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAuthorizationTokenResponse{
		GetAuthorizationTokenOutput: r.Request.Data.(*GetAuthorizationTokenOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAuthorizationTokenResponse is the response type for the
// GetAuthorizationToken API operation.
type GetAuthorizationTokenResponse struct {
	*GetAuthorizationTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAuthorizationToken request.
func (r *GetAuthorizationTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
