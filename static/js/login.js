$(document).ready(function() {
    $('#btn').click(function() {
        var username = $('#username').val();
        var password = $('#password').val();
        $.ajax({
            url: '/dologin', // 后端控制器路由地址
            type: 'POST',
            dataType: 'json',
            contentType: 'application/json',
            data: JSON.stringify({uname: username, upwd: password}),
            success: function (response) {
                // 处理返回的JSON数据
                alert("返回的数据：" + response.message)
                console.log(response)
                if (response.success) {
                    $('#result').html('<p>登录成功！</p>');
                } else {
                    $('#result').html('<p>登录失败：' + response.message + '</p>');
                }
            },
            error: function (xhr, status, error) {
                console.log('Error occurred: ' + error);
            }
        });
    });
})
