import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { createWebsiteBucket } from './createWebsiteBucket';
import { createGitHubActionsDeploymentRole } from './createGitHubActionsDeploymentRole';
import { getGitHubThumbprint } from './getGitHubThumbprint';

/**
 *  InfrastructureStack is the main stack for the Shuffle Showdown infrastructure.
 *  It sets up the website bucket and the GitHub Actions deployment role.
 */
export class InfrastructureStack extends Stack {

  /**
   *  Constructs a new InfrastructureStack.
   *  @param    {Construct} scope - The scope in which this stack is defined.
   *  @param    {string} id - The identifier for this stack.
   *  @param    {StackProps} [props] - Optional properties for the stack.
   *  @throws   {Error} - If the PUBLIC_SHUFFLE_SHOWDOWN_DOMAIN environment variable is not defined.
   */
  constructor(scope: Construct, id: string, props: StackProps = {}) {
    super(scope, id, props);
  }

  /**
   *  Builds the infrastructure stack.
   *  This method creates the website bucket and sets up the GitHub Actions deployment role.
   *  @throws   {Error} - If the PUBLIC_SHUFFLE_SHOWDOWN_DOMAIN environment variable is not defined.
   *  @returns  {Promise<void>} - A promise that resolves when the stack is built.
   */
  public async build(): Promise<void> {
    const domain = process.env.PUBLIC_SHUFFLE_SHOWDOWN_DOMAIN;
    if (!domain) {
      throw new Error('Domain is not defined in environment variables.');
    }

    createWebsiteBucket(this, domain);
    await this.buildGitHubActionsDeploymentRole();
  }

  /**
   *  Builds the GitHub Actions deployment role.
   *  This method retrieves the GitHub thumbprint and creates the deployment role.
   *  @returns  {Promise<void>} - A promise that resolves when the stack is built.
   */
  private async buildGitHubActionsDeploymentRole(): Promise<void> {
    const thumbprint = await getGitHubThumbprint();
    createGitHubActionsDeploymentRole(this, thumbprint);
  }
};
