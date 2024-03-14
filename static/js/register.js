"use strict";
// function getUrl(e,param){
//     analyticCode.getUrl(param,e,function(url1,url2){
//         e.previousElementSibling.innerHTML=url1
//         // console.log("url1",url1)
//         // console.log("url2",url2)
//         // e.previousElementSibling.src = url2;
//     });
// }

// 是否发送了，防止多次请求
let isRegister = false;
// 绑定到注册按钮
function Register() {
    if (isRegister) return;
    // 用户名
    let userName = $("#UserName").val();
    // 密码
    let password = $("#Password").val();
    // 邀请码
    let inviteID = $("#InviteID").val();


    if (userName === "" || password === "") {
        Toast("提示", "用户名或密码不可为空");
        return;
    }

    isRegister = true;
    ajax({
        method: 'POST',
        url: `${ServerAddr}/register`,
        data: {
            user_name: userName,
            pwd_summary: md5(password),
            invitation_code: parseInt(inviteID) ,
        },
        handle: function (data) {
            Toast("提示", data.msg, '确认', function (inst) {
                inst.close();
                setTimeout(function () {
                    window.location.href = '/static/UserInfo.html';
                });
            });
            isRegister = false;

        },
    })

}


function getInvitationCode(url) {
    let urlObj = new URL(url);
    let params = new URLSearchParams(urlObj.search);
    let invitationCode = params.get('invitation_code');
    return invitationCode;
}


// 网页加载完成后绑定到邀请码输入框

window.onload = function () {
    let invitationCode = getInvitationCode(window.location.href);
    document.getElementById("InviteID").value = invitationCode;
    if (invitationCode !== null) {
        document.getElementById("InviteID").readOnly = true;
        document.getElementById("InviteID").parentElement.classList.add("mdui-textfield-not-empty");
    }
}





