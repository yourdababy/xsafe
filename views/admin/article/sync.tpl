{{ template "admin/base/header.tpl" . }}
{{ template "admin/base/nav.tpl" . }}

<div class="content-wrapper" style="min-height: 1342.88px;">
    <!-- Content Header (Page header) -->
    <section class="content-header">
        <div class="container-fluid">
            <div class="row mb-2">
                <div class="col-sm-6">
                    <h1>同步文章</h1>
                </div>
                <div class="col-sm-6">
                    <ol class="breadcrumb float-sm-right">
                        <li class="breadcrumb-item"><a href="#">Home</a></li>
                        <li class="breadcrumb-item active">同步文章</li>
                    </ol>
                </div>
            </div>
        </div><!-- /.container-fluid -->
    </section>

    <!-- Main content -->
    <section class="content">
        <div class="container-fluid">
            <div class="row">
                <div class="col-md-12">
                    <!-- general form elements -->
                    <div class="card">
                        <div class="card-body">
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">账户信息</label>
                                <div class="col-sm-1">
                                    <input type="text" class="form-control" id="username" placeholder="账号">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="password" placeholder="密码">
                                </div>
                            </div>
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">链接信息</label>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="host" placeholder="主机名">
                                </div>
                                <div class="col-sm-1">
                                    <input type="text" class="form-control" id="port" placeholder="端口号">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="dbname" placeholder="数据库">
                                </div>
                            </div>
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">文章表</label>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="articleTbName" placeholder="数据表名">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="titleField" placeholder="标题字段">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="contentField" placeholder="内容字段">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="classifyField" placeholder="分类字段">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="authorField" placeholder="作者字段">
                                </div>
                            </div>
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">分类表</label>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="classifyTbName" placeholder="数据表名">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="classifyIdField" placeholder="分类ID字段">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="classifyNameField" placeholder="分类名称字段">
                                </div>
                            </div>
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">作者表</label>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="authorTbName" placeholder="数据表名">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="authorIdField" placeholder="作者ID字段">
                                </div>
                                <div class="col-sm-2">
                                    <input type="text" class="form-control" id="authorNameField" placeholder="作者名称字段">
                                </div>
                            </div>
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">同步状态</label>
                                <div class="col-sm-5">
                                    <label class="col-form-label" id="status">等待同步中</label>
                                </div>
                            </div>
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">同步数量</label>
                                <div class="col-sm-5">
                                    <label class="col-form-label" id="syncCounts">0</label> / <label class="col-form-label" id="counts">0</label>
                                </div>
                            </div>
                            <div class="form-group row">
                                <label class="col-sm-1 col-form-label">错误信息</label>
                                <div class="col-sm-5">
                                    <label class="col-form-label" id="errMsg" style="color: red"></label>
                                </div>
                            </div>
                        </div>
                        <!-- /.card-body -->
                        <div class="card-footer">
                            <span class="btn btn-primary" id="doImport">开始同步</span>
                        </div>
                    </div>
                </div>
            </div>
            <!-- /.row -->
        </div><!-- /.container-fluid -->
    </section>
    <!-- /.content -->
</div>

{{ template "admin/base/footer.tpl" . }}

<script>
    var counts = 0;
    var timer;

    $(function () {
        $('#doImport').click(function (e) {
            $('#errMsg').html("");

            // 校验是否已经输入
            var data = checkInput();
            if (!data) return;

            // 执行ajax请求
            $.ajax({
                url: "/admin/article/do-sync",
                type: "POST",
                data: data,
                success: function (res) {
                    if (res.code != 200) {
                        $('#errMsg').html(res.msg)
                    } else {
                        // 正在同步中
                        $('#counts').html(res.data.counts)
                        $('#status').html(res.msg)

                        counts = res.data.counts

                        // 开启定时任务 请求同步状态
                        timer = setInterval(getSyncStatus, 500)
                    }
                },
                error: function (res) {
                    console.log(res)
                }
            })
        });
    });

    function getSyncStatus() {
        // 执行ajax请求
        $.ajax({
            url: "/admin/article/sync/status",
            type: "GET",
            success: function (res) {
                $('#syncCounts').html(res.data.counts)

                // 判断数量是否已经大于等于总数了，如果是 则改变状态
                if (res.data.counts >= counts) {
                    $('#status').html('同步完成！')
                    // 终止定时器
                    clearInterval(timer)
                }
            },
            error: function (res) {
                clearInterval(timer)
                console.log(res)
            }
        })
    }

    function checkInput() {
        var errMsg = $('#errMsg');

        var username = $('#username').val()
        if (!username) { errMsg.html("未输入数据库用户名！"); return false; }

        var password = $('#password').val();
        if (!password) { errMsg.html("未输入数据库密码！"); return false; }

        var host = $('#host').val();
        if (!host) { errMsg.html("未输入数据库主机名！"); return false; }

        var port = $('#port').val();
        if (!port) { errMsg.html("未输入数据库端口！"); return false; }

        var dbname = $('#dbname').val();
        if (!dbname) { errMsg.html("未输入数据库名称！"); return false; }

        var articleTbName = $('#articleTbName').val();
        if (!articleTbName) { errMsg.html("未输入文章表名！"); return false; }

        var titleField = $('#titleField').val();
        if (!titleField) { errMsg.html("未输入文章标题字段！"); return false; }

        var contentField = $('#contentField').val();
        if (!contentField) { errMsg.html("未输入文章内容字段！"); return false; }

        var classifyField = $('#classifyField').val();
        if (!classifyField) { errMsg.html("未输入分类字段！"); return false; }

        var authorField = $('#authorField').val();
        if (!authorField) { errMsg.html("未输入未输入作者字段！"); return false; }

        var classifyTbName = $('#classifyTbName').val();
        if (!classifyTbName) { errMsg.html("未输入分类数据表名！"); return false; }

        var classifyIdField = $('#classifyIdField').val();
        if (!classifyIdField) { errMsg.html("未输入分类数据表ID字段！"); return false; }

        var classifyNameField = $('#classifyNameField').val();
        if (!classifyNameField) { errMsg.html("未输入分类数据表名称字段！"); return false; }

        var authorTbName = $('#authorTbName').val();
        if (!authorTbName) { errMsg.html("未输入作者数据表名！"); return false; }

        var authorIdField = $('#authorIdField').val();
        if (!authorIdField) { errMsg.html("未输入作者ID字段！"); return false; }

        var authorNameField = $('#authorNameField').val();
        if (!authorNameField) { errMsg.html("未输入作者名称字段！"); return false; }

        return {
            username: username,
            password: password,
            host: host,
            port: port,
            dbname: dbname,
            articleTbName: articleTbName,
            titleField: titleField,
            contentField: contentField,
            classifyField: classifyField,
            authorField: authorField,
            classifyTbName: classifyTbName,
            classifyIdField: classifyIdField,
            classifyNameField: classifyNameField,
            authorTbName: authorTbName,
            authorIdField: authorIdField,
            authorNameField: authorNameField,
        };
    }

</script>

