# 不出以外的话
# 应该支持的python版本如下：
# 3.10, 3.11. 3.12

import subprocess
import argparse

parser = argparse.ArgumentParser(description="builder input args")
parser.add_argument("--setup", action=argparse.BooleanOptionalAction, default=False, help="setup develop environment")
parser.add_argument("--agent", action=argparse.BooleanOptionalAction, default=False, help="build agent")
parser.add_argument("--agentcentral", action=argparse.BooleanOptionalAction, default=False, help="build agentcentral")
args = parser.parse_args()

def _install_swag():
    print("install swag")


def setup():
    _install_swag()


def build_agent():
    subprocess.run(
        args=[
            "buildx",
            "build",
            "-f",
            "tools/dockerimages/agent/Dockerfile.agent",
            "-t",
            "linda-agent",
            "."
        ],
    )


if __name__ == "__main__":
    print("agent", args.agent)
    print("agent central", args.agentcentral)
    if args.agent:
        build_agent()