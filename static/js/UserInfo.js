"use strict";
isLogin();
const params = new URLSearchParams(window.location.search);
const orderID = params.get('order_id');
let TransactionInvitation = $("#TransactionInvitationTemplate");
let TransactionInvitationData = TransactionInvitation.find("#Data");
$.ajax({
    method: 'GET',
    url: `${ServerAddr}/getToast`,
    data: {
        order_id: orderID,
    },
    success: function (data) {
        data = JSON.parse(data);
        $.each(data,function(uuid,v){
            Toast('通知', v.Msg,"我知道了",function(inst){
                inst.close();
                ajax({
                    method: 'POST',
                    url: `${ServerAddr}/confirmedToast`,
                    data: {
                        uuid:uuid,
                    },
                });
            });
        });
    },
    error:function(xhr){
        var data = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
        Toast('提示', data.msg);
    }
});
setInterval(function(){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/getToast`,
        data: {
            order_id: orderID,
        },
        success: function (data) {
            data = JSON.parse(data);
            $.each(data,function(uuid,v){
                Toast('通知', v.Msg,"我知道了",function(inst){
                    inst.close();
                    ajax({
                        method: 'POST',
                        url: `${ServerAddr}/confirmedToast`,
                        data: {
                            uuid:uuid,
                        },
                    });
                });
            });
        },
        error:function(xhr){
            var data = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
            Toast('提示', data.msg);
        }
    });
},15000)
getPermission(function(data){
    if(data.code === 200){
        if(data.level <= 2){
            $("#GuaranteeOrderTab").hide();
        }
    }else{
        Toast("提示","服务器异常");
    }
});


if(orderID !== null){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getOrderById`,
        data: {
            order_id: orderID,
        },
        success: function (data) {
            data = JSON.parse(data);
            TransactionInvitationData.html("");
            let tr = $(`<tr class="${(data.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
            // 时间
            tr.append(`<td>${convertDateFormat(data.time)}</td>`);
            // 订单ID
            tr.append(`<td>${data.order_id}</td>`);
            // 单价
            tr.append(`<td>${data.prices}</td>`);
            // 数量
            tr.append(`<td>${data.count}</td>`);
            // 总价
            tr.append(`<td>${data.total_price}</td>`);
            TransactionInvitationData.append(tr);
            console.log(TransactionInvitationData.html());
            Toast(`${(data.product_type === 0) ? "金币购买邀请" : "宝石出售邀请"}`, TransactionInvitation.html(), "交易", function (inst) {
                ajax({
                    method: 'POST',
                    url: `${ServerAddr}/transaction/carryTransaction`,
                    data: {
                        order_id: parseInt(data.order_id),
                    },
                    handle: function (data) {
                        inst.close();
                        Toast("提示", data.msg);

                    }
                });
            }, "取消");
        },
        error:function(xhr){
            var data = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
            Toast('提示', data.msg);
        }
    });
}

let InvitationPromotionalCopywriting ="这是一段邀请推广文案，仅作为测试"


// 数据
let UserInfoData;
let AllOrderData;
// 用户信息DOM
let UserName = $("#UserName");
let UserID = $("#UserID");
let CashBalance = $("#CashBalance");
let GoldCoinBalance = $("#GoldCoinBalance");
let GameID = $("#GameID");
let Name = $("#Name");
let BankCardNumber = $("#BankCardNumber");
let WeChat = $("#WeChat");
let Alipay = $("#Alipay");
let PhoneNumber = $("#PhoneNumber");

// 全部订单DOM
let AllOrderList = $("#AllOrder #Data");

// 等候订单DOM
let WaitingOrderList = $("#WaitingOrder #Data");

// 担保订单DOM
let GuaranteeOrderList = $("#GuaranteeOrder #Data");

// 冻结订单DOM
let FrozenOrderList = $("#FrozenOrder #Data");

// 操作按钮
let transactionDisputes = `<button class="mdui-btn mdui-ripple mdui-color-red-500" onclick="transactionDisputesClick(this)">交易纠纷</button>`
function transactionDisputesClick(self) {
    mdui.dialog({
        title: '提示',
        content: '此操作将冻结订单，冻结后请联系客服进行纠纷解冻',
        buttons: [
            {
                text: '取消',
            },
            {
                text: '确认',
                onClick: function () {
                    let orderID = parseInt($(self).parent().parent().parent().find("td").eq(1).text());
                    ajax({
                        method: 'POST',
                        url: `${ServerAddr}/user/createDispute`,
                        data: {
                            order_id: orderID,
                        },
                        handle: function (data) {
                            if (data.code !== 200) {
                                Toast("提示", data.msg);
                            } else {
                                location.reload(true);
                            }

                        }
                    });
                }
            }
        ]
    });

}
let orderShipment = function (order_id) {
    return `<div class="mdui-btn mdui-ripple mdui-color-green-600 mdui-shadow-1"
style="position: relative; overflow: hidden;">
订单发货
<input class="mdui-textfield-input file-input" onChange="orderShipmentChange(this,${order_id})"
    style="opacity: 0; position: absolute; top: 0; left: 0; width: 100%; height: 100%; "
    type="file" name="" value="" required />
</div>`;
}
// 完成订单发货后上传的文件
function orderShipmentChange(self, order_id) {
    self = $(self);
    // 获取文件输入框中的文件对象
    var file = self[0].files[0];

    // 创建一个 FormData 对象
    var formData = new FormData();
    formData.append('file', file); // 添加文件
    formData.append('order_id', order_id); // 添加其他需要的数据

    // 发送 $.ajax 请求
    $.ajax({
        url: `${ServerAddr}/upload`, // 这里替换为你的上传接口的URL
        method: 'POST',
        data: formData,
        processData: false,  // 告诉jQuery不要去处理发送的数据
        contentType: false,  // 告诉jQuery不要去设置Content-Type请求头
        success: function (data) {
            data = JSON.parse(data);
            if (data.code === 200) {
                let parent = self.parent().parent().parent();
                self.parent().remove();
                parent.prev().text("是");
            }
            Toast('提示', data.msg);
        },
        error: function (xhr, status, error) {
            var data = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
            Toast('提示', data.msg);
        }
    });

}
let generateLink = function(id){
    return `<div class="mdui-btn mdui-ripple mdui-color-indigo-accent" onclick="generateLinkClick(${id})">生成链接</div>`;
}
function generateLinkClick(id){
    Toast('订单邀请链接（点击复制）',window.location.origin+"/static/UserInfo.html?order_id="+id,"复制", function(inst){
        inst.close();
        copyTextToClipboard(window.location.origin+"/static/UserInfo.html?order_id="+id)
        Toast('提示','复制成功');
    });
}
let viewInfo = function(id){
    return `<div class="mdui-btn mdui-ripple mdui-color-light-green mdui-text-color-white" onclick="viewInfoClick(${id})">查看</div>`;
}


function viewInfoClick(id){
    let UserInfoDialog =  $("#UserInfoToast").html();
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getUserInfo`,
        data: {
            user_id: id,
        },
        success: function (data) {
            data = JSON.parse(data);
            let dialog = Toast("",UserInfoDialog);
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
            inst.find("#TAlipay").on("click",function(){
                dialog.close();
                Toast("微信",Text2QRCodeTag(data.wxpay_qrcode).html());
            });
            // 微信qr码
            inst.find("#TWxpay").on("click",function(){
                dialog.close();
                Toast("支付宝",Text2QRCodeTag(data.alipay_qrcode).html());
                
            });
        },
        error: function (xhr) {
            var response = xhr.responseText ? JSON.parse(xhr.responseText) : {code:500,msg:'服务器异常'};
            Toast("提示", response.msg);
        }
    })
}
let completeOrder = `<div class="mdui-btn mdui-ripple mdui-color-red-500">完成交易</div>`;


// 用户信息
function UserInfo() {
    
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getSelfInfo`,
        success: function (data) {
            data = JSON.parse(data);
            UserInfoData = data;
            UserName.text(data.user_name);
            UserID.text(`ID:${data.user_id}`);
            CashBalance.text(data.balance);
            GoldCoinBalance.text(data.coin_balance);
            GameID.text(data.game_id);
            Name.text(data.user_lastname);
            BankCardNumber.text(data.bank_card);
            WeChat.text(data.wxpay_qrcode);
            Alipay.text(data.alipay_qrcode);
            PhoneNumber.text(data.phone);
        },
        error: function (xhr) {
            let rsp= JSON.parse(xhr.response)
            Toast("错误", rsp.msg);
        }
    });
}

// 全部订单
function AllOrder() {
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getAllOrder`,
        success: function (data) {
            data = JSON.parse(data);
            AllOrderData = data;
            AllOrderList.html("");
            console.log(data);
            for (let order of data) {
                let TransactionTypeName = getTransactionTypeName(order, UserInfoData);
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                // 时间
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                // 订单ID
                tr.append(`<td>${order.order_id}</td>`);
                // 订单类型
                tr.append(`<td>${TransactionTypeName}</td>`);
                // 单价
                tr.append(`<td>${order.prices}</td>`);
                // 数量
                tr.append(`<td>${order.count}</td>`);
                // 总价
                tr.append(`<td>${order.total_price}</td>`);
                // 游戏ID
                tr.append(`<td>${order.order_id}</td>`);
                // 担保人ID
                tr.append(`<td>${(order.surety_id===0)?"":order.surety_id}</td>`);
                // 冻结原因
                if (order.order_status !== 4) {
                    tr.append(`<td></td>`);
                } else {
                    tr.append(`<td>${order.freeze_reasons}</td>`);
                }
                // 是否发货
                tr.append(`<td>${(order.is_dispatched)?"是":"否"}</td>`);
                if (order.order_status !== 4) {
                    let btnGroup = "";
                    // 操作
                    if (!order.is_dispatched&&order.surety_id !== 0&&(order.purchasers_id !== 0&&order.seller_id !== 0))
                        btnGroup += orderShipment(order.order_id);
                    if(order.purchasers_id === 0||order.seller_id === 0)
                        btnGroup += generateLink(order.order_id);
                    btnGroup += transactionDisputes;
                    tr.append($(`<td><div class="mdui-btn-group">${btnGroup}</div></td>`));
                }else{
                    tr.append($(`<td><div class="mdui-btn-group"></div></td>`));
                }
                AllOrderList.append(tr);
            }

        },
        error: function (xhr) {
            Toast("错误", "服务器异常");
        }
    });
}
// 订单大厅
function OrderHall() {
    setTimeout(function () {
        window.location.href = '/static/orderLobby.html';
    });
}
// 退出登录
function Logout() {
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

// 保存信息
function SavaInfo(){
    ajax({
        method: 'POST',
        url: `${ServerAddr}/user/setSelfInfo`,
        data: {
            alipay_uname:PhoneNumber.text(), // 支付宝用户名
            user_lastname:Name.text(), // 姓名
            alipay_qrcode:Alipay.text(), // 支付宝二维码
            wxpay_qrcode:WeChat.text(), // 微信二维码
            bank_card:BankCardNumber.text(), // 银行卡号
            game_id:parseInt(GameID.text()) , // 游戏ID
            phone:parseInt(PhoneNumber.text()) , // 手机号码
        },
        handle: function (data) {
            if(data.code !== 200){
                Toast("提示",data.msg)
            }else{
                Toast("提示",data.msg)
            }
        }
    });
}
// 等候订单
function WaitingOrder(){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getWaitOrder`,
        success: function (data) {
            data = JSON.parse(data);
            
            WaitingOrderList.html("");
            for (let order of data) {
                let TransactionTypeName = getTransactionTypeName(order, UserInfoData);
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                // 时间
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                // 订单ID
                tr.append(`<td>${order.order_id}</td>`);
                // 订单类型
                tr.append(`<td>${TransactionTypeName}</td>`);
                // 单价
                tr.append(`<td>${order.prices}</td>`);
                // 数量
                tr.append(`<td>${order.count}</td>`);
                // 总价
                tr.append(`<td>${order.total_price}</td>`);
                // 游戏ID
                tr.append(`<td>${order.order_id}</td>`);
                // 担保人ID
                tr.append(`<td>${(order.surety_id===0)?"":order.surety_id}</td>`);
                
                WaitingOrderList.append(tr);
            }

        },
        error: function (xhr) {
            let rsp= JSON.parse(xhr.response)
            Toast("错误", rsp.msg);
        }
    });
}

// 担保人订单
function GuaranteeOrder() {
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/surety/suretyGetSelfOrder`,
        success: function (data) {
            data = JSON.parse(data);
            GuaranteeOrderList.html("");
            for (let order of data) {
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                // 时间
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                // 订单ID
                tr.append(`<td>${order.order_id}</td>`);
                // 单价
                tr.append(`<td>${order.prices}</td>`);
                // 数量
                tr.append(`<td>${order.count}</td>`);
                // 总价
                tr.append(`<td>${order.total_price}</td>`);
                // 是否发货
                tr.append(`<td>${(order.is_dispatched)?"是":"否"}</td>`);
                // 冻结原因
                if (order.order_status !== 4) {
                    tr.append(`<td></td>`);
                } else {
                    tr.append(`<td>${order.freeze_reasons}</td>`);
                }
                
                // 卖家信息
                tr.append($(`<td>${viewInfo(order.seller_id)}</td>`));
                // 买家信息
                tr.append($(`<td>${viewInfo(order.purchasers_id)}</td>`));
                
                GuaranteeOrderList.append(tr);
            }

        },
        error: function (xhr) {
            let rsp= JSON.parse(xhr.response)
            Toast("错误", rsp.msg);
        }
    });
}


