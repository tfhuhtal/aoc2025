#!/bin/bash
set -e

mkdir -p out 2>/dev/null || true

case "${1:-release}" in
  release) make main ;;
  debug)   make debug ;;
  clean)   make clean ;;
  *)       echo "Usage: $0 [release|debug|clean]"; exit 1 ;;
esac
