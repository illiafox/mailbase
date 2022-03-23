function postForm(id) {
    let form = document.getElementById(id)
    document.getElementById("check_password").style.display = "none";
    document.getElementById("check_retype").style.display = "none";
    document.getElementById("check_mail").style.display = "none";
    switch (true) {
        // Mail
        case !/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(form.elements["mail"].value):
            document.getElementById("check_mail").innerHTML = "Wrong Email Format!";
            document.getElementById("check_mail").style.display = "block";
            return;
        // Password
        case !/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,128}$/.test(form.elements["password"].value):
            document.getElementById("check_password").innerHTML = "Wrong Password Format!";
            document.getElementById("check_password").style.display = "block";
            return;
        // Retype
        case (form.elements["password"].value !== form.elements["retype_password"].value):
            document.getElementById("check_retype").innerHTML = "Fields don't match";
            document.getElementById("check_retype").style.display = "block";
            return;
        default:
            form.submit()
            break;
    }
}
