import { App, Stack } from 'aws-cdk-lib';
import { Template, Match } from 'aws-cdk-lib/assertions';
import { createWebsiteBucket } from './createWebsiteBucket';

describe('createWebsiteBucket', () => {
  const bucketName = 'my-test-bucket';
  const stackName = 'TestStack';
  const fakeCertificateArn = 'arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012';

  beforeEach(() => {
    process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN = fakeCertificateArn;
  });

  it('throws an error if the certificate ARN is not defined', () => {
    delete process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN;

    const app = new App();
    const stack = new Stack(app, stackName);

    expect(() => createWebsiteBucket(stack, bucketName)).toThrow(
      'Certificate ARN is not defined in environment variables.'
    );
  });

  it('creates an S3 bucket with website hosting configured', () => {
    const app = new App();
    const stack = new Stack(app, stackName);

    createWebsiteBucket(stack, bucketName);

    const template = Template.fromStack(stack);
    template.hasResourceProperties('AWS::S3::Bucket', {
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

    const template = Template.fromStack(stack);
    // BucketDeployment is implemented as a custom resource.
    // This assertion ensures that one such resource is defined.
    template.resourceCountIs('Custom::CDKBucketDeployment', 1);
  });

  it('creates a CloudFront distribution with the S3 bucket as the origin', () => {
    const app = new App();
    const stack = new Stack(app, stackName);

    createWebsiteBucket(stack, bucketName);

    const template = Template.fromStack(stack);
    template.hasResourceProperties('AWS::CloudFront::Distribution', {
      DistributionConfig: {
        Origins: [
          {
            Id: Match.stringLikeRegexp(`${stackName}WebsiteDistributionOrigin.*`),
          },
        ],
        ViewerCertificate: {
          AcmCertificateArn: fakeCertificateArn,
        },
      },
    });
  });

  it('defines a CfnOutput for the WebsiteUrl', () => {
    const app = new App();
    const stack = new Stack(app, stackName);

    createWebsiteBucket(stack, bucketName);

    const template = Template.fromStack(stack);
    // The output logical id is auto-generated but we can check
    // that one output exists with a value that references the
    // bucket's website URL.
    template.hasOutput('WebsiteUrl', {
      Value: {
        "Fn::GetAtt": [
          Match.stringLikeRegexp("Website.*"),
          "DomainName",
        ],
      },
    });
  });
});