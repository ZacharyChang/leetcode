#!/bin/bash -e

GIT_REPO="${GIT_REPO:-"https://github.com/zacharychang/leetcode"}"
GIT_REPO_PATH="${GIT_REPO_PATH:-"/leetcode/code"}"
GIT_BRANCH="${GIT_BRANCH:-"master"}"
GIT_USERNAME="${GIT_USERNAME:-"Zachary"}"
GIT_EMAIL="${GIT_EMAIL:-"zacharychang.zc@gmail.com"}"

LEETCODE_USERNAME="${LEETCODE_USERNAME}"
LEETCODE_PASSWORD="${LEETCODE_PASSWORD}"

[ -z "$DRY_RUN" ] && git config --global user.name "$GIT_USERNAME"
[ -z "$DRY_RUN" ] && git config --global user.email "$GIT_EMAIL"

rm -rf "$GIT_REPO_PATH" && mkdir -p "$GIT_REPO_PATH" && git clone "$GIT_REPO" "$GIT_REPO_PATH" --depth 1
cd "$GIT_REPO_PATH"
export PYTHONPATH=$GIT_REPO_PATH
nohup python3 -u hack/readme/server.py &
sleep 5s
client --username "${LEETCODE_USERNAME}" --password "${LEETCODE_PASSWORD}"

[ -z "$DRY_RUN" ] || exit

git add .

if ! git diff-index --quiet HEAD --; then
    git commit -m "doc: update readme on $(date +'%Y-%m-%d %H:%M:%S %Z')"
    git push origin "${GIT_BRANCH}"
else
    echo "No change"
fi
