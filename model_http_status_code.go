/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// HttpStatusCode
type HttpStatusCode string

// List of HttpStatusCode
const (
	HTTPSTATUSCODE_CONTINUE                        HttpStatusCode = "Continue"
	HTTPSTATUSCODE_SWITCHING_PROTOCOLS             HttpStatusCode = "SwitchingProtocols"
	HTTPSTATUSCODE_PROCESSING                      HttpStatusCode = "Processing"
	HTTPSTATUSCODE_EARLY_HINTS                     HttpStatusCode = "EarlyHints"
	HTTPSTATUSCODE_OK                              HttpStatusCode = "OK"
	HTTPSTATUSCODE_CREATED                         HttpStatusCode = "Created"
	HTTPSTATUSCODE_ACCEPTED                        HttpStatusCode = "Accepted"
	HTTPSTATUSCODE_NON_AUTHORITATIVE_INFORMATION   HttpStatusCode = "NonAuthoritativeInformation"
	HTTPSTATUSCODE_NO_CONTENT                      HttpStatusCode = "NoContent"
	HTTPSTATUSCODE_RESET_CONTENT                   HttpStatusCode = "ResetContent"
	HTTPSTATUSCODE_PARTIAL_CONTENT                 HttpStatusCode = "PartialContent"
	HTTPSTATUSCODE_MULTI_STATUS                    HttpStatusCode = "MultiStatus"
	HTTPSTATUSCODE_ALREADY_REPORTED                HttpStatusCode = "AlreadyReported"
	HTTPSTATUSCODE_IM_USED                         HttpStatusCode = "IMUsed"
	HTTPSTATUSCODE_AMBIGUOUS                       HttpStatusCode = "Ambiguous"
	HTTPSTATUSCODE_MOVED                           HttpStatusCode = "Moved"
	HTTPSTATUSCODE_REDIRECT                        HttpStatusCode = "Redirect"
	HTTPSTATUSCODE_REDIRECT_METHOD                 HttpStatusCode = "RedirectMethod"
	HTTPSTATUSCODE_NOT_MODIFIED                    HttpStatusCode = "NotModified"
	HTTPSTATUSCODE_USE_PROXY                       HttpStatusCode = "UseProxy"
	HTTPSTATUSCODE_UNUSED                          HttpStatusCode = "Unused"
	HTTPSTATUSCODE_TEMPORARY_REDIRECT              HttpStatusCode = "TemporaryRedirect"
	HTTPSTATUSCODE_PERMANENT_REDIRECT              HttpStatusCode = "PermanentRedirect"
	HTTPSTATUSCODE_BAD_REQUEST                     HttpStatusCode = "BadRequest"
	HTTPSTATUSCODE_UNAUTHORIZED                    HttpStatusCode = "Unauthorized"
	HTTPSTATUSCODE_PAYMENT_REQUIRED                HttpStatusCode = "PaymentRequired"
	HTTPSTATUSCODE_FORBIDDEN                       HttpStatusCode = "Forbidden"
	HTTPSTATUSCODE_NOT_FOUND                       HttpStatusCode = "NotFound"
	HTTPSTATUSCODE_METHOD_NOT_ALLOWED              HttpStatusCode = "MethodNotAllowed"
	HTTPSTATUSCODE_NOT_ACCEPTABLE                  HttpStatusCode = "NotAcceptable"
	HTTPSTATUSCODE_PROXY_AUTHENTICATION_REQUIRED   HttpStatusCode = "ProxyAuthenticationRequired"
	HTTPSTATUSCODE_REQUEST_TIMEOUT                 HttpStatusCode = "RequestTimeout"
	HTTPSTATUSCODE_CONFLICT                        HttpStatusCode = "Conflict"
	HTTPSTATUSCODE_GONE                            HttpStatusCode = "Gone"
	HTTPSTATUSCODE_LENGTH_REQUIRED                 HttpStatusCode = "LengthRequired"
	HTTPSTATUSCODE_PRECONDITION_FAILED             HttpStatusCode = "PreconditionFailed"
	HTTPSTATUSCODE_REQUEST_ENTITY_TOO_LARGE        HttpStatusCode = "RequestEntityTooLarge"
	HTTPSTATUSCODE_REQUEST_URI_TOO_LONG            HttpStatusCode = "RequestUriTooLong"
	HTTPSTATUSCODE_UNSUPPORTED_MEDIA_TYPE          HttpStatusCode = "UnsupportedMediaType"
	HTTPSTATUSCODE_REQUESTED_RANGE_NOT_SATISFIABLE HttpStatusCode = "RequestedRangeNotSatisfiable"
	HTTPSTATUSCODE_EXPECTATION_FAILED              HttpStatusCode = "ExpectationFailed"
	HTTPSTATUSCODE_MISDIRECTED_REQUEST             HttpStatusCode = "MisdirectedRequest"
	HTTPSTATUSCODE_UNPROCESSABLE_ENTITY            HttpStatusCode = "UnprocessableEntity"
	HTTPSTATUSCODE_LOCKED                          HttpStatusCode = "Locked"
	HTTPSTATUSCODE_FAILED_DEPENDENCY               HttpStatusCode = "FailedDependency"
	HTTPSTATUSCODE_UPGRADE_REQUIRED                HttpStatusCode = "UpgradeRequired"
	HTTPSTATUSCODE_PRECONDITION_REQUIRED           HttpStatusCode = "PreconditionRequired"
	HTTPSTATUSCODE_TOO_MANY_REQUESTS               HttpStatusCode = "TooManyRequests"
	HTTPSTATUSCODE_REQUEST_HEADER_FIELDS_TOO_LARGE HttpStatusCode = "RequestHeaderFieldsTooLarge"
	HTTPSTATUSCODE_UNAVAILABLE_FOR_LEGAL_REASONS   HttpStatusCode = "UnavailableForLegalReasons"
	HTTPSTATUSCODE_INTERNAL_SERVER_ERROR           HttpStatusCode = "InternalServerError"
	HTTPSTATUSCODE_NOT_IMPLEMENTED                 HttpStatusCode = "NotImplemented"
	HTTPSTATUSCODE_BAD_GATEWAY                     HttpStatusCode = "BadGateway"
	HTTPSTATUSCODE_SERVICE_UNAVAILABLE             HttpStatusCode = "ServiceUnavailable"
	HTTPSTATUSCODE_GATEWAY_TIMEOUT                 HttpStatusCode = "GatewayTimeout"
	HTTPSTATUSCODE_HTTP_VERSION_NOT_SUPPORTED      HttpStatusCode = "HttpVersionNotSupported"
	HTTPSTATUSCODE_VARIANT_ALSO_NEGOTIATES         HttpStatusCode = "VariantAlsoNegotiates"
	HTTPSTATUSCODE_INSUFFICIENT_STORAGE            HttpStatusCode = "InsufficientStorage"
	HTTPSTATUSCODE_LOOP_DETECTED                   HttpStatusCode = "LoopDetected"
	HTTPSTATUSCODE_NOT_EXTENDED                    HttpStatusCode = "NotExtended"
	HTTPSTATUSCODE_NETWORK_AUTHENTICATION_REQUIRED HttpStatusCode = "NetworkAuthenticationRequired"
)
