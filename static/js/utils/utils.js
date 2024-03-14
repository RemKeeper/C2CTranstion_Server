"use strict";
var $ = mdui.$;
// 弹出对话框的函数
function Toast(title, content, btnName = '确认', fn = undefined) {
    let isClose = true;
    if (fn !== undefined) {
        isClose = false;
    }
    return mdui.dialog({
        title: title,
        content: content,
        buttons: [
            {
                close: isClose,
                text: btnName,
                onClick: fn,

            }
        ]
    });
}



// 抽屉栏切换
function SwitchItem(index) {
    let itemPage = $(".item-page");
    $.each(itemPage, function (_, item) {
        item.style.display = "none";
    })
    itemPage.get(index).style.display = "block";
    drawer.toggle();
}

// 抽屉栏初始化
function SwitchItemInit(initIndex = 0) {
    // 抽屉栏初始化
    window.drawer = new mdui.Drawer('#drawer', { "swipe": true });
    // 抽屉栏切换
    $('#toggle').on('click', function () {
        drawer.toggle();
    });
    // 绑定抽屉栏点击事件
    $(".mdui-drawer > .mdui-list > li").not(".mdui-subheader,.toast").each(function (index, item) {
        item.onclick = function () {
            SwitchItem(index);
        }
    })
    SwitchItem(initIndex);
    drawer.toggle();
}

function ajax(config) {
    $.ajax({
        method: config.method || 'GET',
        url: config.url,
        contentType: config.contentType || 'application/json',
        data: (config.data ? ((config.method !== 'POST')?config.data:JSON.stringify(config.data)): null),
        success: function(response) {
            config.handle&&config.handle(JSON.parse(response));
        },
        error: function(xhr) {
            // 如果服务器响应了，尝试解析响应，否则使用默认错误信息
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : {code:500,msg:'服务器异常'};
            config.handle&&config.handle(response, config);
        }
    });
}

// 把时间转换为日期字符串
function convertDateFormat(dateString) {
    // 解析日期字符串
    const date = new Date(dateString);

    // 提取年、月、日
    const year = date.getFullYear();
    // getMonth() 返回的月份是从0开始的，因此需要+1
    const month = date.getMonth() + 1;
    const day = date.getDate();

    // 格式化日期
    return `${year}-${month}-${day}`;
}

// 获取交易类型名字（购买，售卖）
function getTransactionTypeName(order/*订单*/,userInfo/*用户信息 */){
    // 用户ID是否不等于卖家ID
    if(order.seller_id !== userInfo.user_id){
        return "购买";
    }
    return "售卖";
}

// 是否登录
function isLogin(){
    if(document.cookie === ''){
        setTimeout(function () {
            window.location.href = '/static/login.html';
        });
    }
}
// 获取权限
// 类型	临时用户	正式用户	初级担保人	高级担保人	管理员
// 权限码	1	      2	          3	          4	       5

function getPermission(callback){
    ajax({
        method: 'GET',
        url: `${ServerAddr}/GetPermission`,
        handle: callback 
    });
}

// 全局弹窗
function globalToast() {
    setInterval(function () {
        ajax({
            method: 'GET',
            url: `${ServerAddr}/getToast`,
            handle: function (data) {
                Toast("通知", "")
            }
        })
    }, 15000);
}