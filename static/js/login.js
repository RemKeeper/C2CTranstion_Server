"use strict";
if (document.cookie !== '') {
    getPermission(function (data) {
        if (data.code === 200) {
            if (data.level <= 4) {
                setTimeout(function () {
                    window.location.href = '/static/UserInfo.html';
                });
            } else {
                setTimeout(function () {
                    window.location.href = '/static/admin.html';
                });
            }
        } else {
            Toast("提示", "服务器异常");
        }
    });
}
// 是否发送了，防止多次请求
let _isLogin = false;
// 绑定到注册按钮
function Login() {
    if (_isLogin) return;
    // 用户名
    let userName = $("#UserName").val();
    // 密码
    let password = $("#Password").val();


    if (userName === "" || password === "") {
        Toast("提示", "用户名或密码不可为空");
        return;
    }

    _isLogin = true;
    ajax({
        method: 'POST',
        url: `${ServerAddr}/login`,
        data: {
            user_name: userName,
            pwd_summary: md5(password),
        },
        handle: function (data) {
            if (data.code === 200) {
                getPermission(function (data) {
                    if (data.code === 200) {
                        if (data.level <= 4) {
                            setTimeout(function () {
                                window.location.href = '/static/UserInfo.html';
                            });
                        } else {
                            setTimeout(function () {
                                window.location.href = '/static/admin.html';
                            });
                        }
                    } else {
                        Toast("提示", "服务器异常");
                    }
                });
            } else {
                Toast("提示", data.msg);
            }
            _isLogin = false;

        },
    })

}
