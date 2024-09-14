#/bin/bash
#run this bash in root_dir

docker run --rm -v \
    ${PWD}:/local swaggerapi/swagger-codegen-cli-v3:latest generate \
    -i /local/services/agentcentral/docs/swagger.yaml -l go -o /local/out/go

cp -r out/go/* baselibs/apiscall/swagger/
