///////////////
// CONFIG
///////////////

const config = {
    shipServiceDomain: `http://ship.tiw.io`,
    firebaseApp: {
        apiKey: `AIzaSyA0lTT1L2CqT9X0pnRRPCiq6nTEvuOxISc`,
        authDomain: `mhaddon.firebaseapp.com`
    },
    firebaseUiConfig: {
        credentialHelper: firebaseui.auth.CredentialHelper.NONE,
        signInOptions: [
            firebase.auth.EmailAuthProvider.PROVIDER_ID
        ]
    },
    token: ``
};

firebase.initializeApp(config.firebaseApp);

const ui = new firebaseui.auth.AuthUI(firebase.auth());

///////////////
// GUI METHODS
///////////////

const createElement = (html) => {
    const div = document.createElement(`div`);
    div.innerHTML = html;
    return div.firstChild;
};

const addTextItem = (element, html) => {
    element.appendChild(createElement(html));
};

const updateGUIToSignedIn = (user) => {
    document.querySelectorAll(`.signedIn`).forEach((i) => i.style.display = `flex`);
    document.getElementById(`signedOut`).style.display = `none`;
    document.getElementById(`userEmail`).innerText = user.email;
};

const updateGUIToSignedOut = () => {
    document.querySelectorAll(`.signedIn`).forEach((i) => i.style.display = `none`);
    document.getElementById(`message`).innerText = ``;
    document.getElementById(`userEmail`).innerText = ``;
    document.getElementById(`signedOut`).style.display = `block`;
    ui.start(`#firebaseui-auth-container`, config.firebaseUiConfig);
};

const updateGUIWithNotifications = (notifications) => {
    const section = document.getElementById(`notifications`);
    section.innerHTML = ``;

    notifications.sort((a, b) => a.registration.id > b.registration.id);

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
};

///////////////
// ON EVENT METHODS
///////////////

const onSignedIn = (user) => {
    updateGUIToSignedIn(user);
    user.getIdToken().then((token) => config.token = token);
};

const onSignedOut = () => {
    updateGUIToSignedOut();
    config.token = ``;
};

const onAuthStateChanged = (user) => (user) ? onSignedIn(user) : onSignedOut();

///////////////
// STATE QUERY METHODS
///////////////

const isSignedOut = () => document.getElementById(`signedOut`).style.display !== `none`;

///////////////
// STATE CHANGE METHODS
///////////////

const signOut = () => firebase.auth().signOut();

///////////////
// VISUALISE NOTIFICATIONS
///////////////

const requestNotifications = () =>
    fetch(`${config.shipServiceDomain}/notifications`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    }).then(res => res.json());

const updateNotifications = () => {
    if (isSignedOut()) return false;

    requestNotifications()
        .then(updateGUIWithNotifications)
        .catch(err => {
            document.getElementById(`response`).style.display = `block`;
            document.getElementById(`response`).innerText = `${new Date().toISOString()} - Error pulling notifications: ${err}`;
        });
};

///////////////
// NOTIFICATIONS FORM
///////////////

const getNotificationFromForm = () => ({
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
});

const addNotification = (e) => {
    e.preventDefault();

    document.getElementById(`response`).style.display = `none`;

    const notification = getNotificationFromForm();

    fetch(`${config.shipServiceDomain}/notifications`, {
        method: `PUT`,
        headers: {
            "Content-Type": `application/json`,
            Authorization: `Bearer ${config.token}`
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
};

///////////////
// MAIN
///////////////

const defineTriggers = () => {
    document.getElementById(`signout`).onclick = signOut;
    document.getElementById(`form`).addEventListener(`submit`, addNotification, true);
    firebase.auth().onAuthStateChanged(onAuthStateChanged);

    updateNotifications();
    window.setInterval(updateNotifications, 1500);
};

window.addEventListener(`load`, defineTriggers);
