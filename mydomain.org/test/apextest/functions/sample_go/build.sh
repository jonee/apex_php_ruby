#!/bin/bash

GOPATH=/opt/golang/src/mydomain.org/test/apextest/vendor:/opt/golang GO15VENDOREXPERIMENT=0 go build mydomain.org/test/apextest/functions/sample_go/
