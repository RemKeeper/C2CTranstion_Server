"use strict";
let qcode = new QRCode($("#WeChatServer").get(0),"https://u.wechat.com/MChgWqbIU0r7sw2yFave8_Q")
let qcode2 = new QRCode($("#QQServer").get(0),"https://qm.qq.com/q/L4dTvCstc6")


// 测试用例
let images = [
    qcode._oDrawing._elCanvas.toDataURL("image/png"),
    qcode2 ._oDrawing._elCanvas.toDataURL("image/png"),
];

// 文字转QR码
function Text2QRCodeTag(text){
    let qcode = new QRCode($(`<div/>`).get(0),text)._oDrawing._elCanvas.toDataURL("image/png");
    return $(`<div class="mdui-col-xs-3"><img src="${qcode}" class="mdui-img-fluid"></div>`)
}

function showContactDialog() {
    // images 是一个包含两个图片地址的数组
    // 检查参数是否合法
    if (!Array.isArray(images) || images.length !== 2) {
        console.error("Invalid images array");
        return;
    }

    // 创建对话框的内容
    var content = '<div class="mdui-row">';
    let text = ["微信客服","QQ客服"]
    for (var i = 0; i < 2; i++) {
        content += '<div class="mdui-col-xs-6">';
        content += '<img src="' + images[i] + '" class="mdui-img-fluid">';
        content+= '<div class="mdui-text-center mdui-m-t-1">'+text[i]+'</div>'
        content += '</div>';
    }
    content += '</div>';

    // 调用 mdui.dialog 函数，打开对话框
    mdui.dialog({
        content: content,
        buttons: [
            {
                text: "关闭",
                onClick: function (inst) {
                    inst.close();
                },
            },
        ],
    });
}
