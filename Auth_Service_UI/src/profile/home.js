import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Avatar from '@material-ui/core/Avatar';
import Grid from '@material-ui/core/Grid';
import InputBase from '@material-ui/core/InputBase';
import AccountCircle from '@material-ui/icons/AccountCircle';
import EmailIcon from '@material-ui/icons/Email';
import PhoneIcon from '@material-ui/icons/Phone';
import fakeJsonData2 from './../fakedata2.json';

const styles = theme => ({
    root: {
        marginTop: 40,
        marginLeft: 80,
    },
    bigAvatar: {
        margin: 10,
        width: 150,
        height: 150,
    },
    margin: {
        margin: theme.spacing.unit,
    },
    emailmargin: {
        margin: theme.spacing.unit,
        width: 250,
    },
});

JSON.stringify(fakeJsonData2);

class HomeProfile extends Component {
    state = {
        src: '',
        first_name: fakeJsonData2.firstname,
        last_name: fakeJsonData2.lastname,
        email: fakeJsonData2.email,
        phone: '9944101945'
    };

    render() {
        const {classes} = this.props;

        return (
            <div className={classes.root}>
                <Grid spacing={24} container justify="flex-start" alignItems="flex-start">
                    <Avatar alt={this.state.first_name} src={this.state.src} className={classes.bigAvatar}/>
                </Grid>

                <br/><br/><br/>

                <Grid container spacing={24} justify="flex-start" alignItems="flex-start">
                    <Grid item>
                        <AccountCircle/>
                    </Grid>
                    <Grid item>
                        <InputBase defaultValue={this.state.first_name + "  " + this.state.last_name}/>
                    </Grid>
                </Grid>

                <br/><br/><br/>

                <Grid container spacing={24} justify="flex-start" alignItems="flex-start">
                    <Grid item>
                        <EmailIcon/>
                    </Grid>
                    <Grid item>
                        <InputBase className={classes.emailmargin} defaultValue={this.state.email}/>
                    </Grid>
                </Grid>

                <br/><br/><br/>

                <Grid container spacing={24} justify="flex-start" alignItems="flex-start">
                    <Grid item>
                        <PhoneIcon/>
                    </Grid>
                    <Grid item>
                        <InputBase className={classes.margin} defaultValue={this.state.phone}/>
                    </Grid>
                </Grid>

            </div>
        );
    }

}

HomeProfile.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(HomeProfile);