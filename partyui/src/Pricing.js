import React, { useState, useEffect } from 'react';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import StarIcon from '@material-ui/icons/StarBorder';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import logo from "./logo.svg";
import OnCall from './OnCall';

function Copyright() {
    return (
        <Typography variant="body2" color="textSecondary" align="center">
            {'Copyright © '}
            <Link color="inherit" href="https://knative.dev/">
                The Knative Authors
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    );
}

function Logo() {
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
            </header>
        </div>
    );
}

const useStyles = makeStyles((theme) => ({
    '@global': {
        ul: {
            margin: 0,
            padding: 0,
            listStyle: 'none',
        },
    },
    appBar: {
        borderBottom: `1px solid ${theme.palette.divider}`,
    },
    toolbar: {
        flexWrap: 'wrap',
    },
    toolbarTitle: {
        flexGrow: 1,
    },
    link: {
        margin: theme.spacing(1, 1.5),
    },
    heroContent: {
        padding: theme.spacing(8, 0, 6),
    },
    cardHeader: {
        backgroundColor:
            theme.palette.type === 'light' ? theme.palette.grey[200] : theme.palette.grey[700],
    },
    cardPricing: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'baseline',
        marginBottom: theme.spacing(2),
    },
    footer: {
        borderTop: `1px solid ${theme.palette.divider}`,
        marginTop: theme.spacing(8),
        paddingTop: theme.spacing(3),
    },
}));
//
// const tiers = [{
//   title: 'Serving',  // https://github.com/knative/serving/blob/master/support/COMMUNITY_CONTACTS.md
//   onCall: {
//     name: '@nak3',
//     start: 'Feb 15, 2021',
//     end: 'Feb 19, 2021',
//     github: 'https://github.com/nak3',
//     questions: "#serving-questions",
//     questionsSlack: "https://knative.slack.com/archives/C0186KU7STW",
//   },
// }, {
//   title: 'Eventing', // https://github.com/knative/eventing/blob/master/support/COMMUNITY_CONTACTS.md
//   onCall: {
//     name: '@matzew',
//     start: 'Feb 15, 2021',
//     end: 'Feb 19, 2021',
//     github: 'https://github.com/matzew',
//     questions: "#eventing-questions",
//     questionsSlack: "https://knative.slack.com/archives/C017X0PFC0P",
//   },
// }];

const TOC = { // https://docs.google.com/document/d/1LzOUbTMkMEsCRfwjYm5TKZUWfyXpO589-r9K2rXlHfk/edit#heading=h.jlesqjgc1ij3
  title: "ToC Working Group Update",
  wg: 'Eventing WG',
  date: 'Feb 18, 2021 @ 8:30 – 9:15am PST',
};

export default function Pricing() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState({support:[],events:[]});

    const classes = useStyles();

    // Note: the empty deps array [] means
    // this useEffect will run once
    // similar to componentDidMount()
    useEffect(() => {
        fetch("/now")
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    setItems(result);
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }, [])

    if (error) {
        return (
            <React.Fragment>
                <CssBaseline/>
                {/* Hero unit */}
                <Container maxWidth="sm" component="main" className={classes.heroContent}>
                    <Logo/>
                    <Typography variant="h5" align="center" color="textSecondary" component="p">
                        Error: {error.message}
                    </Typography>
                </Container>
            </React.Fragment>
        )
    } else if (!isLoaded) {
        return (
            <React.Fragment>
                <CssBaseline/>
                {/* Hero unit */}
                <Container maxWidth="sm" component="main" className={classes.heroContent}>
                    <Logo/>
                    <Typography variant="h5" align="center" color="textSecondary" component="p">
                        loading...
                    </Typography>
                </Container>
            </React.Fragment>
        )
    } else {
        const tiers = items.support;
        const events = items.events;

        return (
            <React.Fragment>
                <CssBaseline/>
                {/* Hero unit */}
                <Container maxWidth="sm" component="main" className={classes.heroContent}>
                    <Logo/>
                    <Typography variant="h5" align="center" color="textSecondary" component="p">
                        Time to Party! 🎉🎉
                    </Typography>
                </Container>
                <div></div>
                {/* End hero unit */}
                <Container maxWidth="xlg" component="main">
                    <Grid container spacing={5} alignItems="flex-end">
                        {tiers.map((tier) => (
                            <Grid item key={tier.title} xs={12} sm={6} md={6}>
                                <Card>
                                    <CardHeader
                                        title={tier.title}
                                        subheader={tier.subheader}
                                        titleTypographyProps={{align: 'center'}}
                                        subheaderTypographyProps={{align: 'center'}}
                                        className={classes.cardHeader}
                                    />
                                    <CardContent>
                                        <OnCall onCall={tier.onCall}/>
                                    </CardContent>
                                    <CardActions>
                                        <Button fullWidth variant={tier.buttonVariant} color="primary">
                                            {tier.buttonText}
                                        </Button>
                                    </CardActions>
                                </Card>
                            </Grid>
                        ))}
                        {events.map((event) => (
                        <Grid item key={event.title} xs={12} sm={12} md={12}>
                            <Card>
                                <CardHeader
                                    title={event.title}
                                    subheader={event.subheader}
                                    titleTypographyProps={{align: 'center'}}
                                    subheaderTypographyProps={{align: 'center'}}
                                    className={classes.cardHeader}
                                />
                                <CardContent>
                                    <Typography component="h5" variant="h4" align="center" color="textPrimary"
                                                gutterBottom>
                                        {event.wg}
                                    </Typography>
                                    <Typography variant="h7" align="center" color="textSecondary" component="p">
                                        {event.when}
                                    </Typography>
                                </CardContent>
                                <CardActions>
                                    <Button fullWidth variant={event.buttonVariant} color="primary">
                                        {event.buttonText}
                                    </Button>
                                </CardActions>
                            </Card>
                        </Grid>
                        ))}
                    </Grid>
                </Container>
                {/* Footer */}
                <Container maxWidth="md" component="footer" className={classes.footer}>
                    <Box mt={5}>
                        <Copyright/>
                    </Box>
                </Container>
                {/* End footer */}
            </React.Fragment>
        );
    }
}