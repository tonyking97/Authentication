import React, {Component} from 'react';
import TextField from '@material-ui/core/TextField';
import withStyles from "@material-ui/core/es/styles/withStyles";
import PropTypes from 'prop-types';
import FormLabel from '@material-ui/core/FormLabel';
import Button from '@material-ui/core/Button';
import PhoneIcon from '@material-ui/icons/Phone';
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


class PhoneNumber extends Component {
    state = {
        phone_number: '9944101945',
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
                    Recovery phone numbers help you regain access to your account in case you forget your password.
                </FormLabel>

                <br/><br/><br/>

                <TextField
                    disabled
                    id="standard-disabled"
                    label=""
                    defaultValue={this.state.phone_number}
                    className={classes.textField}
                    margin="normal"
                />

                <IconButton className={classes.button} aria-label="Edit">
                    <EditIcon/>
                </IconButton>
                <IconButton className={classes.button} aria-label="Delete">
                    <DeleteIcon/>
                </IconButton>

                <br/><br/><br/>

                <Button variant="contained" color="default" className={classes.button}>
                    Add Phone Number
                    <PhoneIcon className={classes.rightIcon}/>
                </Button>

            </form>

        )
    }

}

PhoneNumber.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(PhoneNumber);