import React from 'react';
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

const tiers = [{
  title: 'Serving',  // https://github.com/knative/serving/blob/master/support/COMMUNITY_CONTACTS.md
  onCall: {
    name: '@julz',
    start: 'Jan 25, 2021',
    end: 'Jan 29, 2021',
    github: 'https://github.com/julz',
    questions: "#serving-questions",
    questionsSlack: "https://knative.slack.com/archives/C0186KU7STW",
  },
}, {
  title: 'Eventing', // https://github.com/knative/eventing/blob/master/support/COMMUNITY_CONTACTS.md
  onCall: {
    name: '@lionelvillard',
    start: 'Jan 25, 2021',
    end: 'Jan 29, 2021',
    github: 'https://github.com/lionelvillard',
    questions: "#eventing-questions",
    questionsSlack: "https://knative.slack.com/archives/C017X0PFC0P",
  },
}];

const TOC = { // https://docs.google.com/document/d/1LzOUbTMkMEsCRfwjYm5TKZUWfyXpO589-r9K2rXlHfk/edit#heading=h.jlesqjgc1ij3
  title: "ToC Working Group Update",
  wg: 'Scaling WG', 
  date: 'Jan 28, 2021 @ 8:30 – 9:15am PST',
};

export default function Pricing() {
    const classes = useStyles();

    return (
        <React.Fragment>
            <CssBaseline />
            {/* Hero unit */}
            <Container maxWidth="sm" component="main" className={classes.heroContent}>
                <Logo />
                <Typography variant="h5" align="center" color="textSecondary" component="p">
                    Time to Party! 🎉🎉
                </Typography>
            </Container>
            {/* End hero unit */}
            <Container maxWidth="xlg" component="main">
                <Grid container spacing={5} alignItems="flex-end">
                    {tiers.map((tier) => (
                        <Grid item key={tier.title} xs={12} sm={6} md={6}>
                            <Card>
                                <CardHeader
                                    title={tier.title}
                                    subheader={tier.subheader}
                                    titleTypographyProps={{ align: 'center' }}
                                    subheaderTypographyProps={{ align: 'center' }}
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
                    <Grid item key={TOC.title} xs={12} sm={12} md={12}>
                        <Card>
                            <CardHeader
                                title={TOC.title}
                                subheader={TOC.subheader}
                                titleTypographyProps={{ align: 'center' }}
                                subheaderTypographyProps={{ align: 'center' }}
                                className={classes.cardHeader}
                            />
                            <CardContent>
                              <Typography component="h5" variant="h4" align="center" color="textPrimary" gutterBottom>
                                {TOC.wg}
                              </Typography>
                              <Typography variant="h7" align="center" color="textSecondary" component="p">
                                {TOC.date}
                              </Typography>
                            </CardContent>
                            <CardActions>
                                <Button fullWidth variant={TOC.buttonVariant} color="primary">
                                    {TOC.buttonText}
                                </Button>
                            </CardActions>
                        </Card>
                    </Grid>
                </Grid>
            </Container>
            {/* Footer */}
            <Container maxWidth="md" component="footer" className={classes.footer}>
                <Box mt={5}>
                    <Copyright />
                </Box>
            </Container>
            {/* End footer */}
        </React.Fragment>
    );
}