#!/bin/bash

cd /opt/golang/src/mydomain.org/test/apextest/
apex deploy -p <awsprofile> sample_go_lambda_proxy
