const shipServiceDomain = `http://ship.tiw.io`;

function requestNotifications(callback) {
    fetch(`${shipServiceDomain}/notifications`)
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
