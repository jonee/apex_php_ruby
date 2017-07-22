#!/bin/bash

cd /opt/golang/src/mydomain.org/test/apextest/
apex logs -p <awsprofile> -f sample_go_lambda_proxy
