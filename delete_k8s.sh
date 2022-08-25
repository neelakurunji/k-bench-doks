#!/bin/bash

doctl kubernetes cluster delete --force $1 --access-token $2 --api-url $3