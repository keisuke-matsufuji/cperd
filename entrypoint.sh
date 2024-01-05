#!/bin/bash

json_escape () {
    printf '%s' "$1" | python3 -c 'import json,sys; print(json.dumps(sys.stdin.read().rstrip("\n")))'
}

pwd1=$(pwd)
echo "pwd1 $pwd1"
cd "$GITHUB_WORKSPACE"
pwd2=$(pwd)
echo "pwd2 $pwd2"
ls=$(ls)
echo "ls $ls"
gitversion=$(git version)
echo "gitversion $gitversion"


TARGET_DIR="."
if [ "$RUN_LOCAL" = "true" ]; then
  TARGET_DIR="./testdata"
  GITHUB_REF="refs/heads/main"
#   REVIEWDOG_COMMAND="reviewdog -efm=\"%f:%l:%c: %m\" -diff=\"git diff ${GITHUB_REF}\""
  REVIEWDOG_COMMAND="reviewdog -efm=\"%f:%l:%c: %m\" -filter-mode=nofilter"
else
  REVIEWDOG_COMMAND="reviewdog -efm=\"%f:%l:%c: %m\" -reporter=github-pr-review"
fi

# mainブランチとの差分ファイルを取得
diff_files=$(git diff --name-only main)
echo "diff_files:  $diff_files"

# 差分ファイルの中からphp.iniと.htaccessを見つけ出す
files=""
for file in $diff_files
do
    if [[ $file == *'php.ini'* ]] || [[ $file == *'.htaccess'* ]]; then
        files+="$file "
    fi
done

# 各ファイルから必要な情報を取得し、JSONに変換
for file in $files
do
    if [[ $file == *'php.ini'* ]]; then
        php_ini_info=$(awk '/error_reporting/{print NR, index($0, "error_reporting"), $0}' "$file")
        read -r line column value <<< "$(echo "$php_ini_info")"
        json_data=("{\"file\":$(json_escape "$file"),\"line\":$line,\"column\":$column,\"value\":$(json_escape "$value")}")
    elif [[ $file == *'.htaccess'* ]]; then
        htaccess_info=$(awk '/php_value error_reporting/{print NR, index($0, "php_value error_reporting"), $0}' "$file")
        read -r line column value <<< "$(echo "$htaccess_info")"
        json_data=("{\"file\":$(json_escape "$file"),\"line\":$line,\"column\":$column,\"value\":$(json_escape "$value")}")
    fi
    OUTPUT=$(go run main.go "$json_data")
    echo "OUTPUT :: $OUTPUT"
    echo "$OUTPUT" | jq -r '. | "\(.file):\(.line):\(.column): \(.value)"' | eval ${REVIEWDOG_COMMAND}
done
