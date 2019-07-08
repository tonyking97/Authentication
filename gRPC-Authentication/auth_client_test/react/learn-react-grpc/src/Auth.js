
import React, { Component } from 'react';
// import logo from './logo.svg';
import './App.css';

const { AuthServiceClient } = require('./auth_grpc_web_pb');
const { CheckServerStatusRequest, CheckServerStatusResponse } = require('./auth_pb');

var client = new AuthServiceClient('http://localhost:9090', null, null);

class Auth extends Component {

    callGrpcService = () => {
        const request = new CheckServerStatusRequest();

        client.checkServerStatus(request, {}, (err, response) => {
            if (response == null) {
                console.log(err)
            }else {
                console.log(response.getStatus())
            }
        });
    };

    render() {
        return (
            <div className="Auth">
                <button type="submit" className="btn btn-primary" onClick={this.callGrpcService}>Click for grpc request</button>
                <div id="wrapper">
                    <div className="container">
                        <div className="row">
                            <article className="col-md-12">
                                <h2 className="text-center">Auth Server Details</h2>
                                <div id="app" className="app-container" />
                            </article>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default Auth;