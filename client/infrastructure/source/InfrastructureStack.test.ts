import { App, Stack, Aws } from 'aws-cdk-lib';
import { InfrastructureStack } from './InfrastructureStack';
import * as createWebsiteBucket from './createWebsiteBucket';
import * as createGitHubActionsDeploymentRole from './createGitHubActionsDeploymentRole';
import * as getGitHubThumbprint from './getGitHubThumbprint';

jest.mock('aws-cdk-lib');
const createWebsiteBucketSpy = jest.spyOn(createWebsiteBucket, 'createWebsiteBucket');
const createGitHubActionsDeploymentRoleSpy = jest.spyOn(createGitHubActionsDeploymentRole, 'createGitHubActionsDeploymentRole');
const getGitHubThumbprintSpy = jest.spyOn(getGitHubThumbprint, 'getGitHubThumbprint');

describe('InfrastructureStack', () => {
  const dummyDomain = 'some.domain.com';
  const dummyAccount = '123456789012';
  const dummyRegion = 'us-east-1';
  const dummyThumbprint = 'abcdefgh1234567890';
  const dummyStackProperties = {
    env: {
      account: dummyAccount,
      region: dummyRegion,
    },
  };

  beforeEach(() => {
    jest.resetAllMocks();
    getGitHubThumbprintSpy.mockResolvedValue(dummyThumbprint);
    process.env.SHUFFLE_SHOWDOWN_DOMAIN = dummyDomain;
  });

  it('throws an error if the domain is not defined', async () => {
    delete process.env.SHUFFLE_SHOWDOWN_DOMAIN;
    const stack = new InfrastructureStack(new App(), 'TestStack', dummyStackProperties);

    const promise = stack.build();

    await expect(promise).rejects.toThrow('Domain is not defined in environment variables.');
  });

  it('creates a stack with the provided properties', () => {
    new InfrastructureStack(new App(), 'TestStack', dummyStackProperties);

    expect(Stack).toHaveBeenCalledWith(expect.any(App), 'TestStack', dummyStackProperties);
  })

  it('creates a stack without the provided properties', () => {
    new InfrastructureStack(new App(), 'TestStack');

    expect(Stack).toHaveBeenCalledWith(expect.any(App), 'TestStack', {});
  })

  it('creates a website bucket with a proper name', async () => {
    const stack = new InfrastructureStack(new App(), 'TestStack', dummyStackProperties);
    await stack.build();

    expect(createWebsiteBucketSpy).toHaveBeenCalledWith(expect.any(Stack), dummyDomain);
  });

  it('creates a GitHub Actions deployment role', async () => {
    const stack = new InfrastructureStack(new App(), 'TestStack', dummyStackProperties);
    await stack.build();

    expect(createGitHubActionsDeploymentRoleSpy).toHaveBeenCalledWith(
      expect.any(Stack),
      dummyThumbprint,
    );
  });
});