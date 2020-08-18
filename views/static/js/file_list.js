let layer;
let $;
let file_name = "";
let file_format = "";
let tableIns;
let data_count;
let curr_page;
let page_limit;

layui.use(['table', 'layer'], function () {
    let table = layui.table;
    layer = layui.layer;
    $ = layui.jquery;

    // path = App.GetPathFromLocalStorage()
    // depth = App.GetPathDepth()
    // fetch('/api/get_file_list?path=' + path + "&depth=" + depth, {
    //     method: 'get',
    //第一个实例
    tableIns = table.render({
        elem: '#file_list'
        , url: '/api/list' //数据接口
        , method: 'get'
        , limit: 20
        , loading: true
        , title: "文件列表"
        , page: true //开启分页
        , limits: [10, 15, 20, 30, 50, 80, 100]
        , toolbar: '#toolbarDemo'
        , request: {
            pageName: 'page' //页码的参数名称，默认：page
            , limitName: 'limit' //每页数据量的参数名，默认：limit
        }
        , where: {
            path: function () {
                return App.GetPathFromLocalStorage()
            },
            depth: function () {
                return 1
            },
            name: function () {
                return file_name
            },
            format: function () {
                return file_format
            },
        }
        , response: {
            statusName: 'code' //规定数据状态的字段名称，默认：code
            , statusCode: 200 //规定成功的状态码，默认：0
        }
        , parseData: function (res) { //res 即为原始返回的数据
            return {
                "code": res.code, //解析接口状态
                "msg": res.message, //解析提示文本
                "count": res.total, //解析数据长度
                "data": res.data //解析数据列表
            };
        }
        , cols: [[ //表头
            {type: 'numbers', fixed: 'left'}
            , {type: 'checkbox', fixed: 'left'}
            , {type: 'id', fixed: 'left', hide: true}
            , {field: 'file_name', title: '文件名称', align: 'center', fixed: 'left'}
            // , {field: 'file_path', title: '文件路径', align: 'center', fixed: 'left'}
            , {field: 'file_format', title: '文件格式', align: 'center', fixed: 'left'}
            , {field: 'file_perm', title: '文件权限', align: 'center', fixed: 'left'}
            , {field: 'size', title: '大小', align: 'center', fixed: 'left'}
            , {field: 'etc', title: '备注', align: 'center', fixed: 'left', edit: 'text'}
            , {field: 'created_at', title: '创建时间', align: 'center', fixed: 'left'}
            , {fixed: 'left', title: '操作', toolbar: '#barDemo', width: 180, align: 'center'}
        ]]
        , done: function (res, curr, count) {
            $("#file_list_path").val(res.data[0]["file_path"])
            page_limit = tableIns.config.limit;
            data_count = count;
            curr_page = curr;
        }
    });

    //监听行工具事件
    table.on('tool(file_list_table)', function (obj) {
        let data = obj.data;
        if (obj.event === 'update') {
            let request_json = []
            let item = {}
            item.xid = data.xid;
            item.file_name = data.file_name;
            request_json.push(item)
            bunchSave(request_json);
            reloadTable(curr_page)
        }

        if (obj.event === 'download') {
            let request_json = []
            let item = {}
            item.xid = data.xid;
            request_json.push(item)
            bunchDownload(request_json);
        }

        if (obj.event === 'delete') {
            let request_json = []
            let item = {}
            item.xid = data.xid;
            request_json.push(item)
            bunchDelete(request_json);
            reloadTable(curr_page)
        }
    });

    table.on('toolbar(file_list_table)', function (obj) {
        let event = obj.event;
        let checkStatus = table.checkStatus(obj.config.id);
        let data = checkStatus.data;
        let loop = data.length;
        switch (event) {
            case 'bunchSave' : {
                if (loop <= 0) {
                    layer.msg("请选中数据行！")
                } else {
                    let request_json = [];
                    for (let i = 0; i < loop; i++) {
                        let item = {}
                        item.uuid = data[i].id;
                        item.file_name = data[i].file_name;
                        request_json.push(item)
                    }
                    bunchSave(request_json);
                    reloadTable(curr_page)
                }
                break;
            }
            case 'bunchDelete' : {
                if (loop <= 0) {
                    layer.msg("请选中数据行！")
                } else {
                    if (window.confirm("确定删除选中的固定版型吗？(将同时删除该固定版型下关联的所有数据！)") === true) {
                        let request_json = [];
                        for (i = 0; i < loop; i++) {
                            let item = {};
                            item.id = data[i].id;
                            request_json.push(item)
                        }
                        bunchDelete(request_json, loop);
                    } else {
                        break;
                    }
                }

                break;
            }
        }
    })
});

function bunchSave(data) {
    let form_data = new FormData();
    form_data.append('data', JSON.stringify(data));

    fetch('/api/update', {
        method: 'post',
        body: form_data,
    }).then((req) => {
        req.json(() => {
        }).then((msg) => {
            if (msg.code !== 200) {
                // 0感叹号 1打钩 2打叉 3问号 4灰色的锁 5红色哭脸 6绿色笑脸
                layer.msg(msg.message, {icon: 5})
            } else {
                layer.msg("保存成功!", {icon: 1})
            }
        }, () => {
            layer.msg('请求失败', {icon: 5});
        })
    }, () => {
        layer.msg('请求失败', {icon: 5});
    })
}

function bunchDelete(data, loop) {
    let form_data = new FormData();
    form_data.append('data', JSON.stringify(data));

    fetch('/api/delete', {
        method: 'post',
        body: form_data,
    }).then((req) => {
        req.json(() => {
        }).then((msg) => {
            if (msg.code !== 200) {
                // 0感叹号 1打钩 2打叉 3问号 4灰色的锁 5红色哭脸 6绿色笑脸
                layer.msg(msg.message, {icon: 5})
            } else {
                layer.msg("删除成功！", {icon: 1});
                let page = curr_page;
                if (curr_page > 1) {
                    if (((data_count - loop) % page_limit) === 0) {
                        page = curr_page - 1;
                    }
                }
                reloadTable(page);
            }
        }, () => {
            layer.msg('请求失败', {icon: 5});
        })
    }, () => {
        layer.msg('请求失败', {icon: 5});
    })
}
function bunchDownload(data,loop){
    let form_data = new FormData();
    form_data.append('data', JSON.stringify(data));

    fetch('/api/download', {
        method: 'post',
        body: form_data,
    }).then((req) => {
        req.json(() => {
        }).then((msg) => {
            if (msg.code !== 200) {
                // 0感叹号 1打钩 2打叉 3问号 4灰色的锁 5红色哭脸 6绿色笑脸
                layer.msg(msg.message, {icon: 5})
            } else {
                layer.msg("下载成功！", {icon: 1});
                let page = curr_page;
                if (curr_page > 1) {
                    if (((data_count - loop) % page_limit) === 0) {
                        page = curr_page - 1;
                    }
                }
                reloadTable(page);
            }
        }, () => {
            layer.msg('请求失败', {icon: 5});
        })
    }, () => {
        layer.msg('请求失败', {icon: 5});
    })
}
function search() {
    file_name = $("#file_name").val();
    file_format = $("#file_format").val();
    reloadTable(1)
}

function reload() {
    file_name = "";
    file_format = "";
    $("#file_name").val("");
    $("#file_format").val("");
    reloadTable(1)
}

// 重载页面
function reloadTable(page) {
    tableIns.reload({
        page: {
            curr: page      // 重载
        }
    });
}

function get_path() {
    path =$(".breadcrumb > a").val()
    console.log("path=",path)
}

