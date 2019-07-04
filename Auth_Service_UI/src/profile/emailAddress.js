import React, {Component} from 'react';
import TextField from '@material-ui/core/TextField';
import withStyles from "@material-ui/core/es/styles/withStyles";
import PropTypes from 'prop-types';
import {Typography} from "@material-ui/core";
import FormLabel from '@material-ui/core/FormLabel';
import Button from '@material-ui/core/Button';
import EmailIcon from '@material-ui/icons/Email';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import EditIcon from '@material-ui/icons/Edit';


const styles = theme => ({
    root: {
        margin: theme.spacing.unit,
        //...theme.mixins.gutters(),
        paddingTop: theme.spacing.unit,
        paddingBottom: theme.spacing.unit,
    },
    textField: {
        marginLeft: theme.spacing.unit,
        marginRight: theme.spacing.unit,
        width: 300,
    },
    menu: {
        width: 200,
    },
    radio: {
        margin: theme.spacing.unit,
        flexDirection: 'row',
    },
    dropdown: {
        margin: theme.spacing.unit,
    },
    rightIcon: {
        marginLeft: theme.spacing.unit,
    },
    button: {
        marginLeft: theme.spacing.unit,
    },
});


class EmailAddress extends Component {
    state = {
        first_name: 'Gokul',
        last_name: 'GK',
        gender: 'male',
        primaryemail: 'gokul.vellingiri@gmail.com',
        secondaryemail: 'gokulv@terawe.com',
        language: 'EN',
        country: 'IN',
        time_zone: '+5:30',
    }

    handleChange = name => event => {
        this.setState({
            [name]: event.target.value,
        });
    };

    render() {
        const {classes} = this.props;

        return (
            <form className={classes.root} noValidate autoComplete="off">
                <FormLabel className={classes.root}>
                    Update your email preferences of primary/secondary addresses. If you forget your password, we'll
                    send reset instructions to your email.
                </FormLabel>

                <br/><br/><br/>

                <Typography variant="title" className={classes.root} gutterBottom>
                    Primary Email Address
                </Typography>

                <TextField
                    disabled={true}
                    id="standard-disabled"
                    label=""
                    defaultValue={this.state.primaryemail}
                    className={classes.textField}
                    margin="normal"
                />

                <IconButton className={classes.button} aria-label="Edit">
                    <EditIcon/>
                </IconButton>

                <br></br><br></br>

                <Typography variant="title" className={classes.root} gutterBottom>
                    Secondary Email Address
                </Typography>

                <TextField
                    disabled
                    id="standard-disabled"
                    label=""
                    defaultValue={this.state.secondaryemail}
                    className={classes.textField}
                    margin="normal"
                />

                <IconButton className={classes.button} aria-label="Edit">
                    <EditIcon/>
                </IconButton>
                <IconButton className={classes.button} aria-label="Delete">
                    <DeleteIcon/>
                </IconButton>
                <IconButton className={classes.button} aria-label="Make Primary">
                    <EmailIcon/>
                </IconButton>

                <br/><br/><br/>

                <Button variant="contained" color="default" className={classes.button}>
                    Add Email
                    <EmailIcon className={classes.rightIcon}/>
                </Button>

            </form>

        )
    }

}

EmailAddress.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(EmailAddress);