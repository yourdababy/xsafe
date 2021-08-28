{{ template "admin/base/header.tpl" . }}
{{ template "admin/base/nav.tpl" . }}

<style>
  .code pre {
    background: #E8EAED;
  }
</style>

<!-- Content Wrapper. Contains page content -->
<div class="content-wrapper">
  <!-- Content Header (Page header) -->
  <section class="content-header">
    <div class="container-fluid">
      <div class="row mb-2">
        <div class="col-sm-6">
          <h1>文章详情</h1>
        </div>
        <div class="col-sm-6">
          <ol class="breadcrumb float-sm-right">
            <li class="breadcrumb-item"><a href="#">Home</a></li>
            <li class="breadcrumb-item active">文章详情</li>
          </ol>
        </div>
      </div>
    </div><!-- /.container-fluid -->
  </section>

  <!-- Main content -->
  <section class="content">
    <div class="container-fluid">
      <div class="row">
        <!-- /.col -->
      <div class="col-md-12">
        <div class="card card-primary card-outline">
          <div class="card-header">
            <h3 class="card-title">{{ .article.Title }}</h3>
          </div>
          <!-- /.card-header -->
          <div class="card-body p-0">
            <div class="mailbox-read-info">
              <h6>作者 :&nbsp; <span style="font-size: 14px">哈哈哈</span></h6>
              <h6>
                外链 :&nbsp; <a target="_blank" href="https://www.xsafe.org">Xsafe - 信息安全社区</a>
                <span class="mailbox-read-time float-right">修改时间 {{ .article.UpdateAt | Date }}</span>
              </h6>
            </div>
            <!-- /.mailbox-controls -->
            <div class="mailbox-read-message code" style="min-height: 500px">
              {{ .article.Content | MarkdownToHtml | str2html }}
            </div>
            <!-- /.mailbox-read-message -->
          </div>
        </div>
        <!-- /.card -->
      </div>
      <!-- /.col -->
    </div>
    <!-- /.row -->
    </div><!-- /.container-fluid -->
  </section>
  <!-- /.content -->
</div>

{{ template "admin/base/footer.tpl" . }}