// 冻结订单
function FrozenOrder(){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getFreezeOrder`,
        success: function (data) {
            data = JSON.parse(data);
            FrozenOrderList.html("");
            for (let order of data) {
                let TransactionTypeName = getTransactionTypeName(order, UserInfoData);
                let tr = $(`<tr class="${(order.product_type === 0) ? "gold-coins" : "gem"}"></tr>`);
                // 时间
                tr.append(`<td>${convertDateFormat(order.time)}</td>`);
                // 订单ID
                tr.append(`<td>${order.order_id}</td>`);
                // 订单类型
                tr.append(`<td>${TransactionTypeName}</td>`);
                // 单价
                tr.append(`<td>${order.prices}</td>`);
                // 数量
                tr.append(`<td>${order.count}</td>`);
                // 总价
                tr.append(`<td>${order.total_price}</td>`);
                // 游戏ID
                tr.append(`<td>${order.order_id}</td>`);
                // 担保人ID
                tr.append(`<td>${(order.surety_id===0)?"":order.surety_id}</td>`);
                // 冻结原因
                tr.append(`<td>${order.freeze_reasons}</td>`);

                // 是否发货
                tr.append(`<td>${(order.is_dispatched)?"是":"否"}</td>`);
                FrozenOrderList.append(tr);
            }

        },
        error: function (xhr) {
            let rsp= JSON.parse(xhr.response)
            Toast("错误", rsp.msg);
        }
    });
}
SwitchItemInit();
UserInfo();


function InviteNewUsers() {
    var list = $('#inviteUserList');
    // 清空容器
    list.html("");
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/user/getInvitedUser`,
        success: function (data) {
            data = JSON.parse(data);
            $.each(data,function(k,v){
                list.append(`<li class="mdui-list-item mdui-ripple">${v}</li>`);
            });
        },
        error: function (xhr) {
            let rsp = JSON.parse(xhr.response)
            Toast("错误", rsp.msg);
        }
    });
}


