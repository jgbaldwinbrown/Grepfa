#!/bin/bash
set -e

grepfa 'ATG.{,1000}(TAG|TAA|TGA)|(TCA|TTA|CTA).{,1000}CAT' "$@"
