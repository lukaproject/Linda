root_dir=$(pwd)
export env="test"
export test_script_dir="$root_dir/tools/testscripts/"

alias gonow=go$(cat tools/goversion.txt)
alias gotoroot="cd $root_dir"