function InviteNewUsersToast(){
    Toast('新用户邀请链接（点击复制）',window.location.origin+"/static/register.html?invitation_code="+UserInfoData.user_id,"复制", function(inst){
        inst.close();
        copyTextToClipboard(InvitationPromotionalCopywriting+window.location.origin+"/static/register.html?invitation_code="+UserInfoData.user_id)
        Toast('提示','复制成功');
    });
}


function copyTextToClipboard(text) {
    var textArea = document.createElement("textarea");
    textArea.value = text;
    document.body.appendChild(textArea);
    textArea.select();

    try {
        var successful = document.execCommand('copy');
        var msg = successful ? 'successful' : 'unsuccessful';
        console.log('Copying text command was ' + msg);
    } catch (err) {
        console.log('Oops, unable to copy');
    }

    document.body.removeChild(textArea);
}



function EditDialog(self) {
    mdui.prompt('修改',
        function (value) {
            if (value != null && value !== "") {
                $(self).parent().parent().find("div > div").eq(2).text(value);
            }
        },
    );
}


// 检测手机号码是否有效
function CheckPhoneNumber(phoneNumber) {
    var regex = /^1[3-9]\d{9}$/;
    if (!regex.test(phoneNumber)) {
        return false;
    } else {
        return true;
    }
}
function getWxUrl(e, param) {
    analyticCode.getUrl(param, e, function (url, _) {
        if (url === "Failed to load the image"||url === "error decoding QR Code") {
            Toast("提示", "错误的二维码");
        } else {
            Toast("提示", "已选择地址为：" + url);
            WeChat.text(url);
        }
    });
}

