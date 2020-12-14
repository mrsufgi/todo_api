#!/bin/sh
# wait-for-pg.sh

set -e

TIMEOUT=15
DELAY=10

host="$1"
shift
shift
cmd="$@"

sleep $DELAY

for i in `seq $TIMEOUT` ; do
  # TODO: if host is empty, remove endpoint-url 
  if !</dev/tcp/db/5432 > /dev/null 2>&1; then 
    >&2 echo "Kinesis is up - executing command"
    exec $cmd
    exit 0
  else 
    >&2 echo "Kinesis is unavailable - sleeping"
    sleep 1
  fi
done

echo "Operation timed out" >&2
exit 1