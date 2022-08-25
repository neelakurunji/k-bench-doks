#!/bin/bash

doctl kubernetes cluster kubeconfig save $1 --access-token $2 --api-url $3