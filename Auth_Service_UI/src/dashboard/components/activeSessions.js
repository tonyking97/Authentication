import React, {Component} from 'react';
import {BrowserRouter, Route, Switch} from "react-router-dom";
// import ActiveSessions from '../../autentication/activeSessions';
// import AuthSideBar from '../../tab/sidebarauthentication'
import TopTab from "../../tab/tab";

// import App from "../../App";

class DashboardActiveSessions extends Component {
    render() {
        return (
            <TopTab value={2}/>
            //<AuthSideBar/>
            // <ActiveSessions />
        );
    }
}

export default () => (
    <BrowserRouter>
        <Switch>
            {/*<Route path="/activesessions" exact render={ props => <DashboardActiveSessions {...props} />}/>*/}
            <Route path="/activesessions" exact render={props => <DashboardActiveSessions {...props} />}/>
        </Switch>
    </BrowserRouter>
);

// export default (DashboardActiveSessions);