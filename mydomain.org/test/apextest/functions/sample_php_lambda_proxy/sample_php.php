<?php
    $body = '';
    while (FALSE !== ($line = fgets(STDIN))) {
		$body = $body . $line;
    }

	fwrite(STDERR, $body . " \n"); // use STDOUT for your returns, STDERR for your logs

    $event = json_decode($body, true);


	/* gather environment - Input Format of a Lambda Function for Proxy Integration */
	// http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-set-up-simple-proxy.html#api-gateway-simple-proxy-for-lambda-input-format

	$resource = (string)$event["resource"];
	fwrite(STDERR, "resource: $resource \n"); // use STDOUT for your returns, STDERR for your logs

	$path = (string)$event["path"];
	fwrite(STDERR, "path: $path \n");

	$httpMethod = (string)$event["httpMethod"];
	fwrite(STDERR, "httpMethod: $httpMethod \n");

	$headers = $event["headers"];
	fwrite(STDERR, "headers: " . print_r($headers, true));

	$queryStringParameters = $event["queryStringParameters"];
	fwrite(STDERR, "queryStringParameters: " . print_r($queryStringParameters, true));

	$pathParameters = $event["pathParameters"];
	fwrite(STDERR, "pathParameters: " . print_r($pathParameters, true));

	$stageVariables = $event["stageVariables"];
	fwrite(STDERR, "stageVariables: " . print_r($stageVariables, true));

	$stagename = $stageVariables["stagename"];
	fwrite(STDERR, "stagename: $stagename \n");

	$requestContext = $event["requestContext"];
	fwrite(STDERR, "requestContext: " . print_r($requestContext, true));

	$body = (string)$event["body"];
	fwrite(STDERR, "body: $body \n");

	$isBase64Encoded = (bool)$event["isBase64Encoded"];
	fwrite(STDERR, "isBase64Encoded: $isBase64Encoded \n");


	// test mongodb connection
	// unfortunately we can only use the mongo legacy driver + php5 (not compatible with php7) until this is fixed https://jira.mongodb.org/browse/PHPC-759
	// mongo legacy driver is this http://php.net/manual/en/book.mongo.php

	// would have been great to use the new mongodb extension with a library on top of it- please see https://packagist.org/packages/mongodb/mongodb https://docs.mongodb.com/php-library/master/

	// $mongodb_uri = "mongodb://<user>:<password>@<publicip>:27017/<authdb>"; // you have to make sure the security groups have this port open
	$mongodb_uri = "mongodb://<user>:<password>@<privateip>:27017/<authdb>"; // you have to configure your vpc config in lambda

	$useDB = "db_dev";
	if ($stagename == "prod") {
		$useDB = "db_prod";
	}
	fwrite(STDERR, "useDB: $useDB \n");

	// connect
	$m = new MongoClient($mongodb_uri);

	// select a database
	$db = $m->{$useDB}; // mongo legacy seems to need auth in this db as well

	$massCollectionCol = $db->mass_collection;

	$obj = $massCollectionCol->findOne();

	fwrite(STDERR, "document obj: " . print_r($obj, true));

	$m->close(); // if working in replica set might need parameters but we might just be working on a single node config


	/* Output Format of a Lambda Function for Proxy Integration */
	// http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-set-up-simple-proxy.html#api-gateway-simple-proxy-for-lambda-output-format

	$ret = array();
	$ret["isBase64Encoded"] = false;
	$ret["statusCode"] = 200;
	$ret["headers"] = array("headerName"=>"headerValue");
	$ret["body"] = "sample_php"; // can be json string

	fwrite(STDOUT, json_encode($ret)); // use STDOUT for your returns, STDERR for your logs

	// fwrite(STDERR, phpinfo());

