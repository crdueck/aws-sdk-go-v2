// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detaches a policy from a target root, organizational unit (OU), or account. If
// the policy being detached is a service control policy (SCP), the changes to
// permissions for Identity and Access Management (IAM) users and roles in affected
// accounts are immediate. Every root, OU, and account must have at least one SCP
// attached. If you want to replace the default FullAWSAccess policy with an SCP
// that limits the permissions that can be delegated, you must attach the
// replacement SCP before you can remove the default SCP. This is the authorization
// strategy of an " allow list (https://docs.aws.amazon.com/organizations/latest/userguide/SCP_strategies.html#orgs_policies_allowlist)
// ". If you instead attach a second SCP and leave the FullAWSAccess SCP still
// attached, and specify "Effect": "Deny" in the second SCP to override the
// "Effect": "Allow" in the FullAWSAccess policy (or any other attached SCP),
// you're using the authorization strategy of a " deny list (https://docs.aws.amazon.com/organizations/latest/userguide/SCP_strategies.html#orgs_policies_denylist)
// ". This operation can be called only from the organization's management account.
func (c *Client) DetachPolicy(ctx context.Context, params *DetachPolicyInput, optFns ...func(*Options)) (*DetachPolicyOutput, error) {
	if params == nil {
		params = &DetachPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachPolicy", params, optFns, c.addOperationDetachPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachPolicyInput struct {

	// The unique identifier (ID) of the policy you want to detach. You can get the ID
	// from the ListPolicies or ListPoliciesForTarget operations. The regex pattern (http://wikipedia.org/wiki/regex)
	// for a policy ID string requires "p-" followed by from 8 to 128 lowercase or
	// uppercase letters, digits, or the underscore character (_).
	//
	// This member is required.
	PolicyId *string

	// The unique identifier (ID) of the root, OU, or account that you want to detach
	// the policy from. You can get the ID from the ListRoots ,
	// ListOrganizationalUnitsForParent , or ListAccounts operations. The regex pattern (http://wikipedia.org/wiki/regex)
	// for a target ID string requires one of the following:
	//   - Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//   letters or digits.
	//   - Account - A string that consists of exactly 12 digits.
	//   - Organizational unit (OU) - A string that begins with "ou-" followed by from
	//   4 to 32 lowercase letters or digits (the ID of the root that the OU is in). This
	//   string is followed by a second "-" dash and from 8 to 32 additional lowercase
	//   letters or digits.
	//
	// This member is required.
	TargetId *string

	noSmithyDocumentSerde
}

type DetachPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDetachPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDetachPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DetachPolicy",
	}
}
