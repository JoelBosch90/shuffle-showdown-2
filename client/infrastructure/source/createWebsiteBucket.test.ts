const mockCreateCloudFrontDistribution = jest.fn();

import { App, Stack } from 'aws-cdk-lib';
import { Template, Match } from 'aws-cdk-lib/assertions';
import { createWebsiteBucket } from './createWebsiteBucket';

jest.mock('./createCloudFrontDistribution', () => ({
  createCloudFrontDistribution: mockCreateCloudFrontDistribution,
}));

const isLocalEnvironmentSpy = jest.spyOn(require('./isLocalEnvironment'), 'isLocalEnvironment');

describe('createWebsiteBucket', () => {
  const bucketName = 'my-test-bucket';
  const stackName = 'TestStack';
  const fakeCertificateArn = 'arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012';
  const mockDomainName = 'example.com';
  const mockDistribution = { domainName: mockDomainName };

  beforeEach(() => {
    jest.clearAllMocks();

    mockCreateCloudFrontDistribution.mockReturnValue(mockDistribution);
    isLocalEnvironmentSpy.mockReturnValue(false);

    process.env.PRIVATE_SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN = fakeCertificateArn;
  });

  it('creates an S3 bucket with website hosting configured', () => {
    const app = new App();
    const stack = new Stack(app, stackName);

    createWebsiteBucket(stack, bucketName);

    Template.fromStack(stack).hasResourceProperties('AWS::S3::Bucket', {
      BucketName: bucketName,
      WebsiteConfiguration: {
        IndexDocument: 'index.html',
      },
    });
  });

  it('creates a BucketDeployment resource', () => {
    const app = new App();
    const stack = new Stack(app, stackName);

    createWebsiteBucket(stack, bucketName);

    // BucketDeployment is implemented as a custom resource.
    // This assertion ensures that one such resource is defined.
    Template.fromStack(stack).resourceCountIs('Custom::CDKBucketDeployment', 1);
  });

  it('creates a CloudFront distribution if the environment is remote', () => {
    const app = new App();
    const stack = new Stack(app, stackName);

    createWebsiteBucket(stack, bucketName);

    expect(mockCreateCloudFrontDistribution).toHaveBeenCalledWith(
      stack,
      expect.objectContaining({ physicalName: bucketName.replace(/:/g, '-') }),
      bucketName
    );
    Template.fromStack(stack).hasOutput('WebsiteUrl', {
      Value: mockDomainName,
    });
  });

  it('creates no CloudFront distribution and uses bucket url if the environment is local', () => {
    isLocalEnvironmentSpy.mockReturnValue(true);
    const app = new App();
    const stack = new Stack(app, stackName);

    createWebsiteBucket(stack, bucketName);

    expect(mockCreateCloudFrontDistribution).not.toHaveBeenCalledWith();
    Template.fromStack(stack).hasOutput('WebsiteUrl', {
      Value: {
        "Fn::GetAtt": [
          Match.stringLikeRegexp("Website.*"),
          "WebsiteURL",
        ],
      }
    });
  });
});