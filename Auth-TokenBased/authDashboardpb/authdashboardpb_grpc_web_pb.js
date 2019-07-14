/**
 * @fileoverview gRPC-Web generated client stub for authdashboardpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.authdashboardpb = require('./authdashboardpb_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.authdashboardpb.AuthDashboardServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.authdashboardpb.AuthDashboardServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authdashboardpb.CheckTokenRequest,
 *   !proto.authdashboardpb.CheckTokenResponse>}
 */
const methodInfo_AuthDashboardService_CheckToken = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authdashboardpb.CheckTokenResponse,
  /** @param {!proto.authdashboardpb.CheckTokenRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authdashboardpb.CheckTokenResponse.deserializeBinary
);


/**
 * @param {!proto.authdashboardpb.CheckTokenRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authdashboardpb.CheckTokenResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authdashboardpb.CheckTokenResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authdashboardpb.AuthDashboardServiceClient.prototype.checkToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/CheckToken',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_CheckToken,
      callback);
};


/**
 * @param {!proto.authdashboardpb.CheckTokenRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authdashboardpb.CheckTokenResponse>}
 *     A native promise that resolves to the response
 */
proto.authdashboardpb.AuthDashboardServicePromiseClient.prototype.checkToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/CheckToken',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_CheckToken);
};


module.exports = proto.authdashboardpb;

