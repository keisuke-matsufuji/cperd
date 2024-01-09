#!/bin/bash

json_escape () {
    printf '%s' "$1" | python3 -c 'import json,sys; print(json.dumps(sys.stdin.read().rstrip("\n")))'
}

TARGET_DIR="."
if [ "$RUN_LOCAL" = "true" ]; then
    TARGET_DIR="./cperd/testdata"
    REVIEWDOG_COMMAND="reviewdog -efm=\"%f:%l:%c: %m\" -filter-mode=nofilter"

    target_files=$(ls -a $TARGET_DIR)
else
    git config --global --add safe.directory "$PWD"
    REVIEWDOG_COMMAND="reviewdog -efm=\"%f:%l:%c: %m\" -reporter=github-pr-review"

    # Get the default branch
    default_branch=$(git remote show origin | sed -n '/HEAD branch/s/.*: //p')
    echo "default_branch $default_branch"

    # Fetch the information of the default branch
    git fetch origin $default_branch:$default_branch

    # Get the difference files with the main branch
    target_files=$(git diff --name-only "refs/heads/$default_branch")
fi

echo "target_files:  $target_files"

# Find php.ini and .htaccess from the difference files
files=""
for file in $target_files
do
    if [[ $file == *'php.ini'* ]] || [[ $file == *'.htaccess'* ]]; then
        files+="$file "
    fi
done

# Get the necessary information from each file and convert it to JSON
for file in $files
do
    target=$TARGET_DIR/$file
    if [[ $file == *'php.ini'* ]]; then
        php_ini_info=$(awk '/error_reporting/{print NR, index($0, "error_reporting"), $0}' "$target")
        read -r line column value <<< "$(echo "$php_ini_info")"
        json_data=("{\"file\":$(json_escape "$file"),\"line\":$line,\"column\":$column,\"value\":$(json_escape "$value")}")
    elif [[ $file == *'.htaccess'* ]]; then
        htaccess_info=$(awk '/php_value error_reporting/{print NR, index($0, "php_value error_reporting"), $0}' "$target")
        read -r line column value <<< "$(echo "$htaccess_info")"
        json_data=("{\"file\":$(json_escape "$file"),\"line\":$line,\"column\":$column,\"value\":$(json_escape "$value")}")
    fi
    OUTPUT=$(go run /main.go "$json_data")
    echo "OUTPUT :: $OUTPUT"
    echo "$OUTPUT" | jq -r '. | "\(.file):\(.line):\(.column): \(.value)"' | eval ${REVIEWDOG_COMMAND}
done
