#!/usr/bin/env sh
# Validates every package of this source: manifests, dependencies, instructions, skills, migrations.
set -eu
cd "$(dirname "$0")/.."
exec mise exec -- go run ./scripts/check "$@"
