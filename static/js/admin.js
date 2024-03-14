"use strict";
isLogin();
// 限制输入框只能最大输入0-100
$(".limit-100").on("input", function() {
    if (this.value > 100) {
        this.value = 100;
    }
    if (this.value < 0) {
        this.value = 0;
    }
});

SwitchItemInit();
// 所有用户筛选
$("#UserSelect").on("change", function () {
    $("#LevelSelect").next().find(".mdui-select-selected").text("所有权限");
    if($(this).val() === "所有用户"){
        UserManager();
    }else if($(this).val() === "正常用户"){
        UserManager(function(data){
            data = data.filter((v) => {
                return !v.is_freeze;
            });
            for (let user of data) {
                let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                UserManagerData.append(tr);
                tr.append(`<td>${user.user_name}</td>`);
                tr.append(`<td>${user.user_id}</td>`);
                tr.append(`<td>${user.balance}</td>`);
                tr.append(`<td>${user.coin_balance}</td>`);
                tr.append(`<td>${user.quota}</td>`);
                tr.append(`<td>${user.pri_label}</td>`);
                tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
            }
        });
    }else if($(this).val() === "冻结用户"){
        UserManager(function(data){
            data = data.filter((v) => {
                return v.is_freeze;
            });
            for (let user of data) {
                let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                UserManagerData.append(tr);
                tr.append(`<td>${user.user_name}</td>`);
                tr.append(`<td>${user.user_id}</td>`);
                tr.append(`<td>${user.balance}</td>`);
                tr.append(`<td>${user.coin_balance}</td>`);
                tr.append(`<td>${user.quota}</td>`);
                tr.append(`<td>${user.pri_label}</td>`);
                tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
            }
        });
    }
});

// 所有用户筛选
$("#LevelSelect").on("change", function () {
    $("#UserSelect").next().find(".mdui-select-selected").text("所有用户");
    if($(this).val() === "所有权限"){
        UserManager();
    }else if($(this).val() === "临时用户"){
        UserManager(function(data){
            data = data.filter((v) => {
                return v.pri_label === 1;
            });
            for (let user of data) {
                let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                UserManagerData.append(tr);
                tr.append(`<td>${user.user_name}</td>`);
                tr.append(`<td>${user.user_id}</td>`);
                tr.append(`<td>${user.balance}</td>`);
                tr.append(`<td>${user.coin_balance}</td>`);
                tr.append(`<td>${user.quota}</td>`);
                tr.append(`<td>${user.pri_label}</td>`);
                tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
            }
        });
    }else if($(this).val() === "正式用户"){
        UserManager(function(data){
            data = data.filter((v) => {
                return v.pri_label === 2;
            });
            for (let user of data) {
                let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                UserManagerData.append(tr);
                tr.append(`<td>${user.user_name}</td>`);
                tr.append(`<td>${user.user_id}</td>`);
                tr.append(`<td>${user.balance}</td>`);
                tr.append(`<td>${user.coin_balance}</td>`);
                tr.append(`<td>${user.quota}</td>`);
                tr.append(`<td>${user.pri_label}</td>`);
                tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
            }
        });
    }else if($(this).val() === "初级担保人"){
        UserManager(function(data){
            data = data.filter((v) => {
                return v.pri_label === 3;
            });
            for (let user of data) {
                let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                UserManagerData.append(tr);
                tr.append(`<td>${user.user_name}</td>`);
                tr.append(`<td>${user.user_id}</td>`);
                tr.append(`<td>${user.balance}</td>`);
                tr.append(`<td>${user.coin_balance}</td>`);
                tr.append(`<td>${user.quota}</td>`);
                tr.append(`<td>${user.pri_label}</td>`);
                tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
            }
        });
    }else if($(this).val() === "高级担保人"){
        UserManager(function(data){
            data = data.filter((v) => {
                return v.pri_label === 4;
            });
            for (let user of data) {
                let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                UserManagerData.append(tr);
                tr.append(`<td>${user.user_name}</td>`);
                tr.append(`<td>${user.user_id}</td>`);
                tr.append(`<td>${user.balance}</td>`);
                tr.append(`<td>${user.coin_balance}</td>`);
                tr.append(`<td>${user.quota}</td>`);
                tr.append(`<td>${user.pri_label}</td>`);
                tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
            }
        });
    }else if($(this).val() === "管理员"){
        UserManager(function(data){
            data = data.filter((v) => {
                return v.pri_label === 5;
            });
            for (let user of data) {
                let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                UserManagerData.append(tr);
                tr.append(`<td>${user.user_name}</td>`);
                tr.append(`<td>${user.user_id}</td>`);
                tr.append(`<td>${user.balance}</td>`);
                tr.append(`<td>${user.coin_balance}</td>`);
                tr.append(`<td>${user.quota}</td>`);
                tr.append(`<td>${user.pri_label}</td>`);
                tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
            }
        });
    }
});
// 用户管理数据
let UserManagerData = $("#UserManagerData");
// 订单管理数据
let OrderManagerData = $("#OrderManagerData");

let deleteBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-red" onclick="deleteEvent(this,${id})">删除</button>`;
}
// 订单筛选
$("#PriceSort").on("change", function () {
    $("#UserStatus").next().find(".mdui-select-selected").text("所有状态");
    if($(this).val() === "默认排序"){
        OrderManager();
    }else if($(this).val() === "按总价升序"){
        OrderManager(function(data){
            data = data.sort((a,b) => {
                return a.total_price-b.total_price;
            });
            for (let order of data) {
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                OrderManagerData.append(tr);
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                tr.append(`<td>${order.order_id}</td>`);
                tr.append(`<td>${order.purchasers_id}</td>`);
                tr.append(`<td>${order.seller_id}</td>`);
                tr.append(`<td>${(order.surety_id !== 0)?order.surety_id:""}</td>`);
                switch (order.order_status) {
                    case 0:
                        tr.append(`<td>已上架</td>`);
                        break;
                    case 1:
                        tr.append(`<td>等待担保人中</td>`);
                        break;
                    case 2:
                        tr.append(`<td>已完成</td>`);
                        break;
                    case 3:
                        tr.append(`<td>无担保人发单</td>`);
                        break;
                    case 4:
                        tr.append(`<td>已冻结</td>`);
                        break;
                    case 5:
                        tr.append(`<td>交易中</td>`);
                        break;
                }
                // 冻结原因
                tr.append(`<td>${order.freeze_reasons}</td>`);
                tr.append(`<td>${order.count}</td>`);
                tr.append(`<td>${order.prices}</td>`);
                tr.append(`<td>${order.total_price}</td>`);
                tr.append(`
                <td><div class="mdui-btn-group">${(order.order_status === 4) ? orderUnfrozenBtn(order.order_id) : orderFrozenBtn(order.order_id)}
                ${orderDeleteBtn(order.order_id)}</div></td>
            `);
            }
        });
    }else if($(this).val() === "按总价降序"){
        OrderManager(function(data){
            data = data.sort((a,b) => {
                return b.total_price-a.total_price;
            });
            for (let order of data) {
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                OrderManagerData.append(tr);
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                tr.append(`<td>${order.order_id}</td>`);
                tr.append(`<td>${order.purchasers_id}</td>`);
                tr.append(`<td>${order.seller_id}</td>`);
                tr.append(`<td>${(order.surety_id !== 0)?order.surety_id:""}</td>`);
                switch (order.order_status) {
                    case 0:
                        tr.append(`<td>已上架</td>`);
                        break;
                    case 1:
                        tr.append(`<td>等待担保人中</td>`);
                        break;
                    case 2:
                        tr.append(`<td>已完成</td>`);
                        break;
                    case 3:
                        tr.append(`<td>无担保人发单</td>`);
                        break;
                    case 4:
                        tr.append(`<td>已冻结</td>`);
                        break;
                    case 5:
                        tr.append(`<td>交易中</td>`);
                        break;
                }
                // 冻结原因
                tr.append(`<td>${order.freeze_reasons}</td>`);
                tr.append(`<td>${order.count}</td>`);
                tr.append(`<td>${order.prices}</td>`);
                tr.append(`<td>${order.total_price}</td>`);
                tr.append(`
                <td><div class="mdui-btn-group">${(order.order_status === 4) ? orderUnfrozenBtn(order.order_id) : orderFrozenBtn(order.order_id)}
                ${orderDeleteBtn(order.order_id)}</div></td>
            `);
            }
        });
    }
});
// 订单筛选
$("#UserStatus").on("change", function () {
    $("#PriceSort").next().find(".mdui-select-selected").text("默认排序");
    if($(this).val() === "所有状态"){
        OrderManager();
    }else if($(this).val() === "正常订单"){
        OrderManager(function(data){
            data = data.filter((v) => {
                return v.order_status !== 4;
            });
            for (let order of data) {
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                OrderManagerData.append(tr);
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                tr.append(`<td>${order.order_id}</td>`);
                tr.append(`<td>${order.purchasers_id}</td>`);
                tr.append(`<td>${order.seller_id}</td>`);
                tr.append(`<td>${(order.surety_id !== 0)?order.surety_id:""}</td>`);
                switch (order.order_status) {
                    case 0:
                        tr.append(`<td>已上架</td>`);
                        break;
                    case 1:
                        tr.append(`<td>等待担保人中</td>`);
                        break;
                    case 2:
                        tr.append(`<td>已完成</td>`);
                        break;
                    case 3:
                        tr.append(`<td>无担保人发单</td>`);
                        break;
                    case 4:
                        tr.append(`<td>已冻结</td>`);
                        break;
                    case 5:
                        tr.append(`<td>交易中</td>`);
                        break;
                }
                // 冻结原因
                tr.append(`<td>${order.freeze_reasons}</td>`);
                tr.append(`<td>${order.count}</td>`);
                tr.append(`<td>${order.prices}</td>`);
                tr.append(`<td>${order.total_price}</td>`);
                tr.append(`
                <td><div class="mdui-btn-group">${(order.order_status === 4) ? orderUnfrozenBtn(order.order_id) : orderFrozenBtn(order.order_id)}
                ${orderDeleteBtn(order.order_id)}</div></td>
            `);
            }
        });
    }else if($(this).val() === "冻结订单"){
        OrderManager(function(data){
            data = data.filter((v) => {
                return v.order_status === 4;
            });
            for (let order of data) {
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                OrderManagerData.append(tr);
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                tr.append(`<td>${order.order_id}</td>`);
                tr.append(`<td>${order.purchasers_id}</td>`);
                tr.append(`<td>${order.seller_id}</td>`);
                tr.append(`<td>${(order.surety_id !== 0)?order.surety_id:""}</td>`);
                switch (order.order_status) {
                    case 0:
                        tr.append(`<td>已上架</td>`);
                        break;
                    case 1:
                        tr.append(`<td>等待担保人中</td>`);
                        break;
                    case 2:
                        tr.append(`<td>已完成</td>`);
                        break;
                    case 3:
                        tr.append(`<td>无担保人发单</td>`);
                        break;
                    case 4:
                        tr.append(`<td>已冻结</td>`);
                        break;
                    case 5:
                        tr.append(`<td>交易中</td>`);
                        break;
                }
                // 冻结原因
                tr.append(`<td>${order.freeze_reasons}</td>`);
                tr.append(`<td>${order.count}</td>`);
                tr.append(`<td>${order.prices}</td>`);
                tr.append(`<td>${order.total_price}</td>`);
                tr.append(`
                <td><div class="mdui-btn-group">${(order.order_status === 4) ? orderUnfrozenBtn(order.order_id) : orderFrozenBtn(order.order_id)}
                ${orderDeleteBtn(order.order_id)}</div></td>
            `);
            }
        });
    }
});
function deleteEvent(self,id){
    mdui.dialog({
        title: '提示',
        content: '此操作将删除此用户且无法撤销',
        buttons: [
            {
                text: '取消',
            },
            {
                text: '确认',
                onClick: function () {
                    ajax({
                        method: 'POST',
                        url: `${ServerAddr}/admin/deleteUser`,
                        data: {
                            user_id: id,
                        },
                        handle: function (data) {
                            Toast("提示", data.msg);
                            if (data.code !== 200) {
                                return;
                            }
                            $(self).parent().parent().remove();

                        }
                    });
                }
            }
        ]
    });
}

let frozenBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-orange" onclick="frozenEvent(this,${id})">冻结</button>`
};

function frozenEvent(self,id){
    ajax({
        method: 'POST',
        url: `${ServerAddr}/admin/freezeUser`,
        data:{
            user_id:id,
        },
        handle: function (data) {
            Toast("提示",data.msg);
            if(data.code !== 200){
                return;
            }
            $(self).removeClass("mdui-color-orange").addClass("mdui-color-green").attr("onclick",`unfrozenEvent(this,${id})`).
            text("解冻").parent().parent().removeClass("user-normal").addClass("user-frozen");
            
        }
    });
}
let unfrozenBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-green" onclick="unfrozenEvent(this,${id})">解冻</button>`;
}
function unfrozenEvent(self,id){
    ajax({
        method: 'POST',
        url: `${ServerAddr}/admin/unFreezeUser`,
        data:{
            user_id:id,
        },
        handle: function (data) {
            Toast("提示",data.msg);
            if(data.code !== 200){
                return;
            }
            $(self).removeClass("mdui-color-green").addClass("mdui-color-orange").attr("onclick",`frozenEvent(this,${id})`).
            text("冻结").parent().parent().removeClass("user-frozen").addClass("user-normal");
            
        }
    });
}
// 用户管理
function UserManager(callback=undefined){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/admin/getAllUser`,
        success: function (data) {
            data = JSON.parse(data);
            UserManagerData.html("");
            if (callback === undefined) {
                for (let user of data) {
                    let tr = $(`<tr class="${(user.is_freeze) ? "user-frozen" : "user-normal"}"></tr>`);
                    UserManagerData.append(tr);
                    tr.append(`<td>${user.user_name}</td>`);
                    tr.append(`<td>${user.user_id}</td>`);
                    tr.append(`<td>${user.balance}</td>`);
                    tr.append(`<td>${user.coin_balance}</td>`);
                    tr.append(`<td>${user.quota}</td>`);
                    tr.append(`<td>${user.pri_label}</td>`);
                    tr.append(`
                    <td><div class="mdui-btn-group">${(user.is_freeze) ? unfrozenBtn(user.user_id) : frozenBtn(user.user_id)}
                    ${deleteBtn(user.user_id)}</div></td>
                `);
                }
            }else{
                callback(data);
            }
        },
        error: function (xhr) {
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : {code:500,msg:'服务器异常'};
            Toast("提示", response.msg);
        }
    });
}
let orderDeleteBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-red" onclick="orderDeleteEvent(this,${id})">删除</button>`;
}
function orderDeleteEvent(self,id){
    mdui.dialog({
        title: '提示',
        content: '此操作将删除此订单且无法撤销',
        buttons: [
            {
                text: '取消',
            },
            {
                text: '确认',
                onClick: function () {
                    ajax({
                        method: 'POST',
                        url: `${ServerAddr}/admin/deleteOrder`,
                        data: {
                            order_id: id,
                        },
                        handle: function (data) {
                            Toast("提示", data.msg);
                            if (data.code !== 200) {
                                return;
                            }
                            $(self).parent().parent().remove();

                        }
                    });
                }
            }
        ]
    });
}

let orderFrozenBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-orange" onclick="orderFrozenEvent(this,${id})">冻结</button>`
};

function orderFrozenEvent(self,id){
    ajax({
        method: 'POST',
        url: `${ServerAddr}/admin/freezeOrder`,
        data:{
            order_id:id,
        },
        handle: function (data) {
            Toast("提示",data.msg);
            if(data.code !== 200){
                return;
            }
            $(self).removeClass("mdui-color-orange").addClass("mdui-color-green").attr("onclick",`orderUnfrozenEvent(this,${id})`).
            text("解冻").parent().parent().removeClass("user-normal").addClass("user-frozen");
            
        }
    });
}
let orderUnfrozenBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-green" onclick="orderUnfrozenEvent(this,${id})">解冻</button>`;
}
function orderUnfrozenEvent(self,id){
    ajax({
        method: 'POST',
        url: `${ServerAddr}/admin/unFreezeOrder`,
        data:{
            order_id:id,
        },
        handle: function (data) {
            Toast("提示",data.msg);
            if(data.code !== 200){
                return;
            }
            $(self).removeClass("mdui-color-green").addClass("mdui-color-orange").attr("onclick",`orderFrozenEvent(this,${id})`).
            text("冻结").parent().parent().removeClass("user-frozen").addClass("user-normal");
            
        }
    });
}
// 订单管理
function OrderManager(callback=undefined){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/admin/getAllOrder`,
        success: function (data) {
            data = JSON.parse(data);
            OrderManagerData.html("");
            if (callback === undefined) {
                for (let order of data) {
                    let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                    OrderManagerData.append(tr);
                    tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                    tr.append(`<td>${order.order_id}</td>`);
                    tr.append(`<td>${order.purchasers_id}</td>`);
                    tr.append(`<td>${order.seller_id}</td>`);
                    tr.append(`<td>${(order.surety_id !== 0)?order.surety_id:""}</td>`);
                    switch (order.order_status) {
                        case 0:
                            tr.append(`<td>已上架</td>`);
                            break;
                        case 1:
                            tr.append(`<td>等待担保人中</td>`);
                            break;
                        case 2:
                            tr.append(`<td>已完成</td>`);
                            break;
                        case 3:
                            tr.append(`<td>无担保人发单</td>`);
                            break;
                        case 4:
                            tr.append(`<td>已冻结</td>`);
                            break;
                        case 5:
                            tr.append(`<td>交易中</td>`);
                            break;
                    }
                    // 冻结原因
                    tr.append(`<td>${order.freeze_reasons}</td>`);
                    tr.append(`<td>${order.count}</td>`);
                    tr.append(`<td>${order.prices}</td>`);
                    tr.append(`<td>${order.total_price}</td>`);
                    tr.append(`
                    <td><div class="mdui-btn-group">${(order.order_status === 4) ? orderUnfrozenBtn(order.order_id) : orderFrozenBtn(order.order_id)}
                    ${orderDeleteBtn(order.order_id)}</div></td>
                `);
                }
            }else{
                callback(data);
            }
        },
        error: function (xhr) {
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : {code:500,msg:'服务器异常'};
            Toast("提示", response.msg);
        }
    });
}
// 数据统计
function DataStatistics(){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/admin/getStatistic`,
        success: function (data) {
            data = JSON.parse(data);
            console.log(data);
            $("#TotalDividend").text(data.total_dividend+"元");
            $("#CoinDividend").text(data.coin_dividend+"元");
            $("#CoinNumberOfTransactions").text(data.coin_number_of_transactions+"个");
            $("#GemNumberOfTransactions").text(data.gem_number_of_transactions+"个");
            $("#PlatformCumulativeProfit").text(data.platform_cumulative_profit+"元");
        },
        error: function (xhr) {
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : {code:500,msg:'服务器异常'};
            Toast("提示", response.msg);
        }
    });
}
// 取消
let refuseBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-red" onclick="refuseClick(this,${id})">拒绝</button>`;
}
function refuseClick(self,id){
    mdui.dialog({
        title: '提示',
        content: '是否拒绝提现',
        buttons: [
            {
                text: '取消',
            },
            {
                text: '确认',
                onClick: function () {
                    ajax({
                        method: 'POST',
                        url: `${ServerAddr}/admin/refusalWithDraw`,
                        data:{
                            withdraw_id:id,
                        },
                        handle:function(data){
                            if(data.code === 200){
                                $(self).parent().parent().remove();
                            }
                            Toast("提示",data.msg);
                        }
                    });
                }
            }
        ]
    });
    
}
// 完成
let completeBtn = function(id){
    return `<button class="mdui-btn mdui-ripple mdui-color-green-600" onclick="completeClick(this,${id})">完成</button>`;
}
let userInfoBtn = function(id){
    return `<div class="mdui-btn mdui-ripple mdui-color-light-green mdui-text-color-white" onclick="userInfoClick(${id})">查看</div>`;
}

