import React, { useContext } from 'react'; // Make sure to import useContext
import './Message.scss';
import UserContext from '../../utils/UserContext';

function Message(props) {
    const { userName} = useContext(UserContext);

    const isMessageFromCurrentUser = props.userName === userName;
    const messageClass = isMessageFromCurrentUser ? 'currentUser' : 'otherUser';

    return (
        <div className={`Message ${messageClass}`}>
            <div className='userName'>{props.userName}</div> 
            <div className='message'>{props.message}</div>
        </div>
    );
}

export default Message;
