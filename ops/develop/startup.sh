#!/bin/bash

echo "Starting app..."

exec ior -listen ":80" -binary run-with-dlv "$@"