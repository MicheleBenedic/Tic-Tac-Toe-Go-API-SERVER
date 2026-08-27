#!/usr/bin/env bash
# set=shell settings
# -euo: e=exit on error, u=unset variables as errors
# o pipefail=returns the last comand execution with a
# value different than 0 instead of the last script command
set -euo pipefail

BRANCH="${1:-main}"

while true; do
    # -rp: r=escape disabled, p=prompt
    read -rp "Messaggio di commit: " commit_msg

    # -z: z=no text submitted
    if [ -z "$commit_msg" ]; then
        echo "Il messaggio non può essere vuoto."
        continue
    fi

    read -rp "Confermi il messaggio: \"$commit_msg\"? [s/N] " confirm
    case "$confirm" in
        # input cases
        [sS]|[sS][iI]|[yY]|[yY][eE][sS])
            break
            ;;
        *)
            echo "Ok, reinserisci il messaggio."
            ;;
    esac
done

git add .
git commit -m "$commit_msg"
git push -u origin "$BRANCH"