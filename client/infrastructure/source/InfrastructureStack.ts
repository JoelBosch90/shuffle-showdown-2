import { Stack, StackProps, Aws } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { createWebsiteBucket } from './createWebsiteBucket';
import { createGitHubActionsDeploymentRole } from './createGitHubActionsDeploymentRole';
import { getGitHubThumbprint } from './getGitHubThumbprint';

export class InfrastructureStack extends Stack {
  constructor(scope: Construct, id: string, props: StackProps = {}) {
    super(scope, id, props);
  }

  public async build() {
    createWebsiteBucket(this, `clientstack-website-${Aws.ACCOUNT_ID}-${Aws.REGION}`);
    await this.buildGitHubActionsDeploymentRole();
  }

  private async buildGitHubActionsDeploymentRole() {
    const thumbprint = await getGitHubThumbprint();
    createGitHubActionsDeploymentRole(this, thumbprint);
  }
};
