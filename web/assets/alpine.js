Sentry.onLoad(function () {
    Sentry.init({
        integrations: [
            Sentry.replayIntegration({
                maskAllText: false,
                blockAllMedia: false,
            }),
        ],
        // Session Replay
        replaysSessionSampleRate: 0.1, // This sets the sample rate at 10%. You may want to change it to 100% while in development and then sample at a lower rate in production.
        replaysOnErrorSampleRate: 1.0, // If you're not already sampling the entire session, change the sample rate to 100% when sampling sessions where errors occur.
    });
});

const urlInput = document.getElementById('url');
urlInput.addEventListener("keydown", function (event) {
    const urlValid = document.getElementById('url-valid');
    if (urlInput.validity.patternMismatch) {
        urlInput.setCustomValidity("Invalid URL");
        urlValid.textContent = "Invalid URL";
    } else {
        urlInput.setCustomValidity("");
        urlValid.textContent = "";
    }
});



// Alpine thing 

var userName = document.getElementById('user-name').innerText()

console.log(userName)



document.addEventListener("alpine:init", () => {
    Alpine.store("user", {
        id: "",
        username: "",
        setUsername(username) {
            this.username = username
        },
        setId(id) {
            this.id = id
        }
    })
}
)
