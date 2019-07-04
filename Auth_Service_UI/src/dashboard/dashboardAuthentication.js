import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';
// import { Link } from 'react-router-dom';
import {withStyles} from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Group from '@material-ui/icons/Group';
import VpnKey from '@material-ui/icons/VpnKey';
import History from '@material-ui/icons/History';
import VerifiedUser from '@material-ui/icons/VerifiedUser';
import Android from '@material-ui/icons/Android';
import Devices from '@material-ui/icons/Devices';
import AccountBox from '@material-ui/icons/AccountBox';
import Email from '@material-ui/icons/Email';
import blue from '@material-ui/core/colors/blue';
import Avatar from "@material-ui/core/Avatar";
import Tooltip from '@material-ui/core/Tooltip';
import DashboardActiveSessions from "./components/activeSessions";
import DashboardActiveAuthtokens from "./components/activeAuthtokens";
// import Routes from './components/activeSessions';
// import { Redirect } from 'react-router-dom'

const session = blue[800];

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        padding: theme.spacing.unit,
        textAlign: 'center',
        color: theme.palette.text.primary,
        height: 170,
        width: 140,
    },
    button: {
        padding: theme.spacing.unit,
        textAlign: 'center',
        color: theme.palette.text.primary,
        height: 180,
        width: 160,
    },
    icon: {
        marginTop: theme.spacing.unit * 3.5,
        marginLeft: theme.spacing.unit * 1,
        height: 80,
        width: 80,
        color: session,
    },
    grid: {
        marginLeft: 100,
        marginTop: 10,
    },
    bigAvatar: {
        margin: 10,
        width: 150,
        height: 150,
    },
});

class DashboardAuthentication extends Component {
    state = {
        src: '',
        first_name: 'Gokul',
        last_name: 'GK',
        email: 'gokul.vellingiri@gmail.com',
        phone: '9944101945',
    };

    constructor(props) {
        super(props);
        this.state = {
            src: '',
            first_name: 'Gokul',
            last_name: 'GK',
        };
    }


    setASRedirect = () => {
        this.setState({
            redirect: true,
        })
    }
    renderASRedirect = () => {
        if (this.state.redirect) {
            ReactDOM.render(
                <DashboardActiveSessions/>,
                document.getElementById('root')
            );
        }
    }

    setATRedirect = () => {
        this.setState({
            redirect: true,
        })
    }
    renderATRedirect = () => {
        if (this.state.redirect) {
            ReactDOM.render(
                <DashboardActiveAuthtokens/>,
                document.getElementById('root')
            );
        }
    }


    render() {
        const {classes} = this.props;

        return (
            <div className={classes.root}>
                <Grid container spacing={24} className={classes.grid}>

                    <Grid item xs={4}>
                        {
                            <Tooltip title="Manage all your concurrent active Signed In browser Sessions."
                                     placement="bottom">
                                <button className={classes.button} variant="contained" onClick={this.setASRedirect}>
                                    {this.renderASRedirect()}
                                    <Group className={classes.icon}/>
                                    <h3>Active Sessions</h3>
                                </button>
                            </Tooltip>
                        }
                    </Grid>

                    <Grid item xs={4}>
                        <Tooltip title="Manage all the active authtokens authorised to different scopes."
                                 placement="bottom">
                            <button className={classes.button} variant="contained" onClick={this.setATRedirect}>
                                {this.renderATRedirect()}
                                <VpnKey className={classes.icon}/>
                                <h3>Active Authtokens</h3>
                            </button>
                        </Tooltip>
                    </Grid>

                    <Grid item xs={4}>
                        <Tooltip title="Monitor all your account access." placement="bottom">
                            <button className={classes.button} variant="contained">
                                <History className={classes.icon}/>
                                <h3>Activity History</h3>
                            </button>
                        </Tooltip>
                    </Grid>

                    <Grid item xs={4}>
                        <Tooltip title="Manage all the IOT devices authenticated." placement="bottom">
                            <button className={classes.button} variant="contained">
                                <Devices className={classes.icon}/>
                                <h3>Connected Devices</h3>
                            </button>
                        </Tooltip>
                    </Grid>

                    <Grid item xs={4}>
                        <Avatar alt={this.state.first_name} src={this.state.src} className={classes.bigAvatar}/>
                    </Grid>

                    <Grid item xs={4}>
                        <Tooltip title="Manage all the accounts linked to the platform." placement="bottom">
                            <button className={classes.button} variant="contained">
                                <VerifiedUser className={classes.icon}/>
                                <h3>Connected Accounts</h3>
                            </button>
                        </Tooltip>
                    </Grid>

                    <Grid item xs={4}>
                        <Tooltip title="Manage your Personal Information and edit them." placement="bottom">
                            <button className={classes.button} variant="contained">
                                <AccountBox className={classes.icon}/>
                                <h3>Personal Information</h3>
                            </button>
                        </Tooltip>
                    </Grid>

                    <Grid item xs={4}>
                        <Tooltip title="Update your email preference of primary and secondary email address."
                                 placement="bottom">
                            <button className={classes.button} variant="contained">
                                <Email className={classes.icon}/>
                                <h3>Email Address</h3>
                            </button>
                        </Tooltip>
                    </Grid>

                    <Grid item xs={4}>
                        <Tooltip title="Manage List of Mobile/Desktop Apps Sessions to your FogFind account."
                                 placement="bottom">
                            <button className={classes.button} variant="contained">
                                <Android className={classes.icon}/>
                                <h3>App Logins</h3>
                            </button>
                        </Tooltip>
                    </Grid>

                </Grid>

            </div>


        );
    }
}

DashboardAuthentication.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(DashboardAuthentication);