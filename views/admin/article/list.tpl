{{ template "admin/base/header.tpl" . }}
{{ template "admin/base/nav.tpl" . }}

<!-- Content Wrapper. Contains page content -->
<div class="content-wrapper">
  <!-- Content Header (Page header) -->
  <section class="content-header">
    <div class="container-fluid">
      <div class="row mb-2">
        <div class="col-sm-6">
          <h1>{{ .classifyName }}</h1>
        </div>
        <div class="col-sm-6">
          <ol class="breadcrumb float-sm-right">
            <li class="breadcrumb-item"><a href="/admin">首页</a></li>
            <li class="breadcrumb-item active">文章列表</li>
          </ol>
        </div>
      </div>
    </div><!-- /.container-fluid -->
  </section>

  <!-- Main content -->
  <section class="content">

    <!-- Default box -->
    <div class="card">
      <div class="card-header">
        <h3 class="card-title">文章列表</h3>

        <div class="card-tools">
          <button type="button" class="btn btn-tool" data-card-widget="collapse" title="Collapse">
            <i class="fas fa-minus"></i>
          </button>
          <button type="button" class="btn btn-tool" data-card-widget="remove" title="Remove">
            <i class="fas fa-times"></i>
          </button>
        </div>
      </div>
      <div class="card-body p-0">
        <table class="table table-striped projects">
          <thead>
          <tr>
            <th style="width: 1%">
              #
            </th>
            <th style="width: 20%">
              标题
            </th>
            <th style="width: 30%">
              内容
            </th>
            <th>
              最后修改时间
            </th>
            <th style="width: 20%">
            </th>
          </tr>
          </thead>
          <tbody>
          {{ range .articles }}
          <tr>
            <td>
              #
            </td>
            <td>
              <a href="/admin/article/{{ .Id }}">
                {{ .Title }}
              </a>
            </td>
            <td class="content" style="max-width: 500px; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;">
              {{ .Content }}
            </td>
            <td class="update_at">
              {{ .UpdateAt | Date }}
            </td>
            <td class="project-actions text-right">
              <a class="btn btn-primary btn-sm" href="/admin/article/{{ .Id }}">
                <i class="fas fa-folder">
                </i>
                查看
              </a>
              <a class="btn btn-info btn-sm" href="/admin/article/{{ .Id }}/edit">
                <i class="fas fa-pencil-alt">
                </i>
                编辑
              </a>
              <a class="btn btn-danger btn-sm" href="javascript:void(0);" onclick="delArticle({{ .Id }})">
                <i class="fas fa-trash">
                </i>
                删除
              </a>
            </td>
          </tr>
          {{ end }}
          <tr>
              <td colspan="2">
                共 <span>{{ .counts }}</span> 条数据
              </td>
              <td colspan="3">
                <ul class="pagination" style="float: right; margin: 0 auto">
                  {{ range .pager }}
                    <li class="paginate_button page-item {{ if .active }} active {{ end }} {{ if .disable }} disabled {{ end }} ">
                      <a href="{{ .href }}" class="page-link">{{ str2html .name }}</a>
                    </li>
                  {{ end }}
                </ul>
              </td>
          </tr>
          </tbody>
        </table>

      </div>
      <!-- /.card-body -->
    </div>
    <!-- /.card -->

  </section>
  <!-- /.content -->
</div>
<!-- /.content-wrapper -->

<script>
  function delArticle(id) {
    if (confirm("确定删除该文章吗？") == true){
      // 执行ajax请求
      $.ajax({
        url: "/admin/article/"+id+"/del",
        type: "POST",
        success: function (res) {
          if (res.code == 200) {
            alert("删除成功！")
            window.location.href = document.URL
          } else {
            alert(res.msg)
          }
        },
        error: function (res) {
          alert("请求错误！状态码：" + res.status)
        }
      });
    }
  }
</script>

{{ template "admin/base/footer.tpl" . }}
