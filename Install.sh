#!/bin/sh
set -ev

## Install dependencies
goinstall github.com/kless/godef-ioctl/ioctl

## Build
cd term; make install

## Clean
make clean

## Install succeeded!

