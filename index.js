let sessid = 0;
async function request(url) {
    let res = await fetch(url, {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            headers: {
                'Content-Type': 'application/json',
                'Session-Id' : sessid,
            },
            redirect: 'follow',
            referrerPolicy: 'no-referrer',
        });

    let inf = await res;
    console.log(inf);
    let jso = await res.text();
    let msg = jso.split(";");
    if (msg[0] === "666") {
        Swal.fire({
            title: 'Request failed',
            text: msg[1],
            icon: 'error',
            confirmButtonText: 'Continue'
        })
    } else {
        Swal.fire({
            title: 'Request success!',
            text: jso,
            icon: 'success',
            confirmButtonText: 'Continue'
        })
    }

}


async function login(url, data) {
    let res = await fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: data,
    });

    let inf = await res;
    console.log(inf);

    let jso = await res.text();
    Swal.fire({
        title: 'Request',
        text: jso,
        icon: 'success',
        confirmButtonText: 'Continue'
    })
    sessid = parseInt(jso)

}

$(document).ready(function(){
    $("#getbut").click(function(){
        request("http://127.0.0.1:5000/data");
    });

    $("#logbut").click(function(){
        login("http://127.0.0.1:5000/login", JSON.stringify({login: $("#login").val(), pass: $("#pass").val()}));
    });
});