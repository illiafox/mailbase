function postForm(id) {
    let form = document.getElementById(id)
    document.getElementById("check_mail").style.display = "none";
    switch (true) {
        // Mail
        case !/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(form.elements["mail"].value):
            document.getElementById("check_mail").innerHTML = "Wrong Email Format!";
            document.getElementById("check_mail").style.display = "block";
            return;
        default:
            form.submit()
            break;
    }
}

function redirect() {

}


