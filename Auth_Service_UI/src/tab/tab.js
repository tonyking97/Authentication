import React from 'react';
import * as PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Security from '@material-ui/icons/Security';
import Dashboard from '@material-ui/icons/Dashboard';
import AuthSideBar from './sidebarauthentication'
import ProfileSideBar from './sidebarprofile'
import DashboardAuthentication from '../dashboard/dashboardAuthentication'

// import { Link, BrowserRouter } from 'react-router-dom'


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

const styles = theme => ({
    root: {
        flexGrow: 1,
        width: '100%',
        backgroundColor: theme.palette.background.paper,
    },
});

class TopTab extends React.Component {

    state = {
        value: this.props.value || 0,
    };

    constructor(props) {
        super(props);
        console.log(props)
    }

    handleChange = (event, value) => {
        this.setState({value});
    };

    render() {
        const {classes} = this.props;
        const {value} = this.state;

        return (
            <div className={classes.root}>
                <AppBar position="sticky" color="default">
                    <Tabs
                        value={value}
                        onChange={this.handleChange}
                        variant="scrollable"
                        scrollButtons="on"
                        indicatorColor="primary"
                        textColor="primary"
                    >
                        <Tab label="Dashboard" icon={<Dashboard/>}/>
                        <Tab label="Profile" icon={<AccountCircle/>}/>
                        <Tab label="Authentication" icon={<Security/>}/>


                        {/*<Tab label="Dashboard" icon={<Dashboard />} href="dashboard"/>*/}
                        {/*<Tab label="Profile" icon={<AccountCircle />} href="profile"/>*/}
                        {/*<Tab label="Authentication" icon={<Security />} href="authentication"/>*/}

                    </Tabs>
                </AppBar>
                {value === 0 && <TabContainer> <DashboardAuthentication/> </TabContainer>}
                {value === 1 && <TabContainer> <ProfileSideBar/> </TabContainer>}
                {value === 2 && <TabContainer> <AuthSideBar/> </TabContainer>}

                {/*{value === 0 && <TabContainer> <DashboardAuthentication/> <Link to={"dashboard"}/>  </TabContainer>}*/}
                {/*{value === 1 && <TabContainer> <ProfileSideBar/> <Link to={"profile"}/>  </TabContainer>}*/}
                {/*{value === 2 && <TabContainer> <AuthSideBar/> <Link to={"authentication"}/>  </TabContainer>}*/}

                {/*{value === 0 && <Link to={"/dashboard"}/>}*/}
                {/*{value === 1 && <TabContainer> <ProfileSideBar/> </TabContainer>}*/}
                {/*{value === 2 && <TabContainer> <AuthSideBar/> </TabContainer>}*/}

            </div>
        );
    }
}

TopTab.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(TopTab);