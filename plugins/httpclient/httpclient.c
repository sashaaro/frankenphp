#ifndef DEFINITIONS_H
#define DEFINITIONS_H

#include <errno.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <php.h>
#include <ext/standard/head.h>
//#include <php_variables.h>
//#include <php_output.h>
#include <Zend/zend_alloc.h>
#include <Zend/zend_types.h>
//#include <Zend/zend_exceptions.h>
//#include <Zend/zend_interfaces.h>

#include "httpclient_arginfo.h"

//#include "_cgo_export.h"

#endif

#if defined(PHP_WIN32) && defined(ZTS)
ZEND_TSRMLS_CACHE_DEFINE()
#endif

PHP_FUNCTION(frankenphp_client_send_request) {
    zend_string *request;

    ZEND_PARSE_PARAMETERS_START(1, 1)
        Z_PARAM_STR(request)
    ZEND_PARSE_PARAMETERS_END();

    const char *response = "";//go_frankenphp_client_send_request(request->val);
    RETURN_STRING(response);
}


static zend_module_entry frankenphp_httpclient_module = {
    STANDARD_MODULE_HEADER,
    "frankenphp-http-client",
    ext_http_functions,	/* function table */
    NULL,			/* initialization */
    NULL,			/* shutdown */
    NULL,			/* request initialization */
    NULL,			/* request shutdown */
    NULL,			/* information */
    "dev",
    STANDARD_MODULE_PROPERTIES
};
