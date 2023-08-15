// initiate handshake
var socket = new WebSocket('ws://localhost:9000/start')


// init connect function
let connect = (cb) => {
    console.log("connecting")


    // sets up an event listener on the socket
    socket.onopen =  () => {
        console.log("socket starts listening")
    }


    // receive message
    socket.onmessage = (msg) => {
        console.log("socket message from webs:", msg)

        // do something with the message using the callback function
        cb(msg);
    }

    socket.onclose = (event) => {
        console.log("socket closed connection:", event)
    }

    socket.onerror = (err) => {
        console.log("socket error:", err)
    }
}

let sendMsg = (msg) => {
    console.log("socket sending message:",msg)
    socket.send(msg)
}

export {connect, sendMsg}