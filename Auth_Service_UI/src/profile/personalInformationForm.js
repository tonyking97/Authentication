import React, {Component} from 'react';
import TextField from '@material-ui/core/TextField';
import withStyles from "@material-ui/core/es/styles/withStyles";
import PropTypes from 'prop-types';
import Radio from '@material-ui/core/Radio';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormLabel from '@material-ui/core/FormLabel';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import SaveIcon from '@material-ui/icons/Save';
import fakeJsonData2 from './../fakedata2.json';


const styles = theme => ({
    root: {
        ...theme.mixins.gutters(),
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
});

JSON.stringify(fakeJsonData2);

class PersonalInformationForm extends Component {
    state = {
        first_name: fakeJsonData2.firstname,
        last_name: fakeJsonData2.lastname,
        gender: 'male',
        email: fakeJsonData2.email,
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
                <FormLabel className={classes.radio}>
                    Manage your personal details, contact information, language, country and timezone settings.
                </FormLabel>
                <br></br><br></br>
                <TextField
                    id="outlined-firstname"
                    name="first_name"
                    label="First Name"
                    className={classes.textField}
                    value={this.state.first_name}
                    onChange={this.handleChange('first_name')}
                    margin="normal"
                    variant="outlined"
                />
                <TextField
                    id="outlined-lastname"
                    name="last_name"
                    label="Last Name"
                    className={classes.textField}
                    value={this.state.last_name}
                    onChange={this.handleChange('last_name')}
                    margin="normal"
                    variant="outlined"
                />

                <br></br><br></br>

                <FormLabel component="legend" className={classes.radio}>Gender</FormLabel>
                <RadioGroup
                    id="gender"
                    name="gender"
                    aria-label="Gender"
                    className={classes.radio}
                    value={this.state.gender}
                    onChange={this.handleChange('gender')}
                >
                    <FormControlLabel value="female" control={<Radio/>} label="Female"/>
                    <FormControlLabel value="male" control={<Radio/>} label="Male"/>
                    <FormControlLabel value="other" control={<Radio/>} label="Other"/>
                </RadioGroup>

                <TextField
                    id="outlined-email"
                    name="email"
                    label="Contact Details"
                    className={classes.textField}
                    value={this.state.email}
                    onChange={this.handleChange('email')}
                    margin="normal"
                />

                <br></br><br></br>

                <InputLabel htmlFor="language" className={classes.dropdown}>Language </InputLabel>

                <Select
                    className={classes.dropdown}
                    value={this.state.language}
                    onChange={this.handleChange('language')}
                    inputProps={{
                        name: 'language',
                        id: 'language',
                    }}
                >
                    <MenuItem value={'EN'}>English</MenuItem>
                    <MenuItem value={'FR'}>French</MenuItem>
                    <MenuItem value={'DE'}>German</MenuItem>
                </Select>

                <br></br><br></br>

                <InputLabel htmlFor="country" className={classes.dropdown}>Country</InputLabel>

                <Select
                    className={classes.dropdown}
                    value={this.state.country}
                    onChange={this.handleChange('country')}
                    inputProps={{
                        name: 'country',
                        id: 'country',
                    }}
                >
                    <MenuItem value={'IN'}>India</MenuItem>
                    <MenuItem value={'US'}>United States</MenuItem>
                    <MenuItem value={'UK'}>United Kingdom</MenuItem>
                </Select>

                <br></br><br></br>

                <InputLabel htmlFor="time_zone" className={classes.dropdown}>Time Zone</InputLabel>

                <Select
                    className={classes.dropdown}
                    value={this.state.time_zone}
                    onChange={this.handleChange('time_zone')}
                    inputProps={{
                        name: 'time_zone',
                        id: 'time_zone',
                    }}
                >
                    <MenuItem value={'+5:30'}>( GMT + 5:30 ) India Standard Time (Asia/Kolkata)</MenuItem>
                    <MenuItem value={'-6:00'}>( GMT - 6:00 ) Central Standard Time (America/Guatemala)</MenuItem>
                    <MenuItem value={'+1:00'}>( GMT + 1:00 ) British Summer Time (Europe/London)</MenuItem>
                </Select>

                <br></br><br></br>

                <Button type="submit" variant="contained" color="default" className={classes.button}>
                    Save
                    <SaveIcon className={classes.rightIcon}/>
                </Button>

            </form>

        )
    }

}

PersonalInformationForm.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(PersonalInformationForm);