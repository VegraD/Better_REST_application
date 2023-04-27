// syntax highlighter for json response was copied from:
// https://stackoverflow.com/questions/4810841/pretty-print-json-using-javascript
function syntaxHighlight(json) {
    if (typeof json != 'string') {
        json = JSON.stringify(json, undefined, 2);
    }
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        let cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}

// Get all the links in the nav element except the last one
//const links = document.querySelectorAll('nav li:not(:last-child) a');

// Add a click event listener to each link
links.forEach(link => {
    link.addEventListener('click', event => {
        // Prevent the default link behavior
        event.preventDefault();

        // Get the URL from the link href attribute
        const url = link.getAttribute('href');

        // Make a GET request to the URL
        fetch(url)
            .then(response => response.json())
            .then(data => {
                // Update the response container with the formatted response data
                const responseContainer = document.querySelector('#response-container');
                const formattedJson = JSON.stringify(data, null, 4);
                responseContainer.innerHTML = `<pre>${syntaxHighlight(formattedJson)}</pre>`;
            });
    });
});

