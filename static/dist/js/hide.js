function GetSite() {
    var data = $.ajax({
        url: "/v1/get_site",
        async: false
    })
    return data.responseJSON;
}

function GetLog(id) {
    var data = $.ajax({
        url: "/v1/get_log?id=" + id,
        async: false
    })
    return data.responseJSON;
}

function GetTemplate() {
    var data = $.ajax({
        url: "/v1/get_template",
        async: false
    })
    return data.responseJSON;
}