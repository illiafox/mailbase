function postForm() {
    let form = document.getElementById("loginform")
    document.getElementById("check_password").style.display = "none";
    switch (true) {
        // Password
        case !/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,128}$/.test(form.elements["password"].value):
            document.getElementById("check_password").innerHTML = "Wrong Password Format!";
            document.getElementById("check_password").style.display = "block";
            return;
        default:
            form.submit()
            break;
    }
}

