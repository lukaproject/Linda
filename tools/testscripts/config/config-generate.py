#!/bin/python3
# python 3.10+

import argparse
import enum
import os.path
import pathlib
import json

base_dir = pathlib.Path(os.path.dirname(__file__))
templates_dir = base_dir / "templates"


class Env(enum.Enum):
    NONE = "none"
    DEV = "dev"
    PROD = "prod"


class Constants:
    DEFAULT_PGSQL_DSN = "host=172.17.0.1 user=dxyinme password=123456 dbname=linda port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    OUTPUT_DIR = pathlib.Path(os.getenv("root_dir")) / "output" / "devconfig"


def GetEnv(string: str) -> Env:
    if string.lower() == "dev":
        return Env.DEV
    if string.lower() == "prod":
        return Env.PROD
    return Env.NONE


parser = argparse.ArgumentParser(description="config-generate input args")
parser.add_argument("--env", type=str, default="dev", help="environments")
parser.add_argument(
    "--agentcentral",
    action=argparse.BooleanOptionalAction,
    default=False,
    help="generate agentcentral config",
)
parser.add_argument(
    "--pgsql_dsn",
    type=str,
    default="",
    help="if you have bring-you-own pgsql dsn, you can use this parameter to confirm it",
)
args = parser.parse_args()


def replace(content: str, template: str, value: str) -> str:
    return content.replace(template, value)


def prep_agentcentral_config():
    env = GetEnv(args.env)
    conf: dict = None
    content = ""
    with open(templates_dir / "agentcentral.template.json") as f:
        content = f.read()

    content = replace(
        content,
        "{{PGSQL_DSN}}",
        Constants.DEFAULT_PGSQL_DSN,
    )
    content = replace(content, '"{{Redis.Addrs}}"', '"redis:6379"')
    content = replace(content, "{{Redis.Password}}", "123456")

    conf = json.loads(content)
    print(f"env = {env}")
    print(f"config = \n{json.dumps(conf, indent=4)}")
    print(f"output to {Constants.OUTPUT_DIR}")
    configfile = Constants.OUTPUT_DIR / "agentcentral.json"
    Constants.OUTPUT_DIR.mkdir(parents=True, exist_ok=True)
    configfile.write_text(json.dumps(conf, indent=4))


if __name__ == "__main__":
    if args.agentcentral:
        prep_agentcentral_config()
