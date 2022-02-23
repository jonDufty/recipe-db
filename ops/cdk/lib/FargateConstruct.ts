import { Duration, Stack, StackProps } from 'aws-cdk-lib';
import { ApplicationLoadBalancedFargateService } from 'aws-cdk-lib/aws-ecs-patterns';
import { Construct } from 'constructs';
import { SecurityGroup, Vpc } from 'aws-cdk-lib/aws-ec2';
import {
    Cluster,
    ContainerImage,
    FargateTaskDefinition,
} from 'aws-cdk-lib/aws-ecs';
import { Role, ServicePrincipal } from 'aws-cdk-lib/aws-iam';
import { HostedZone } from 'aws-cdk-lib/aws-route53';
import { aws_kms as kms } from 'aws-cdk-lib';

export interface FargateProps {
    serviceName: string;
    imageTag: string;
    command: string[];
}

export class FargateConstruct extends Construct {
    constructor(scope: Construct, id: string, props: FargateProps) {
        super(scope, id);

        const vpc = Vpc.fromLookup(this, 'SanboxVPC', {
            vpcName: 'vpc/sandbox',
        });

        const kmsKey = kms.Key.fromLookup(this, 'ChamberKey', {
            aliasName: 'parameter_store_key',
        })

        kmsKey.keyArn

        const cluster = Cluster.fromClusterAttributes(scope, 'Cluster', {
            clusterName: 'test',
            vpc: vpc,
            securityGroups: [],
        });

        const hostedZone = HostedZone.fromLookup(scope, 'HostedZone', {
            domainName: 'sandbox.99test.com',
        });

        // Create new role for our task
        const role = new Role(scope, 'FargateTaskRole', {
            assumedBy: new ServicePrincipal('ecs-tasks.amazonaws.com'),
        });

        

        const taskDefinition = new FargateTaskDefinition(this, 'ECSTask', {
            cpu: 256,
            memoryLimitMiB: 512,
            family: 'test-service',
            executionRole: role,
            taskRole: role,
        });

        const container = taskDefinition.addContainer('recipes', {
            image: ContainerImage.fromRegistry(props.imageTag),
            entryPoint: ['/bin/recipes'],
            command: props.command,
            cpu: 256,
            memoryLimitMiB: 512,
            portMappings: [
                {
                    containerPort: 80,
                },
            ],
        });

        // Use default VPC security group in addition to any others provided.
        // const securityGroups = [
        //     new SecurityGroup(this, 'SecurityGroup', { vpc }),
        // ];

        const fargate = new ApplicationLoadBalancedFargateService(
            this,
            'FargateService',
            {
                serviceName: props.serviceName,
                assignPublicIp: false,
                cluster,
                taskDefinition,
                domainZone: hostedZone,
                domainName: `test-service.sandbox.99test.com`,
                desiredCount: 1,
                cpu: 256,
                memoryLimitMiB: 512,
                // cloudMapOptions: {
                //     cloudMapNamespace: new PublicDnsNamespace(this, 'DNSNamespace', {
                //         name: 'sandbox.99test.com',
                //     }),
                //     name: 'test-service',
                // }
            }
        );

        fargate.targetGroup.configureHealthCheck({
            path: "/",
            timeout: Duration.seconds(10),
            unhealthyThresholdCount: 3,
        });

        fargate.targetGroup.setAttribute(
            "deregistration_delay.timeout_seconds",
            "30"
          );

    }
}
