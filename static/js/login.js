$(document).ready(function () {
    $("form").submit(function (e) {
        e.preventDefault();
        var username = $("#username").val();
        var password = $("#password").val();
        var data = {
            username: username,
            password: password
        };

        // Perform AJAX request
        $.ajax({
            type: "POST",
            contentType: "application/json",
            url: "/api/user/login",
            data: JSON.stringify(data),
            success: function (response) {
                try {
                    // Parse the JSON response
                    if (response.code === 200) {
                        // Response indicates success, extract the 'data' field
                        var data = response.data;
                        // store the token & redirect to home page
                        localStorage.setItem("token", data);
                        window.location.href = "/page/user/home";
                    } else {
                        showToast(response.message)
                    }
                } catch (Exception) {
                    showToast(Exception)
                }
            },
            error: function (response) {
                // Parse the JSON response
                if (response.responseJSON != null) {
                    showToast(response.responseJSON.message)
                } else {
                    showToast(response.responseText)
                }
            }
        });
    });
});