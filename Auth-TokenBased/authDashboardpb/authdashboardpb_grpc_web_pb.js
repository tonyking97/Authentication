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


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authdashboardpb.UserDetailsRequest,
 *   !proto.authdashboardpb.UserDetailsResponse>}
 */
const methodInfo_AuthDashboardService_UserDetails = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authdashboardpb.UserDetailsResponse,
  /** @param {!proto.authdashboardpb.UserDetailsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authdashboardpb.UserDetailsResponse.deserializeBinary
);


/**
 * @param {!proto.authdashboardpb.UserDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authdashboardpb.UserDetailsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authdashboardpb.UserDetailsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authdashboardpb.AuthDashboardServiceClient.prototype.userDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/UserDetails',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_UserDetails,
      callback);
};


/**
 * @param {!proto.authdashboardpb.UserDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authdashboardpb.UserDetailsResponse>}
 *     A native promise that resolves to the response
 */
proto.authdashboardpb.AuthDashboardServicePromiseClient.prototype.userDetails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/UserDetails',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_UserDetails);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authdashboardpb.CurrentSessionDetailsRequest,
 *   !proto.authdashboardpb.CurrentSessionDetailsResponse>}
 */
const methodInfo_AuthDashboardService_CurrentSessionDetails = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authdashboardpb.CurrentSessionDetailsResponse,
  /** @param {!proto.authdashboardpb.CurrentSessionDetailsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authdashboardpb.CurrentSessionDetailsResponse.deserializeBinary
);


/**
 * @param {!proto.authdashboardpb.CurrentSessionDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authdashboardpb.CurrentSessionDetailsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authdashboardpb.CurrentSessionDetailsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authdashboardpb.AuthDashboardServiceClient.prototype.currentSessionDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/CurrentSessionDetails',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_CurrentSessionDetails,
      callback);
};


/**
 * @param {!proto.authdashboardpb.CurrentSessionDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authdashboardpb.CurrentSessionDetailsResponse>}
 *     A native promise that resolves to the response
 */
proto.authdashboardpb.AuthDashboardServicePromiseClient.prototype.currentSessionDetails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/CurrentSessionDetails',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_CurrentSessionDetails);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authdashboardpb.SessionDetailsRequest,
 *   !proto.authdashboardpb.SessionDetailsResponse>}
 */
const methodInfo_AuthDashboardService_SessionDetails = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authdashboardpb.SessionDetailsResponse,
  /** @param {!proto.authdashboardpb.SessionDetailsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authdashboardpb.SessionDetailsResponse.deserializeBinary
);


/**
 * @param {!proto.authdashboardpb.SessionDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authdashboardpb.SessionDetailsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authdashboardpb.SessionDetailsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authdashboardpb.AuthDashboardServiceClient.prototype.sessionDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/SessionDetails',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_SessionDetails,
      callback);
};


/**
 * @param {!proto.authdashboardpb.SessionDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authdashboardpb.SessionDetailsResponse>}
 *     A native promise that resolves to the response
 */
proto.authdashboardpb.AuthDashboardServicePromiseClient.prototype.sessionDetails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/SessionDetails',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_SessionDetails);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authdashboardpb.LogoutRequest,
 *   !proto.authdashboardpb.LogoutResponse>}
 */
const methodInfo_AuthDashboardService_Logout = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authdashboardpb.LogoutResponse,
  /** @param {!proto.authdashboardpb.LogoutRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authdashboardpb.LogoutResponse.deserializeBinary
);


/**
 * @param {!proto.authdashboardpb.LogoutRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authdashboardpb.LogoutResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authdashboardpb.LogoutResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authdashboardpb.AuthDashboardServiceClient.prototype.logout =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/Logout',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_Logout,
      callback);
};


/**
 * @param {!proto.authdashboardpb.LogoutRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authdashboardpb.LogoutResponse>}
 *     A native promise that resolves to the response
 */
proto.authdashboardpb.AuthDashboardServicePromiseClient.prototype.logout =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authdashboardpb.AuthDashboardService/Logout',
      request,
      metadata || {},
      methodInfo_AuthDashboardService_Logout);
};


module.exports = proto.authdashboardpb;

