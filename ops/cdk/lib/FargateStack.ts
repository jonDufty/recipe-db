import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { FargateConstruct, FargateProps } from './FargateConstruct';
import { Vpc } from 'aws-cdk-lib/aws-ec2';
import { ARecord } from 'aws-cdk-lib/aws-route53';

interface FargateStackProps extends StackProps, FargateProps {

}

export class FargateStack extends Stack {
  constructor(scope: Construct, id: string, props: FargateStackProps) {
    super(scope, id, props);

    new FargateConstruct(this, 'FargateService', {
      serviceName: props.serviceName,
      command: props.command,
      imageTag: props.imageTag
    })
  
  }
}
