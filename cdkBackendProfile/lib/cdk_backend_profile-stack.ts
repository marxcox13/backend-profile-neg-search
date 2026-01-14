import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import * as apigw from 'aws-cdk-lib/aws-apigatewayv2';
import * as integrations from 'aws-cdk-lib/aws-apigatewayv2-integrations';
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class CdkBackendProfileStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const goLambda = new lambda.Function(this, 'UE1DESABACKENDPROFILENEG', {
      runtime: lambda.Runtime.PROVIDED_AL2023,
      handler: 'bootstrap',
      architecture: lambda.Architecture.ARM_64,
      code: lambda.Code.fromAsset('../',{
        bundling: {
          image: lambda.Runtime.PROVIDED_AL2023.bundlingImage,
          environment: {
            GOCACHE: '/tmp/go-cache',
            GOMODCACHE: '/tmp/go-mod',
          },
          command: [
            'bash', '-c',
            `
            mkdir -p /tmp/go-cache /tmp/go-mod &&
            GOOS=linux GOARCH=arm64 go build -buildvcs=false -o /asset-output/bootstrap
            `
          ],
        }
      })
    });

     const api = new apigw.HttpApi(this, 'HttpApi');

    api.addRoutes({
      path: '/{proxy+}',
      methods: [apigw.HttpMethod.ANY],
      integration: new integrations.HttpLambdaIntegration(
        'LambdaIntegration',
        goLambda
      ),
    });
  }
}
