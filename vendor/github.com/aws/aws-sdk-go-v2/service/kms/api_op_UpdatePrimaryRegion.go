// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the primary key of a multi-Region key. This operation changes the
// replica key in the specified Region to a primary key and changes the former
// primary key to a replica key. For example, suppose you have a primary key in
// us-east-1 and a replica key in eu-west-2 . If you run UpdatePrimaryRegion with
// a PrimaryRegion value of eu-west-2 , the primary key is now the key in eu-west-2
// , and the key in us-east-1 becomes a replica key. For details, see Updating the
// primary Region (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-update)
// in the Key Management Service Developer Guide. This operation supports
// multi-Region keys, an KMS feature that lets you create multiple interoperable
// KMS keys in different Amazon Web Services Regions. Because these KMS keys have
// the same key ID, key material, and other metadata, you can use them
// interchangeably to encrypt data in one Amazon Web Services Region and decrypt it
// in a different Amazon Web Services Region without re-encrypting the data or
// making a cross-Region call. For more information about multi-Region keys, see
// Multi-Region keys in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
// in the Key Management Service Developer Guide. The primary key of a multi-Region
// key is the source for properties that are always shared by primary and replica
// keys, including the key material, key ID (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id)
// , key spec (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-spec)
// , key usage (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-usage)
// , key material origin (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-origin)
// , and automatic key rotation (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
// . It's the only key that can be replicated. You cannot delete the primary key (https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html)
// until all replica keys are deleted. The key ID and primary Region that you
// specify uniquely identify the replica key that will become the primary key. The
// primary Region must already have a replica key. This operation does not create a
// KMS key in the specified Region. To find the replica keys, use the DescribeKey
// operation on the primary key or any replica key. To create a replica key, use
// the ReplicateKey operation. You can run this operation while using the affected
// multi-Region keys in cryptographic operations. This operation should not delay,
// interrupt, or cause failures in cryptographic operations. Even after this
// operation completes, the process of updating the primary Region might still be
// in progress for a few more seconds. Operations such as DescribeKey might
// display both the old and new primary keys as replicas. The old and new primary
// keys have a transient key state of Updating . The original key state is restored
// when the update is complete. While the key state is Updating , you can use the
// keys in cryptographic operations, but you cannot replicate the new primary key
// or perform certain management operations, such as enabling or disabling these
// keys. For details about the Updating key state, see Key states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide. This operation does not return
// any output. To verify that primary key is changed, use the DescribeKey
// operation. Cross-account use: No. You cannot use this operation in a different
// Amazon Web Services account. Required permissions:
//   - kms:UpdatePrimaryRegion on the current primary key (in the primary key's
//     Region). Include this permission primary key's key policy.
//   - kms:UpdatePrimaryRegion on the current replica key (in the replica key's
//     Region). Include this permission in the replica key's key policy.
//
// Related operations
//   - CreateKey
//   - ReplicateKey
func (c *Client) UpdatePrimaryRegion(ctx context.Context, params *UpdatePrimaryRegionInput, optFns ...func(*Options)) (*UpdatePrimaryRegionOutput, error) {
	if params == nil {
		params = &UpdatePrimaryRegionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePrimaryRegion", params, optFns, c.addOperationUpdatePrimaryRegionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePrimaryRegionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePrimaryRegionInput struct {

	// Identifies the current primary key. When the operation completes, this KMS key
	// will be a replica key. Specify the key ID or key ARN of a multi-Region primary
	// key. For example:
	//   - Key ID: mrk-1234abcd12ab34cd56ef1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/mrk-1234abcd12ab34cd56ef1234567890ab
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey .
	//
	// This member is required.
	KeyId *string

	// The Amazon Web Services Region of the new primary key. Enter the Region ID,
	// such as us-east-1 or ap-southeast-2 . There must be an existing replica key in
	// this Region. When the operation completes, the multi-Region key in this Region
	// will be the primary key.
	//
	// This member is required.
	PrimaryRegion *string

	noSmithyDocumentSerde
}

type UpdatePrimaryRegionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePrimaryRegionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePrimaryRegion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePrimaryRegion{}, middleware.After)
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
	if err = addOpUpdatePrimaryRegionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePrimaryRegion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePrimaryRegion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "UpdatePrimaryRegion",
	}
}
