

$().ready(() => {

    let layer
    layui.use(['table', 'layer'], function () {
        layer = layui.layer;
    })

    var GetFileList = () => {
        path = App.GetPathFromLocalStorage()
        depth = App.GetPathDepth()
        fetch('/api/list_depth?path=' + path + "&depth=" + depth, {
            method: 'get',
        }).then((req) => {
            req.json(() => {
            }).then((msg) => {
                if (msg.code !== 200) {
                    // 0感叹号 1打钩 2打叉 3问号 4灰色的锁 5红色哭脸 6绿色笑脸
                    layer.msg(msg.message, {icon: 5})
                } else {
                    let d = msg.data
                    console.log("file_list\n", d)

                    $('#fileList').html("");

                    for (var i = 0; i < d.length; i++) {

                        var item = ""

                        if (d[i]["is_dir"]) {
                            item = "<tr class ='item'>" +
                                "<td>" + (i + 1) + "</td>" +
                                "<td><i class='layui-icon  layui-icon-list dir' style='font-size: 30px;'></i> " + d[i]["file_name"] + "</td>" +
                                "<td>" + d[i]["file_path"] + " </td>" +
                                "<td>" + d[i]["file_format"] + "</td>" +
                                "<td>" + d[i]["file_perm"] + "</td>" +
                                "<td>" + d[i]["file_size"] + "</td>" +
                                "<td>" + d[i]["created_at"] + "</td>" +
                                "</tr>";
                        } else {
                            item = "<tr class ='item'>" +
                                "<td>" + (i + 1) + "</td>" +
                                "<td><i class='layui-icon  layui-icon-file' style='font-size: 30px;'></i> " + d[i]["file_name"] + "</td>" +
                                "<td>" + d[i]["file_path"] + " </td>" +
                                "<td>" + d[i]["file_format"] + "</td>" +
                                "<td>" + d[i]["file_perm"] + "</td>" +
                                "<td>" + d[i]["file_size"] + "</td>" +
                                "<td>" + d[i]["created_at"] + "</td>" +
                                "</tr>";
                        }


                        $('#fileList').append(item);
                    }

                }
            }, () => {
                layer.msg('请求失败', {icon: 5});
            })
        }, () => {
            layer.msg('请求失败', {icon: 5});
        })
    }

    // 监听
    $('#path').change(() => {
        let path = $('#path').val()
        App.SetPathToLocalStorage(path)
    })


    $('#clearPath').click(() => {
        App.ClearPathValue()
    })

    $("#checkValid").click(() => {
        if (!App.CheckPathValid()) {
            if (App.os == "windows") {
                layer.msg("保存路径不合法,目录以'/'或者'\\\\'或者'\\'结尾", {icon: 5});
            } else {
                layer.msg("保存路径不合法,目录以'/'结尾", {icon: 5});
            }

            return
        }

        layer.msg('路径合法！', {icon: 1});
        // layer.tips('路径合法！', "#path");
    })

    $("#listFile").click(() => {
        GetFileList()
    })


})