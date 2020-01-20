const shipServiceDomain = `http://ship.tiw.io`;
firebase.initializeApp({
    apiKey: "AIzaSyA0lTT1L2CqT9X0pnRRPCiq6nTEvuOxISc",
    authDomain: "mhaddon.firebaseapp.com",
});

const uiConfig = {
    credentialHelper: firebaseui.auth.CredentialHelper.NONE,
    signInOptions: [
        firebase.auth.EmailAuthProvider.PROVIDER_ID
    ]
};

const ui = new firebaseui.auth.AuthUI(firebase.auth());

function signedIn(user) {
    document.querySelectorAll('.signedIn').forEach((i) => i.style.display = 'flex');
    document.getElementById('signedOut').style.display = 'none';
    document.getElementById('userEmail').innerText = user.email;

    user.getIdToken().then((token) => window.token = token);
}

function signedOut() {
    document.querySelectorAll('.signedIn').forEach((i) => i.style.display = 'none');
    document.getElementById('message').innerText = '';
    document.getElementById('userEmail').innerText = '';
    document.getElementById('signedOut').style.display = 'block';
    ui.start('#firebaseui-auth-container', uiConfig);
}

window.addEventListener('load', function() {
    document.getElementById('signout').onclick = function() {
        firebase.auth().signOut();
    };

    firebase.auth().onAuthStateChanged((user) => {
        if (user) {
            signedIn(user);
        } else {
            signedOut();
        }
    });
});

function requestNotifications(callback) {
    fetch(`${shipServiceDomain}/notifications`, {
        headers: {
            Authorization: 'Bearer ' +  window.token
        }
    })
        .then(res => res.json())
        .then(callback)
        .catch(err => {
            document.getElementById(`response`).style.display = `block`;
            document.getElementById(`response`).innerText = `${new Date().toISOString()} - Error pulling notifications: ${err}`;
        });
}

function createElement(html) {
    const div = document.createElement('div');
    div.innerHTML = html;
    return div.firstChild;
}

function addTextItem(element, html) {
    element.appendChild(createElement(html));
}

function updateNotifications() {
    if (document.getElementById('signedOut').style.display !== 'none') {
        return false;
    }

    requestNotifications((notifications) => {
        const section = document.getElementById(`notifications`);
        section.innerHTML = "";

        notifications.sort(function (a, b) {
            return a.registration.id > b.registration.id;
        });

        notifications.map(notification => {
            const p = document.createElement(`p`);
            addTextItem(p, `<span>ID: ${notification.registration.id}</span>`);
            addTextItem(p, `<span>Name: ${notification.registration.name}</span>`);
            addTextItem(p, `<span>Captain: ${notification.registration.captain}</span>`);
            addTextItem(p, `<span>Lat: ${notification.status.position.lat}</span>`);
            addTextItem(p, `<span>Long: ${notification.status.position.long}</span>`);
            addTextItem(p, `<span>Velocity: ${notification.status.velocity}</span>`);
            section.appendChild(p);
        });
    });
}

document.getElementById(`form`).addEventListener("submit", addNotification, true);

function addNotification(e) {
    e.preventDefault();

    document.getElementById(`response`).style.display = `none`;

    const notification = {
        registration: {
            id: document.getElementById(`id`).value,
            name: document.getElementById(`name`).value,
            captain: document.getElementById(`captain`).value
        },
        status: {
            position: {
                lat: document.getElementById(`lat`).value,
                long: document.getElementById(`lat`).value
            },
            velocity: document.getElementById(`velocity`).value
        }
    };

    fetch(`${shipServiceDomain}/notifications`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            Authorization: 'Bearer ' +  window.token
        },
        body: JSON.stringify(notification),
    })
        .then((response) => response.json())
        .then((data) => {
            document.getElementById(`response`).style.display = `block`;
            document.getElementById(`response`).innerText = `${new Date().toISOString()} - Success: ${data}`;
        })
        .catch((error) => {
            document.getElementById(`response`).style.display = `block`;
            document.getElementById(`response`).innerText = `${new Date().toISOString()} - Error saving notifications: ${error}`;
        });
}

updateNotifications();
window.setInterval(updateNotifications, 1500);
