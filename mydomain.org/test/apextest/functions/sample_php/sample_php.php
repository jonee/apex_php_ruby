<?php
    $body = '';
    while (FALSE !== ($line = fgets(STDIN))) {
		$body = $body . $line;
    }

	fwrite(STDERR, $body . " \n"); // use STDOUT for your returns, STDERR for your logs

    $event = json_decode($body, true);


	$ret = array();
	$ret["hi"] = "hello";

	fwrite(STDOUT, json_encode($ret)); // use STDOUT for your returns, STDERR for your logs



