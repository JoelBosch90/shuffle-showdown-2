import { Stack, StackProps, CfnOutput, RemovalPolicy, Aws } from 'aws-cdk-lib'
import { CfnDisk } from 'aws-cdk-lib/aws-lightsail'
import { BlockPublicAccess, Bucket } from 'aws-cdk-lib/aws-s3'
import { BucketDeployment, Source } from 'aws-cdk-lib/aws-s3-deployment'
import { Construct } from 'constructs'

export class InfrastructureStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props)

    const bucket = new Bucket(this, 'Website', {
      bucketName: `clientstack-website-${Aws.ACCOUNT_ID}-${Aws.REGION}`,
      websiteIndexDocument: 'index.html',
      websiteErrorDocument: 'error.html',
      publicReadAccess: true,
      blockPublicAccess: BlockPublicAccess.BLOCK_ACLS,
      removalPolicy: RemovalPolicy.DESTROY,
      autoDeleteObjects: true,
    })

    new BucketDeployment(this, 'DeployWebsite', {
      sources: [Source.asset('../source/assets')],
      destinationBucket: bucket,
    })

    new CfnOutput(this, 'WebsiteUrl', {
      value: bucket.bucketWebsiteUrl,
    })
  }
}
