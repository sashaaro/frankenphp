<?php

require_once __DIR__ . '/_executor.php';
require "vendor/autoload.php";

use GuzzleHttp\Promise\PromiseInterface;

return function () {
    echo "print frankenphp_client_send_request return<br/>";
    echo frankenphp_client_send_request(
        "GET /headers HTTP/2.0
Host: httpbin.org\r\n\r\n"
    );
    echo "<br/>";
    echo "<br/>";

    $guzzle = new \GuzzleHttp\Client([
        'handler' => $stack = \GuzzleHttp\HandlerStack::create(new \App\FrankenGuzzleHandler()),
    ]);

    $request = new \GuzzleHttp\Psr7\Request('GET', 'https://httpbin.org/headers');
    $request = $request->withHeader('Content-Length', $request->getBody()->getSize());


    echo "print guzzle send return<br/>";
    $response = $guzzle->send($request);
    echo "<pre>";
    var_dump($response);
    echo "<pre>";
};
