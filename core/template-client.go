package core

import "github.com/spf13/viper"

func HeaderTemplateClient() string {
	base_url := viper.GetString("base_url")
	html := `
	<!doctype html>
	<html class="no-js h-100" lang="en">
	  <head>
		<!-- Google Tag Manager -->
		<script>dataLayer = [{ 'ga-tracking-id' : 'UA-115105611-1' }];</script>
		<script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
			new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
			j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
			'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
			})(window,document,'script','dataLayer','GTM-WGLBNC8');</script>
		<!-- End Google Tag Manager -->
		<meta charset="utf-8">
		<meta http-equiv="x-ua-compatible" content="ie=edge">
		<title>Permadani rias</title>
		<meta name="description" content="A premium collection of beautiful hand-crafted Bootstrap 4 admin dashboard templates and dozens of custom components built for data-driven applications.">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<link href="https://use.fontawesome.com/releases/v5.0.6/css/all.css" rel="stylesheet">
		<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
		<link rel="stylesheet" id="main-stylesheet" data-version="1.3.1" href="` + base_url + `asset/shard/styles/shards-dashboards.1.1.0.min.css">
		<link rel="stylesheet" href="` + base_url + `asset/shard/styles/extras.1.1.0.min.css">
		<link rel="stylesheet" href="` + base_url + `asset/asset/css/global.css">
		<script async defer src="https://buttons.github.io/buttons.js"></script>
		<style>
			.text-info{
				color: #428fe1;
			}
			.bg-info{
				background-color: #428fe1;
				color: white;
			}
			.btn-white:not([disabled]):not(.disabled).active{
				box-shadow: none!important;
				background-color: #428fe1;
				border-color: #428fe1;
				color: #fff;
			}
		</style>
	  </head>
	  <body class="h-100">
		<div class="container-fluid">
		  <div class="row">
			<main class="main-content col-lg-12 col-md-12 col-sm-12 p-0">
			  <div class="main-navbar  bg-white">
				<div class="container p-0">
				  <!-- Main Navbar -->
				  <nav class="navbar align-items-stretch navbar-light flex-md-nowrap p-0">
					<a class="navbar-brand" href="#" style="line-height: 25px;">
					  <div class="d-table m-auto">
						<img id="main-logo" class="d-inline-block align-top mr-1 ml-3" style="max-width: 80px;" src="` + base_url + `asset/asset/image/permadani-rias-logo.svg" alt="Shards Dashboard">
						<!-- <span class="d-none d-md-inline ml-1">Shards Dashboard</span> -->
					  </div>
					</a>
					<form action="#" class="main-navbar__search w-100 d-none d-md-flex d-lg-flex">
					  <div class="input-group input-group-seamless ml-3">
						<div class="input-group-prepend">
						  <div class="input-group-text">
							<i class="fas fa-search"></i>
						  </div>
						</div>
						<input class="navbar-search form-control" type="text" placeholder="Search for something..." aria-label="Search">
					  </div>
					</form>
					<ul class="navbar-nav border-left flex-row border-right ml-auto">
					  <li class="nav-item border-right dropdown notifications">
						<a class="nav-link nav-link-icon text-center ml-0" href="#" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
						  <div class="nav-link-icon__wrapper">
							<i class="material-icons">&#xE7F4;</i>
							<span class="badge badge-pill badge-danger">2</span>
						  </div>
						</a>
						<div class="dropdown-menu dropdown-menu-small" aria-labelledby="dropdownMenuLink">
						  <a class="dropdown-item" href="#">
							<div class="notification__icon-wrapper">
							  <div class="notification__icon">
								<i class="material-icons">&#xE6E1;</i>
							  </div>
							</div>
							<div class="notification__content">
							  <span class="notification__category">Analytics</span>
							  <p>Your website’s active users count increased by <span class="text-success text-semibold">28%</span> in the last week. Great job!</p>
							</div>
						  </a>
						  <a class="dropdown-item" href="#">
							<div class="notification__icon-wrapper">
							  <div class="notification__icon">
								<i class="material-icons">&#xE8D1;</i>
							  </div>
							</div>
							<div class="notification__content">
							  <span class="notification__category">Sales</span>
							  <p>Last week your store’s sales count decreased by <span class="text-danger text-semibold">5.52%</span>. It could have been worse!</p>
							</div>
						  </a>
						  <a class="dropdown-item notification__all text-center" href="#"> View all Notifications </a>
						</div>
					  </li>
					  <li class="nav-item dropdown m-auto">
					  	<a class="nav-link  text-nowrap px-3 ml-0"  href="` + base_url + `" role="button" >
							<i class="fa fa-user"></i><span class="d-none d-md-inline-block" style="font-size:15px !important;font-weight:bold !important">Login</span>
						</a>
					  	<!-- <a class="nav-link dropdown-toggle text-nowrap px-3" data-toggle="dropdown" href="#" role="button" aria-haspopup="true" aria-expanded="false">
							<img class="user-avatar rounded-circle mr-2" src="` + base_url + `asset/shard/images/avatars/0.jpg" alt="User Avatar"> 
							<span class="d-none d-md-inline-block">Sierra Brooks</span>
						</a> -->
						<div class="dropdown-menu dropdown-menu-small">
							<a class="dropdown-item" href="user-profile.html"><i class="material-icons">&#xE7FD;</i> Profile</a>
							<a class="dropdown-item" href="edit-user-profile.html"><i class="material-icons">&#xE8B8;</i> Edit Profile</a>
							<a class="dropdown-item" href="file-manager-cards.html"><i class="material-icons">&#xE2C7;</i> Files</a>
							<a class="dropdown-item" href="transaction-history.html"><i class="material-icons">&#xE896;</i> Transactions</a>
							<div class="dropdown-divider"></div>
							<a class="dropdown-item text-danger" href="'.$base_url.'auth/logout">
							<i class="material-icons text-danger">&#xE879;</i> Logout </a>
						</div>
					  </li>
					</ul>
					<nav class="nav">
					  <a href="#" class="nav-link nav-link-icon toggle-sidebar  d-inline d-lg-none text-center " data-toggle="collapse" data-target=".header-navbar" aria-expanded="false" aria-controls="header-navbar">
						<i class="material-icons">&#xE5D2;</i>
					  </a>
					</nav>
				  </nav>
				</div> <!-- / .container -->
			  </div> <!-- / .main-navbar -->
			  <div class="header-navbar collapse d-lg-flex p-0 bg-white border-top">
				<div class="container">
				  <div class="row">
					<div class="col">
					  <ul class="nav nav-tabs border-0 flex-column flex-lg-row" id="container-menu">
						<li class="nav-item">
							<a href="` + base_url + `client" class="nav-link"><i class="material-icons">home</i> Home</a>
						</li>

						<li class="nav-item">
						  <a href="` + base_url + `client/catalog" class="nav-link"><i class="material-icons">view_day</i> Catalog</a>
						</li>

						<li class="nav-item">
						  <a href="` + base_url + `client/status" class="nav-link"><i class="material-icons">&#xE251;</i> Status</a>
						</li>

						<li class="nav-item">
						  <a href="` + base_url + `client/transaksi" class="nav-link"><i class="material-icons"></i> Transaksi</a>
						</li>

						<li class="nav-item">
						  <a href="` + base_url + `client/jadwal" class="nav-link"><i class="material-icons"></i> Jadwal</a>
						</li>

						<li class="nav-item dropdown">
						  <a class="nav-link" data-toggle="dropdown"><i class="material-icons">&#xE8B9;</i> User Account</a>
						  <div class="dropdown-menu dropdown-menu-small">
							<a href="file-manager-list.html" class="dropdown-item">User Profile</a>
							<a href="file-manager-cards.html" class="dropdown-item">Edit User Profile</a>
							<a href="login.html" class="dropdown-item">Login</a>
							<a href="register.html" class="dropdown-item">Register</a>
							<a href="forgot-password.html" class="dropdown-item">Forgot Password</a>
							<a href="change-password.html" class="dropdown-item">Change Password</a>
						  </div>
						</li>
						<li class="nav-item">
						  <a href="errors.html" class="nav-link"><i class="material-icons">error</i> Error</a>
						</li>
					  </ul>
					</div>
				  </div>
				</div>
			  </div>
	`
	return html
}

