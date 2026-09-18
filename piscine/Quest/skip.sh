#!/bin/bash
ls -l | tail -n +2 | awk 'NR % 2 == 1'