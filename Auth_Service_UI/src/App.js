import React, {Component} from 'react';
import './App.css';
import TopTab from "./tab/tab";
// import {Router, Route, } from 'react-router';
// import AuthSideBar from './tab/sidebarauthentication'
// import ProfileSideBar from './tab/sidebarprofile'
// import DashboardAuthentication from './dashboard/dashboardAuthentication'
import axios from "axios";


// import {createBrowserHistory} from 'history';

// const browserHistory = createBrowserHistory();

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            pdata: [],
        };
    }

    componentDidMount() {
        axios
            .get("")
            .then(response => {
                // create an array of contacts only with relevant data
                const newData = response.data.map(c => {
                    return {
                        logged_Time: c.logged_Time,
                        exipry_Time: c.exipry_Time,
                        ip: c.ip_details.ip,
                        status: c.status,
                        token: c.token,
                        referer: c.referer,
                        service: 'Authentication',
                        description: '',
                        browser_Name: c.browser_Name,
                        scope: 'Authentication Service',
                    };
                });

                // create a new "State" object without mutating
                // the original State object.
                const newState = Object.assign({}, this.state, {
                    data: newData
                });

                // store the new state object in the component's state
                this.setState(newState);
            })
            .catch(error => console.log(error));
    }

    componentDidMount1() {
        axios
            .get("")
            .then(response => {
                // create an array of contacts only with relevant data
                const newPdata = response.data.map(p => {
                    return {
                        firstname: p.firstname,
                        lastname: p.lastname,
                        email: p.email,
                        username: p.username,
                        createdtime: p.createdtime,
                        updatedtime: p.updatedtime,
                    };
                });

                // create a new "State" object without mutating
                // the original State object.
                const newState1 = Object.assign({}, this.state, {
                    pdata: newPdata
                });

                // store the new state object in the component's state
                this.setState(newState1);
            })
            .catch(error => console.log(error));
    }


    // render() {
    //     return (
    //         <Router history={browserHistory}>
    //             <Route path={"/"} component={TopTab}>
    //                 {/*<Route path={"dashboard"} component={DashboardAuthentication}></Route>*/}
    //                 {/*<Route path={"profile"} component={ProfileSideBar}/>*/}
    //                 {/*<Route path={"authentication"} component={AuthSideBar}/>*/}
    //             </Route>
    //             <Route path={"dashboard"} component={DashboardAuthentication}></Route>
    //             <Route path={"profile"} component={ProfileSideBar}/>
    //
    //
    //         </Router>
    //
    // );
    // }

    render() {
        return (
            <TopTab/>
        );
    }
}

export default (App);