<!-- Main Sidebar Container -->
<aside class="main-sidebar sidebar-dark-primary elevation-4">
    <!-- Brand Logo -->
    <a href="/admin" class="brand-link">
        <img src="/static/dist/img/AdminLTELogo.png" alt="AdminLTE Logo" class="brand-image img-circle elevation-3" style="opacity: .8">
        <span class="brand-text font-weight-light">{{ .WebSite }}</span>
    </a>

    <!-- Sidebar -->
    <div class="sidebar">
        <!-- Sidebar user panel (optional) -->
        <div class="user-panel mt-3 pb-3 mb-3 d-flex">
            <div class="image">
                <img src="/static/dist/img/user2-160x160.jpg" class="img-circle elevation-2" alt="User Image">
            </div>
            <div class="info">
                <a href="#" class="d-block">Admin</a>
            </div>
        </div>

{{/*        <!-- SidebarSearch Form -->*/}}
{{/*        <div class="form-inline">*/}}
{{/*            <div class="input-group" data-widget="sidebar-search">*/}}
{{/*                <input class="form-control form-control-sidebar" type="search" placeholder="Search" aria-label="Search">*/}}
{{/*                <div class="input-group-append">*/}}
{{/*                    <button class="btn btn-sidebar">*/}}
{{/*                        <i class="fas fa-search fa-fw"></i>*/}}
{{/*                    </button>*/}}
{{/*                </div>*/}}
{{/*            </div>*/}}
{{/*        </div>*/}}

        <!-- Sidebar Menu -->
        <nav class="mt-2">
            <ul class="nav nav-pills nav-sidebar flex-column" data-widget="treeview" role="menu" data-accordion="false">
                <!-- Add icons to the links using the .nav-icon class
                     with font-awesome or any other icon font library -->
                <li class="nav-header">漏洞文章</li>
                <li class="nav-item menu-is-opening menu-open">
                    <a href="#" class="nav-link">
                        <i class="nav-icon fa fa-list"></i>
                        <p>
                            分类
                            <i class="fas fa-angle-left right"></i>
                        </p>
                    </a>
                    <ul class="nav nav-treeview">
                        {{ range .Classify}}
                        <li class="nav-item">
                            <a href="/admin/article-list/{{ .Id }}" class="nav-link {{if eq .Id $.ActiveId }} {{ print "active" }} {{ end }}">
                                <i class="far fa-circle nav-icon"></i>
                                <p>{{ .Name }}</p>
                            </a>
                        </li>
                        {{ end }}
                    </ul>
                </li>
                <li class="nav-item">
                    <a href="#" class="nav-link">
                        <i class="nav-icon fa fa-cog"></i>
                        <p>
                            工具
                            <i class="fas fa-angle-left right"></i>
                        </p>
                    </a>
                    <ul class="nav nav-treeview">
                        <li class="nav-item">
                            <a href="/admin/article/import" class="nav-link">
                                <i class="fa fa-upload nav-icon"></i>
                                <p>导入文章</p>
                            </a>
                        </li>
                        <li class="nav-item">
                            <a href="/admin/article/sync" class="nav-link">
                                <i class="fa fa-retweet nav-icon"></i>
                                <p>同步文章</p>
                            </a>
                        </li>
                    </ul>
                </li>
            </ul>
        </nav>
        <!-- /.sidebar-menu -->
    </div>
    <!-- /.sidebar -->
</aside>
