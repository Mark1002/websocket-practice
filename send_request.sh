#! /bin/bash

# -i: Include the HTTP response headers in the output.
# -N: Prevent curl from buffering the output.
# -H: Specify an HTTP header.


set -euo pipefail

curl -i -N -H "Connection: Upgrade" \
    -H "Upgrade: websocket" \
    -H "Host: 127.0.0.1" \
    -H "Origin: http://127.0.0.1" \
    -H "Sec-WebSocket-Key: x3JJHMbDL1EzLkh9GBhXDw==" \
    -H "Sec-WebSocket-Version: 13" \
    http://127.0.0.1:8080/ws
