#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib'
import { InfrastructureStack } from '../lib/InfrastructureStack'
import { getGitHubThumbprint } from '../lib/getGitHubThumbprint'

const thumbprint = await getGitHubThumbprint()
const app = new cdk.App()
new InfrastructureStack(app, 'ClientStack', thumbprint)