#!/bin/bash

doctl kubernetes cluster create $1 --1-clicks $2 --count $3 --region $4  --size $5 --access-token $6 --api-url $7