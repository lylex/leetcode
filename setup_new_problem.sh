#!/bin/bash -e

usage="Usage: $(basename "$0") [-h] [-n <integer>] [-t <string>]

Script to initialize a leetcode problem in this project.

Options:

  -n the leetcode problem id
  -t the leetcode problem title
"

LEETCODE_ID=""
LEETCODE_TITLE=""

while getopts :hn:t: OPTION; do
    case $OPTION in
        h) echo "$usage"
           exit
           ;;
        n) LEETCODE_ID=$OPTARG
           ;;
        t) LEETCODE_TITLE=$OPTARG
           ;;
        :) printf "missing argument for -%s\\n" "$OPTARG" >&2
           echo "$usage" >&2
           exit 1
           ;;
        \?) printf "illegal option: -%s\\n" "$OPTARG" >&2
            echo "$usage" >&2
            exit 1
            ;;
    esac
done

if [[ ! "$LEETCODE_ID" =~ ^[0-9]+$ ]]; then
    echo "valid Leetcode ID required"
    exit 1
fi

if [ -z "$LEETCODE_TITLE" ]; then
    echo "valid leetcode title required"
    exit 1
fi

work_dir=${LEETCODE_ID}_$(echo $LEETCODE_TITLE | sed -E 's/[\ .-]/_/g')
mkdir -p $work_dir

cat > $work_dir/${LEETCODE_ID}.md << EOF
# Question:

**Note**:

**Example**:
\`\`\`
here goes the example
\`\`\`

# Analysis

# Tips

# Code

EOF

cat > $work_dir/main.go << EOF
package main

import (
        "fmt"
)

func main() {
        fmt.Println("hello world!")
}
EOF

cat >> README.md << EOF
* [${LEETCODE_ID}. $(echo ${LEETCODE_TITLE} | sed 's/[-_.]/ /g; s/\<./\U&/g')]($work_dir/${LEETCODE_ID}.md)
EOF
