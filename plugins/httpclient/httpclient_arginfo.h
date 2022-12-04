#ifndef DEFINITIONS_H
#define DEFINITIONS_H

#include <stdint.h>
#include <Zend/zend_types.h>

#endif

ZEND_BEGIN_ARG_WITH_RETURN_TYPE_INFO_EX(arginfo_frankenphp_client_send_request, 0, 1, IS_STRING, 0)
	ZEND_ARG_TYPE_INFO(0, request, IS_STRING, 0)
ZEND_END_ARG_INFO()
ZEND_FUNCTION(frankenphp_client_send_request);

static const zend_function_entry ext_http_functions[] = {
    ZEND_FE(frankenphp_client_send_request, arginfo_frankenphp_client_send_request)
	ZEND_FE_END
};
