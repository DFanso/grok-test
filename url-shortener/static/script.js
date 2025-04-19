async function shortenURL() {
    const urlInput = document.getElementById('urlInput');
    const resultDiv = document.getElementById('result');
    const url = urlInput.value.trim();

    if (!url) {
        resultDiv.innerHTML = '<p style="color: red;">Please enter a URL to shorten.</p>';
        return;
    }

    try {
        const response = await fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url }),
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        resultDiv.innerHTML = `
            <p>Shortened URL: <a href="${data.short_url}" target="_blank">${data.short_url}</a></p>
        `;
        urlInput.value = ''; // Clear input after successful shortening
    } catch (error) {
        resultDiv.innerHTML = `<p style="color: red;">Error: ${error.message}</p>`;
    }
}

// Add event listener for Enter key
document.getElementById('urlInput').addEventListener('keypress', function(e) {
    if (e.key === 'Enter') {
        shortenURL();
    }
});
