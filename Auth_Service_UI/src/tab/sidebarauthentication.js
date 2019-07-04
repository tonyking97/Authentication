import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import AppBar from '@material-ui/core/AppBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Group from '@material-ui/icons/Group';
import VpnKey from '@material-ui/icons/VpnKey';
import History from '@material-ui/icons/History';
import VerifiedUser from '@material-ui/icons/VerifiedUser';
import Android from '@material-ui/icons/Android';
import Devices from '@material-ui/icons/Devices';
import ActiveSessions from '../autentication/activeSessions';
import ActiveAuthtokens from '../autentication/activeAuthtokens';
import ActivityDetails from '../autentication/activityDetails';
import ConnectedAccounts from '../autentication/connectedAccounts';
import ConnectedDevices from '../autentication/connectedDevices';


const drawerWidth = 240;

const styles = theme => ({
    root: {
        //marginTop: 1,
        display: 'flex',
    },
    appBar: {
        marginTop: 74,
        zIndex: theme.zIndex.drawer + 1,
    },
    drawer: {
        width: drawerWidth,
        flexShrink: 0,
    },
    drawerPaper: {
        marginTop: 74,
        width: drawerWidth,
    },
    content: {
        flexGrow: 1,
        padding: theme.spacing.unit * 3,
    },
    toolbar: theme.mixins.toolbar,
});

function TabContainer(props) {
    return (
        <Typography component="div" style={{padding: 8 * 3}}>
            {props.children}
        </Typography>
    );
}

TabContainer.propTypes = {
    children: PropTypes.node.isRequired,
};


class AuthSideBar extends Component {
    state = {
        value: 0
    };
    checkEvent0 = () => {
        this.setState({value: 0})
    }
    checkEvent1 = () => {
        this.setState({value: 1})
    }
    checkEvent2 = () => {
        this.setState({value: 2})
    }
    checkEvent3 = () => {
        this.setState({value: 3})
    }
    checkEvent4 = () => {
        this.setState({value: 4})
    }
    checkEvent5 = () => {
        this.setState({value: 5})
    }

    handleChange = (event, value) => {
        this.setState({value})
    }

    render() {
        const {classes} = this.props;
        const {value} = this.state;

        return (
            <div className={classes.root}>
                <CssBaseline/>
                <AppBar position="fixed" className={classes.appBar}>
                    <Toolbar>
                        <Typography variant="h6" color="inherit" noWrap>
                            Autentication Summary
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Drawer
                    className={classes.drawer}
                    variant="permanent"
                    classes={{
                        paper: classes.drawerPaper,
                    }}
                >
                    <div className={classes.toolbar}/>
                    <List>

                        <ListItem button onClick={this.checkEvent0} onChange={this.handleChange}>
                            <ListItemIcon> <Group/> </ListItemIcon>
                            <ListItemText primary='Active Sessions'/>
                        </ListItem>

                        <ListItem button onClick={this.checkEvent1} onChange={this.handleChange}>
                            <ListItemIcon> <VpnKey/> </ListItemIcon>
                            <ListItemText primary='Active Authtokens'/>
                        </ListItem>

                        <ListItem button onClick={this.checkEvent2} onChange={this.handleChange}>
                            <ListItemIcon> <History/> </ListItemIcon>
                            <ListItemText primary='Activity History'/>
                        </ListItem>

                        <ListItem button onClick={this.checkEvent3} onChange={this.handleChange}>
                            <ListItemIcon> <Android/> </ListItemIcon>
                            <ListItemText primary='App Logins'/>
                        </ListItem>

                    </List>

                    <Divider/>

                    <List>

                        <ListItem button onClick={this.checkEvent4} onChange={this.handleChange}>
                            <ListItemIcon> <VerifiedUser/> </ListItemIcon>
                            <ListItemText primary='Connected Accounts'/>
                        </ListItem>

                        <ListItem button onClick={this.checkEvent5} onChange={this.handleChange}>
                            <ListItemIcon> <Devices/> </ListItemIcon>
                            <ListItemText primary='Connected Devices'/>
                        </ListItem>

                    </List>

                </Drawer>
                <main className={classes.content}>
                    <div className={classes.toolbar}/>
                    {value === 0 && <TabContainer> <ActiveSessions/> </TabContainer>}
                    {value === 1 && <TabContainer> <ActiveAuthtokens/> </TabContainer>}
                    {value === 2 && <TabContainer> <ActivityDetails/> </TabContainer>}
                    {value === 3 && <TabContainer>This is a list of Mobile/Desktop Apps Sessions to your FogFind
                        account.</TabContainer>}
                    {value === 4 && <TabContainer> <ConnectedAccounts/> </TabContainer>}
                    {value === 5 && <TabContainer> <ConnectedDevices/> </TabContainer>}

                </main>
            </div>
        );
    }
}

AuthSideBar.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(AuthSideBar);