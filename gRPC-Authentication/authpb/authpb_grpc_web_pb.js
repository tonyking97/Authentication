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
    function (hostname, credentials, options) {
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
    function (hostname, credentials, options) {
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
 *   !proto.authpb.CheckServerStatusRequest,
 *   !proto.authpb.CheckServerStatusResponse>}
 */
const methodInfo_AuthService_CheckServerStatus = new grpc.web.AbstractClientBase.MethodInfo(
    proto.authpb.CheckServerStatusResponse,
    /** @param {!proto.authpb.CheckServerStatusRequest} request */
    function (request) {
        return request.serializeBinary();
    },
    proto.authpb.CheckServerStatusResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.CheckServerStatusRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.CheckServerStatusResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.CheckServerStatusResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.checkServerStatus =
    function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/authpb.AuthService/CheckServerStatus',
            request,
            metadata || {},
            methodInfo_AuthService_CheckServerStatus,
            callback);
    };


/**
 * @param {!proto.authpb.CheckServerStatusRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.CheckServerStatusResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.checkServerStatus =
    function (request, metadata) {
        return this.client_.unaryCall(this.hostname_ +
            '/authpb.AuthService/CheckServerStatus',
            request,
            metadata || {},
            methodInfo_AuthService_CheckServerStatus);
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
    function (request) {
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
    function (request, metadata, callback) {
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
    function (request, metadata) {
        return this.client_.unaryCall(this.hostname_ +
            '/authpb.AuthService/CheckUsername',
            request,
            metadata || {},
            methodInfo_AuthService_CheckUsername);
    };


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.RegisterAccountRequest,
 *   !proto.authpb.RegisterAccountResponse>}
 */
const methodInfo_AuthService_RegisterAccount = new grpc.web.AbstractClientBase.MethodInfo(
    proto.authpb.RegisterAccountResponse,
    /** @param {!proto.authpb.RegisterAccountRequest} request */
    function (request) {
        return request.serializeBinary();
    },
    proto.authpb.RegisterAccountResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.RegisterAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.RegisterAccountResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.RegisterAccountResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.registerAccount =
    function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/authpb.AuthService/RegisterAccount',
            request,
            metadata || {},
            methodInfo_AuthService_RegisterAccount,
            callback);
    };


/**
 * @param {!proto.authpb.RegisterAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.RegisterAccountResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.registerAccount =
    function (request, metadata) {
        return this.client_.unaryCall(this.hostname_ +
            '/authpb.AuthService/RegisterAccount',
            request,
            metadata || {},
            methodInfo_AuthService_RegisterAccount);
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
    function (request) {
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
    function (request, metadata, callback) {
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
    function (request, metadata) {
        return this.client_.unaryCall(this.hostname_ +
            '/authpb.AuthService/Login',
            request,
            metadata || {},
            methodInfo_AuthService_Login);
    };


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.UpdateAccountRequest,
 *   !proto.authpb.UpdateAccountResponse>}
 */
const methodInfo_AuthService_UpdateAccount = new grpc.web.AbstractClientBase.MethodInfo(
    proto.authpb.UpdateAccountResponse,
    /** @param {!proto.authpb.UpdateAccountRequest} request */
    function (request) {
        return request.serializeBinary();
    },
    proto.authpb.UpdateAccountResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.UpdateAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.UpdateAccountResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.UpdateAccountResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.updateAccount =
    function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/authpb.AuthService/UpdateAccount',
            request,
            metadata || {},
            methodInfo_AuthService_UpdateAccount,
            callback);
    };


/**
 * @param {!proto.authpb.UpdateAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.UpdateAccountResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.updateAccount =
    function (request, metadata) {
        return this.client_.unaryCall(this.hostname_ +
            '/authpb.AuthService/UpdateAccount',
            request,
            metadata || {},
            methodInfo_AuthService_UpdateAccount);
    };


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.GetAccountDetailsRequest,
 *   !proto.authpb.GetAccountDetailsResponse>}
 */
const methodInfo_AuthService_GetAccountDetails = new grpc.web.AbstractClientBase.MethodInfo(
    proto.authpb.GetAccountDetailsResponse,
    /** @param {!proto.authpb.GetAccountDetailsRequest} request */
    function (request) {
        return request.serializeBinary();
    },
    proto.authpb.GetAccountDetailsResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.GetAccountDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.GetAccountDetailsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.GetAccountDetailsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.getAccountDetails =
    function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/authpb.AuthService/GetAccountDetails',
            request,
            metadata || {},
            methodInfo_AuthService_GetAccountDetails,
            callback);
    };


