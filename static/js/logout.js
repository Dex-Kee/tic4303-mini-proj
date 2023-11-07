$(document).ready(function () {
    $("form").submit(function (e) {
        e.preventDefault();
        $.ajax({
            type: "PUT",
            contentType: "application/json",
            url: "/api/user/logout",
            success: function (response) {
                try {
                    // Parse the JSON response
                    if (response.code === 200) {
                        window.location.href = "/page/user/login";
                    } else {
                        showToast(response.message)
                    }
                } catch (e) {
                    console.log("Logout failed: " + e);
                    showToast("Logout failed");
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