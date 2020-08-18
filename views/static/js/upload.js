$().ready(() => {

    layui.use(['upload', 'element', 'layer'], function () {
        // $ = layui.jquery;
        let upload = layui.upload;
        let layer = layui.layer;
        var element = layui.element;


        var process_i = 1
        var process_files_total = 1

        //文件上传
        //多文件列表示例
        var demoListView = $('#demoList')
            , uploadListIns = upload.render({
                elem: '#testList'
                , url: '/api/upload/'
                , data: {
                    path: () => {
                        return App.GetPathValue()
                    },
                    create_dir: () => {
                        return App.GetFlagCreateDir()
                    }
                }
                , accept: 'file'
                , multiple: true
                , auto: false
                , bindAction: '#testListAction'
                , number: 10
                , choose: function (obj) {
                    var files = this.files = obj.pushFile(); //将每次选择的文件追加到文件队列
                    //读取本地文件
                    obj.preview(function (index, file, result) {
                        var tr = $(['<tr id="upload-' + index + '">'
                            , '<td>' + file.name + '</td>'
                            , '<td>' + (file.size / 1024).toFixed(1) + 'kb</td>'
                            , '<td>等待上传</td>'
                            , '<td>'
                            , '<button class="layui-btn layui-btn-xs demo-reload layui-hide">重传</button>'
                            , '<button class="layui-btn layui-btn-xs layui-btn-danger demo-delete">删除</button>'
                            , '</td>'
                            , '</tr>'].join(''));

                        //单个重传
                        tr.find('.demo-reload').on('click', function () {
                            obj.upload(index, file);
                        });

                        //删除
                        tr.find('.demo-delete').on('click', function () {
                            delete files[index]; //删除对应的文件
                            tr.remove();
                            uploadListIns.config.elem.next()[0].value = ''; //清空 input file 值，以免删除后出现同名文件不可选
                        });

                        demoListView.append(tr);
                    });
                }
                , before: function (obj) { //obj参数包含的信息，跟 choose回调完全一致，可参见上文。
                    layer.load(); //上传loading
                    process_i = 1
                    process_files_total = Object.keys(this.files).length
                }
                , done: function (res, index, upload) {

                    layer.closeAll('loading');
                    if (res.code == 200) { //上传成功

                        let per = 100 / process_files_total * process_i
                        if (per >= 100) {
                            per = 100
                        }
                        per = Math.ceil(per)
                        element.progress('demo', per + '%');
                        process_i++;


                        errcode = res.errcode;
                        if (errcode == 0) {
                            var tr = demoListView.find('tr#upload-' + index)
                                , tds = tr.children();
                            tds.eq(2).html('<span style="color: #5FB878;">上传成功</span>');
                            tds.eq(3).html(''); //清空操作
                            return delete this.files[index]; //删除文件队列已经上传成功的文件

                        } else if (errcode == 1) {
                            var tr = demoListView.find('tr#upload-' + index)
                                , tds = tr.children();
                            tds.eq(2).html('<span style="color: #FFB800;">文件已存在</span>');
                            // tds.eq(3).html(''); //清空操作
                            // return delete this.files[index]; //删除文件队列已经上传成功的文件
                            return
                        }

                    }

                    this.error(index, upload);
                }
                ,
                allDone: function (obj) {
                    console.log("file")
                    console.log(" total:", obj.total); //得到总文件数
                    console.log(" success:", obj.successful); //请求成功的文件数
                    console.log(" fail:", obj.aborted); //请求失败的文件数
                    window.App.GetFileList()
                }
                ,
                error: function (index, upload) {
                    layer.closeAll('loading');
                    var tr = demoListView.find('tr#upload-' + index)
                        , tds = tr.children();
                    tds.eq(2).html('<span style="color: #FF5722;">上传失败</span>');
                    tds.eq(3).find('.demo-reload').removeClass('layui-hide'); //显示重传
                }
            })
        ;


        //多张图上传
        var uploadInst2 = upload.render({
            elem: '#upload_images'
            , url: '/api/upload/images/'
            , auto: false
            , bindAction: '#start_ocr_imgs'
            , accept: 'images'
            , multiple: true
            //, size: 1024 //KB
            , number: 10
            , choose: function (obj) {

                obj.preview(function (index, file, result) {
                    let img = `<img src="` + result + `" alt="` + file.name + `" class="layui-upload-img" style="height:100px;padding-right: 10px">`
                    // img_list = img_list + img
                    $("#images-list").append(img)
                });

                $('#start_ocr_imgs').attr('class', 'layui-btn layui-btn-normal')

            }
            , progress: function (n, elem) {
                var percent = n + '%' //获取进度百分比
                element.progress('demo', percent); //可配合 layui 进度条元素使用

                //以下系 layui 2.5.6 新增
                console.log(elem); //得到当前触发的元素 DOM 对象。可通过该元素定义的属性值匹配到对应的进度条。
            }
            , allDone: function (obj) {

                console.log("images")
                console.log(" total:", obj.total); //得到总文件数
                console.log(" success:", obj.successful); //请求成功的文件数
                console.log(" fail:", obj.aborted); //请求失败的文件数
                window.App.GetFileList()

            },
            done: function (res, index, upload) {
                layer.closeAll('loading'); //关闭loading
                //如果上传失败

                //如果上传失败
                if (res.code !== 200) {
                    layer.msg(res.message, {icon: 5})
                } else {
                    layer.msg('上传成功！', {icon: 1});
                }
            }
            , error: function () {
                //演示失败状态，并实现重传
                var demoText = $('#start_ocr_imgs');
                demoText.html('<a class="layui-btn layui-btn-normal demo-reload" style="float:left;margin-right: 15px"id="start_ocr_imgs">重试</a>');
                demoText.on('click', function () {
                    clear_result()
                    uploadInst2.upload();
                });
            }
        })
    })
})


// 重置upload images
$("#reset_upload_images").click(() => {
    html = `  
                <div class="layui-upload">
                    <button type="button" class="layui-btn" id="upload_images">多图片上传</button><input class="layui-upload-file" type="file" accept="undefined" name="file" multiple="">
                    <blockquote class="layui-elem-quote layui-quote-nm" style="margin-top: 10px;">
                        预览图：
                        <div id="images-list"></div>
                    </blockquote>
                </div>`
    setTimeout(() => {
        $('#start_ocr_imgs').attr('class', 'layui-btn layui-btn-disabled')
        $('#upload_images.layui-upload').html(html)
    }, 500)
})