function userInfoClick(id) {
    let UserInfoDialog = $("#UserInfoToast").html();
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getUserInfo`,
        data: {
            user_id: id,
        },
        success: function (data) {
            data = JSON.parse(data);
            let dialog = Toast("", UserInfoDialog);
            let inst = $(dialog.$element[0]);
            // 用户名
            inst.find("#TUserName").text(data.user_name);
            // 游戏ID
            inst.find("#TGameID").text(data.game_id);
            // 电话
            inst.find("#TPhone").text(data.phone);
            // 姓名
            inst.find("#TName").text(data.user_lastname);
            // 支付宝qr码
            inst.find("#TAlipay").on("click", function () {
                dialog.close();
                Toast("微信", Text2QRCodeTag(data.wxpay_qrcode).html());
            });
            // 微信qr码
            inst.find("#TWxpay").on("click", function () {
                dialog.close();
                Toast("支付宝", Text2QRCodeTag(data.alipay_qrcode).html());

            });
        },
        error: function (xhr) {
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
            Toast("提示", response.msg);
        }
    });
}
function completeClick(self,id){
    mdui.dialog({
        title: '提示',
        content: '是否完成提现',
        buttons: [
            {
                text: '取消',
            },
            {
                text: '确认',
                onClick: function () {
                    ajax({
                        method: 'POST',
                        url: `${ServerAddr}/admin/deleteWithDraw`,
                        data:{
                            withdraw_id:id,
                        },
                        handle:function(data){
                            if(data.code === 200){
                                $(self).parent().parent().remove();
                            }
                            Toast("提示",data.msg);
                        }
                    });
                }
            }
        ]
    });
   
}
let WithdrawalManagerData = $("#WithdrawalManagerData");
// 提现管理
function WithdrawalManager(){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/admin/getAllWithDraw`,
        success: function (data) {
            data = JSON.parse(data);
            WithdrawalManagerData.html("");
            for (let withdrawal of data) {
                let tr = $(`<tr></tr>`);
                WithdrawalManagerData.append(tr);
                tr.append(`<td>${convertDateFormat(withdrawal.time)}</td>`);
                tr.append(`<td>${withdrawal.withdraw_id}</td>`);
                tr.append(`<td>${withdrawal.user_id}</td>`);
                tr.append(`<td>${withdrawal.balance}</td>`);
                tr.append(`<td>${withdrawal.cash_withdrawal}</td>`);
                tr.append(`<td>${userInfoBtn(withdrawal.user_id)}</td>`);
                tr.append(`
                <td><div class="mdui-btn-group">${completeBtn(withdrawal.withdraw_id)}${refuseBtn(withdrawal.withdraw_id)}</div></td>
            `);
            }
        },
        error: function (xhr) {
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
            Toast("提示", response.msg);
        }
    });
}
let bounty_percentage = $("#bounty_percentage");
let platform_commission = $("#platform_commission");
let min_coin_transaction = $("#min_coin_transaction");
let min_gem_transaction = $("#min_gem_transaction");
let dividend_percentage_I = $("#dividend_percentage_I");
let dividend_percentage_II = $("#dividend_percentage_II");
let coin_dividend_percentage_I = $("#coin_dividend_percentage_I");
let coin_dividend_percentage_II = $("#coin_dividend_percentage_II");
let bulletin = $("#bulletin");
let save = $("#Save");
save.on("click",function(){
    ajax({
        method: 'POST',
        url: `${ServerAddr}/admin/saveSettings`,
        data:{
            bounty_percentage: parseFloat(bounty_percentage.val()),
            platform_commission: parseFloat(platform_commission.val()),
            min_coin_transaction: parseInt(min_coin_transaction.val()),
            min_gem_transaction: parseInt(min_gem_transaction.val()),
            dividend_percentage_I: parseFloat(dividend_percentage_I.val()),
            dividend_percentage_II: parseFloat(dividend_percentage_II.val()),
            coin_dividend_percentage_I: parseFloat(coin_dividend_percentage_I.val()),
            coin_dividend_percentage_II: parseFloat(coin_dividend_percentage_II.val()),
            bulletin: bulletin.val(),
        },
        handle:function(data){
            Toast("提示",data.msg);
        }
    });
});
// 综合设置
function Settings(){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/admin/getSettings`,
        success: function (data) {
            data = JSON.parse(data);
            bounty_percentage.val(data.bounty_percentage).trigger("input");
            platform_commission.val(data.platform_commission).trigger("input");
            min_coin_transaction.val(data.min_coin_transaction).trigger("input");
            min_gem_transaction.val(data.min_gem_transaction).trigger("input");
            dividend_percentage_I.val(data.dividend_percentage_I).trigger("input");
            dividend_percentage_II.val(data.dividend_percentage_II).trigger("input");
            coin_dividend_percentage_I.val(data.coin_dividend_percentage_I).trigger("input");
            coin_dividend_percentage_II.val(data.coin_dividend_percentage_II).trigger("input");
            bulletin.val(data.bulletin).trigger("input");
        },
        error: function (xhr) {
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
            Toast("提示", response.msg);
        }
    });
}
function Logout(){
    mdui.dialog({
        title: '提示',
        content: '是否退出登录',
        buttons: [
            {
                text: '取消',
            },
            {
                text: '确认',
                onClick: function () {
                    ajax({
                        method: 'GET',
                        url: `${ServerAddr}/user/logout`,
                        handle: function (data) {
                            if (data.code !== 200) {
                                Toast("提示", "服务器异常");
                                return;
                            }
                            Toast("提示", data.msg, '确认', function (inst) {
                                inst.close();
                            });
                            setTimeout(function () {
                                window.location.href = '/static/login.html';
                            });
                           
                        }
                    });
                }
            }
        ]
    });
}
UserManager();