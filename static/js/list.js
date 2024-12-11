$(document).ready(function(){
    // 初始化弹框
    $.ajax({
        url: "/getall",
        type: "GET",
        dataType: "json",
        success: function(data) {
            $.each(data, function(index, user) {
                // 修改这里，创建一个包含用户ID、用户姓名和用户性别的列表项
                var listItem = $("<li></li>");
                var userId = $("<span></span>").text("用户ID: " + user.id);
                var userName = $("<span></span>").text("用户姓名: " + user.uname);
                var userGender = $("<span></span>").text("用户性别: " + user.usex);
                listItem.append(userId);
                listItem.append(userName);
                listItem.append(userGender);
                // 添加删除按钮，并设置其点击事件
                var deleteButton = $("<button></button>").text("删除").click(function() {
                    $.ajax({
                        url: "/deleteuser/" + user.id,
                        type: "DELETE",
                        success: function() {
                            // 如果删除成功，则移除对应的列表项
                            listItem.remove();
                        }
                    });
                });

                // 添加修改按钮，并设置其点击事件
                var updateButton = $("<button></button>").text("修改").click(function() {
                    window.location.href="/Updatepage/"+user.id
                    // $.ajax({
                    //     url: "/findById/" + user.id,
                    //     type: "GET",
                    //     dataType: "json",
                    //     success: function(user) {
                    //         // 如果更新成功，则弹框显示用户信息
                    //         alert(JSON.stringify(user));
                    //         // window.location.href="/updateUser"
                    //     }
                    // });
                });

                // 将按钮和用户信息添加到列表项中
                listItem.append(deleteButton);
                listItem.append(updateButton);

                // 将列表项添加到页面中
                $("#userList").append(listItem);
            });
        }
    });
});

// 创建用户列表项
function createListItem(user) {
    var listItem = $("<li></li>");
    var userId = $("<span></span>").text("用户ID: " + user.id);
    var userName = $("<span></span>").text("用户姓名: " + user.uname);
    var userGender = $("<span></span>").text("用户性别: " + user.usex);

    listItem.append(userId);
    listItem.append(userName);
    listItem.append(userGender);
    listItem.append(createButtons(user.id));

    return listItem;
}