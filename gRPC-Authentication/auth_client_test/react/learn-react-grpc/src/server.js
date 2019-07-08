var grpc = require('grpc');
var pingPongProto = grpc.load('ping_pong.proto');

var authProto = grpc.load('auth.proto');
var server = new grpc.Server();

// server.addService(pingPongProto.pingpong.PingPongService.service, {
//     pingPong: function(call,callback) {
//         console.log("Request");
//         return callback(null,{pong:"Pong"})
//     }
// });

server.addService(authProto.auth.AuthService.service,{
    checkServerStatus: function (call, callback) {
        return callback(null,{server_name:"Auth Server",host:"",port:"", time:"",status:"good :)"})
    }
});

server.bind('localhost:8080',grpc.ServerCredentials.createInsecure());
server.start();