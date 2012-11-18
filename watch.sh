#!/bin/bash
while /bin/true; do echo -n `date +%T`" " ;curl -s http://localhost:8080/debug/vars | ./debug.pl ; sleep 1; done

