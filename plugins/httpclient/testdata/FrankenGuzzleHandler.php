<?php

namespace App;

use GuzzleHttp\Promise\PromiseInterface;
use Psr\Http\Message\RequestInterface;

class FrankenGuzzleHandler {
    public function __invoke(RequestInterface $request, array $options): PromiseInterface
    {
        $message = $this->transformRequestHeadersToString($request);
        if ($request->getBody()->isSeekable()) {
            $request->getBody()->rewind();
        }
        $content = $request->getBody()->getContents();

        $rawRequest = $message . $content;
        $rawResponse = frankenphp_client_send_request($rawRequest);

        return new \GuzzleHttp\Promise\FulfilledPromise($this->createResponse($rawResponse));
    }

    private function transformRequestHeadersToString(\Psr\Http\Message\RequestInterface $request): string
    {
        $message = vsprintf('%s %s HTTP/%s', [
                strtoupper($request->getMethod()),
                $request->getRequestTarget(),
                $request->getProtocolVersion(),
            ]) . "\r\n";

        foreach ($request->getHeaders() as $name => $values) {
            $message .= $name . ': ' . implode(', ', $values) . "\r\n";
        }

        $message .= "\r\n";

        return $message;
    }

    private function createResponse(string $rawRequest): \Psr\Http\Message\ResponseInterface
    {
        $headers = [];
        foreach (explode("\n", $rawRequest) as $line) {
            if ('' === rtrim($line)) {
                break;
            }
            $headers[] = trim($line);
        }

        $header = array_shift($headers);
        $parts = null !== $header ? explode(' ', $header, 3) : [];

        if (count($parts) <= 1) {
            throw new \InvalidArgumentException('Cannot read the response');
        }
        $reason = null;

        $protocol = substr($parts[0], -3);
        $status = $parts[1];

        if (isset($parts[2])) {
            $reason = $parts[2];
        }

        $responseHeaders = [];

        foreach ($headers as $header) {
            $headerParts = explode(':', $header, 2);

            if (!array_key_exists(trim($headerParts[0]), $responseHeaders)) {
                $responseHeaders[trim($headerParts[0])] = [];
            }

            $responseHeaders[trim($headerParts[0])][] = isset($headerParts[1])
                ? trim($headerParts[1])
                : '';
        }
        $response = new \GuzzleHttp\Psr7\Response((int)$status, $responseHeaders, null, $protocol, $reason);

        $body = substr($rawRequest, strpos($rawRequest, "\r\n\r\n") + 4);
        $response = $response->withBody(\GuzzleHttp\Psr7\Utils::streamFor($body));

        return $response;
    }
};
