package core

import "github.com/spf13/viper"

func HeaderTemplateAdmin() string {
	base_url := viper.GetString("base_url")
	html := `
	<!doctype html>
	<html class="no-js h-100" lang="en">
	  <head>
		<meta charset="utf-8">
		<meta http-equiv="x-ua-compatible" content="ie=edge">
		<title>Semen Indonesia Goup</title>
		<meta name="description" content="A high-quality &amp; free Bootstrap admin dashboard template pack that comes with lots of templates and components.">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<link href="https://use.fontawesome.com/releases/v5.0.6/css/all.css" rel="stylesheet">
		<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
		<link rel="stylesheet" id="main-stylesheet" data-version="1.1.0" href="` + base_url + `asset/shard/styles/shards-dashboards.1.1.0.min.css">
		<link rel="stylesheet" href="` + base_url + `asset/shard/styles/extras.1.1.0.min.css">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta2/css/all.min.css" integrity="sha512-YWzhKL2whUzgiheMoBFwW8CKV4qpHQAEuvilg9FAn5VJUDwKZZxkJNuGM4XkWuk94WCrrwslk8yWNGmY1EduTA==" crossorigin="anonymous" referrerpolicy="no-referrer" />
		<link rel="stylesheet" href="` + base_url + `asset/shard/datatable/dataTables.bootstrap4.min.css">
		<link rel="stylesheet" href="` + base_url + `asset/shard/datatable/responsive.dataTables.min.css">
		<link rel="stylesheet" href="` + base_url + `asset/shard/datatable/fixedColumns.dataTables.min.css">
		<script src="https://unpkg.com/sweetalert/dist/sweetalert.min.js"></script>
		<script async defer src="https://buttons.github.io/buttons.js"></script>
		
	  </head>
	  <body class="h-100">
		<div class="container-fluid">
		<div class="row">
		  <!-- Main Sidebar -->
		  <aside class="main-sidebar col-12 col-md-3 col-lg-2 px-0">
			<div class="main-navbar">
			  <nav class="navbar align-items-stretch navbar-light bg-white flex-md-nowrap border-bottom p-0">
			  <a class="navbar-brand w-100 mr-0" href="#" style="line-height: 25px;">
				<div class="d-flex justify-content-center align-items-center w-100">
				  <div class="mr-2">
					<img id="main-logo" class="d-inline-block align-top mr-1" style="max-width: 25px;" src="` + base_url + `asset/asset/admin/images/LogoJust.png" alt="Shards Dashboard">
				  </div>
				  <!-- <div class="d-flex" style="flex-direction: column; line-height: 16px;">
					<span>Permadani</span>
					<span>Rias</span>
				  </div> -->
				  <span class="d-none d-md-inline ml-1">Permadani Rias</span>
				</div>
			  </a>
			  <a class="toggle-sidebar d-sm-inline d-md-none d-lg-none">
				<i class="material-icons">&#xE5C4;</i>
			  </a>
			  </nav>
			</div>
			<form action="#" class="main-sidebar__search w-100 border-right d-sm-flex d-md-none d-lg-none">
			  <div class="input-group input-group-seamless ml-3">
			  <div class="input-group-prepend">
				<div class="input-group-text">
				<i class="fas fa-search"></i>
				</div>
			  </div>
			  <input class="navbar-search form-control" type="text" placeholder="Search for something..." aria-label="Search"> </div>
			</form>
			<div class="nav-wrapper">
			 
			  <ul class="nav flex-column">
				<ul class="nav flex-column">
				  <li class="nav-item">
					<a class="nav-link active" href="` + base_url + `admin">
					  <i class="material-icons">dashboard</i>
					  <span>Dashboard</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link active" href="` + base_url + `admin/data-penyewaan">
					  <i class="material-icons">dashboard</i>
					  <span>Dashboard</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link active" href="` + base_url + `admin/transaksi">
					  <i class="material-icons">dashboard</i>
					  <span>Transaksi</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link active" href="` + base_url + `admin/data-jadwal">
					  <i class="material-icons">dashboard</i>
					  <span>Manage jadwal</span>
					</a>
				  </li>
				  <li class="nav-item dropdown">
					<a class="nav-link dropdown-toggle" href="#" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
					  <i class="material-icons">vertical_split</i>
					  <span class="list-menu">Manage Barang</span>
					</a>
					<div class="dropdown-menu dropdown-menu-small" role="menu" aria-labelledby="dLabel">
					  <a class="dropdown-item" href="` + base_url + `admin/data-panggung">
						Data Panggung
					  </a>
					  <a class="dropdown-item" href="` + base_url + `admin/data-dekorasi">
						Data Dekorasi
					  </a>
					</div> 
				  </li>
				  <li class="nav-item dropdown">
					<a class="nav-link dropdown-toggle" href="#" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
					  <i class="material-icons">vertical_split</i>
					  <span class="list-menu">Keuangan</span>
					</a>
					<div class="dropdown-menu dropdown-menu-small" role="menu" aria-labelledby="dLabel">
					  <a class="dropdown-item" href="` + base_url + `admin/pemasukan">
						Pemasukan
					  </a>
					  <a class="dropdown-item" href="` + base_url + `admin/pengeluaran">
						Pengeluaran
					  </a>
					</div> 
				  </li>
				  
				  <li class="nav-item dropdown">
					<a class="nav-link dropdown-toggle" href="#" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
					  <i class="material-icons">settings</i>
					  <span class="list-menu">Setting</span>
					</a>
					<div class="dropdown-menu dropdown-menu-small" role="menu" aria-labelledby="dLabel">
					  <a class="dropdown-item" href="<?=base_url()?>admin/setting">
						Setting User
					  </a>
					  <a class="dropdown-item" href="http://localhost:8011/setting/user">
						User
					  </a>
					</div> 
				  </li>
				  <li class="nav-item">
					<a class="nav-link " href="transaction-history.html">
					  <i class="material-icons"></i>
					  <span>Transaction History</span>
					</a>
				  </li>
				
	
				  <!-- <li class="nav-item">
					<a class="nav-link active" href="<?=base_url()?>admin">
					  <i class="material-icons">settings</i>
					  <span>Setting User</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link " href="components-blog-posts.html">
					  <i class="material-icons">vertical_split</i>
					  <span>Blog Posts</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link " href="add-new-post.html">
					  <i class="material-icons">note_add</i>
					  <span>Add New Post</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link " href="form-components.html">
					  <i class="material-icons">view_module</i>
					  <span>Forms &amp; Components</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link " href="tables.html">
					  <i class="material-icons">table_chart</i>
					  <span>Tables</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link " href="user-profile-lite.html">
					  <i class="material-icons">person</i>
					  <span>User Profile</span>
					</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link " href="errors.html">
					  <i class="material-icons">error</i>
					  <span>Errors</span>
					</a>
				  </li> -->
				</ul>
			  </ul>
			</div>
		  </aside>
		  <!-- End Main Sidebar -->
		  <!-- ///////////////////////////////////////////////////////// -->
		  <main class="main-content col-lg-10 col-md-9 col-sm-12 p-0 offset-lg-2 offset-md-3">
	`
	return html
}

func FooterTemplateAdmin() string {
	base_url := viper.GetString("base_url")
	html := `
		</body>
		</html>
		<script src="https://code.jquery.com/jquery-3.3.1.min.js" integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8=" crossorigin="anonymous"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
			<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.1/Chart.min.js"></script>
			<script src="https://unpkg.com/shards-ui@latest/dist/js/shards.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/Sharrre/2.0.1/jquery.sharrre.min.js"></script>
			<script src="` + base_url + `asset/shard/scripts/extras.1.1.0.min.js"></script>
			<script src="` + base_url + `asset/shard/scripts/shards-dashboards.1.1.0.min.js"></script>
			<script src="` + base_url + `asset/shard/scripts/app/app-blog-overview.1.1.0.js"></script>
			<script src="https://cdn.datatables.net/1.10.19/js/jquery.dataTables.min.js"></script>
			<script src="https://cdn.datatables.net/1.10.20/js/dataTables.bootstrap4.min.js"></script>
			<script src="https://cdn.datatables.net/fixedcolumns/4.0.0/js/dataTables.fixedColumns.min.js"></script>
	`
	return html
}
