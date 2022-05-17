
function setStatus(text, className ) {
    // Show loading
    const statusIndicator = document.getElementById(`status-indicator`);

    if ( !!statusIndicator ) {
        // Clear all previous status-es
        for ( let [x, k] of statusIndicator.classList.entries() ) {
            statusIndicator.classList.remove(k);
        }

        // Mark as loading
        statusIndicator.textContent = text;
        statusIndicator.classList.add(className);
    }

}

function updatePost(postName) {

    const postContentElement = document.getElementById(`post_content`);

    if ( !!postContentElement && !!postContentElement.value ) {
        const postContent = String(postContentElement.value);
        if ( !!postContent && postContent.length > 0 ) {

            setStatus(`Loading...`, `has-text-warning`);

            // Generate the form data
            let formData = new FormData();
            formData.append('name',    postName);
            formData.append('content', postContent);

            // Send out a fetch request
            fetch("/api/post", {
                method: "post",
                body: formData,
            })
            .then( result => {

                if ( result.status < 200 || result.status > 299 ) {
                    if ( result.status == 429) {
                        setStatus(`Too many requests, please wait`, `has-text-danger`);
                    } else {
                        setStatus(`Failed to Save`, `has-text-danger`);
                    }
                } else {
                    setStatus(`Saved`, `has-text-success`);
                }

                console.log(result);

            })
            .catch( error => {
                console.error(error);
                alert(error);
                setStatus(`Failed to Save`, `has-text-danger`);
            })
        }
    }

}

/**
 * @location /
 * @role Searching
 */
function goToPost() {
    // Get the post name element
    const postNameElement = document.getElementById(`postName`);

    // Check if the element exists
    if ( !!postNameElement ) {
        // Get the post name string
        const postName = String(postNameElement.value);
        // Check if the post name is valid
        if ( !!postName && postName.length > 0 && postName.length <= 256 ) {
            // Change the location
            window.location.href = `/${postName}`;
        }
    }
}

function getQr(link = `https://justkato.me/`) {
    return new Promise((_r, _e) => {
        let formData = new FormData();
        formData.append('link', link);
        
        // Send out a fetch request
        fetch("/api/qr", {
            method: "post",
            body: formData,
        })
        .then( result => {
            result.json()
            .then( rez => {
                return _r(rez);
            })
        })
        .catch( error => {
            console.error(error);
            alert(error);
            return _e(error);
        })
    })
}