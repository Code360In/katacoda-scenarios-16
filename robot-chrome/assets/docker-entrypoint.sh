#!/bin/bash

# turn on bash's job control
set -m

# start tinit to control all child
/tini -g -s -- /usr/sbin/sshd -D 
