import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import QuestionAnswerIcon from '@material-ui/icons/QuestionAnswer';

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        maxWidth: 360,
        backgroundColor: theme.palette.background.paper,
    },
}));

export default function FolderList(props) {
    const classes = useStyles();
    const onCall = props.onCall;

    console.log(`${onCall.github}.png?size=60`)
    return (
        <List className={classes.root}>
            <ListItem button onClick={() => window.open(onCall.github)}>
                <ListItemAvatar>
                    <Avatar src={`${onCall.github}.png?size=60`} />
                </ListItemAvatar>
                <ListItemText primary={onCall.name} secondary={`On-Call through ${onCall.end}`} />
            </ListItem>
            <ListItem button onClick={() => window.open(onCall.questionsSlack)}>
                <ListItemAvatar>
                    <Avatar>
                        <QuestionAnswerIcon />
                    </Avatar>
                </ListItemAvatar>
                <ListItemText primary={onCall.questions} secondary="Knative Slack" />
            </ListItem>
        </List>
    );
}