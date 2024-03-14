"use strict";
isLogin();
const params = new URLSearchParams(window.location.search);
const orderID = params.get('order_id');
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
                        handle: function (data) {
                            Toast("提示",data.msg)
                        }
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
let level = 0;
getPermission(function(data){
    if(data.code === 200){
        level = data.level;
    }else{
        Toast("提示","服务器异常");
    }
});

// 担保抢单
let GuaranteedGrabOrderBtn = function(id){
    return `<button class="mdui-btn mdui-color-green" onclick="GuaranteedGrabOrderClick(this,${id})">担保抢单</button>`;
}
function GuaranteedGrabOrderClick(self,id){
    self = $(self);
    ajax({
        method: 'POST',
        url: `${ServerAddr}/surety/carrySurety`,
        data:{
            order_id: id,
        },
        handle:function(data){
            if(data.code === 200){
                self.parent().parent().parent().remove();
            }
            Toast("提示",data.msg);
        }
    });
    

}
// 交易抢单
let TradeGrabOrderBtn = function(id){
    return `<button class="mdui-btn mdui-color-red" onclick="TradeGrabOrderClick(this,${id})">交易抢单</button>`;
}
function TradeGrabOrderClick(self,id){
    self = $(self);
    ajax({
        method: 'POST',
        url: `${ServerAddr}/transaction/carryTransaction`,
        data:{
            order_id: id,
        },
        handle:function(data){
            if(data.code === 200){
                self.parent().parent().parent().remove();
            }
            Toast("提示",data.msg);
        }
    });
}
// 订单大厅
function GrabOrder(){
    $.ajax({
        method: 'GET',
        url: `${ServerAddr}/transaction/getAllNormalOrder`,
        success: function (data) {
            data = JSON.parse(data);
            DataList.html("");
            for(let order of data){
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
                let btnGroup = "";
                // 操作
                // 担保抢单
                if (level > 2)
                    btnGroup += GuaranteedGrabOrderBtn(order.order_id);
                // 交易抢单
                btnGroup += TradeGrabOrderBtn(order.order_id);
                tr.append($(`<td><div class="mdui-btn-group">${btnGroup}</div></td>`));
                DataList.append(tr);
            }
        },
        error:function(xhr){
            var data = xhr.responseText ? JSON.parse(xhr.responseText) : { code: 500, msg: '服务器异常' };
            Toast('提示', data.msg);
        }
    });
}

// 订单大厅列表
let DataList = $("#Data");

$("#Quantity").on("input", function () {
    let quantity = $(this).val();
    let price = $("#Price").val();
    $("#Total").val(quantity * price);
});

SwitchItemInit();
// 把数值限制在100的倍数
function roundTo100 (self){
    var value = parseInt(self.value, 10);
    if (value % 100 !== 0) {
        self.value = Math.round(value / 100) * 100;
    }
}
// 创建订单
function CreateOrder() {
    // 订单类型，0为金币,1为宝石
    let orderType = 1;
    let inst = Toast('订单创建', $("#CreateOrderTemplate").html(), '创建', function (inst) {
        
        inst.close();
        let dialog = $(inst.$element);
        // 物品数量
        let quantity = dialog.find("#Quantity").val();
        // 物品单价
        let price = dialog.find("#Price").val();
        if (quantity <= 0 || price <= 0) {
            Toast('提示', "数量或单价不得小于0");
            return;
        } else {
            
            ajax({
                method: 'POST',
                url: `${ServerAddr}/transaction/createOrder`,
                data: {
                    product_type: orderType,
                    prices: parseInt(price),
                    count: parseInt(quantity),
                },
                handle: function (data) {
                    Toast('提示', data.msg);
                }
            });
        }
    }, false);
    let dialog = $(inst.$element);
    // 选择物品类型
    dialog.find("#CreateOrderSelect").on("change", function () {
        // select1是宝石，0是金币
        if($(this).val() === "select1"){
            orderType = 1;
            dialog.find("#Quantity").off("blur");
        }else{
            orderType = 0;
            roundTo100(dialog.find("#Quantity")[0]);
            dialog.find("#Quantity").on("blur", function () {   
                roundTo100(this);
            });
        }
    });
    mdui.mutation();
}
// 用户信息
function UserInfo(){
    setTimeout(function () {
        window.location.href = '/static/UserInfo.html';
    });
}
GrabOrder()