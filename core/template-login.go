package core

func headerAuth()string{
	html:=`
		<!doctype html>
		<html lang="en">
		<head>
			<!-- Required meta tags -->
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<link href="https://fonts.googleapis.com/css?family=Roboto:300,400&display=swap" rel="stylesheet">
		
			<link rel="stylesheet" href="<?=base_url('shard/')?>login/fonts/icomoon/style.css">
		
			<link rel="stylesheet" href="<?=base_url('shard/')?>login/css/owl.carousel.min.css">
		
			<!-- Bootstrap CSS -->
			<link rel="stylesheet" href="<?=base_url('shard/')?>login/css/bootstrap.min.css">
			
			<!-- Style -->
			<link rel="stylesheet" href="<?=base_url('shard/')?>login/css/style.css">
		
			<title>Login #8</title>
		</head>
		<body>
	`

	return html
}

func footerAuth()  string{
	html:= `
			<script src="<?=base_url('shard/')?>login/js/jquery-3.3.1.min.js"></script>
			<script src="<?=base_url('shard/')?>login/js/popper.min.js"></script>
			<script src="<?=base_url('shard/')?>login/js/bootstrap.min.js"></script>
			<script src="<?=base_url('shard/')?>login/js/main.js"></script>
			<script src="https://unpkg.com/sweetalert/dist/sweetalert.min.js"></script>
		</body>
		</html>
	`

	return html
}