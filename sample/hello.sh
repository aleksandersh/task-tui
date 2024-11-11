#!/usr/bin/env bash

echo "Hi! What is your name?"
read -r name || exit 1
if [ -n "$name" ]; then
    echo "Nice to meet you, $name"
else
    echo "Nice to meet you"
fi
