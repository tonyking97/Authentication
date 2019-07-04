import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import AppBar from '@material-ui/core/AppBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import AccountBox from '@material-ui/icons/AccountBox';
import Email from '@material-ui/icons/Email';
import Phone from '@material-ui/icons/Phone';
import Home from '@material-ui/icons/Home';
import PersonalInformationForm from '../profile/personalInformationForm';
import EmailAddress from '../profile/emailAddress';
import PhoneNumber from '../profile/phoneNumber';
import HomeProfile from '../profile/home';

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
        padding: theme.spacing.unit,
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


class ProfileSideBar extends Component {
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
                            Your Profile Information
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
                            <ListItemIcon> <Home/> </ListItemIcon>
                            <ListItemText primary='Home'/>
                        </ListItem>

                        <ListItem button onClick={this.checkEvent1} onChange={this.handleChange}>
                            <ListItemIcon> <AccountBox/> </ListItemIcon>
                            <ListItemText primary='Personal Information'/>
                        </ListItem>

                        <ListItem button onClick={this.checkEvent2} onChange={this.handleChange}>
                            <ListItemIcon> <Email/> </ListItemIcon>
                            <ListItemText primary='Email Address'/>
                        </ListItem>

                        <ListItem button onClick={this.checkEvent3} onChange={this.handleChange}>
                            <ListItemIcon> <Phone/> </ListItemIcon>
                            <ListItemText primary='Phone Numbers'/>
                        </ListItem>

                    </List>

                </Drawer>

                <main className={classes.content}>
                    <div className={classes.toolbar}/>
                    {value === 0 && <TabContainer> <HomeProfile/> </TabContainer>}
                    {value === 1 && <TabContainer> <PersonalInformationForm/> </TabContainer>}
                    {value === 2 && <TabContainer> <EmailAddress/> </TabContainer>}
                    {value === 3 && <TabContainer> <PhoneNumber/> </TabContainer>}
                </main>

            </div>
        );
    }
}

ProfileSideBar.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ProfileSideBar);