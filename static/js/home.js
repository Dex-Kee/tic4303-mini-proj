$(document).ready(function () {
    // Fetch country data
    var select = $("#country");
    $.ajax({
        url: "/api/info/country",
        dataType: "json",
        success: function(resp) {
            // Iterate through the country data and populate the <select> element
            $.each(resp.data, function(k, v) {
                var option = $("<option>");
                option.val(v);
                option.text(k);
                select.append(option);
            });
            console.log(resp.data);
        },
        error: function(error) {
            showToast("Failed to fetch country data.")
        }
    });


    // Fetch user data from the backend
    $.ajax({
        type: "GET",
        url: "/api/user/profile",
        headers: {
            "token": localStorage.getItem("token")
        },
        success: function (resp) {
            // Populate the form fields with fetched data
            console.log(resp);
            $("#name").val(resp.data.username);
            $("#email").val(resp.data.email);
            $("#phone").val(resp.data.phone);
            $("#country").val(resp.data.country);
            $("input[name='gender'][value='" + resp.data.gender + "']").prop("checked", true);
            $("#qualification").val(resp.data.qualification);
        },
        error: function () {
            // Handle errors (e.g., display an error message)
            showToast("Failed to fetch user profile data.");
        }
    });

    $("#profile-form").submit(function (e) {
        e.preventDefault();

        var gender = "M";
        genderSelector = $('input[name="gender"]:checked');
        if (genderSelector.length > 0) {
            gender = genderSelector.val();
        }

        // Serialize form data
        var data = {
            username: $("#name").val(),
            email: $("#email").val(),
            phone: $("#phone").val(),
            country: $("#country").val(),
            gender: gender,
            qualification: $("#qualification").val(),
        };

        // Send data to the backend
        $.ajax({
            type: "PUT",
            url: "/api/user/update",
            contentType: "application/json",
            headers: {
                "token": localStorage.getItem("token")
            },
            data: JSON.stringify(data),
            success: function (response) {
                // Handle the response from the backend (e.g., show a success message)
                console.log(response);
                if (response.code === 200) {
                    window.location.href = "/page/user/logout";
                } else {
                    showToast(response.message)
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