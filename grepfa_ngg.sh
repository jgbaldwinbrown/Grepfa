#!/bin/bash
set -e

fagrep '.{0,50}((.[Gg][Gg]){5,}|([Cc][Cc].){5,}).{0,50}' "$@"
