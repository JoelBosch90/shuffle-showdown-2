import { App, Stack } from 'aws-cdk-lib';
import { Template, Match } from 'aws-cdk-lib/assertions';
import { createWebsiteBucket } from './createWebsiteBucket';

describe('createWebsiteBucket', () => {
  const bucketName = 'my-test-bucket';

  it('creates an S3 bucket with website hosting configured', () => {
    const app = new App();
    const stack = new Stack(app, 'TestStack');

    createWebsiteBucket(stack, bucketName);

    const template = Template.fromStack(stack);
    template.hasResourceProperties('AWS::S3::Bucket', {
      BucketName: bucketName,
      WebsiteConfiguration: {
        IndexDocument: 'index.html',
        ErrorDocument: 'error.html',
      },
    });
  });

  it('creates a BucketDeployment resource', () => {
    const app = new App();
    const stack = new Stack(app, 'TestStack');

    createWebsiteBucket(stack, bucketName);

    const template = Template.fromStack(stack);
    // BucketDeployment is implemented as a custom resource.
    // This assertion ensures that one such resource is defined.
    template.resourceCountIs('Custom::CDKBucketDeployment', 1);
  });

  it('defines a CfnOutput for the WebsiteUrl', () => {
    const app = new App();
    const stack = new Stack(app, 'TestStack');

    createWebsiteBucket(stack, bucketName);

    const template = Template.fromStack(stack);
    // The output logical id is auto-generated but we can check
    // that one output exists with a value that references the
    // bucket's website URL.
    template.hasOutput('WebsiteUrl', {
      Value: {
        "Fn::GetAtt": [
          Match.stringLikeRegexp("Website.*"),
          "WebsiteURL",
        ],
      },
    });
  });
});