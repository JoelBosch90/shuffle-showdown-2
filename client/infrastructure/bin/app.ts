#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { InfrastructureStack } from '../source/InfrastructureStack';

const app = new cdk.App();
const stack = new InfrastructureStack(app, 'ClientStack');
stack.build();