/**
 * @param {!proto.authpb.GetAccountDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.GetAccountDetailsResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.getAccountDetails =
    function (request, metadata) {
        return this.client_.unaryCall(this.hostname_ +
            '/authpb.AuthService/GetAccountDetails',
            request,
            metadata || {},
            methodInfo_AuthService_GetAccountDetails);
    };


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.ChangePasswordRequest,
 *   !proto.authpb.ChangePasswordResponse>}
 */
const methodInfo_AuthService_ChangePassword = new grpc.web.AbstractClientBase.MethodInfo(
    proto.authpb.ChangePasswordResponse,
    /** @param {!proto.authpb.ChangePasswordRequest} request */
    function (request) {
        return request.serializeBinary();
    },
    proto.authpb.ChangePasswordResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.ChangePasswordRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authpb.ChangePasswordResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.ChangePasswordResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.changePassword =
    function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/authpb.AuthService/ChangePassword',
            request,
            metadata || {},
            methodInfo_AuthService_ChangePassword,
            callback);
    };


/**
 * @param {!proto.authpb.ChangePasswordRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authpb.ChangePasswordResponse>}
 *     A native promise that resolves to the response
 */
proto.authpb.AuthServicePromiseClient.prototype.changePassword =
    function (request, metadata) {
        return this.client_.unaryCall(this.hostname_ +
            '/authpb.AuthService/ChangePassword',
            request,
            metadata || {},
            methodInfo_AuthService_ChangePassword);
    };


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.GetAllAccountsRequest,
 *   !proto.authpb.GetAllAccountsResponse>}
 */
const methodInfo_AuthService_GetAllAccounts = new grpc.web.AbstractClientBase.MethodInfo(
    proto.authpb.GetAllAccountsResponse,
    /** @param {!proto.authpb.GetAllAccountsRequest} request */
    function (request) {
        return request.serializeBinary();
    },
    proto.authpb.GetAllAccountsResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.GetAllAccountsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.GetAllAccountsResponse>}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.getAllAccounts =
    function (request, metadata) {
        return this.client_.serverStreaming(this.hostname_ +
            '/authpb.AuthService/GetAllAccounts',
            request,
            metadata || {},
            methodInfo_AuthService_GetAllAccounts);
    };


/**
 * @param {!proto.authpb.GetAllAccountsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.GetAllAccountsResponse>}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServicePromiseClient.prototype.getAllAccounts =
    function (request, metadata) {
        return this.client_.serverStreaming(this.hostname_ +
            '/authpb.AuthService/GetAllAccounts',
            request,
            metadata || {},
            methodInfo_AuthService_GetAllAccounts);
    };


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authpb.GetAllSessionDetailsRequest,
 *   !proto.authpb.GetAllSessionDetailsResponse>}
 */
const methodInfo_AuthService_GetAllSessionDetails = new grpc.web.AbstractClientBase.MethodInfo(
    proto.authpb.GetAllSessionDetailsResponse,
    /** @param {!proto.authpb.GetAllSessionDetailsRequest} request */
    function (request) {
        return request.serializeBinary();
    },
    proto.authpb.GetAllSessionDetailsResponse.deserializeBinary
);


/**
 * @param {!proto.authpb.GetAllSessionDetailsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.GetAllSessionDetailsResponse>}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServiceClient.prototype.getAllSessionDetails =
    function (request, metadata) {
        return this.client_.serverStreaming(this.hostname_ +
            '/authpb.AuthService/GetAllSessionDetails',
            request,
            metadata || {},
            methodInfo_AuthService_GetAllSessionDetails);
    };


/**
 * @param {!proto.authpb.GetAllSessionDetailsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.authpb.GetAllSessionDetailsResponse>}
 *     The XHR Node Readable Stream
 */
proto.authpb.AuthServicePromiseClient.prototype.getAllSessionDetails =
    function (request, metadata) {
        return this.client_.serverStreaming(this.hostname_ +
            '/authpb.AuthService/GetAllSessionDetails',
            request,
            metadata || {},
            methodInfo_AuthService_GetAllSessionDetails);
    };


module.exports = proto.authpb;

