#!/bin/bash

cd /opt/golang/src/mydomain.org/test/apextest/
apex deploy -p <awsprofile> sample_php_lambda_proxy
