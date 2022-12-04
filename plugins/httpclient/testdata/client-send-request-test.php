<?php

require_once __DIR__ . '/_executor.php';
require "vendor/autoload.php";

use GuzzleHttp\Promise\PromiseInterface;

return function () {
    $guzzle = new \GuzzleHttp\Client([
        'handler' => $stack = \GuzzleHttp\HandlerStack::create(new \App\FrankenGuzzleHandler()),
    ]);

    $address = $_GET['address'];
    echo $address . '<br/>';
    $request = new \GuzzleHttp\Psr7\Request('GET', $address);
    $request = $request->withHeader('Content-Length', $request->getBody()->getSize());

    echo "response from external service requested via http2 without tls<br/>";
    $response = $guzzle->send($request);
    echo "<pre>";
    var_dump($response->getBody()->getContents());
    echo "<pre>";
};
