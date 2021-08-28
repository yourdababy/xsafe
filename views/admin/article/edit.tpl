{{ template "admin/base/header.tpl" . }}
{{ template "admin/base/nav.tpl" . }}

<!-- Content Wrapper. Contains page content -->
<div class="content-wrapper" style="min-height: 1156.88px;">
    <!-- Content Header (Page header) -->
    <section class="content-header">
        <div class="container-fluid">
            <div class="row mb-2">
                <div class="col-sm-10">
                    <h1 style="text-overflow: ellipsis; overflow: hidden; white-space: nowrap;">编辑：{{ .Article.Title }}</h1>
                </div>
                <div class="col-sm-2">
                    <button class="btn btn-primary float-sm-right" id="save" style="margin-right: 20px">保存文章</button>
                </div>
                <div class="card-body">
                    <div class="form-group">
                        <label>标题</label>
                        <input type="text" value="{{ .Article.Title }}" id="title" class="form-control form-control-lg" placeholder="请输入标题">
                    </div>
                    <div class="form-group" style="height: 800px">
                        <label>内容</label>
                        <link rel="stylesheet" href="/static/markdown/css/editormd.css" />
                        <div id="test-editor">
                            <textarea id="content" style="display:none;">{{ .Article.Content }}</textarea>
                        </div>
                        <script src="https://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
                        <script src="/static/markdown/editormd.min.js"></script>
                        <script type="text/javascript">
                            $(function() {
                                var editor = editormd("test-editor", {
                                    width  : "100%",
                                    height : "100%",
                                    path   : "/static/markdown/lib/"
                                });
                            });
                        </script>
                    </div>
                </div>
            </div>
        </div><!-- /.container-fluid -->
    </section>
</div>

<script>
    $(function () {
        $('#save').click(function () {
            data = {
                title: $('#title').val(),
                content: $('#content').val()
            };

            // 执行ajax请求
            $.ajax({
                url: "/admin/article/{{ .Article.Id }}/update",
                type: "POST",
                data: data,
                success: function (res) {
                    if (res.code == 200) {
                        window.location.href = "/admin/article/{{ .Article.Id }}"
                    } else {
                        alert(res.msg)
                    }
                },
                error: function (res) {
                    alert("请求错误！状态码：" + res.status)
                }
            })
        });
    });
</script>

{{ template "admin/base/footer.tpl" . }}