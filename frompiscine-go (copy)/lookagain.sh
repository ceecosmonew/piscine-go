#!/bin/bash

# Find all .sh files, print their names without .sh, sort descending
find . -type f -name "*.sh" -printf "%f\n" | sed 's/\.sh$//' | sort -r
