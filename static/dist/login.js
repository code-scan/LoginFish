function onclick_input(){
    if(isMobile()){
        return;
    }
    window.location.href="/down?t="+Date.now()
    return false;
}
function login_btn(){
    if(isMobile()){
        return;
    }
    alert("请输入帐号密码")
}
function isMobile(){
    if (/Mobi|Android|iPhone/i.test(navigator.userAgent)) {
        // 当前设备是移动设备
        alert("暂不支持手机访问");
        return true;
    }
    return false;
}


