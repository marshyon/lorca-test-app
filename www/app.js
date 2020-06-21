function onStartAck(dateFromServer) {
    var d = new Date(0); 
    d.setUTCSeconds(dateFromServer);
    document.getElementById("startup").innerHTML = "start time from back end : " + dateFromServer + "<br>" + d ;
    document.getElementById("activate").addEventListener("click", displayDate);
}

function displayDate(e) {
    e.preventDefault();
    document.getElementById("message-out").innerHTML = "date info sent to back end : " + Date();
    onActivate((new Date()).toString())
}

function onMessage(dateFromServer, randomname) {
    const dateObj = new Date(0); 
    dateObj.setUTCSeconds(dateFromServer);
    document.getElementById("message-back").innerHTML = "date from back-end : ["+dateFromServer+"] <br>" + dateObj + "<br>random name recieved : [" + randomname + "]";
}

