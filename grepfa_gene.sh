#!/bin/bash
set -e

fagrep 'ATG.{,1000}(TAG|TAA|TGA)|(TCA|TTA|CTA).{,1000}CAT' "$@"
