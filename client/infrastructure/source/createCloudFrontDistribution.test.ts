const mockCreateCertificate = jest.fn();
const mockCreateApiGateWayOrigin = jest.fn();

import { Stack } from 'aws-cdk-lib';
import { Bucket } from 'aws-cdk-lib/aws-s3';
import { createCloudFrontDistribution } from "./createCloudFrontDistribution";
import { HttpOrigin } from 'aws-cdk-lib/aws-cloudfront-origins';
import { Match, Template } from 'aws-cdk-lib/assertions';
import { ViewerProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';

jest.mock('./createCertificate', () => ({ createCertificate: mockCreateCertificate }));
jest.mock('./createApiGatewayOrigin', () => ({ createApiGatewayOrigin: mockCreateApiGateWayOrigin }));

const isLocalEnvironmentSpy = jest.spyOn(require('./isLocalEnvironment'), 'isLocalEnvironment');

describe('createCloudFrontDistribution', () => {
  const distributionName = 'WebsiteDistribution';
  const mockDomainName = 'example.com';
  const mockApiDomainName = 'example.com/api';
  const mockBucketName = 'mock-bucket';
  const mockCertificate = { certificateArn: 'arn:aws:acm:us-east-1:account-id:certificate/certificate-id' };
  const mockHttpOrigin = new HttpOrigin(mockApiDomainName);

  beforeEach(() => {
    jest.clearAllMocks();

    mockCreateCertificate.mockReturnValue(mockCertificate);
    mockCreateApiGateWayOrigin.mockReturnValue(mockHttpOrigin);
    isLocalEnvironmentSpy.mockReturnValue(false);
  });

  it('creates a CloudFront distribution with the correct name', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    const distribution = createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    expect(distribution.node.id).toEqual(distributionName);
  });

  it('creates a CloudFront distribution with the correct url', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        Aliases: Match.arrayWith([mockDomainName]),
      },
    });
  });

  it('creates a CloudFront distribution using createCertificate to get a certificate', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    expect(mockCreateCertificate).toHaveBeenCalledWith(mockStack);
  });

  it('creates a CloudFront distribution with a correct https redirect', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        CacheBehaviors: Match.arrayWith([Match.objectLike({
          ViewerProtocolPolicy: ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
        })]),
        DefaultCacheBehavior: Match.objectLike({
          ViewerProtocolPolicy: ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
        }),
      },
    });
  });

  it('creates a CloudFront distribution with the correct API path', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);
    const expectedPath = 'api';

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        CacheBehaviors: Match.arrayWith([Match.objectLike({
          PathPattern: `/${expectedPath}/*`,
        })]),
        Origins: Match.arrayWith([Match.objectLike({
          DomainName: `${mockDomainName}/${expectedPath}`,
        })]),
      },
    });
  });

  it('creates a CloudFront distribution using compression', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    expect(mockCreateApiGateWayOrigin).toHaveBeenCalled();
    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        DefaultCacheBehavior: {
          Compress: true,
        },
      },
    });
  });

  it('sets the default behavior with the correct origin', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    expect(mockCreateApiGateWayOrigin).toHaveBeenCalled();
    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        CacheBehaviors: Match.arrayWith([Match.objectLike({
          OriginRequestPolicyId: Match.stringLikeRegexp('.*'),
          TargetOriginId: Match.stringLikeRegexp(`${distributionName}Origin.*`),
        })]),
        DefaultCacheBehavior: {
          TargetOriginId: Match.stringLikeRegexp(`${distributionName}Origin.*`),
        },
      },
    });
  });

  it('sets the default behavior for the correct cache configurations', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        CacheBehaviors: Match.arrayWith([Match.objectLike({
          CachePolicyId: Match.stringLikeRegexp('.*'),
        })]),
        DefaultCacheBehavior: {
          CachePolicyId: Match.stringLikeRegexp('.*'),
        },
      },
    });
  });

  it('creates a CloudFront distribution allowing all methods', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        CacheBehaviors: Match.arrayWith([Match.objectLike({
          AllowedMethods: ['GET', 'HEAD', 'OPTIONS', 'PUT', 'PATCH', 'POST', 'DELETE'],
        })]),
        DefaultCacheBehavior: {
          AllowedMethods: ['GET', 'HEAD'],
        },
      },
    });
  });

  it('sets no additional behaviors in local environment', () => {
    isLocalEnvironmentSpy.mockReturnValue(true);
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    expect(mockCreateApiGateWayOrigin).not.toHaveBeenCalled();
    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        AdditionalBehaviors: Match.absent(),
        CacheBehaviors: Match.absent(),
      },
    });
  });

  it('creates a CloudFront distribution with a 403 error response', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        CustomErrorResponses: Match.arrayWith([Match.objectLike({
          ErrorCode: 403,
          ResponsePagePath: '/index.html',
          ErrorCachingMinTTL: 0,
          ResponseCode: 200
        })]),
      },
    });
  });

  it('creates a CloudFront distribution with a 404 error response', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        CustomErrorResponses: Match.arrayWith([Match.objectLike({
          ErrorCode: 404,
          ResponsePagePath: '/index.html',
          ErrorCachingMinTTL: 0,
          ResponseCode: 200
        })]),
      },
    });
  });

  it('uses the provided certificate for the distribution', () => {
    const mockStack = new Stack();
    const mockBucket = new Bucket(mockStack, mockBucketName);

    createCloudFrontDistribution(mockStack, mockBucket, mockDomainName);

    Template.fromStack(mockStack).hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        ViewerCertificate: Match.objectLike({ AcmCertificateArn: mockCertificate.certificateArn }),
      },
    });
  });
});