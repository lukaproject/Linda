#!/bin/bash

GO=go

# For local test, to check out current local go version.
if [ $# == 1 ]; then
    GO=go$1
fi

GoVersion=`$GO version`
Root=$(pwd)
CoverageOutputPath=$Root/output/cover/
mkdir -p $CoverageOutputPath

echo "run tests in $GoVersion"

function RunModTests {
    local path=$Root/$1
    local outname=$2
    local outpath=$CoverageOutputPath/$outname.out
    echo "test in $path, outpath=$outpath"
    cd $path
    $GO test ./... -coverprofile=$outpath
    cd $Root
    echo "test finished for $path"
}

function InstallGocovmerge {
    command -v gocovmerge >/dev/null 2>&1
    if [ $? != 0 ]; then
        echo "gocovmerge not exists, installing..."
        $GO install github.com/alexfalkowski/gocovmerge@latest
    fi
}

RunModTests agent agent
RunModTests services/agentcentral agentcentral
RunModTests baselibs/multifs multifs
RunModTests baselibs/abstractions abstractions

InstallGocovmerge

gocovmerge $CoverageOutputPath/agent.out $CoverageOutputPath/agentcentral.out \
           $CoverageOutputPath/multifs.out $CoverageOutputPath/abstractions.out \
           > $CoverageOutputPath/coverageall.out