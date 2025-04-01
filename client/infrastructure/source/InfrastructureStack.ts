import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { createWebsiteBucket } from './createWebsiteBucket';
import { createGitHubActionsDeploymentRole } from './createGitHubActionsDeploymentRole';
import { getGitHubThumbprint } from './getGitHubThumbprint';

export class InfrastructureStack extends Stack {
  constructor(scope: Construct, id: string, props: StackProps = {}) {
    super(scope, id, props);
  }

  public async build() {
    const domain = process.env.SHUFFLE_SHOWDOWN_DOMAIN;
    if (!domain) {
      throw new Error('Domain is not defined in environment variables.');
    }

    createWebsiteBucket(this, domain);
    await this.buildGitHubActionsDeploymentRole();
  }

  private async buildGitHubActionsDeploymentRole() {
    const thumbprint = await getGitHubThumbprint();
    createGitHubActionsDeploymentRole(this, thumbprint);
  }
};
