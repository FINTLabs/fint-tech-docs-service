#!/bin/bash

message=$(<$1)
shared_key=$2

echo -n "$message" | openssl dgst -sha1 -hmac "${shared_key}"