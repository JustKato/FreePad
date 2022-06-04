class PadSocket {

    ws      = null;
    padName = null;
    state   = null;

    /**
     * Create a new PadSocket
     * @param {string} padName The name of the pad
     * @param {string} connUrl The URL to the websocket
     */
    constructor(padName, connUrl = null) {

        // Check if a connection URL was mentioned
        if ( connUrl == null ) {
            // Try and connect to the local websocket
            connUrl = `ws://` + window.location.host + "/ws/get";
        }

        // Connect to the websocket
        const ws = new WebSocket(connUrl);
        ws.onopen = () => {
            this.state = 'active';
        }

        // Bind the onMessage function
        ws.onmessage = this.handleMessage;

        // Assign the websocket
        this.ws = ws;
        // Assign the pad name
        this.padName = padName;
    }

    /**
     * @description Send a message to the server
     * @param {string} eventType The type of event, this can be anything really, it's just used for routing by the server
     * @param {Object} message The message to send out to the server, this can only be of format string but JSON is parsed.
     */
    sendMessage = (eventType, message) => {

        if ( this.state != 'active' ) {
            throw new Error(`The websocket connection is not active`);
        }

        // Check if the message is a string
        if ( typeof message == 'string' ) {
            // Convert the message into a map[string]interface{}
            message = {
                "message": message,
            };
        }

        this.ws.send( JSON.stringify({
            eventType,
            padName: this.padName,
            message,
        }))

    }

    handleMessage = ev => {
        console.log(ev);
    }

}

// wait for the whole window to load
window.addEventListener(`load`, e => {
    window.socket = new PadSocket(padTitle);
})