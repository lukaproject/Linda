# /bin/bash
# pls run this command in root dir

dir="$test_script_dir"


$dir/down-agent-cluster.sh
sleep 3
$dir/down-agentcentral.sh
