import React, {Component} from 'react';
import {BrowserRouter as Router, Redirect, Route, Switch} from 'react-router-dom';

const PrivateRoute = ({component, exact = false, path, authenticated}) => (
    <Route
        exact={exact}
        path={path}
        render={props =>
            authenticated ? (
                React.createElement(component, props)
            ) : (
                <Redirect
                    key="login"
                    to={'/login'}
                />
            )
        }
    />
)

class Login extends Component {
    login = () => {
        setTimeout(() => this.props.onLogin, 1000)
    }

    render() {
        return (
            <div>
                <button onClick={this.login}>Login</button>
            </div>
        )
    }
}

const Welcome = () => <div>Private Page</div>

class Routing extends Component {
    state = {loggedIn: false}
    handleLogin = () => {
        this.setState({loggedIn: true})
    }

    routes() {
        const {loggedIn} = this.state
        const routeArr = [
            <Route
                exact
                path="/"
                component={() =>
                    loggedIn ? (
                        <Redirect to="/welcomuser"/>
                    ) : (
                        <Redirect to="/login"/>
                    )}
            />,
            <PrivateRoute
                authenticated={loggedIn}
                key="main"
                path="/welcomeuser"
                component={Welcome}
            />,
            <Route key="login" exact path="/login" component={() => <Login onLogin={this.handleLogin}/>}/>,
            <Route component={NotFound}/>,
        ];
        if (!loggedIn) {
            routeArr.push(<Redirect key="loginRedirect" to="/login"/>);
        }
        return routeArr;
    }

    render() {
        return (
            <Router>
                <Switch>{this.routes()}</Switch>
            </Router>
        )
    }
}

const App = () => (
    <div>
        <Routing/>
    </div>
)
