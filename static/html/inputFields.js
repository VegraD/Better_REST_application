// Get the input fields container and search button elements
const inputFieldsContainer = document.querySelector('#input-fields');
const searchButton = document.querySelector('#search-button');

// Set the minimum and maximum values for the begin and end input fields
document.querySelector('#begin').min = 1950;
document.querySelector('#begin').max = 2023;
document.querySelector('#end').min = 1950;
document.querySelector('#end').max = 2023;

// Set the default values for the begin and end input fields
document.querySelector('#begin').value = 2000;
document.querySelector('#end').value = 2005;

// Center the input fields container and search button in the top right quadrant
inputFieldsContainer.style.position = 'absolute';
inputFieldsContainer.style.top = '50%';
inputFieldsContainer.style.right = '50%';
inputFieldsContainer.style.transform = 'translate(50%, -50%)';
searchButton.style.position = 'absolute';
searchButton.style.top = 'calc(50% + 150px)';
searchButton.style.right = '50%';
searchButton.style.transform = 'translateX(50%)';

// Get all the links in the nav element except the last one
const links = document.querySelectorAll('nav li:not(:last-child) a');

// Add a click event listener to each link
links.forEach(link => {
    // Keep track of whether the link has been clicked or not
    let clicked = false;
    console.log(clicked);

    link.addEventListener('click', event => {
        // Prevent the default link behavior
        event.preventDefault();

        // Toggle the display of the input fields and search button depending on whether the link has been clicked or not
        if (clicked) {
            console.log(clicked);
            inputFieldsContainer.style.display = 'none';
            searchButton.style.display = 'none';
            clicked = false;
        } else {
            if (link.textContent === 'Current percentage of renewables') {
                inputFieldsContainer.style.display = 'block';
                document.querySelector('#country').parentElement.style.display = 'block';
                document.querySelector('#neighbours').parentElement.style.display = 'block';
                document.querySelector('#begin').parentElement.style.display = 'none';
                document.querySelector('#end').parentElement.style.display = 'none';
                document.querySelector('#sortByValue').parentElement.style.display = 'none';
                searchButton.style.display = 'inline-block';
            } else if (link.textContent === 'Historical percentages of renewables') {
                inputFieldsContainer.style.display = 'block';
                document.querySelector('#country').parentElement.style.display = 'block';
                document.querySelector('#neighbours').parentElement.style.display = 'none';
                document.querySelector('#begin').parentElement.style.display = 'block';
                document.querySelector('#end').parentElement.style.display = 'block';
                document.querySelector('#sortByValue').parentElement.style.display = 'block';
                searchButton.style.display = 'inline-block';
            } else {
                inputFieldsContainer.style.display = 'none';
                searchButton.style.display = 'none';

                // Make a GET request to the URL specified in the link href attribute
                fetch(link.getAttribute('href'))
                    .then(response => response.json())
                    .then(data => {
                        // Update the response container with the formatted response data
                        const responseContainer = document.querySelector('#response-container');
                        const formattedJson = JSON.stringify(data, null, 2);
                        responseContainer.innerHTML = `<pre>${syntaxHighlight(formattedJson)}</pre>`;
                    });
            }
            searchButton.setAttribute('data-href', link.getAttribute('href'));
            clicked = true;
        }
    });
});

// Add a click event listener to the search button
searchButton.addEventListener('click', () => {
    // Get the values from the input fields
    const inputFields = {
        country: document.querySelector('#country').value,
        neighbours: document.querySelector('#neighbours').checked,
        begin: parseInt(document.querySelector('#begin').value),
        end: parseInt(document.querySelector('#end').value),
        sortByValue: document.querySelector('#sortByValue').checked
    };

    // Construct the URL with the entered parameters
    let url = searchButton.getAttribute('data-href');
    if (url.includes('/energy/v1/renewables/current')) {
        url += `/?country=${inputFields.country}&neighbours=${inputFields.neighbours}`;
    } else if (url.includes('/energy/v1/renewables/history')) {
        url += `/?country=${inputFields.country}&begin=${inputFields.begin}&end=${inputFields.end}&sortByValue=${inputFields.sortByValue}`;
    }

    // Make a GET request to the URL
    fetch(url)
        .then(response => response.json())
        .then(data => {
            // Update the response container with the formatted response data
            const responseContainer = document.querySelector('#response-container');
            const formattedJson = JSON.stringify(data, null, 2);
            responseContainer.innerHTML = `<pre>${syntaxHighlight(formattedJson)}</pre>`;
        });
});
