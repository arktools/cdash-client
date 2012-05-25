#!/bin/bash
path=`dirname $0`
ctest -j4 -S $path/machine.ctest -VV
