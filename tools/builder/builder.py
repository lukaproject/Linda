#!/bin/python3
# python 3.10+

import subprocess
import argparse

parser = argparse.ArgumentParser(description="builder input args")
parser.add_argument(
    "--setup",
    action=argparse.BooleanOptionalAction,
    default=False,
    help="setup develop environment",
)
parser.add_argument(
    "--agent", action=argparse.BooleanOptionalAction, default=False, help="build agent"
)
parser.add_argument(
    "--agentcentral",
    action=argparse.BooleanOptionalAction,
    default=False,
    help="build agentcentral",
)
parser.add_argument(
    "--fileservicefe",
    action=argparse.BooleanOptionalAction,
    default=False,
    help="build fileservicefe",
)
parser.add_argument(
    "--buildtool",
    type=str,
    choices=["buildx", "docker-buildx", "docker"],
    default="docker-buildx",
    help="build tool, default is 'docker-buildx'",
)
args = parser.parse_args()


def _build_base_image():
    print("build base image")
    subprocess.run(
        args=["./tools/builder/build_base_image.sh", "goproxy.cn"],
    )


def setup():
    _build_base_image()


def build_agent():
    # cleanup
    subprocess.run(
        args=[
            "docker",
            "rmi",
            "linda-agent:latest",
        ],
    )
    subprocess.run(
        args=[
            "docker",
            "buildx",
            "build",
            "-f",
            "tools/dockerimages/agent/Dockerfile.agent",
            "-t",
            "linda-agent",
            ".",
        ],
    )


def build_agentcentral():
    # cleanup
    subprocess.run(
        args=[
            "docker",
            "rmi",
            "linda-agentcentral:latest",
        ],
    )
    subprocess.run(
        args=[
            "docker",
            "buildx",
            "build",
            "-f",
            "tools/dockerimages/services/agentcentral/Dockerfile.agentcentral",
            "-t",
            "linda-agentcentral",
            ".",
        ],
    )


def build_fileservicefe():
    # cleanup
    subprocess.run(
        args=[
            "docker",
            "rmi",
            "linda-fileservicefe:latest",
        ],
    )
    subprocess.run(
        args=[
            "docker",
            "buildx",
            "build",
            "-f",
            "tools/dockerimages/services/fileservicefe/Dockerfile.fileservicefe",
            "-t",
            "linda-fileservicefe",
            ".",
        ],
    )


if __name__ == "__main__":
    print("agent", args.agent)
    print("agent central", args.agentcentral)
    print("fileservicefe", args.fileservicefe)
    if args.setup:
        setup()
    if args.agent:
        build_agent()
    if args.agentcentral:
        build_agentcentral()
    if args.fileservicefe:
        build_fileservicefe()
