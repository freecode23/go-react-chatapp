
// init a connection to a chatroom
let initChatroomSocket = (chatroom, cb) => {

    // Create the WebSocket for the specific chatroom.
    let socket = new WebSocket(`ws://localhost:9000/start/${chatroom}`);
    

    // Sets up rules on what happen on various events on this socket
    // case1: sets up an event listener on the socket
    socket.onopen =  () => {
        console.log("socket starts listening")
    }

    // case2: receive message
    socket.onmessage = (msg) => {
        // console.log("socket message from webs:", msg)

        // do something with the message using the callback function
        cb(msg);
    }

    // case3: close socket
    socket.onclose = (event) => {
        console.log("socket closed connection:", event)
    }

     // case4: socket error
    socket.onerror = (err) => {
        console.log("socket error:", err)
    }

    // return an object chatroom socket
    // that can call socketSendMsg function
    return {
        socketSendMsg: (msg) => {
            console.log("socket sending message:", msg);
            socket.send(JSON.stringify(msg));
        },
        // Any other functions that operate on this socket can be added here.
    };
}



export {initChatroomSocket}