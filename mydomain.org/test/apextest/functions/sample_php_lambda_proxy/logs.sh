#!/bin/bash

cd /opt/golang/src/mydomain.org/test/apextest/
apex logs -p <awsprofile> -f sample_php_lambda_proxy
