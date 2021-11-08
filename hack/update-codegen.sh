#!/bin/bash
#
#Copyright 2019 hatech Authors
#

set -e

if [[ -z $GOPATH ]]; then
	echo "Setting GOPATH to ~/go"
	GOPATH=~/go
fi

if [[ ! -d "${GOPATH}/src/k8s.io/code-generator" ]]; then
  echo ">>>>>> k8s.io/code-generator of v0.15.12 missing from GOPATH"
  echo ">>>>>> Cloning https://github.com/kubernetes/code-generator with tag v0.15.12 under '${GOPATH}/src/k8s.io'"
  git clone -b v0.15.12 https://github.com/kubernetes/code-generator ${GOPATH}/src/k8s.io/code-generator
fi
# Switching to v0.15.12 if already cloned
git --git-dir=${GOPATH}/src/k8s.io/code-generator/.git  --work-tree=${GOPATH}/src/k8s.io/code-generator checkout v0.15.12

${GOPATH}/src/k8s.io/code-generator/generate-groups.sh client,lister,informer \
  github.com/vossss/cnbrchaos/pkg/client github.com/vossss/cnbrchaos/pkg/apis \
  cnbrchaos:v1alpha1

#${GOPATH}/src/k8s.io/code-generator/generate-groups.sh client,lister,informer \
#  client apis cnbrchaos:v1alpha1