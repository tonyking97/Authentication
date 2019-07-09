/**
 * @fileoverview gRPC-Web generated client stub for authpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.authpb = require('./authpb_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.authpb.AuthServiceClient =
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
proto.authpb.AuthServicePromiseClient =
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
 *   !proto.authpb.LoginRequest,
 *   !proto.authpb.LoginResponse>}
 */
const methodInfo_AuthService_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authpb.LoginResponse,
  /** @param {!proto.authpb.LoginRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authpb.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.LoginResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.LoginResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authpb.AuthService/Login',
      request,
      metadata || {},
      methodInfo_AuthService_Login,
      callback);
};


/**
 * @param {!proto.authpb.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.LoginResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authpb.AuthService/Login',
      request,
      metadata || {},
      methodInfo_AuthService_Login);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.CheckUsernameRequest,
 *   !proto.authpb.CheckUsernameResponse>}
 */
const methodInfo_AuthService_CheckUsername = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authpb.CheckUsernameResponse,
  /** @param {!proto.authpb.CheckUsernameRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authpb.CheckUsernameResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.CheckUsernameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.CheckUsernameResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.CheckUsernameResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.checkUsername =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authpb.AuthService/CheckUsername',
      request,
      metadata || {},
      methodInfo_AuthService_CheckUsername,
      callback);
};


/**
 * @param {!proto.authpb.CheckUsernameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.CheckUsernameResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.checkUsername =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authpb.AuthService/CheckUsername',
      request,
      metadata || {},
      methodInfo_AuthService_CheckUsername);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.CheckUsernameAvailabilityRequest,
 *   !proto.authpb.CheckUsernameAvailabilityResponse>}
 */
const methodInfo_AuthService_CheckUsernameAvailability = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authpb.CheckUsernameAvailabilityResponse,
  /** @param {!proto.authpb.CheckUsernameAvailabilityRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authpb.CheckUsernameAvailabilityResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.CheckUsernameAvailabilityRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.CheckUsernameAvailabilityResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.CheckUsernameAvailabilityResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.checkUsernameAvailability =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authpb.AuthService/CheckUsernameAvailability',
      request,
      metadata || {},
      methodInfo_AuthService_CheckUsernameAvailability,
      callback);
};


/**
 * @param {!proto.authpb.CheckUsernameAvailabilityRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.CheckUsernameAvailabilityResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.checkUsernameAvailability =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authpb.AuthService/CheckUsernameAvailability',
      request,
      metadata || {},
      methodInfo_AuthService_CheckUsernameAvailability);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.CheckEmailAvailabilityRequest,
 *   !proto.authpb.CheckEmailAvailabilityResponse>}
 */
const methodInfo_AuthService_CheckEmailAvailability = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authpb.CheckEmailAvailabilityResponse,
  /** @param {!proto.authpb.CheckEmailAvailabilityRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authpb.CheckEmailAvailabilityResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.CheckEmailAvailabilityRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.CheckEmailAvailabilityResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.CheckEmailAvailabilityResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.checkEmailAvailability =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authpb.AuthService/CheckEmailAvailability',
      request,
      metadata || {},
      methodInfo_AuthService_CheckEmailAvailability,
      callback);
};


/**
 * @param {!proto.authpb.CheckEmailAvailabilityRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.CheckEmailAvailabilityResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.checkEmailAvailability =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authpb.AuthService/CheckEmailAvailability',
      request,
      metadata || {},
      methodInfo_AuthService_CheckEmailAvailability);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.SignupRequest,
 *   !proto.authpb.SignupResponse>}
 */
const methodInfo_AuthService_Signup = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authpb.SignupResponse,
  /** @param {!proto.authpb.SignupRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.authpb.SignupResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.SignupRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.SignupResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.SignupResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.signup =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authpb.AuthService/Signup',
      request,
      metadata || {},
      methodInfo_AuthService_Signup,
      callback);
};


/**
 * @param {!proto.authpb.SignupRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.SignupResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.signup =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authpb.AuthService/Signup',
      request,
      metadata || {},
      methodInfo_AuthService_Signup);
};


module.exports = proto.authpb;

