#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { FargateStack } from '../lib/FargateStack';

const app = new cdk.App();
new FargateStack(app, 'FargateStackV1', {
    env: {
        account: '022732888589',
        region: 'ap-southeast-2',
    },
    serviceName: 'test-service',
    imageTag: 'jonathandufty/recipes:latest',
    command: ['serve', 'auth']
});

new FargateStack(app, 'FargateStackV2', {
    env: {
        account: '022732888589',
        region: 'ap-southeast-2',
    },
    serviceName: 'test-service-v2',
    imageTag: 'jonathandufty/recipes:latest',
    command: ['serve','auth']
});