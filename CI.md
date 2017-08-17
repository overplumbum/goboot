```bash
set -xue -o pipefail

AWS_REGION=eu-west-1
PROJECT_NAME=${BUILDKITE_PIPELINE_SLUG}
REPO=xxx.dkr.ecr.eu-west-1.amazonaws.com/${PROJECT_NAME}

BUILDTAG=${BUILDKITE_PIPELINE_SLUG}-${BUILDKITE_BUILD_NUMBER}
docker build --pull --build-arg "COMMIT=`git rev-parse HEAD`" -t $BUILDTAG .
docker tag $BUILDTAG $REPO:b${BUILDKITE_BUILD_NUMBER}


aws ecr get-login --region ${AWS_REGION} > ./ecr.sh
bash ./ecr.sh


docker push $REPO:b${BUILDKITE_BUILD_NUMBER}


wget -P ~ -nc https://raw.githubusercontent.com/silinternational/ecs-deploy/3.2/ecs-deploy

RELEASE_STAGE=stage
time bash ~/ecs-deploy -r "${AWS_REGION}" -c api -n ${PROJECT_NAME}-${RELEASE_STAGE} -i $REPO:b${BUILDKITE_BUILD_NUMBER} --min 1 -t 600 --max-definitions 10
```
