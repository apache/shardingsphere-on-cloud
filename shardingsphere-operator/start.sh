#!/bin/sh

arch=$(uname -m)
case $arch in
  x86_64)
    echo "arch: "$arch
    exec /operator-bin-amd64 $@
    ;;
  aarch64)
    echo "arch: "$arch
    exec /operator-bin-arm64 $@
    ;;
  *)
    echo "Unsupported architecture: $arch"
    exit 1
    ;;
esac
