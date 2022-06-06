class PadSocket {

    /**
     * @type {WebSocket}
     */
    ws      = null;
    /**
     * @type {String}
     */
    padName = null;

    /**
     * The actual textarea you write in
     * @type {HTMLTextAreaElement}
     */
    padContents = null;
    /**
     * The <code> of the preview
     * @type {HTMLElement}
     */
    padPreview = null;

    /**
     * Create a new PadSocket
     * @param {string} padName The name of the pad
     * @param {string} connUrl The URL to the websocket
     */
    constructor(padName, connUrl = null) {
        // Assign the pad name
        this.padName = padName;

        // Check if a connection URL was mentioned
        if ( connUrl == null ) {
            // Try and connect to the local websocket
            connUrl = `ws://` + window.location.host + `/ws/get/${padName}`;
        }

        // Connect to the websocket
        const ws = new WebSocket(connUrl);

        // Bind the onMessage function
        ws.onmessage = this.handleMessage;

        ws.onopen = () => {
            updateStatus(`Established`, `text-success`);
        }

        // Try and reconnect on failure
        ws.onclose = connectSocket;
        ws.onerror = connectSocket;
        
        // Assign the websocket
        this.ws = ws;

        // Get all relevant references from the HTML
        this.padContents = document.getElementById(`pad-content`);
        this.padPreview  = document.getElementById(`textarea-preview`);
    }

    /**
     * @description Send a message to the server
     * @param {string} eventType The type of event, this can be anything really, it's just used for routing by the server
     * @param {Object} message The message to send out to the server, this can only be of format string but JSON is parsed.
     */
    sendMessage = (eventType, message) => {

        if ( this.ws.readyState !== WebSocket.OPEN ) {
            throw new Error(`The websocket connection is not active`);
        }

        // Check if the message is a string
        if ( typeof message !== 'object' ) {
            // Convert the message into a map[string]interface{}
            message = {
                "message": message,
            };
        }

        // TODO: Compress the message, usually we will be sending the whole body of the pad from the client to the server or vice-versa.
        this.ws.send( JSON.stringify({
            eventType,
            padName: this.padName,
            message,
        }))

    }

    /**
     * Handle the message from the socket based on the message type
     * @param {MessageEvent} e The websocket message
     */
    handleMessage = ev => {
        updateStatus(`Catching Message`, `text-white`);

        // Check if the message has valid data
        if ( !!ev.data ) {
            // Try and parse the data
            let parsedData = null;

            try { 
                parsedData = JSON.parse(ev.data);
            } catch ( err ) {
                console.error(`Failed to parse the WebSocket data`,err);
                updateStatus(`Parse Fail`, `text-warning`);
            }

            if ( !!!parsedData['message'] ) {
                console.error(`Failed to find the message`)
                updateStatus(`Message Fail`, `text-warning`);
                return;
            }

            // Check if this is a pad Content Update
            if ( parsedData['eventType'] === `padUpdate`) {
                // Pass on the parsed data
                this.onPadUpdate(parsedData);
            } // Check if this is a pad  Status Update 
            else if ( parsedData['eventType'] === `statusUpdate`) {
                // Pass on the parsed data
                this.onStatusUpdate(parsedData);
            }

            updateStatus(`Established`, `text-success`);
        }
    }

    /**
     * Whenever a pad update is trigered, run this function
     * @param {Object} The response from the server
     */
    onPadUpdate = data => {
        // Check that the content is clear
        if ( !!data['message']['content'] ) {
            // Send over the new content to be updated.
            updatePadContent(data['message']['content']);
        }
    }

    onStatusUpdate = data => {
        // Check that the content is clear
        if ( !!data['message']['currentViewers'] ) {
            // Get the amount of viewers reported by the server
            const viewerCount = Number(data['message']['currentViewers']);
            // Check if this is a valid number
            if ( Number.isNaN(viewerCount) ) {
                // Looks like this is a malformed message
                return console.error(`Malformed Message`, data);
            }

            // Send over the new content to be updated.
            updatePadViewers(viewerCount);
        }
    }

    /**
     * Sending a pad update for each keystroke to the server.
     * @param {String} msg The new contents of the pad
     */
    sendPadUpdate = msg => {
        // Get the contents of the pad
        const padContents = this.padContents.value;

        // Send the data over the webSocket
        this.sendMessage(`padUpdate`, {
            "content": padContents,
        });

        updatePadContent(padContents, false);
    }

}

/**
 * Update the contents of the pad
 * @param {String} newContent 
 */
function updatePadContent(newContent, textArea = true) {
    // Update the textarea
    if ( textArea ) document.getElementById(`pad-content`).value = newContent;
    // Update the preview
    document.getElementById(`textarea-preview`).innerHTML = newContent;
    // TODO: Re-run the syntax highlight

}

function updatePadViewers(vc) {
    // Get the reference to the viewers count inputElement
    /**
     * @type {HTMLInputElement}
     */
    const viewerCount = document.getElementById(`currentViewers`);

    // Get the amount of total viewers
    const totalViews = viewerCount.value.split("|")[1].trim();

    // Set back the real value
    viewerCount.value = `${vc} | ${totalViews}`;
}

function connectSocket() {
    // Check if the socket is established
    if ( !!!window.socket || window.socket.readyState !== WebSocket.OPEN ) {
        updateStatus(`Connecting...`, `text-warning`);
        // Connect the socket
        window.socket = new PadSocket(padTitle);
    }

}

// TODO: Test if this is actually necessary or the DOMContentLoaded event would suffice
// wait for the whole window to load
window.addEventListener(`load`, e => {
    connectSocket()
})