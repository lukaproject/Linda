# /bin/bash
# pls run this command in root dir

number_of_agents=3

while getopts n: flag
do
    case "${flag}" in
        n) number_of_agents=${OPTARG};;
    esac
done

dir="$test_script_dir"
echo "Generate config"
python3 $dir/config/config-generate.py --env dev --agentcentral

echo "Number of agents: $number_of_agents"

$dir/setup-agentcentral.sh
$dir/setup-agent-cluster.sh $number_of_agents
