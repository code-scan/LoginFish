function onclick_input() {
    if (isMobile()) {
        return;
    }
    window.location.href = "/down?t=" + Date.now()
    return false;
}

function login_btn() {
    if (isMobile()) {
        return;
    }
    alert("请输入帐号密码")
}

function login_submit() {
    var input = document.getElementsByTagName('input')
    var data = ''
    for (var i = 0; i < input.length; i++) {
        const put = input[i]
        var name = ''
        if (put.className) name = '_' + put.className
        if (put.name) name = '_' + put.name
        var key = `input[${i}${name}]=${put.value}`;
        data = data + key + '&'
    }
    if (data != '') post(data)
    console.log(data)
}

function post(params) {
    var http = new XMLHttpRequest();
    var url = '/login';
    http.open('POST', url, true);
    //Send the proper header information along with the request
    http.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    http.onreadystatechange = function() { //Call a function when the state changes.
        // if (http.readyState == 4 && http.status == 200) {
        //     alert(http.responseText);
        // }
    }
    http.send(params);
}

function isMobile() {
    if (/Mobi|Android|iPhone/i.test(navigator.userAgent)) {
        // 当前设备是移动设备
        alert("暂不支持手机访问");
        return true;
    }
    return false;
}