function getAlipayUrl(e, param) {
    analyticCode.getUrl(param, e, function (url, _) {
        if (url === "Failed to load the image"||url === "error decoding QR Code") {
            Toast("提示", "错误的二维码");
        } else {
            Toast("提示", "已选择地址为：" + url);
            Alipay.text(url);
        }
    });
}

function getAliUrl(e, param) {
    console.log(this);
    analyticCode.getUrl(param, e, function (url, _) {
        if (url === "Failed to load the image") {
            Toast("提示", "错误的二维码");
        } else {
            Toast("提示", "已选择地址为：" + url);
            Alipay.text(url);

        }

    });
}


// 编辑框手机号处理
function ShowEditPhoneNumberDialog(self) {
    mdui.prompt('修改',
        function (value) {
            if (!CheckPhoneNumber(value)) {
                // 如果手机号码不符合规则，弹出对话框
                mdui.alert('请输入有效的中国手机号码', '手机号码错误');
                return;
            }
            if (value != null && value !== "") {
                PhoneNumber.text(value);
            }
            
        },
    );
}
// 金币提现
function GoldCoinWithdrawal(){
    mdui.prompt('请输入要提现的金币', '金币提现',
        function (value) {
            value = parseInt(value);
            if(isNaN(value)){
                mdui.alert('提现数额应该为数字');
                return;
            }

            ajax({
                method: 'POST',
                url: `${ServerAddr}/user/createWithDraw`,
                data:{
                    // 金币
                    product_type:0,
                    amount:value,
                },
                handle: function (data) {
                    if(data.code === 200){
                        UserInfoData.coin_balance -= value;
                        GoldCoinBalance.text(UserInfoData.coin_balance);
                    }
                    Toast("提示",data.msg);
                }
            });
        }
    );
}

