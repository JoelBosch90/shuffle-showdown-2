import { Stack, StackProps, CfnOutput, RemovalPolicy } from 'aws-cdk-lib'
import { BlockPublicAccess, Bucket } from 'aws-cdk-lib/aws-s3'
import { BucketDeployment, Source } from 'aws-cdk-lib/aws-s3-deployment'
import { Construct } from 'constructs'

export class InfrastructureStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props)

    const bucket = new Bucket(this, 'Website', {
      websiteIndexDocument: 'index.html',
      websiteErrorDocument: 'error.html',
      publicReadAccess: true,
      blockPublicAccess: BlockPublicAccess.BLOCK_ACLS,
      removalPolicy: RemovalPolicy.DESTROY,
      autoDeleteObjects: true,
    })

    new BucketDeployment(this, 'DeployWebsite', {
      sources: [Source.asset('../src/assets')],
      destinationBucket: bucket,
    })

    new CfnOutput(this, 'WebsiteUrl', {
      value: bucket.bucketWebsiteUrl,
    })
  }
}
