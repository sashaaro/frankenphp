<?php

require_once __DIR__.'/_executor.php';

return function () {
    echo frankenphp_client_send_request("GET /headers HTTP/1.1
Host: httpbin.org");
};