// 金币转账
function GoldCoinTransaction(){
    Toast('金币余额转账',$("#TransactionTemplate").html(),"确定",
        function (inst) {
            inst.close();
            let el = ($(inst.$element[0]));
            let id = el.find("#TUserID").val();
            let value = parseInt(el.find("#TValue").val());
            if(id === ""){
                mdui.alert('ID不应该为空');
                return;
            }
            id = parseInt(id);
            if(isNaN(value)){
                mdui.alert('转账数额应该为数字');
                return;
            }

            ajax({
                method: 'POST',
                url: `${ServerAddr}/user/transferCoin`,
                data:{
                    to_user_id:id,
                    amount:value,
                },
                handle: function (data) {
                    if(data.code === 200){
                        UserInfoData.coin_balance -= value;
                        GoldCoinBalance.text(UserInfoData.coin_balance);
                    }
                    Toast("提示",data.msg);
                }
            });
        }
    );
}
// 余额提现
function BalanceWithdrawal(){
    mdui.prompt('请输入要提现的余额', '余额提现',
        function (value) {
            value = parseFloat(value);
            if(isNaN(value)){
                mdui.alert('提现数额应该为数字');
                return;
            }

            ajax({
                method: 'POST',
                url: `${ServerAddr}/user/createWithDraw`,
                data:{
                    // 余额
                    product_type:1,
                    amount:value,
                },
                handle: function (data) {
                    if(data.code === 200){
                        UserInfoData.balance -= value;
                        CashBalance.text(UserInfoData.balance);
                    }
                    Toast("提示",data.msg);
                }
            });
        }
    );
}
// 余额转账
function BalanceTransaction(){
    Toast('账户余额转账',$("#TransactionTemplate").html(),"确定",
        function (inst) {
            inst.close();
            let el = ($(inst.$element[0]));
            let id = el.find("#TUserID").val();
            let value = parseFloat(el.find("#TValue").val());
            if(id === ""){
                mdui.alert('ID不应该为空');
                return;
            }
            id = parseInt(id);
            if(isNaN(value)){
                mdui.alert('转账数额应该为数字');
                return;
            }

            ajax({
                method: 'POST',
                url: `${ServerAddr}/user/transferBalance`,
                data:{
                    to_user_id:id,
                    amount:value,
                },
                handle: function (data) {
                    if(data.code === 200){
                        UserInfoData.balance -= value;
                        CashBalance.text(UserInfoData.balance);
                    }
                    Toast("提示",data.msg);
                }
            });
        }
    );
}