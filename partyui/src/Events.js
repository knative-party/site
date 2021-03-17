import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import { Link as LinkRoute } from "react-router-dom";
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import logo from "./logo.svg";
import OnCall from './OnCall';
import GitHubIcon from '@material-ui/icons/GitHub';
import ForwardIcon from '@material-ui/icons/ArrowForward';
import BackIcon from '@material-ui/icons/ArrowBack';

function Copyright() {
    return (
        <React.Fragment>
          <Typography variant="body2" color="textSecondary" align="center">
              {'Copyright © '}
              <Link color="inherit" href="https://knative.dev/">
                  The Knative Authors
              </Link>{' '}
              {new Date().getFullYear()}
              {'.'}
          </Typography>
          <Typography variant="body2" color="textSecondary" align="center">
              <Link color="inherit" href="https://github.com/knative-party/site">
                <GitHubIcon />
              </Link>
          </Typography>
      </React.Fragment>
    );
}

function Logo() {
    return (
        <div className="App">
            <LinkRoute to="/party">
                <header className="App-header">
                    <img src={logo} className="App-logo" alt="logo" />
                </header>
            </LinkRoute>
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
    cardEvent: {
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

export default function Events() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);

    const [nowOn, setNowOn] = useState(new Date());

    const [items, setItems] = useState({support:[],events:[]});

    const classes = useStyles();

    const loadNow = (nowQuery) => {
        console.log("/now?"+nowQuery)
        fetch("/now?"+nowQuery)
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
    };

    const doForward = () => {
        let now = new Date(nowOn.setDate(nowOn.getDate()+7));
        setNowOn(now)
        loadNow("on="+now.toISOString())
    };

    const doBack = () => {
        let now = new Date(nowOn.setDate(nowOn.getDate()-7));
        setNowOn(now)
        loadNow("on="+now.toISOString())
    };

    // Note: the empty deps array [] means
    // this useEffect will run once
    // similar to componentDidMount()
    useEffect(loadNow, [])

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
                            <Grid item key={tier.title} xs={12} sm={12} md={4} lg={4}>
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
                        <Grid item key="back" xs={1} sm={1} md={1}>
                            <Button color="inherit" onClick={doBack}>
                                <BackIcon />
                            </Button>
                        </Grid>
                        <Grid item key="display" xs={10} sm={10} md={10}>
                            <Typography variant="h7" align="center" color="textSecondary" component="p">
                                Displaying week of {nowOn.toLocaleDateString()}
                            </Typography>
                        </Grid>
                        <Grid item key="next" xs={1} sm={1} md={1}>
                            <Button color="inherit" onClick={doForward}>
                                <ForwardIcon />
                            </Button>
                        </Grid>
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