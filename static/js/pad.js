class Pad {

    title = '';
    content = '';
    timestmap = '';

    constructor(t, ts) {
        this.title = t;
        this.content = document.getElementById(`pad-content`).value;
        this.timestmap = ts;
    }

    downloadPadContents() {
        // Create a new blob of the contents of the pad
        var blob = new Blob([ document.getElementById(`pad-content`).value ], { type: "text/plain;charset=utf-8" });

        // Save the blob as
        saveAs(blob, `${this.title}.txt`);
    }

}