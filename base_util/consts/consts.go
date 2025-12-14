package consts

const (
	SwaggerUITemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name="description" content="SwaggerUI"/>
	<title>X Video Translator</title>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.10.5/swagger-ui.min.css" />
	<style>		
		#tokenInput {
			width: 400px;  /* 设置宽度 */
			height: 40px;  /* 设置高度 */
			font-size: 16px; /* 设置字体大小 */
			margin-bottom: 20px; /* 设置底部间距 */
			padding: 5px; /* 添加内边距 */
		}
	</style>
</head>
<body>
	<div>
		<label for="tokenInput">输入Authorization:</label>
		<input type="text" id="tokenInput" />
	</div>
	<div id="swagger-ui"></div>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.10.5/swagger-ui-bundle.js" crossorigin></script>
	<script>
		window.onload = () => {
			window.ui = SwaggerUIBundle({
				url: '{SwaggerUIDocUrl}',
				dom_id: '#swagger-ui',
				requestInterceptor: (request) => {
					const token = document.getElementById('tokenInput').value; // 获取输入的 token
					if (token) {
						request.headers['Authorization'] = token;
					}
					return request;
				},
			});
		};
	</script>
</body>
</html>
`
)
