#!/bin/sh
# wait-for-pg.sh

set -e

TIMEOUT=15
DELAY=10

host="${1:-db}"
shift
shift
cmd="$@"

sleep $DELAY

for i in `seq $TIMEOUT` ; do
  # change to psql
  if /dev/tcp/${host}/35432 >/dev/null 2>&1; then 
    >&2 echo "pg is up - executing command"
    exec $cmd
    exit 0
  else 
    >&2 echo "pg is unavailable - sleeping"
    sleep 1
  fi
done

echo "Operation timed out" >&2
exit 1