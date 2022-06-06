
function sendMyData(el) {
    const formData = new FormData();

    // Check if the writing watch was sending something already
    if (!!window.writingWatch) {
        // Clear old timeout
        clearTimeout(window.writingWatch);
    }

    if (el.value.length > maximumPadSize) {
        let err = new Error(`Your Pad is too big! Please keep it limited to ${maximumPadSize} characters!`);
        alert(err);
        throw err;
    }

    el.setAttribute(`readonly`, `1`);

    const textareaPreview = document.getElementById(`textarea-preview`)
    if (!!textareaPreview) {
        textareaPreview.textContent = el.value;

        hljs.highlightElement(document.getElementById(`textarea-preview`));
    }

    formData.set("content", el.value);

    updateStatus(`Attempting to save...`, `text-warning`);

    fetch(window.location.href.toString(), {
        body: formData,
        method: "post",
    })
        .then(resp => {
            resp.json()
                .then(e => {
                    document.getElementById(`last_modified_`).value = e.pad.last_modified;
                    updateStatus(`Succesfully Saved`, `text-success`);
                })
                .catch(err => {
                    updateStatus(`Failed to Save`, `text-danger`);
                    console.error(err);
                })
        })
        .catch(err => {
            updateStatus(`Failed to Save`, `text-danger`);
            console.error(err);
        })
        .finally(() => {
            el.removeAttribute(`readonly`);
        })
}

function toggleWritingWatch(el) {

    // Check if the writing watch was sending something already
    if (!!window.writingWatch) {
        // Clear old timeout
        clearTimeout(window.writingWatch);
    }

    // Set a timeout for the action
    window.writingWatch = setTimeout(() => {
        // Send out the data
        sendMyData(el)
    }, 750)

}

function updateStatus(txt, cls) {

    const loading_status = document.getElementById(`loading_status`)

    loading_status.value = txt;
    loading_status.classList.remove("text-danger", "text-warning", "text-success", "text-white", "text-primary");
    loading_status.classList.add(cls);
}

function getLocalArchives() {

    let a = localStorage.getItem(`${padTitle}_archives`);

    // Check if we had anything in storage for the archives
    if (a == null) {
        // There were nothing in storage
        return [];
    }

    try {
        // Try and parse the json
        a = JSON.parse(a);
    } catch (err) {
        // Return null of the fail
        return [];
    }

    return a;
}

function storeArchives(archives) {

    // Check if the provided list is an array
    if (!Array.isArray(archives)) return;

    // Set the current archives
    localStorage.setItem(`${padTitle}_archives`, JSON.stringify(archives));
}

function renderArchivesSelection() {

    // Get the archives selection
    const archivesSelection = document.getElementById(`archives-selection`);
    const rowTemplate = document.getElementById(`archive-selection-example`);
    // Clear any old optiosn
    archivesSelection.querySelectorAll(`.dropdown-item:not(#do-archive-button):not(#archive-selection-example)`).forEach(el => {
        // Remove the element
        el.remove();
    })

    // Get the current list of available archives
    for (let a of getLocalArchives()) {
        // Clone the template row
        const row = rowTemplate.cloneNode(true);

        // Remove the id from the row
        row.removeAttribute(`id`);
        // Append the row to the selection menu
        archivesSelection.appendChild(row);

        const ts = new Date(a.ts);

        // Update the display date
        row.querySelector(`.archive-date`).textContent = ts.toLocaleString();

        // Add an event listener
        row.addEventListener(`click`, e => {

            let resp = confirm("Load contents of pad from memory? This will overwrite the current pad for everyone.");

            if (!!resp) {
                // Update visually for the client
                updatePadContent(a.content);
                // Send the update
                window.socket.sendPadUpdate();
            }
        })

    }

}

function saveLocalArchive() {
    let resp = confirm("Save a local copy of the current Pad?");

    if (!resp) {
        // Do not
        return;
    }

    // Get all of the previous archives, append this one to them
    let myArchives = getLocalArchives();

    myArchives.push({
        ts: new Date().getTime(),
        content: document.getElementById(`pad-content`).value,
    });

    // Store the archives
    storeArchives(myArchives);

    // Re-Render the archives selection
    renderArchivesSelection();

    // Save
    alert(`Saved`);

}

function generateQRCode() {
    var qrcodeContainer = document.getElementById(`qrcode`)
    // Remove old contents
    qrcodeContainer.innerHTML = "";
    // Add new qr
    new QRCode(qrcodeContainer, {
        text: window.location.toString(),
        width: 256,
        height: 256,
        colorDark: "#555273",
        colorLight: "#ffffff",
        correctLevel: QRCode.CorrectLevel.H
    });

    // Open the modal
    MicroModal.show(`qrmodal`)
}

function toggleTextareaPreview() {
    setTextareaPreview(!document.getElementById(`pad-content-toggler`).classList.contains(`read-only`))
}

// t == true - Read Only
// t == false - Edit mode
function setTextareaPreview(t = true) {
    const prev = document.getElementById(`textarea-preview`)
    const textarea = document.getElementById(`pad-content`);
    const toggler = document.getElementById(`pad-content-toggler`);
    const padContentArea = document.getElementById(`pad-content-area`);

    if (t) {
        // Toggle read only
        prev.classList.remove(`hidden`)
        toggler.classList.add(`read-only`);

        padContentArea.classList.add(`read-only-content`);

        prev.scrollTop = prev.scrollHeight;

        textarea.classList.add(`hidden`);
    } else {
        // Toggle edit mode
        prev.classList.add(`hidden`)
        toggler.classList.remove(`read-only`);

        padContentArea.classList.remove(`read-only-content`);


        textarea.classList.remove(`hidden`);
        // Focus
        textarea.focus();
        // Scroll
        textarea.scrollTop = textarea.scrollHeight;
        // Move cursor
        textarea.setSelectionRange(textarea.value.length, textarea.value.length);
    }

}

document.addEventListener(`DOMContentLoaded`, e => {

    { // Textarea Handling
        const textarea = document.getElementById(`pad-content`);
        setTextareaPreview(!!textarea.value);

        // Make sure tabs are taken into consideration
        textarea.addEventListener('keydown', function (e) {
            if (e.key == 'Tab') {
                e.preventDefault();
                const start = this.selectionStart;
                const end = this.selectionEnd;

                // set textarea value to: text before caret + tab + text after caret
                this.value = this.value.substring(0, start) +
                    "\t" + this.value.substring(end);

                // put caret at right position again
                this.selectionStart =
                    this.selectionEnd = start + 1;
            }
        });
    }

    try { // highlights
        hljs.highlightElement(document.getElementById(`textarea-preview`));
    } catch ( err ) {
        console.err(err);
    }

    { // Archives
        renderArchivesSelection()
    }

})