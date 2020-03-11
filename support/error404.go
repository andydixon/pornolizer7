package support

func Error404() string {
	return `<html>
<head>
<title>The Pornolizer::Error</title>
<style>*{transition:all .6s}html{height:100%}body{font-family:Lato,sans-serif;color:#888;margin:0}#main{display:table;width:100%;height:100vh;text-align:center}.fof{display:table-cell;vertical-align:middle}.fof h1{font-size:50px;display:inline-block;padding-right:12px;animation:type .5s alternate infinite}@keyframes type{from{box-shadow:inset -3px 0 0 #888}to{box-shadow:inset -3px 0 0 transparent}}</style>
</head>
<body>
<div id="main">
    	<div class="fof">
        		<h1>Error 404</h1>
    	</div>
</div>
</body>
</html>
`
}