func FooterTemplateClient() string {
	base_url := viper.GetString("base_url")
	html := `
			<footer class="main-footer d-flex p-2 px-3 bg-white border-top">
			<ul class="nav">
			<li class="nav-item">
				<a class="nav-link" href="#">Home</a>
			</li>
			<li class="nav-item">
				<a class="nav-link" href="#">Services</a>
			</li>
			<li class="nav-item">
				<a class="nav-link" href="#">About</a>
			</li>
			<li class="nav-item">
				<a class="nav-link" href="#">Products</a>
			</li>
			<li class="nav-item">
				<a class="nav-link" href="#">Blog</a>
			</li>
			</ul>
			<span class="copyright ml-auto my-auto mr-2">Copyright © 2019 DesignRevision</span>
		</footer>
		</main>
		</div>
		</div>
		<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
		<script src="https://code.jquery.com/jquery-3.3.1.min.js" integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.1/Chart.min.js"></script>
		<script src="https://unpkg.com/shards-ui@latest/dist/js/shards.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/Sharrre/2.0.1/jquery.sharrre.min.js"></script>
		<script src="` + base_url + `asset/shard/scripts/extras.1.1.0.min.js"></script>
		<script src="` + base_url + `asset/shard/scripts/shards-dashboards.1.1.0.min.js"></script>
		<script src="` + base_url + `asset/asset/js/globalClient.js"></script>
		</body>
		</html>
	`

	return html
}
